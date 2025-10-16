// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenPortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInternetOpenPortResponseBodyDataList) *DescribeInternetOpenPortResponseBody
	GetDataList() []*DescribeInternetOpenPortResponseBodyDataList
	SetPageInfo(v *DescribeInternetOpenPortResponseBodyPageInfo) *DescribeInternetOpenPortResponseBody
	GetPageInfo() *DescribeInternetOpenPortResponseBodyPageInfo
	SetRequestId(v string) *DescribeInternetOpenPortResponseBody
	GetRequestId() *string
}

type DescribeInternetOpenPortResponseBody struct {
	DataList []*DescribeInternetOpenPortResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	PageInfo *DescribeInternetOpenPortResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// F3637663-991B-547F-9163-1A5AC367****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetOpenPortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenPortResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenPortResponseBody) GetDataList() []*DescribeInternetOpenPortResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetOpenPortResponseBody) GetPageInfo() *DescribeInternetOpenPortResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInternetOpenPortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetOpenPortResponseBody) SetDataList(v []*DescribeInternetOpenPortResponseBodyDataList) *DescribeInternetOpenPortResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetOpenPortResponseBody) SetPageInfo(v *DescribeInternetOpenPortResponseBodyPageInfo) *DescribeInternetOpenPortResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInternetOpenPortResponseBody) SetRequestId(v string) *DescribeInternetOpenPortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInternetOpenPortResponseBodyDataList struct {
	// example:
	//
	// 5
	DetailNum *int32 `json:"DetailNum,omitempty" xml:"DetailNum,omitempty"`
	// example:
	//
	// 1456536639.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 117200.0
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// test
	ProbRisk *string `json:"ProbRisk,omitempty" xml:"ProbRisk,omitempty"`
	// example:
	//
	// desc
	ProbRiskDesc *string `json:"ProbRiskDesc,omitempty" xml:"ProbRiskDesc,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 10
	PublicIpNum *int32 `json:"PublicIpNum,omitempty" xml:"PublicIpNum,omitempty"`
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test
	RiskReason      *string   `json:"RiskReason,omitempty" xml:"RiskReason,omitempty"`
	ServiceNameList []*string `json:"ServiceNameList,omitempty" xml:"ServiceNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	SuggestLevel *int32 `json:"SuggestLevel,omitempty" xml:"SuggestLevel,omitempty"`
	// example:
	//
	// 621404
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// example:
	//
	// 10
	TrafficPercent1Day *string `json:"TrafficPercent1Day,omitempty" xml:"TrafficPercent1Day,omitempty"`
	// example:
	//
	// 48
	TrafficPercent30Day *string `json:"TrafficPercent30Day,omitempty" xml:"TrafficPercent30Day,omitempty"`
	// example:
	//
	// 30
	TrafficPercent7Day *string   `json:"TrafficPercent7Day,omitempty" xml:"TrafficPercent7Day,omitempty"`
	UnknownReason      []*string `json:"UnknownReason,omitempty" xml:"UnknownReason,omitempty" type:"Repeated"`
}

func (s DescribeInternetOpenPortResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenPortResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetDetailNum() *int32 {
	return s.DetailNum
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetProbRisk() *string {
	return s.ProbRisk
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetProbRiskDesc() *string {
	return s.ProbRiskDesc
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetPublicIpNum() *int32 {
	return s.PublicIpNum
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetRiskReason() *string {
	return s.RiskReason
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetServiceNameList() []*string {
	return s.ServiceNameList
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetSuggestLevel() *int32 {
	return s.SuggestLevel
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetTrafficPercent1Day() *string {
	return s.TrafficPercent1Day
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetTrafficPercent30Day() *string {
	return s.TrafficPercent30Day
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetTrafficPercent7Day() *string {
	return s.TrafficPercent7Day
}

func (s *DescribeInternetOpenPortResponseBodyDataList) GetUnknownReason() []*string {
	return s.UnknownReason
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetDetailNum(v int32) *DescribeInternetOpenPortResponseBodyDataList {
	s.DetailNum = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetInBytes(v int64) *DescribeInternetOpenPortResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetOutBytes(v int64) *DescribeInternetOpenPortResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetPort(v int32) *DescribeInternetOpenPortResponseBodyDataList {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetProbRisk(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.ProbRisk = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetProbRiskDesc(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.ProbRiskDesc = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetProtocol(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.Protocol = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetPublicIpNum(v int32) *DescribeInternetOpenPortResponseBodyDataList {
	s.PublicIpNum = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetRiskLevel(v int32) *DescribeInternetOpenPortResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetRiskReason(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.RiskReason = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetServiceNameList(v []*string) *DescribeInternetOpenPortResponseBodyDataList {
	s.ServiceNameList = v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetSuggestLevel(v int32) *DescribeInternetOpenPortResponseBodyDataList {
	s.SuggestLevel = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetTotalBytes(v int64) *DescribeInternetOpenPortResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetTrafficPercent1Day(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.TrafficPercent1Day = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetTrafficPercent30Day(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.TrafficPercent30Day = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetTrafficPercent7Day(v string) *DescribeInternetOpenPortResponseBodyDataList {
	s.TrafficPercent7Day = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) SetUnknownReason(v []*string) *DescribeInternetOpenPortResponseBodyDataList {
	s.UnknownReason = v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeInternetOpenPortResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 39
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInternetOpenPortResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenPortResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInternetOpenPortResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) SetPageSize(v int32) *DescribeInternetOpenPortResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInternetOpenPortResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInternetOpenPortResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
