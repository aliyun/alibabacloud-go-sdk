// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfiguredDomainNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeConfiguredDomainNamesRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeConfiguredDomainNamesRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeConfiguredDomainNamesRequest
	GetDomainName() *string
	SetGroupName(v string) *DescribeConfiguredDomainNamesRequest
	GetGroupName() *string
	SetLang(v string) *DescribeConfiguredDomainNamesRequest
	GetLang() *string
	SetPageSize(v string) *DescribeConfiguredDomainNamesRequest
	GetPageSize() *string
	SetSourceCode(v string) *DescribeConfiguredDomainNamesRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeConfiguredDomainNamesRequest
	GetSourceIp() *string
}

type DescribeConfiguredDomainNamesRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
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
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// example:
	//
	// 1.202.149.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeConfiguredDomainNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDomainNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDomainNamesRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeConfiguredDomainNamesRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeConfiguredDomainNamesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeConfiguredDomainNamesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeConfiguredDomainNamesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeConfiguredDomainNamesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeConfiguredDomainNamesRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeConfiguredDomainNamesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeConfiguredDomainNamesRequest) SetCurrentPage(v string) *DescribeConfiguredDomainNamesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetDirection(v string) *DescribeConfiguredDomainNamesRequest {
	s.Direction = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetDomainName(v string) *DescribeConfiguredDomainNamesRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetGroupName(v string) *DescribeConfiguredDomainNamesRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetLang(v string) *DescribeConfiguredDomainNamesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetPageSize(v string) *DescribeConfiguredDomainNamesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetSourceCode(v string) *DescribeConfiguredDomainNamesRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) SetSourceIp(v string) *DescribeConfiguredDomainNamesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeConfiguredDomainNamesRequest) Validate() error {
	return dara.Validate(s)
}
