// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCorePackageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCorePackageListResponseBody
	GetCode() *string
	SetCorePackageInfo(v *DescribeCorePackageListResponseBodyCorePackageInfo) *DescribeCorePackageListResponseBody
	GetCorePackageInfo() *DescribeCorePackageListResponseBodyCorePackageInfo
	SetMessage(v string) *DescribeCorePackageListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCorePackageListResponseBody
	GetRequestId() *string
	SetTraceId(v string) *DescribeCorePackageListResponseBody
	GetTraceId() *string
}

type DescribeCorePackageListResponseBody struct {
	Code            *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	CorePackageInfo *DescribeCorePackageListResponseBodyCorePackageInfo `json:"CorePackageInfo,omitempty" xml:"CorePackageInfo,omitempty" type:"Struct"`
	Message         *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId         *string                                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeCorePackageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCorePackageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCorePackageListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCorePackageListResponseBody) GetCorePackageInfo() *DescribeCorePackageListResponseBodyCorePackageInfo {
	return s.CorePackageInfo
}

func (s *DescribeCorePackageListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCorePackageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCorePackageListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeCorePackageListResponseBody) SetCode(v string) *DescribeCorePackageListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCorePackageListResponseBody) SetCorePackageInfo(v *DescribeCorePackageListResponseBodyCorePackageInfo) *DescribeCorePackageListResponseBody {
	s.CorePackageInfo = v
	return s
}

func (s *DescribeCorePackageListResponseBody) SetMessage(v string) *DescribeCorePackageListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCorePackageListResponseBody) SetRequestId(v string) *DescribeCorePackageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCorePackageListResponseBody) SetTraceId(v string) *DescribeCorePackageListResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeCorePackageListResponseBody) Validate() error {
	if s.CorePackageInfo != nil {
		if err := s.CorePackageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCorePackageListResponseBodyCorePackageInfo struct {
	CorePackageList           []*DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList `json:"CorePackageList,omitempty" xml:"CorePackageList,omitempty" type:"Repeated"`
	SummaryRemainingCoreHours *float32                                                             `json:"SummaryRemainingCoreHours,omitempty" xml:"SummaryRemainingCoreHours,omitempty"`
}

func (s DescribeCorePackageListResponseBodyCorePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCorePackageListResponseBodyCorePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfo) GetCorePackageList() []*DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	return s.CorePackageList
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfo) GetSummaryRemainingCoreHours() *float32 {
	return s.SummaryRemainingCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfo) SetCorePackageList(v []*DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) *DescribeCorePackageListResponseBodyCorePackageInfo {
	s.CorePackageList = v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfo) SetSummaryRemainingCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfo {
	s.SummaryRemainingCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfo) Validate() error {
	if s.CorePackageList != nil {
		for _, item := range s.CorePackageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList struct {
	AssignedCoreHours     *float32                                                                             `json:"AssignedCoreHours,omitempty" xml:"AssignedCoreHours,omitempty"`
	DeductionInstanceList []interface{}                                                                        `json:"DeductionInstanceList,omitempty" xml:"DeductionInstanceList,omitempty" type:"Repeated"`
	ExpireTime            *string                                                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId            *string                                                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PeriodEndTime         *string                                                                              `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	PeriodStartTime       *string                                                                              `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	ProductType           *string                                                                              `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RemainingCoreHours    *float32                                                                             `json:"RemainingCoreHours,omitempty" xml:"RemainingCoreHours,omitempty"`
	RemainingPeriods      []*DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods `json:"RemainingPeriods,omitempty" xml:"RemainingPeriods,omitempty" type:"Repeated"`
	StartTime             *string                                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCoreHours        *float32                                                                             `json:"TotalCoreHours,omitempty" xml:"TotalCoreHours,omitempty"`
	UnassignedCoreHours   *float32                                                                             `json:"UnassignedCoreHours,omitempty" xml:"UnassignedCoreHours,omitempty"`
	UsedCoreHours         *float32                                                                             `json:"UsedCoreHours,omitempty" xml:"UsedCoreHours,omitempty"`
}

func (s DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GoString() string {
	return s.String()
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetAssignedCoreHours() *float32 {
	return s.AssignedCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetDeductionInstanceList() []interface{} {
	return s.DeductionInstanceList
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetRemainingCoreHours() *float32 {
	return s.RemainingCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetRemainingPeriods() []*DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	return s.RemainingPeriods
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetStatus() *string {
	return s.Status
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetTotalCoreHours() *float32 {
	return s.TotalCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetUnassignedCoreHours() *float32 {
	return s.UnassignedCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) GetUsedCoreHours() *float32 {
	return s.UsedCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetAssignedCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.AssignedCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetDeductionInstanceList(v []interface{}) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.DeductionInstanceList = v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetExpireTime(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetInstanceId(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.InstanceId = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetPeriodEndTime(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.PeriodEndTime = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetPeriodStartTime(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.PeriodStartTime = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetProductType(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.ProductType = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetRemainingCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.RemainingCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetRemainingPeriods(v []*DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.RemainingPeriods = v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetStartTime(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.StartTime = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetStatus(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.Status = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetTotalCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.TotalCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetUnassignedCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.UnassignedCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) SetUsedCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList {
	s.UsedCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageList) Validate() error {
	if s.RemainingPeriods != nil {
		for _, item := range s.RemainingPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods struct {
	PeriodEndTime      *string  `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	PeriodStartTime    *string  `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	RemainingCoreHours *float32 `json:"RemainingCoreHours,omitempty" xml:"RemainingCoreHours,omitempty"`
	Status             *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCoreHours     *float32 `json:"TotalCoreHours,omitempty" xml:"TotalCoreHours,omitempty"`
	UsedCoreHours      *float32 `json:"UsedCoreHours,omitempty" xml:"UsedCoreHours,omitempty"`
}

func (s DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GoString() string {
	return s.String()
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GetRemainingCoreHours() *float32 {
	return s.RemainingCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GetStatus() *string {
	return s.Status
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GetTotalCoreHours() *float32 {
	return s.TotalCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) GetUsedCoreHours() *float32 {
	return s.UsedCoreHours
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) SetPeriodEndTime(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	s.PeriodEndTime = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) SetPeriodStartTime(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	s.PeriodStartTime = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) SetRemainingCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	s.RemainingCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) SetStatus(v string) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	s.Status = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) SetTotalCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	s.TotalCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) SetUsedCoreHours(v float32) *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods {
	s.UsedCoreHours = &v
	return s
}

func (s *DescribeCorePackageListResponseBodyCorePackageInfoCorePackageListRemainingPeriods) Validate() error {
	return dara.Validate(s)
}
