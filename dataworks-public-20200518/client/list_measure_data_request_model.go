// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMeasureDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentCode(v string) *ListMeasureDataRequest
	GetComponentCode() *string
	SetDomainCode(v string) *ListMeasureDataRequest
	GetDomainCode() *string
	SetEndTime(v int64) *ListMeasureDataRequest
	GetEndTime() *int64
	SetStartTime(v int64) *ListMeasureDataRequest
	GetStartTime() *int64
}

type ListMeasureDataRequest struct {
	// The measurement component. Valid values:
	//
	// 	- Count: phone call-based alerts and text message-based alerts
	//
	// This parameter is required.
	//
	// example:
	//
	// Count
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The measurement item. Valid values:
	//
	// 	- DideAlarmPhone: phone call-based alerts
	//
	// 	- DideAlarmSms: text message-based alerts
	//
	// This parameter is required.
	//
	// example:
	//
	// DideAlarmPhone
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// The end timestamp of the measurement period, in milliseconds. The measurement period is calculated in days. You can query only the data within the previous 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1717430400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start timestamp of the measurement period, in milliseconds. The measurement period is calculated in days. You can query only the data within the previous 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1717344000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListMeasureDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *ListMeasureDataRequest) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *ListMeasureDataRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *ListMeasureDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMeasureDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMeasureDataRequest) SetComponentCode(v string) *ListMeasureDataRequest {
	s.ComponentCode = &v
	return s
}

func (s *ListMeasureDataRequest) SetDomainCode(v string) *ListMeasureDataRequest {
	s.DomainCode = &v
	return s
}

func (s *ListMeasureDataRequest) SetEndTime(v int64) *ListMeasureDataRequest {
	s.EndTime = &v
	return s
}

func (s *ListMeasureDataRequest) SetStartTime(v int64) *ListMeasureDataRequest {
	s.StartTime = &v
	return s
}

func (s *ListMeasureDataRequest) Validate() error {
	return dara.Validate(s)
}
