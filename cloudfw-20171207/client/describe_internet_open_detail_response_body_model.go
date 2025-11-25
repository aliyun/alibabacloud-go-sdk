// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInternetOpenDetailResponseBodyDataList) *DescribeInternetOpenDetailResponseBody
	GetDataList() []*DescribeInternetOpenDetailResponseBodyDataList
	SetPageInfo(v *DescribeInternetOpenDetailResponseBodyPageInfo) *DescribeInternetOpenDetailResponseBody
	GetPageInfo() *DescribeInternetOpenDetailResponseBodyPageInfo
	SetRequestId(v string) *DescribeInternetOpenDetailResponseBody
	GetRequestId() *string
}

type DescribeInternetOpenDetailResponseBody struct {
	DataList []*DescribeInternetOpenDetailResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	PageInfo *DescribeInternetOpenDetailResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 3495E758-BB4B-5F5C-8AE0-897489F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetOpenDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenDetailResponseBody) GetDataList() []*DescribeInternetOpenDetailResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetOpenDetailResponseBody) GetPageInfo() *DescribeInternetOpenDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInternetOpenDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetOpenDetailResponseBody) SetDataList(v []*DescribeInternetOpenDetailResponseBodyDataList) *DescribeInternetOpenDetailResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetOpenDetailResponseBody) SetPageInfo(v *DescribeInternetOpenDetailResponseBodyPageInfo) *DescribeInternetOpenDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInternetOpenDetailResponseBody) SetRequestId(v string) *DescribeInternetOpenDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBody) Validate() error {
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

type DescribeInternetOpenDetailResponseBodyDataList struct {
	// example:
	//
	// i-bp1ix9w22kv6aew9****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// example:
	//
	// launch-advisor-2023****
	AssetsName *string `json:"AssetsName,omitempty" xml:"AssetsName,omitempty"`
	// example:
	//
	// EcsEIP
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// example:
	//
	// 1123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 0.0
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 3389
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
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
	// 39.101.167.XX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
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
	// Redis
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// 20
	SuggestLevel *int32 `json:"SuggestLevel,omitempty" xml:"SuggestLevel,omitempty"`
	// example:
	//
	// 0
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

func (s DescribeInternetOpenDetailResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenDetailResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetAssetsName() *string {
	return s.AssetsName
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetAssetsType() *string {
	return s.AssetsType
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetProbRisk() *string {
	return s.ProbRisk
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetProbRiskDesc() *string {
	return s.ProbRiskDesc
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetRiskReason() *string {
	return s.RiskReason
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetSuggestLevel() *int32 {
	return s.SuggestLevel
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetTrafficPercent1Day() *string {
	return s.TrafficPercent1Day
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetTrafficPercent30Day() *string {
	return s.TrafficPercent30Day
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetTrafficPercent7Day() *string {
	return s.TrafficPercent7Day
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) GetUnknownReason() []*string {
	return s.UnknownReason
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetAssetsInstanceId(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetAssetsName(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.AssetsName = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetAssetsType(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.AssetsType = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetId(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetInBytes(v int64) *DescribeInternetOpenDetailResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetOutBytes(v int64) *DescribeInternetOpenDetailResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetPort(v int32) *DescribeInternetOpenDetailResponseBodyDataList {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetProbRisk(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.ProbRisk = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetProbRiskDesc(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.ProbRiskDesc = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetProtocol(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.Protocol = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetPublicIp(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetRegionNo(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetRiskLevel(v int32) *DescribeInternetOpenDetailResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetRiskReason(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.RiskReason = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetServiceName(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetSuggestLevel(v int32) *DescribeInternetOpenDetailResponseBodyDataList {
	s.SuggestLevel = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetTotalBytes(v int64) *DescribeInternetOpenDetailResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetTrafficPercent1Day(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.TrafficPercent1Day = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetTrafficPercent30Day(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.TrafficPercent30Day = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetTrafficPercent7Day(v string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.TrafficPercent7Day = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) SetUnknownReason(v []*string) *DescribeInternetOpenDetailResponseBodyDataList {
	s.UnknownReason = v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeInternetOpenDetailResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInternetOpenDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInternetOpenDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribeInternetOpenDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInternetOpenDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInternetOpenDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
