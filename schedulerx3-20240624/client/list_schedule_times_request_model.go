// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListScheduleTimesRequest
	GetAppName() *string
	SetCalendar(v string) *ListScheduleTimesRequest
	GetCalendar() *string
	SetClusterId(v string) *ListScheduleTimesRequest
	GetClusterId() *string
	SetTimeExpression(v string) *ListScheduleTimesRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *ListScheduleTimesRequest
	GetTimeType() *int32
	SetTimeZone(v string) *ListScheduleTimesRequest
	GetTimeZone() *string
}

type ListScheduleTimesRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the calendar to use for scheduling, such as a business day calendar.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time expression, such as a cron expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. The only supported type is cron.
	//
	// - 1: cron
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone used to evaluate the time expression.
	//
	// example:
	//
	// Asia/Beijing
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ListScheduleTimesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTimesRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleTimesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListScheduleTimesRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *ListScheduleTimesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListScheduleTimesRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *ListScheduleTimesRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *ListScheduleTimesRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListScheduleTimesRequest) SetAppName(v string) *ListScheduleTimesRequest {
	s.AppName = &v
	return s
}

func (s *ListScheduleTimesRequest) SetCalendar(v string) *ListScheduleTimesRequest {
	s.Calendar = &v
	return s
}

func (s *ListScheduleTimesRequest) SetClusterId(v string) *ListScheduleTimesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListScheduleTimesRequest) SetTimeExpression(v string) *ListScheduleTimesRequest {
	s.TimeExpression = &v
	return s
}

func (s *ListScheduleTimesRequest) SetTimeType(v int32) *ListScheduleTimesRequest {
	s.TimeType = &v
	return s
}

func (s *ListScheduleTimesRequest) SetTimeZone(v string) *ListScheduleTimesRequest {
	s.TimeZone = &v
	return s
}

func (s *ListScheduleTimesRequest) Validate() error {
	return dara.Validate(s)
}
