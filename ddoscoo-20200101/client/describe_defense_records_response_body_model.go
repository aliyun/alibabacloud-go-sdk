// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseRecords(v []*DescribeDefenseRecordsResponseBodyDefenseRecords) *DescribeDefenseRecordsResponseBody
	GetDefenseRecords() []*DescribeDefenseRecordsResponseBodyDefenseRecords
	SetRequestId(v string) *DescribeDefenseRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDefenseRecordsResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseRecordsResponseBody struct {
	// An array that consists of details of the log of an advanced mitigation session.
	DefenseRecords []*DescribeDefenseRecordsResponseBodyDefenseRecords `json:"DefenseRecords,omitempty" xml:"DefenseRecords,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of advanced mitigation sessions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsResponseBody) GetDefenseRecords() []*DescribeDefenseRecordsResponseBodyDefenseRecords {
	return s.DefenseRecords
}

func (s *DescribeDefenseRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseRecordsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseRecordsResponseBody) SetDefenseRecords(v []*DescribeDefenseRecordsResponseBodyDefenseRecords) *DescribeDefenseRecordsResponseBody {
	s.DefenseRecords = v
	return s
}

func (s *DescribeDefenseRecordsResponseBody) SetRequestId(v string) *DescribeDefenseRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBody) SetTotalCount(v int64) *DescribeDefenseRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBody) Validate() error {
	if s.DefenseRecords != nil {
		for _, item := range s.DefenseRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDefenseRecordsResponseBodyDefenseRecords struct {
	// The peak attack traffic. Unit: bit/s.
	//
	// example:
	//
	// 6584186000
	AttackPeak *int64 `json:"AttackPeak,omitempty" xml:"AttackPeak,omitempty"`
	// The end time of the advanced mitigation session. This value is a UNIX timestamp. Units: miliseconds.
	//
	// example:
	//
	// 1583683200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of attacks.
	//
	// example:
	//
	// 2
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time of the advanced mitigation session. This value is a UNIX timestamp. Units: miliseconds.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the advanced mitigation session. Valid values:
	//
	// 	- **0**: The advanced mitigation session is being used.
	//
	// 	- **1**: The advanced mitigation session is used.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDefenseRecordsResponseBodyDefenseRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRecordsResponseBodyDefenseRecords) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) GetAttackPeak() *int64 {
	return s.AttackPeak
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) GetEventCount() *int32 {
	return s.EventCount
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetAttackPeak(v int64) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.AttackPeak = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetEndTime(v int64) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetEventCount(v int32) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.EventCount = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetInstanceId(v string) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetStartTime(v int64) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetStatus(v int32) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) Validate() error {
	return dara.Validate(s)
}
