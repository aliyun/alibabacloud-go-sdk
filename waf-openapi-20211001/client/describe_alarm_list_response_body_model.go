// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarms(v []*DescribeAlarmListResponseBodyAlarms) *DescribeAlarmListResponseBody
	GetAlarms() []*DescribeAlarmListResponseBodyAlarms
	SetRequestId(v string) *DescribeAlarmListResponseBody
	GetRequestId() *string
}

type DescribeAlarmListResponseBody struct {
	Alarms []*DescribeAlarmListResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// example:
	//
	// 8D8EBFB7-E1EB-5236-952A-092EDC72***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAlarmListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmListResponseBody) GetAlarms() []*DescribeAlarmListResponseBodyAlarms {
	return s.Alarms
}

func (s *DescribeAlarmListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlarmListResponseBody) SetAlarms(v []*DescribeAlarmListResponseBodyAlarms) *DescribeAlarmListResponseBody {
	s.Alarms = v
	return s
}

func (s *DescribeAlarmListResponseBody) SetRequestId(v string) *DescribeAlarmListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmListResponseBody) Validate() error {
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

type DescribeAlarmListResponseBodyAlarms struct {
	// example:
	//
	// 4count
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// example:
	//
	// 1605600798
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12000
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// example:
	//
	// 10000
	Spec *int64 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// 1605600767
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// qps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAlarmListResponseBodyAlarms) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmListResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *DescribeAlarmListResponseBodyAlarms) GetCause() *string {
	return s.Cause
}

func (s *DescribeAlarmListResponseBodyAlarms) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlarmListResponseBodyAlarms) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribeAlarmListResponseBodyAlarms) GetSpec() *int64 {
	return s.Spec
}

func (s *DescribeAlarmListResponseBodyAlarms) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlarmListResponseBodyAlarms) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAlarmListResponseBodyAlarms) GetType() *string {
	return s.Type
}

func (s *DescribeAlarmListResponseBodyAlarms) SetCause(v string) *DescribeAlarmListResponseBodyAlarms {
	s.Cause = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) SetEndTime(v int64) *DescribeAlarmListResponseBodyAlarms {
	s.EndTime = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) SetMaxQps(v int64) *DescribeAlarmListResponseBodyAlarms {
	s.MaxQps = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) SetSpec(v int64) *DescribeAlarmListResponseBodyAlarms {
	s.Spec = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) SetStartTime(v int64) *DescribeAlarmListResponseBodyAlarms {
	s.StartTime = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) SetStatus(v int32) *DescribeAlarmListResponseBodyAlarms {
	s.Status = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) SetType(v string) *DescribeAlarmListResponseBodyAlarms {
	s.Type = &v
	return s
}

func (s *DescribeAlarmListResponseBodyAlarms) Validate() error {
	return dara.Validate(s)
}
