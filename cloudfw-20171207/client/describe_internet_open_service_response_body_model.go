// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInternetOpenServiceResponseBodyDataList) *DescribeInternetOpenServiceResponseBody
	GetDataList() []*DescribeInternetOpenServiceResponseBodyDataList
	SetPageInfo(v *DescribeInternetOpenServiceResponseBodyPageInfo) *DescribeInternetOpenServiceResponseBody
	GetPageInfo() *DescribeInternetOpenServiceResponseBodyPageInfo
	SetRequestId(v string) *DescribeInternetOpenServiceResponseBody
	GetRequestId() *string
}

type DescribeInternetOpenServiceResponseBody struct {
	DataList []*DescribeInternetOpenServiceResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	PageInfo *DescribeInternetOpenServiceResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 99A65AA0-C5B5-5092-BFCF-8111B436****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetOpenServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenServiceResponseBody) GetDataList() []*DescribeInternetOpenServiceResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetOpenServiceResponseBody) GetPageInfo() *DescribeInternetOpenServiceResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInternetOpenServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetOpenServiceResponseBody) SetDataList(v []*DescribeInternetOpenServiceResponseBodyDataList) *DescribeInternetOpenServiceResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetOpenServiceResponseBody) SetPageInfo(v *DescribeInternetOpenServiceResponseBodyPageInfo) *DescribeInternetOpenServiceResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInternetOpenServiceResponseBody) SetRequestId(v string) *DescribeInternetOpenServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBody) Validate() error {
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

type DescribeInternetOpenServiceResponseBodyDataList struct {
	// example:
	//
	// 5
	DetailNum *int32 `json:"DetailNum,omitempty" xml:"DetailNum,omitempty"`
	// example:
	//
	// 447458.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 1123
	OutBytes *int64    `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	PortList []*string `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ProbRisk *string `json:"ProbRisk,omitempty" xml:"ProbRisk,omitempty"`
	// example:
	//
	// test
	ProbRiskDesc *string `json:"ProbRiskDesc,omitempty" xml:"ProbRiskDesc,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 3
	PublicIpNum *int32 `json:"PublicIpNum,omitempty" xml:"PublicIpNum,omitempty"`
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test
	RiskReason *string `json:"RiskReason,omitempty" xml:"RiskReason,omitempty"`
	// example:
	//
	// SMB
	ServiceName *int32 `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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
	// 0
	TrafficPercent1Day *string `json:"TrafficPercent1Day,omitempty" xml:"TrafficPercent1Day,omitempty"`
	// example:
	//
	// 0
	TrafficPercent30Day *string `json:"TrafficPercent30Day,omitempty" xml:"TrafficPercent30Day,omitempty"`
	// example:
	//
	// 77
	TrafficPercent7Day *string   `json:"TrafficPercent7Day,omitempty" xml:"TrafficPercent7Day,omitempty"`
	UnknownReason      []*string `json:"UnknownReason,omitempty" xml:"UnknownReason,omitempty" type:"Repeated"`
}

func (s DescribeInternetOpenServiceResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenServiceResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetDetailNum() *int32 {
	return s.DetailNum
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetPortList() []*string {
	return s.PortList
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetProbRisk() *string {
	return s.ProbRisk
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetProbRiskDesc() *string {
	return s.ProbRiskDesc
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetPublicIpNum() *int32 {
	return s.PublicIpNum
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetRiskReason() *string {
	return s.RiskReason
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetServiceName() *int32 {
	return s.ServiceName
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetSuggestLevel() *int32 {
	return s.SuggestLevel
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetTrafficPercent1Day() *string {
	return s.TrafficPercent1Day
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetTrafficPercent30Day() *string {
	return s.TrafficPercent30Day
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetTrafficPercent7Day() *string {
	return s.TrafficPercent7Day
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) GetUnknownReason() []*string {
	return s.UnknownReason
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetDetailNum(v int32) *DescribeInternetOpenServiceResponseBodyDataList {
	s.DetailNum = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetInBytes(v int64) *DescribeInternetOpenServiceResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetOutBytes(v int64) *DescribeInternetOpenServiceResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetPortList(v []*string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.PortList = v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetProbRisk(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.ProbRisk = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetProbRiskDesc(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.ProbRiskDesc = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetProtocol(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.Protocol = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetPublicIpNum(v int32) *DescribeInternetOpenServiceResponseBodyDataList {
	s.PublicIpNum = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetRiskLevel(v int32) *DescribeInternetOpenServiceResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetRiskReason(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.RiskReason = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetServiceName(v int32) *DescribeInternetOpenServiceResponseBodyDataList {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetSuggestLevel(v int32) *DescribeInternetOpenServiceResponseBodyDataList {
	s.SuggestLevel = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetTotalBytes(v int64) *DescribeInternetOpenServiceResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetTrafficPercent1Day(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.TrafficPercent1Day = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetTrafficPercent30Day(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.TrafficPercent30Day = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetTrafficPercent7Day(v string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.TrafficPercent7Day = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) SetUnknownReason(v []*string) *DescribeInternetOpenServiceResponseBodyDataList {
	s.UnknownReason = v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeInternetOpenServiceResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInternetOpenServiceResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenServiceResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInternetOpenServiceResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) SetPageSize(v int32) *DescribeInternetOpenServiceResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInternetOpenServiceResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInternetOpenServiceResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
