// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDomainRelationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVpcFirewallDomainRelationListResponseBodyDataList) *DescribeVpcFirewallDomainRelationListResponseBody
	GetDataList() []*DescribeVpcFirewallDomainRelationListResponseBodyDataList
	SetDstVpcList(v []*DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) *DescribeVpcFirewallDomainRelationListResponseBody
	GetDstVpcList() []*DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList
	SetRequestId(v string) *DescribeVpcFirewallDomainRelationListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallDomainRelationListResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallDomainRelationListResponseBody struct {
	DataList   []*DescribeVpcFirewallDomainRelationListResponseBodyDataList   `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	DstVpcList []*DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList `json:"DstVpcList,omitempty" xml:"DstVpcList,omitempty" type:"Repeated"`
	// example:
	//
	// C5DDD596-1191-5F36-A504-8733045A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallDomainRelationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainRelationListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) GetDataList() []*DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) GetDstVpcList() []*DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList {
	return s.DstVpcList
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) SetDataList(v []*DescribeVpcFirewallDomainRelationListResponseBodyDataList) *DescribeVpcFirewallDomainRelationListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) SetDstVpcList(v []*DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) *DescribeVpcFirewallDomainRelationListResponseBody {
	s.DstVpcList = v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) SetRequestId(v string) *DescribeVpcFirewallDomainRelationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallDomainRelationListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DstVpcList != nil {
		for _, item := range s.DstVpcList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDomainRelationListResponseBodyDataList struct {
	// example:
	//
	// Google
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 192.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// example:
	//
	// cn-beijing
	DstRegionId *string `json:"DstRegionId,omitempty" xml:"DstRegionId,omitempty"`
	// example:
	//
	// vpc-bp10w5nb30r4jzfyc****
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// example:
	//
	// vpc-****
	DstVpcName *string `json:"DstVpcName,omitempty" xml:"DstVpcName,omitempty"`
	// example:
	//
	// 1767147003
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// Google
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 10
	IpsHitCnt *int64 `json:"IpsHitCnt,omitempty" xml:"IpsHitCnt,omitempty"`
	// example:
	//
	// 1767147003
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// 10000
	RequestBytes *int64 `json:"RequestBytes,omitempty" xml:"RequestBytes,omitempty"`
	// example:
	//
	// 10000
	ResponseBytes *int64 `json:"ResponseBytes,omitempty" xml:"ResponseBytes,omitempty"`
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 192.0.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// cn-hangzhou
	SrcRegionId *string `json:"SrcRegionId,omitempty" xml:"SrcRegionId,omitempty"`
	// example:
	//
	// vpc-t4nlt09olhpazpoeg****
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// example:
	//
	// vpc-****
	SrcVpcName *string `json:"SrcVpcName,omitempty" xml:"SrcVpcName,omitempty"`
	// example:
	//
	// 16287823
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeVpcFirewallDomainRelationListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainRelationListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetBusiness() *string {
	return s.Business
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetDstRegionId() *string {
	return s.DstRegionId
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetDstVpcId() *string {
	return s.DstVpcId
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetDstVpcName() *string {
	return s.DstVpcName
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetIpsHitCnt() *int64 {
	return s.IpsHitCnt
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetRequestBytes() *int64 {
	return s.RequestBytes
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetResponseBytes() *int64 {
	return s.ResponseBytes
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetSrcRegionId() *string {
	return s.SrcRegionId
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetSrcVpcName() *string {
	return s.SrcVpcName
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetBusiness(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.Business = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetDomain(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetDstIP(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.DstIP = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetDstRegionId(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.DstRegionId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetDstVpcId(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.DstVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetDstVpcName(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.DstVpcName = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetFirstTime(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.FirstTime = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetGroupName(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.GroupName = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetIpsHitCnt(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.IpsHitCnt = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetLastTime(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.LastTime = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetRequestBytes(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.RequestBytes = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetResponseBytes(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.ResponseBytes = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetSessionCount(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetSrcIP(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.SrcIP = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetSrcRegionId(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.SrcRegionId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetSrcVpcId(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetSrcVpcName(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.SrcVpcName = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) SetTotalBytes(v int64) *DescribeVpcFirewallDomainRelationListResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList struct {
	// example:
	//
	// vpc-bp10w5nb30r4jzfyc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vpc-****
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) SetVpcId(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) SetVpcName(v string) *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallDomainRelationListResponseBodyDstVpcList) Validate() error {
	return dara.Validate(s)
}
