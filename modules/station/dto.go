package station

type Station struct {
	Id   string `json:"nid"`
	Name string `json:"title"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	Id                 string `json:"nid"`
	Name               string `json:"title"`
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	SchduleLebakBulus  string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	Name string `json:"name"`
	Time string `json:"time"`
}
