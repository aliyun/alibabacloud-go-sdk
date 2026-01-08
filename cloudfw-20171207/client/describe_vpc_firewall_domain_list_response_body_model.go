// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVpcFirewallDomainListResponseBodyDataList) *DescribeVpcFirewallDomainListResponseBody
	GetDataList() []*DescribeVpcFirewallDomainListResponseBodyDataList
	SetRequestId(v string) *DescribeVpcFirewallDomainListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallDomainListResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallDomainListResponseBody struct {
	DataList []*DescribeVpcFirewallDomainListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 133173B9-8010-5DF5-8B93-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainListResponseBody) GetDataList() []*DescribeVpcFirewallDomainListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallDomainListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallDomainListResponseBody) SetDataList(v []*DescribeVpcFirewallDomainListResponseBodyDataList) *DescribeVpcFirewallDomainListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBody) SetRequestId(v string) *DescribeVpcFirewallDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallDomainListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDomainListResponseBodyDataList struct {
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// example:
	//
	// Google
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// example:
	//
	// www.a.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// Google
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 3214
	RequestBytes *int64 `json:"RequestBytes,omitempty" xml:"RequestBytes,omitempty"`
	// example:
	//
	// 4582
	ResponseBytes *int64 `json:"ResponseBytes,omitempty" xml:"ResponseBytes,omitempty"`
	// example:
	//
	// 12
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 2
	SrcIpCount *int64 `json:"SrcIpCount,omitempty" xml:"SrcIpCount,omitempty"`
	// example:
	//
	// 1
	SrcVpcCount *int64 `json:"SrcVpcCount,omitempty" xml:"SrcVpcCount,omitempty"`
	// example:
	//
	// 8111126106
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeVpcFirewallDomainListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDomainListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetBusiness() *string {
	return s.Business
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetRequestBytes() *int64 {
	return s.RequestBytes
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetResponseBytes() *int64 {
	return s.ResponseBytes
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetSrcIpCount() *int64 {
	return s.SrcIpCount
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetSrcVpcCount() *int64 {
	return s.SrcVpcCount
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetApplicationNameList(v []*string) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetBusiness(v string) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.Business = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetDomain(v string) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetGroupName(v string) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.GroupName = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetRequestBytes(v int64) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.RequestBytes = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetResponseBytes(v int64) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.ResponseBytes = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetSessionCount(v int64) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetSrcIpCount(v int64) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.SrcIpCount = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetSrcVpcCount(v int64) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.SrcVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) SetTotalBytes(v int64) *DescribeVpcFirewallDomainListResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVpcFirewallDomainListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
