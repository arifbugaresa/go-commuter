package station

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/arifbugaresa/go-commuter/utils/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScheduleByStation(id string) (response []ScheduleResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second, // Configure the timeout or other settings
		},
	}
}

func (s *service) GetAllStation() (responses []StationResponse, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	if err = json.Unmarshal(byteResponse, &stations); err != nil {
		return
	}

	for _, item := range stations {
		responses = append(responses, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}

func (s *service) CheckScheduleByStation(id string) (response []ScheduleResponse, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var schedules []Schedule
	if err = json.Unmarshal(byteResponse, &schedules); err != nil {
		return
	}

	// filter schedule
	var scheduleSelected Schedule
	for _, schedule := range schedules {
		if schedule.Id == id {
			scheduleSelected = schedule
			break
		}
	}

	if scheduleSelected.Id == "" {
		err = errors.New("station not found")
		return
	}

	// convert schedule to better responses
	response, err = ConvertDataToResponses(scheduleSelected)
	if err != nil {
		return
	}

	return
}

func ConvertDataToResponses(schedule Schedule) (result []ScheduleResponse, err error) {
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI"
	)

	schedulesLebakBulus := schedule.SchduleLebakBulus
	schedulesBundaranHI := schedule.ScheduleBundaranHI

	// parsing to better time golang
	schedulesLebakBulusParsed, err := ConvertScheduleToTimeFormat(schedulesLebakBulus)
	if err != nil {
		return
	}

	schedulesBundaranHIParsed, err := ConvertScheduleToTimeFormat(schedulesBundaranHI)
	if err != nil {
		return
	}

	// convert to better response
	for _, item := range schedulesLebakBulusParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			result = append(result, ScheduleResponse{
				Name: LebakBulusTripName,
				Time: item.Format("15:04"),
			})
		}
	}

	for _, item := range schedulesBundaranHIParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			result = append(result, ScheduleResponse{
				Name: BundaranHITripName,
				Time: item.Format("15:04"),
			})
		}
	}

	return
}

func ConvertScheduleToTimeFormat(schedule string) (result []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schedule, ",")
	)

	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime == "" {
			continue
		}

		parsedTime, err = time.Parse("15:04", trimmedTime)
		if err != nil {
			err = errors.New(fmt.Sprintf("error parsing time %s", trimmedTime))
			return
		}

		result = append(result, parsedTime)
	}

	return
}
