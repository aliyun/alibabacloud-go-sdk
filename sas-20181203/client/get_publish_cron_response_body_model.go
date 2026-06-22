// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishCronResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPublishCronResponseBodyData) *GetPublishCronResponseBody
	GetData() *GetPublishCronResponseBodyData
	SetRequestId(v string) *GetPublishCronResponseBody
	GetRequestId() *string
}

type GetPublishCronResponseBody struct {
	// The publish scheduling configuration.
	Data *GetPublishCronResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublishCronResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublishCronResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublishCronResponseBody) GetData() *GetPublishCronResponseBodyData {
	return s.Data
}

func (s *GetPublishCronResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublishCronResponseBody) SetData(v *GetPublishCronResponseBodyData) *GetPublishCronResponseBody {
	s.Data = v
	return s
}

func (s *GetPublishCronResponseBody) SetRequestId(v string) *GetPublishCronResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublishCronResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPublishCronResponseBodyData struct {
	// The cron expression for the upgrade start time.
	//
	// example:
	//
	// 0 0 7 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The day of the week for the upgrade. Valid values:
	//
	// - **MON**: Monday
	//
	// - **TUE**: Tuesday
	//
	// - **WED**: Wednesday
	//
	// - **THU**: Thursday
	//
	// - **FRI**: Friday
	//
	// - **SAT**: Saturday
	//
	// - **SUN**: Sunday.
	//
	// example:
	//
	// MON
	CronDay *string `json:"CronDay,omitempty" xml:"CronDay,omitempty"`
	// The publish start timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1724522400000
	CronTime *int64 `json:"CronTime,omitempty" xml:"CronTime,omitempty"`
	// The upgrade start cycle type. Valid values:
	//
	// - **day**: every day
	//
	// - **week**: every week.
	//
	// example:
	//
	// day
	CronType *string `json:"CronType,omitempty" xml:"CronType,omitempty"`
	// The upgrade duration. Unit: hours.
	//
	// example:
	//
	// 24
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s GetPublishCronResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPublishCronResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPublishCronResponseBodyData) GetCron() *string {
	return s.Cron
}

func (s *GetPublishCronResponseBodyData) GetCronDay() *string {
	return s.CronDay
}

func (s *GetPublishCronResponseBodyData) GetCronTime() *int64 {
	return s.CronTime
}

func (s *GetPublishCronResponseBodyData) GetCronType() *string {
	return s.CronType
}

func (s *GetPublishCronResponseBodyData) GetDuration() *int32 {
	return s.Duration
}

func (s *GetPublishCronResponseBodyData) SetCron(v string) *GetPublishCronResponseBodyData {
	s.Cron = &v
	return s
}

func (s *GetPublishCronResponseBodyData) SetCronDay(v string) *GetPublishCronResponseBodyData {
	s.CronDay = &v
	return s
}

func (s *GetPublishCronResponseBodyData) SetCronTime(v int64) *GetPublishCronResponseBodyData {
	s.CronTime = &v
	return s
}

func (s *GetPublishCronResponseBodyData) SetCronType(v string) *GetPublishCronResponseBodyData {
	s.CronType = &v
	return s
}

func (s *GetPublishCronResponseBodyData) SetDuration(v int32) *GetPublishCronResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetPublishCronResponseBodyData) Validate() error {
	return dara.Validate(s)
}
