// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationIPDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclCoverage(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetAclCoverage() *string
	SetCurrentPage(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetCurrentPage() *string
	SetDstIP(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetDstIP() *string
	SetEndTime(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetEndTime() *string
	SetIPType(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetIPType() *string
	SetLang(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetLang() *string
	SetNatGatewayId(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetNatGatewayId() *string
	SetOrder(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetPageSize() *string
	SetPrivateIP(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetPrivateIP() *string
	SetPublicIP(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetPublicIP() *string
	SetSort(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetSort() *string
	SetSourceIp(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetStartTime() *string
	SetTagId(v string) *DescribeOutgoingDestinationIPDetailRequest
	GetTagId() *string
}

type DescribeOutgoingDestinationIPDetailRequest struct {
	// example:
	//
	// FullCoverage
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 34.136.111.XXX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1733710383
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// NatPrivate
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// ngw-2zed6z6qkd7ogc****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10.210.0.XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// example:
	//
	// 192.0.XX.XX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// InBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 1.202.193.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749434787
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// FirstFlow
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DescribeOutgoingDestinationIPDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetIPType() *string {
	return s.IPType
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingDestinationIPDetailRequest) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetAclCoverage(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetCurrentPage(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetDstIP(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetEndTime(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetIPType(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.IPType = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetLang(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetNatGatewayId(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetOrder(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetPageSize(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetPrivateIP(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetPublicIP(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetSort(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetSourceIp(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetStartTime(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) SetTagId(v string) *DescribeOutgoingDestinationIPDetailRequest {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailRequest) Validate() error {
	return dara.Validate(s)
}
