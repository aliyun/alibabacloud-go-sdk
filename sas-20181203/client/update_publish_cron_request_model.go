// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishCronRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCron(v string) *UpdatePublishCronRequest
	GetCron() *string
	SetCronDay(v string) *UpdatePublishCronRequest
	GetCronDay() *string
	SetCronTime(v int64) *UpdatePublishCronRequest
	GetCronTime() *int64
	SetCronType(v string) *UpdatePublishCronRequest
	GetCronType() *string
	SetDuration(v int32) *UpdatePublishCronRequest
	GetDuration() *int32
}

type UpdatePublishCronRequest struct {
	// The cron expression for the upgrade start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 5 10 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The day of the week for the upgrade time. Valid values:
	//
	// - **MON**: Monday.
	//
	// - **TUE**: Tuesday.
	//
	// - **WED**: Wednesday.
	//
	// - **THU**: Thursday.
	//
	// - **FRI**: Friday.
	//
	// - **SAT**: Saturday.
	//
	// - **SUN**: Sunday.
	//
	// example:
	//
	// SUN
	CronDay *string `json:"CronDay,omitempty" xml:"CronDay,omitempty"`
	// The start timestamp of the release, in milliseconds.
	//
	// example:
	//
	// 1657407600000
	CronTime *int64 `json:"CronTime,omitempty" xml:"CronTime,omitempty"`
	// The cycle type for the upgrade start time. Valid values:
	//
	// - **day**: Daily.
	//
	// - **week**: Weekly.
	//
	// example:
	//
	// day
	CronType *string `json:"CronType,omitempty" xml:"CronType,omitempty"`
	// The duration of the upgrade, in hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s UpdatePublishCronRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishCronRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublishCronRequest) GetCron() *string {
	return s.Cron
}

func (s *UpdatePublishCronRequest) GetCronDay() *string {
	return s.CronDay
}

func (s *UpdatePublishCronRequest) GetCronTime() *int64 {
	return s.CronTime
}

func (s *UpdatePublishCronRequest) GetCronType() *string {
	return s.CronType
}

func (s *UpdatePublishCronRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *UpdatePublishCronRequest) SetCron(v string) *UpdatePublishCronRequest {
	s.Cron = &v
	return s
}

func (s *UpdatePublishCronRequest) SetCronDay(v string) *UpdatePublishCronRequest {
	s.CronDay = &v
	return s
}

func (s *UpdatePublishCronRequest) SetCronTime(v int64) *UpdatePublishCronRequest {
	s.CronTime = &v
	return s
}

func (s *UpdatePublishCronRequest) SetCronType(v string) *UpdatePublishCronRequest {
	s.CronType = &v
	return s
}

func (s *UpdatePublishCronRequest) SetDuration(v int32) *UpdatePublishCronRequest {
	s.Duration = &v
	return s
}

func (s *UpdatePublishCronRequest) Validate() error {
	return dara.Validate(s)
}
