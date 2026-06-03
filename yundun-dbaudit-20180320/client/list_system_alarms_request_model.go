// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemAlarmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmType(v int32) *ListSystemAlarmsRequest
	GetAlarmType() *int32
	SetBeginDate(v string) *ListSystemAlarmsRequest
	GetBeginDate() *string
	SetEndDate(v string) *ListSystemAlarmsRequest
	GetEndDate() *string
	SetInstanceId(v string) *ListSystemAlarmsRequest
	GetInstanceId() *string
	SetLang(v string) *ListSystemAlarmsRequest
	GetLang() *string
	SetRegionId(v string) *ListSystemAlarmsRequest
	GetRegionId() *string
}

type ListSystemAlarmsRequest struct {
	// example:
	//
	// 17
	AlarmType *int32 `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSystemAlarmsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsRequest) GetAlarmType() *int32 {
	return s.AlarmType
}

func (s *ListSystemAlarmsRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *ListSystemAlarmsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListSystemAlarmsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSystemAlarmsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSystemAlarmsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSystemAlarmsRequest) SetAlarmType(v int32) *ListSystemAlarmsRequest {
	s.AlarmType = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetBeginDate(v string) *ListSystemAlarmsRequest {
	s.BeginDate = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetEndDate(v string) *ListSystemAlarmsRequest {
	s.EndDate = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetInstanceId(v string) *ListSystemAlarmsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetLang(v string) *ListSystemAlarmsRequest {
	s.Lang = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetRegionId(v string) *ListSystemAlarmsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSystemAlarmsRequest) Validate() error {
	return dara.Validate(s)
}
