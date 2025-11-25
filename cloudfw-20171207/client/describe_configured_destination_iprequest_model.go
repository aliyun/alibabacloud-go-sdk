// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfiguredDestinationIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeConfiguredDestinationIPRequest
	GetCurrentPage() *string
	SetDestinationIP(v string) *DescribeConfiguredDestinationIPRequest
	GetDestinationIP() *string
	SetDestinationISP(v string) *DescribeConfiguredDestinationIPRequest
	GetDestinationISP() *string
	SetDestinationRegion(v string) *DescribeConfiguredDestinationIPRequest
	GetDestinationRegion() *string
	SetDirection(v string) *DescribeConfiguredDestinationIPRequest
	GetDirection() *string
	SetGroupName(v string) *DescribeConfiguredDestinationIPRequest
	GetGroupName() *string
	SetLang(v string) *DescribeConfiguredDestinationIPRequest
	GetLang() *string
	SetPageSize(v string) *DescribeConfiguredDestinationIPRequest
	GetPageSize() *string
	SetSourceCode(v string) *DescribeConfiguredDestinationIPRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeConfiguredDestinationIPRequest
	GetSourceIp() *string
}

type DescribeConfiguredDestinationIPRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1.1.1.1
	DestinationIP  *string `json:"DestinationIP,omitempty" xml:"DestinationIP,omitempty"`
	DestinationISP *string `json:"DestinationISP,omitempty" xml:"DestinationISP,omitempty"`
	// example:
	//
	// cn-shenzhen
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ignore
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// example:
	//
	// 123.113.99.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeConfiguredDestinationIPRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDestinationIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDestinationIPRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeConfiguredDestinationIPRequest) GetDestinationIP() *string {
	return s.DestinationIP
}

func (s *DescribeConfiguredDestinationIPRequest) GetDestinationISP() *string {
	return s.DestinationISP
}

func (s *DescribeConfiguredDestinationIPRequest) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *DescribeConfiguredDestinationIPRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeConfiguredDestinationIPRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeConfiguredDestinationIPRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeConfiguredDestinationIPRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeConfiguredDestinationIPRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeConfiguredDestinationIPRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeConfiguredDestinationIPRequest) SetCurrentPage(v string) *DescribeConfiguredDestinationIPRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetDestinationIP(v string) *DescribeConfiguredDestinationIPRequest {
	s.DestinationIP = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetDestinationISP(v string) *DescribeConfiguredDestinationIPRequest {
	s.DestinationISP = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetDestinationRegion(v string) *DescribeConfiguredDestinationIPRequest {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetDirection(v string) *DescribeConfiguredDestinationIPRequest {
	s.Direction = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetGroupName(v string) *DescribeConfiguredDestinationIPRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetLang(v string) *DescribeConfiguredDestinationIPRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetPageSize(v string) *DescribeConfiguredDestinationIPRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetSourceCode(v string) *DescribeConfiguredDestinationIPRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) SetSourceIp(v string) *DescribeConfiguredDestinationIPRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeConfiguredDestinationIPRequest) Validate() error {
	return dara.Validate(s)
}
