// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckRiskStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *GetCheckRiskStatisticsResponseBody
	GetCount() *int32
	SetData(v []*GetCheckRiskStatisticsResponseBodyData) *GetCheckRiskStatisticsResponseBody
	GetData() []*GetCheckRiskStatisticsResponseBodyData
	SetRequestId(v string) *GetCheckRiskStatisticsResponseBody
	GetRequestId() *string
	SetSummary(v *GetCheckRiskStatisticsResponseBodySummary) *GetCheckRiskStatisticsResponseBody
	GetSummary() *GetCheckRiskStatisticsResponseBodySummary
}

type GetCheckRiskStatisticsResponseBody struct {
	// The number of risk scenarios.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// An array consisting of the statistics on check items that are used in risk scenarios.
	Data []*GetCheckRiskStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2C455672-2490-5211-84EC-420C7818****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Historical check item statistics.
	Summary *GetCheckRiskStatisticsResponseBodySummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
}

func (s GetCheckRiskStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRiskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckRiskStatisticsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetCheckRiskStatisticsResponseBody) GetData() []*GetCheckRiskStatisticsResponseBodyData {
	return s.Data
}

func (s *GetCheckRiskStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckRiskStatisticsResponseBody) GetSummary() *GetCheckRiskStatisticsResponseBodySummary {
	return s.Summary
}

func (s *GetCheckRiskStatisticsResponseBody) SetCount(v int32) *GetCheckRiskStatisticsResponseBody {
	s.Count = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBody) SetData(v []*GetCheckRiskStatisticsResponseBodyData) *GetCheckRiskStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetCheckRiskStatisticsResponseBody) SetRequestId(v string) *GetCheckRiskStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBody) SetSummary(v *GetCheckRiskStatisticsResponseBodySummary) *GetCheckRiskStatisticsResponseBody {
	s.Summary = v
	return s
}

func (s *GetCheckRiskStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckRiskStatisticsResponseBodyData struct {
	// The number of high-risk items.
	//
	// example:
	//
	// 43
	HighWarningCount *int32 `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	// The number of low-risk items.
	//
	// example:
	//
	// 3
	LowWarningCount *int32 `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	// The number of medium-risk items.
	//
	// example:
	//
	// 29
	MediumWarningCount *int32 `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
	// The number of passed check items.
	//
	// example:
	//
	// 143
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The name of the risk scenario.
	//
	// example:
	//
	// SECURITY
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The statistics on check items that are used in the risk scenario by baseline type.
	SubStatistics []*GetCheckRiskStatisticsResponseBodyDataSubStatistics `json:"SubStatistics,omitempty" xml:"SubStatistics,omitempty" type:"Repeated"`
	// The total number of check items.
	//
	// example:
	//
	// 219
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCheckRiskStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRiskStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetHighWarningCount() *int32 {
	return s.HighWarningCount
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetLowWarningCount() *int32 {
	return s.LowWarningCount
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetMediumWarningCount() *int32 {
	return s.MediumWarningCount
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetPassCount() *int32 {
	return s.PassCount
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetSceneName() *string {
	return s.SceneName
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetSubStatistics() []*GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	return s.SubStatistics
}

func (s *GetCheckRiskStatisticsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetHighWarningCount(v int32) *GetCheckRiskStatisticsResponseBodyData {
	s.HighWarningCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetLowWarningCount(v int32) *GetCheckRiskStatisticsResponseBodyData {
	s.LowWarningCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetMediumWarningCount(v int32) *GetCheckRiskStatisticsResponseBodyData {
	s.MediumWarningCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetPassCount(v int32) *GetCheckRiskStatisticsResponseBodyData {
	s.PassCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetSceneName(v string) *GetCheckRiskStatisticsResponseBodyData {
	s.SceneName = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetSubStatistics(v []*GetCheckRiskStatisticsResponseBodyDataSubStatistics) *GetCheckRiskStatisticsResponseBodyData {
	s.SubStatistics = v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) SetTotalCount(v int32) *GetCheckRiskStatisticsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyData) Validate() error {
	if s.SubStatistics != nil {
		for _, item := range s.SubStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCheckRiskStatisticsResponseBodyDataSubStatistics struct {
	// The name of the baseline type.
	//
	// example:
	//
	// weak_password
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The number of high-risk items.
	//
	// example:
	//
	// 3
	HighWarningCount *int32 `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	// The number of low-risk items.
	//
	// example:
	//
	// 0
	LowWarningCount *int32 `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	// The number of medium-risk items.
	//
	// example:
	//
	// 0
	MediumWarningCount *int32 `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
	// The number of passed check items.
	//
	// example:
	//
	// 2
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// The total number of check items.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The baseline type.
	//
	// example:
	//
	// weak_password
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetCheckRiskStatisticsResponseBodyDataSubStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRiskStatisticsResponseBodyDataSubStatistics) GoString() string {
	return s.String()
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetAlias() *string {
	return s.Alias
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetHighWarningCount() *int32 {
	return s.HighWarningCount
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetLowWarningCount() *int32 {
	return s.LowWarningCount
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetMediumWarningCount() *int32 {
	return s.MediumWarningCount
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetPassCount() *int32 {
	return s.PassCount
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) GetTypeName() *string {
	return s.TypeName
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetAlias(v string) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.Alias = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetHighWarningCount(v int32) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.HighWarningCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetLowWarningCount(v int32) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.LowWarningCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetMediumWarningCount(v int32) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.MediumWarningCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetPassCount(v int32) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.PassCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetTotalCount(v int32) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.TotalCount = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) SetTypeName(v string) *GetCheckRiskStatisticsResponseBodyDataSubStatistics {
	s.TypeName = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodyDataSubStatistics) Validate() error {
	return dara.Validate(s)
}

type GetCheckRiskStatisticsResponseBodySummary struct {
	// Check items handled today.
	//
	// example:
	//
	// 0
	HandledCheckToday *int32 `json:"HandledCheckToday,omitempty" xml:"HandledCheckToday,omitempty"`
	// A risk item exists.
	//
	// example:
	//
	// 3
	HandledCheckTotal *int32 `json:"HandledCheckTotal,omitempty" xml:"HandledCheckTotal,omitempty"`
	// Total days since check items were handled.
	//
	// example:
	//
	// 365
	HandledDays *int32 `json:"HandledDays,omitempty" xml:"HandledDays,omitempty"`
	// Check items that failed to pass the check.
	//
	// example:
	//
	// 1
	RiskCheckCnt *int32 `json:"RiskCheckCnt,omitempty" xml:"RiskCheckCnt,omitempty"`
	// Days since check items failed.
	//
	// example:
	//
	// 30
	RiskDays *int32 `json:"RiskDays,omitempty" xml:"RiskDays,omitempty"`
	// Risks to be handled.
	//
	// example:
	//
	// 5
	RiskWarningCnt *int32 `json:"RiskWarningCnt,omitempty" xml:"RiskWarningCnt,omitempty"`
}

func (s GetCheckRiskStatisticsResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRiskStatisticsResponseBodySummary) GoString() string {
	return s.String()
}

func (s *GetCheckRiskStatisticsResponseBodySummary) GetHandledCheckToday() *int32 {
	return s.HandledCheckToday
}

func (s *GetCheckRiskStatisticsResponseBodySummary) GetHandledCheckTotal() *int32 {
	return s.HandledCheckTotal
}

func (s *GetCheckRiskStatisticsResponseBodySummary) GetHandledDays() *int32 {
	return s.HandledDays
}

func (s *GetCheckRiskStatisticsResponseBodySummary) GetRiskCheckCnt() *int32 {
	return s.RiskCheckCnt
}

func (s *GetCheckRiskStatisticsResponseBodySummary) GetRiskDays() *int32 {
	return s.RiskDays
}

func (s *GetCheckRiskStatisticsResponseBodySummary) GetRiskWarningCnt() *int32 {
	return s.RiskWarningCnt
}

func (s *GetCheckRiskStatisticsResponseBodySummary) SetHandledCheckToday(v int32) *GetCheckRiskStatisticsResponseBodySummary {
	s.HandledCheckToday = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodySummary) SetHandledCheckTotal(v int32) *GetCheckRiskStatisticsResponseBodySummary {
	s.HandledCheckTotal = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodySummary) SetHandledDays(v int32) *GetCheckRiskStatisticsResponseBodySummary {
	s.HandledDays = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodySummary) SetRiskCheckCnt(v int32) *GetCheckRiskStatisticsResponseBodySummary {
	s.RiskCheckCnt = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodySummary) SetRiskDays(v int32) *GetCheckRiskStatisticsResponseBodySummary {
	s.RiskDays = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodySummary) SetRiskWarningCnt(v int32) *GetCheckRiskStatisticsResponseBodySummary {
	s.RiskWarningCnt = &v
	return s
}

func (s *GetCheckRiskStatisticsResponseBodySummary) Validate() error {
	return dara.Validate(s)
}
