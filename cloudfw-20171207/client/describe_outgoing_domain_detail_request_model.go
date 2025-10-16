// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclCoverage(v string) *DescribeOutgoingDomainDetailRequest
	GetAclCoverage() *string
	SetCurrentPage(v string) *DescribeOutgoingDomainDetailRequest
	GetCurrentPage() *string
	SetDomain(v string) *DescribeOutgoingDomainDetailRequest
	GetDomain() *string
	SetDomainList(v []*string) *DescribeOutgoingDomainDetailRequest
	GetDomainList() []*string
	SetEndTime(v string) *DescribeOutgoingDomainDetailRequest
	GetEndTime() *string
	SetIPType(v string) *DescribeOutgoingDomainDetailRequest
	GetIPType() *string
	SetLang(v string) *DescribeOutgoingDomainDetailRequest
	GetLang() *string
	SetNatGatewayId(v string) *DescribeOutgoingDomainDetailRequest
	GetNatGatewayId() *string
	SetOrder(v string) *DescribeOutgoingDomainDetailRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeOutgoingDomainDetailRequest
	GetPageSize() *string
	SetPrivateIP(v string) *DescribeOutgoingDomainDetailRequest
	GetPrivateIP() *string
	SetPublicIP(v string) *DescribeOutgoingDomainDetailRequest
	GetPublicIP() *string
	SetSort(v string) *DescribeOutgoingDomainDetailRequest
	GetSort() *string
	SetSourceIp(v string) *DescribeOutgoingDomainDetailRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingDomainDetailRequest
	GetStartTime() *string
	SetTagId(v string) *DescribeOutgoingDomainDetailRequest
	GetTagId() *string
}

type DescribeOutgoingDomainDetailRequest struct {
	// example:
	//
	// FullCoverage
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// example.com
	Domain     *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1733450528
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
	// ngw-uf62zzi7000bca7zn****
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
	// 47.96.181.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// OutBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 121.15.137.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1753617600
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// FirstFlow
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DescribeOutgoingDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainDetailRequest) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDomainDetailRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOutgoingDomainDetailRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeOutgoingDomainDetailRequest) GetDomainList() []*string {
	return s.DomainList
}

func (s *DescribeOutgoingDomainDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingDomainDetailRequest) GetIPType() *string {
	return s.IPType
}

func (s *DescribeOutgoingDomainDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingDomainDetailRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeOutgoingDomainDetailRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOutgoingDomainDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOutgoingDomainDetailRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingDomainDetailRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDomainDetailRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeOutgoingDomainDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingDomainDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingDomainDetailRequest) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDomainDetailRequest) SetAclCoverage(v string) *DescribeOutgoingDomainDetailRequest {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetCurrentPage(v string) *DescribeOutgoingDomainDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetDomain(v string) *DescribeOutgoingDomainDetailRequest {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetDomainList(v []*string) *DescribeOutgoingDomainDetailRequest {
	s.DomainList = v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetEndTime(v string) *DescribeOutgoingDomainDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetIPType(v string) *DescribeOutgoingDomainDetailRequest {
	s.IPType = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetLang(v string) *DescribeOutgoingDomainDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetNatGatewayId(v string) *DescribeOutgoingDomainDetailRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetOrder(v string) *DescribeOutgoingDomainDetailRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetPageSize(v string) *DescribeOutgoingDomainDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetPrivateIP(v string) *DescribeOutgoingDomainDetailRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetPublicIP(v string) *DescribeOutgoingDomainDetailRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetSort(v string) *DescribeOutgoingDomainDetailRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetSourceIp(v string) *DescribeOutgoingDomainDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetStartTime(v string) *DescribeOutgoingDomainDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) SetTagId(v string) *DescribeOutgoingDomainDetailRequest {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
