package swe

type ObservedAirTemperature struct {
  Id int        `gorm:"column:id;not null;primary_key"`
  Time time.Time `json:"time"`
  Temperature float64 `json:"temperature"`
}

func (ObservedAirTemperature) TableName() string {
	return "observed_air_temperature"
}