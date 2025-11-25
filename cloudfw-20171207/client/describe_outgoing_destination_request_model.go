// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclCoverage(v string) *DescribeOutgoingDestinationRequest
	GetAclCoverage() *string
	SetApplicationName(v string) *DescribeOutgoingDestinationRequest
	GetApplicationName() *string
	SetCategoryId(v string) *DescribeOutgoingDestinationRequest
	GetCategoryId() *string
	SetCurrentPage(v string) *DescribeOutgoingDestinationRequest
	GetCurrentPage() *string
	SetDstIP(v string) *DescribeOutgoingDestinationRequest
	GetDstIP() *string
	SetEndTime(v string) *DescribeOutgoingDestinationRequest
	GetEndTime() *string
	SetIsAITraffic(v string) *DescribeOutgoingDestinationRequest
	GetIsAITraffic() *string
	SetLang(v string) *DescribeOutgoingDestinationRequest
	GetLang() *string
	SetOrder(v string) *DescribeOutgoingDestinationRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeOutgoingDestinationRequest
	GetPageSize() *string
	SetPort(v string) *DescribeOutgoingDestinationRequest
	GetPort() *string
	SetPrivateIP(v string) *DescribeOutgoingDestinationRequest
	GetPrivateIP() *string
	SetPublicIP(v string) *DescribeOutgoingDestinationRequest
	GetPublicIP() *string
	SetSecuritySuggest(v string) *DescribeOutgoingDestinationRequest
	GetSecuritySuggest() *string
	SetSort(v string) *DescribeOutgoingDestinationRequest
	GetSort() *string
	SetSourceIp(v string) *DescribeOutgoingDestinationRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingDestinationRequest
	GetStartTime() *string
	SetTagId(v string) *DescribeOutgoingDestinationRequest
	GetTagId() *string
}

type DescribeOutgoingDestinationRequest struct {
	// example:
	//
	// FullCoverage
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// AliYun
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 47.100.111XXX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749089441
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// true
	IsAITraffic *string `json:"IsAITraffic,omitempty" xml:"IsAITraffic,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 10.111.53XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// example:
	//
	// 47.96.74.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// pass
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// example:
	//
	// InBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 106.3.198.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749657600
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// FirstFlow
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DescribeOutgoingDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationRequest) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDestinationRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeOutgoingDestinationRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDestinationRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOutgoingDestinationRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeOutgoingDestinationRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingDestinationRequest) GetIsAITraffic() *string {
	return s.IsAITraffic
}

func (s *DescribeOutgoingDestinationRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingDestinationRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOutgoingDestinationRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOutgoingDestinationRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeOutgoingDestinationRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingDestinationRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDestinationRequest) GetSecuritySuggest() *string {
	return s.SecuritySuggest
}

func (s *DescribeOutgoingDestinationRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeOutgoingDestinationRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingDestinationRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingDestinationRequest) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDestinationRequest) SetAclCoverage(v string) *DescribeOutgoingDestinationRequest {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetApplicationName(v string) *DescribeOutgoingDestinationRequest {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetCategoryId(v string) *DescribeOutgoingDestinationRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetCurrentPage(v string) *DescribeOutgoingDestinationRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetDstIP(v string) *DescribeOutgoingDestinationRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetEndTime(v string) *DescribeOutgoingDestinationRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetIsAITraffic(v string) *DescribeOutgoingDestinationRequest {
	s.IsAITraffic = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetLang(v string) *DescribeOutgoingDestinationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetOrder(v string) *DescribeOutgoingDestinationRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetPageSize(v string) *DescribeOutgoingDestinationRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetPort(v string) *DescribeOutgoingDestinationRequest {
	s.Port = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetPrivateIP(v string) *DescribeOutgoingDestinationRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetPublicIP(v string) *DescribeOutgoingDestinationRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetSecuritySuggest(v string) *DescribeOutgoingDestinationRequest {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetSort(v string) *DescribeOutgoingDestinationRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetSourceIp(v string) *DescribeOutgoingDestinationRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetStartTime(v string) *DescribeOutgoingDestinationRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) SetTagId(v string) *DescribeOutgoingDestinationRequest {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationRequest) Validate() error {
	return dara.Validate(s)
}
