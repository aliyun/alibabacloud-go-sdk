// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductionStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDeductionStatisticResponseBodyData) *DescribeDeductionStatisticResponseBody
	GetData() *DescribeDeductionStatisticResponseBodyData
	SetRequestId(v string) *DescribeDeductionStatisticResponseBody
	GetRequestId() *string
}

type DescribeDeductionStatisticResponseBody struct {
	Data *DescribeDeductionStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A1B2C3D4-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeductionStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticResponseBody) GetData() *DescribeDeductionStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeDeductionStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeductionStatisticResponseBody) SetData(v *DescribeDeductionStatisticResponseBodyData) *DescribeDeductionStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeductionStatisticResponseBody) SetRequestId(v string) *DescribeDeductionStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeductionStatisticResponseBodyData struct {
	AvailableCorePackages []*DescribeDeductionStatisticResponseBodyDataAvailableCorePackages `json:"AvailableCorePackages,omitempty" xml:"AvailableCorePackages,omitempty" type:"Repeated"`
	Deductions            []*DescribeDeductionStatisticResponseBodyDataDeductions            `json:"Deductions,omitempty" xml:"Deductions,omitempty" type:"Repeated"`
	Usages                []*DescribeDeductionStatisticResponseBodyDataUsages                `json:"Usages,omitempty" xml:"Usages,omitempty" type:"Repeated"`
}

func (s DescribeDeductionStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticResponseBodyData) GetAvailableCorePackages() []*DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	return s.AvailableCorePackages
}

func (s *DescribeDeductionStatisticResponseBodyData) GetDeductions() []*DescribeDeductionStatisticResponseBodyDataDeductions {
	return s.Deductions
}

func (s *DescribeDeductionStatisticResponseBodyData) GetUsages() []*DescribeDeductionStatisticResponseBodyDataUsages {
	return s.Usages
}

func (s *DescribeDeductionStatisticResponseBodyData) SetAvailableCorePackages(v []*DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) *DescribeDeductionStatisticResponseBodyData {
	s.AvailableCorePackages = v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyData) SetDeductions(v []*DescribeDeductionStatisticResponseBodyDataDeductions) *DescribeDeductionStatisticResponseBodyData {
	s.Deductions = v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyData) SetUsages(v []*DescribeDeductionStatisticResponseBodyDataUsages) *DescribeDeductionStatisticResponseBodyData {
	s.Usages = v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyData) Validate() error {
	if s.AvailableCorePackages != nil {
		for _, item := range s.AvailableCorePackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Deductions != nil {
		for _, item := range s.Deductions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usages != nil {
		for _, item := range s.Usages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeductionStatisticResponseBodyDataAvailableCorePackages struct {
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-12-31 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2024-12-31 23:59:59
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// g-xxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// CoreHour
	GroupResourceType *string `json:"GroupResourceType,omitempty" xml:"GroupResourceType,omitempty"`
	NoLx              *bool   `json:"NoLx,omitempty" xml:"NoLx,omitempty"`
	NoLxSource        *string `json:"NoLxSource,omitempty" xml:"NoLxSource,omitempty"`
	// example:
	//
	// res-xxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// CoreHour
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Active
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalTime *int64  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	UsedTime  *int64  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetGroupResourceType() *string {
	return s.GroupResourceType
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetNoLx() *bool {
	return s.NoLx
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetNoLxSource() *string {
	return s.NoLxSource
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetStatus() *string {
	return s.Status
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetAliUid(v int64) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.AliUid = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetCreateTime(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.CreateTime = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetEndTime(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.EndTime = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetExpiredTime(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetGroupId(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.GroupId = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetGroupResourceType(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.GroupResourceType = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetNoLx(v bool) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.NoLx = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetNoLxSource(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.NoLxSource = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetResourceId(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.ResourceId = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetResourceType(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.ResourceType = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetStartTime(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.StartTime = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetStatus(v string) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.Status = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetTotalTime(v int64) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.TotalTime = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) SetUsedTime(v int64) *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages {
	s.UsedTime = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataAvailableCorePackages) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductionStatisticResponseBodyDataDeductions struct {
	ConsumeSecond *int64 `json:"ConsumeSecond,omitempty" xml:"ConsumeSecond,omitempty"`
	// example:
	//
	// 2024-01-01
	DeductionDate *string `json:"DeductionDate,omitempty" xml:"DeductionDate,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeDeductionStatisticResponseBodyDataDeductions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticResponseBodyDataDeductions) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) GetConsumeSecond() *int64 {
	return s.ConsumeSecond
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) GetDeductionDate() *string {
	return s.DeductionDate
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) SetConsumeSecond(v int64) *DescribeDeductionStatisticResponseBodyDataDeductions {
	s.ConsumeSecond = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) SetDeductionDate(v string) *DescribeDeductionStatisticResponseBodyDataDeductions {
	s.DeductionDate = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) SetResourceType(v string) *DescribeDeductionStatisticResponseBodyDataDeductions {
	s.ResourceType = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataDeductions) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductionStatisticResponseBodyDataUsages struct {
	ConsumeSecond *int64 `json:"ConsumeSecond,omitempty" xml:"ConsumeSecond,omitempty"`
	// example:
	//
	// 2024-01-01
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeDeductionStatisticResponseBodyDataUsages) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticResponseBodyDataUsages) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) GetConsumeSecond() *int64 {
	return s.ConsumeSecond
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) GetPeriod() *string {
	return s.Period
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) SetConsumeSecond(v int64) *DescribeDeductionStatisticResponseBodyDataUsages {
	s.ConsumeSecond = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) SetPeriod(v string) *DescribeDeductionStatisticResponseBodyDataUsages {
	s.Period = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) SetResourceType(v string) *DescribeDeductionStatisticResponseBodyDataUsages {
	s.ResourceType = &v
	return s
}

func (s *DescribeDeductionStatisticResponseBodyDataUsages) Validate() error {
	return dara.Validate(s)
}
