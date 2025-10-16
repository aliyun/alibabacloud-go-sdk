// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeInternetOpenIpResponseBodyDataList) *DescribeInternetOpenIpResponseBody
	GetDataList() []*DescribeInternetOpenIpResponseBodyDataList
	SetPageInfo(v *DescribeInternetOpenIpResponseBodyPageInfo) *DescribeInternetOpenIpResponseBody
	GetPageInfo() *DescribeInternetOpenIpResponseBodyPageInfo
	SetRequestId(v string) *DescribeInternetOpenIpResponseBody
	GetRequestId() *string
}

type DescribeInternetOpenIpResponseBody struct {
	// The data returned.
	DataList []*DescribeInternetOpenIpResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeInternetOpenIpResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6B780BD6-282C-51A9-A8E6-59F636BAFA54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetOpenIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponseBody) GetDataList() []*DescribeInternetOpenIpResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetOpenIpResponseBody) GetPageInfo() *DescribeInternetOpenIpResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInternetOpenIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetOpenIpResponseBody) SetDataList(v []*DescribeInternetOpenIpResponseBodyDataList) *DescribeInternetOpenIpResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetOpenIpResponseBody) SetPageInfo(v *DescribeInternetOpenIpResponseBodyPageInfo) *DescribeInternetOpenIpResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInternetOpenIpResponseBody) SetRequestId(v string) *DescribeInternetOpenIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBody) Validate() error {
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

type DescribeInternetOpenIpResponseBodyDataList struct {
	// The reason why recommended intelligent policies are unavailable. Valid values:
	//
	// 	- No recommended intelligent policies are available.
	//
	// 	- This feature is available only to some users.
	//
	// 	- The policy configuration has been modified. No recommended intelligent policies are available.
	//
	// 	- The recommended intelligent policies have been configured. No new recommended intelligent policies are available.
	//
	// example:
	//
	// No recommended intelligent policies are available.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1ix9w22kv6aew9****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// launch-advisor-2023****
	AssetsName *string `json:"AssetsName,omitempty" xml:"AssetsName,omitempty"`
	// The asset type of the instance.
	//
	// example:
	//
	// EcsEIP
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The total number of ports.
	//
	// example:
	//
	// 5
	DetailNum *int32 `json:"DetailNum,omitempty" xml:"DetailNum,omitempty"`
	// Specifies whether an access control policy is recommended. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The inbound network throughput, which indicates the total number of bytes that are sent inbound. Unit: bytes.
	//
	// example:
	//
	// 235
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 14151892****7022
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The outbound network throughput, which indicates the total number of bytes that are sent outbound. Unit: bytes.
	//
	// example:
	//
	// 1123
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The list of ports.
	PortList []*string `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// The public IP address of the instance.
	//
	// example:
	//
	// 203.0.113.1
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **3**: high risk
	//
	// 	- **2**: medium risk
	//
	// 	- **1**: low risk
	//
	// 	- **0**: no risk
	//
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The reason for the risk.
	//
	// example:
	//
	// Previous traffic is all malicious traffic.
	RiskReason *string `json:"RiskReason,omitempty" xml:"RiskReason,omitempty"`
	// The list of applications.
	ServiceNameList []*string `json:"ServiceNameList,omitempty" xml:"ServiceNameList,omitempty" type:"Repeated"`
	// Number of source IPs.
	//
	// example:
	//
	// 22
	SrcIpCnt *int64 `json:"SrcIpCnt,omitempty" xml:"SrcIpCnt,omitempty"`
	// The total inbound and outbound network throughput, which indicates the total number of bytes that are sent inbound and outbound. Unit: bytes.
	//
	// example:
	//
	// 253023143
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// Outbound traffic in the last 7 days.
	//
	// example:
	//
	// 100000
	TotalReplyBytes *int64 `json:"TotalReplyBytes,omitempty" xml:"TotalReplyBytes,omitempty"`
	// For detailed traffic information, see the TotalBytes field.
	//
	// example:
	//
	// 0
	TrafficPercent1Day *string `json:"TrafficPercent1Day,omitempty" xml:"TrafficPercent1Day,omitempty"`
	// For detailed traffic information, see the TotalBytes field.
	//
	// example:
	//
	// 0
	TrafficPercent30Day *string `json:"TrafficPercent30Day,omitempty" xml:"TrafficPercent30Day,omitempty"`
	// For detailed traffic information, see the TotalBytes field.
	//
	// example:
	//
	// 0
	TrafficPercent7Day *string `json:"TrafficPercent7Day,omitempty" xml:"TrafficPercent7Day,omitempty"`
	// Reasons for not analyzing the protocol when the protocol is identified as Unknown.
	UnknownReason []*string `json:"UnknownReason,omitempty" xml:"UnknownReason,omitempty" type:"Repeated"`
}

func (s DescribeInternetOpenIpResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenIpResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetAclRecommendDetail() *string {
	return s.AclRecommendDetail
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetAssetsName() *string {
	return s.AssetsName
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetAssetsType() *string {
	return s.AssetsType
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetDetailNum() *int32 {
	return s.DetailNum
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetHasAclRecommend() *bool {
	return s.HasAclRecommend
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetPortList() []*string {
	return s.PortList
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetRiskReason() *string {
	return s.RiskReason
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetServiceNameList() []*string {
	return s.ServiceNameList
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetSrcIpCnt() *int64 {
	return s.SrcIpCnt
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetTotalReplyBytes() *int64 {
	return s.TotalReplyBytes
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetTrafficPercent1Day() *string {
	return s.TrafficPercent1Day
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetTrafficPercent30Day() *string {
	return s.TrafficPercent30Day
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetTrafficPercent7Day() *string {
	return s.TrafficPercent7Day
}

func (s *DescribeInternetOpenIpResponseBodyDataList) GetUnknownReason() []*string {
	return s.UnknownReason
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAclRecommendDetail(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAssetsInstanceId(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAssetsName(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AssetsName = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAssetsType(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AssetsType = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetDetailNum(v int32) *DescribeInternetOpenIpResponseBodyDataList {
	s.DetailNum = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetHasAclRecommend(v bool) *DescribeInternetOpenIpResponseBodyDataList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetInBytes(v int64) *DescribeInternetOpenIpResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetMemberUid(v int64) *DescribeInternetOpenIpResponseBodyDataList {
	s.MemberUid = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetOutBytes(v int64) *DescribeInternetOpenIpResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetPortList(v []*string) *DescribeInternetOpenIpResponseBodyDataList {
	s.PortList = v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetPublicIp(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetRegionNo(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetRiskLevel(v int32) *DescribeInternetOpenIpResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetRiskReason(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.RiskReason = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetServiceNameList(v []*string) *DescribeInternetOpenIpResponseBodyDataList {
	s.ServiceNameList = v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetSrcIpCnt(v int64) *DescribeInternetOpenIpResponseBodyDataList {
	s.SrcIpCnt = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTotalBytes(v int64) *DescribeInternetOpenIpResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTotalReplyBytes(v int64) *DescribeInternetOpenIpResponseBodyDataList {
	s.TotalReplyBytes = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTrafficPercent1Day(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.TrafficPercent1Day = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTrafficPercent30Day(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.TrafficPercent30Day = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTrafficPercent7Day(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.TrafficPercent7Day = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetUnknownReason(v []*string) *DescribeInternetOpenIpResponseBodyDataList {
	s.UnknownReason = v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeInternetOpenIpResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInternetOpenIpResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenIpResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInternetOpenIpResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) SetPageSize(v int32) *DescribeInternetOpenIpResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInternetOpenIpResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
