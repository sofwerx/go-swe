package swe

type ObservedBlob struct {
	Id   int    `gorm:"column:id;not null;primary_key"`
	Blob []byte `json:"blob"`
}

func (ObservedBlob) TableName() string {
	return "observed_blob"
}
