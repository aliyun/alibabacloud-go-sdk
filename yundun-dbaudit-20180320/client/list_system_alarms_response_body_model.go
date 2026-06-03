// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemAlarmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarms(v []*ListSystemAlarmsResponseBodyAlarms) *ListSystemAlarmsResponseBody
	GetAlarms() []*ListSystemAlarmsResponseBodyAlarms
	SetRequestId(v string) *ListSystemAlarmsResponseBody
	GetRequestId() *string
}

type ListSystemAlarmsResponseBody struct {
	Alarms []*ListSystemAlarmsResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSystemAlarmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsResponseBody) GetAlarms() []*ListSystemAlarmsResponseBodyAlarms {
	return s.Alarms
}

func (s *ListSystemAlarmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemAlarmsResponseBody) SetAlarms(v []*ListSystemAlarmsResponseBodyAlarms) *ListSystemAlarmsResponseBody {
	s.Alarms = v
	return s
}

func (s *ListSystemAlarmsResponseBody) SetRequestId(v string) *ListSystemAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemAlarmsResponseBody) Validate() error {
	if s.Alarms != nil {
		for _, item := range s.Alarms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSystemAlarmsResponseBodyAlarms struct {
	AlarmDetail *string `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty"`
	// example:
	//
	// 1****
	AlarmId *int32 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// 17
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// 2019-06-06 05:03:17
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	ReadMark *int32 `json:"ReadMark,omitempty" xml:"ReadMark,omitempty"`
}

func (s ListSystemAlarmsResponseBodyAlarms) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAlarmsResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsResponseBodyAlarms) GetAlarmDetail() *string {
	return s.AlarmDetail
}

func (s *ListSystemAlarmsResponseBodyAlarms) GetAlarmId() *int32 {
	return s.AlarmId
}

func (s *ListSystemAlarmsResponseBodyAlarms) GetAlarmType() *string {
	return s.AlarmType
}

func (s *ListSystemAlarmsResponseBodyAlarms) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSystemAlarmsResponseBodyAlarms) GetReadMark() *int32 {
	return s.ReadMark
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetAlarmDetail(v string) *ListSystemAlarmsResponseBodyAlarms {
	s.AlarmDetail = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetAlarmId(v int32) *ListSystemAlarmsResponseBodyAlarms {
	s.AlarmId = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetAlarmType(v string) *ListSystemAlarmsResponseBodyAlarms {
	s.AlarmType = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetCreateTime(v string) *ListSystemAlarmsResponseBodyAlarms {
	s.CreateTime = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetReadMark(v int32) *ListSystemAlarmsResponseBodyAlarms {
	s.ReadMark = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) Validate() error {
	return dara.Validate(s)
}
