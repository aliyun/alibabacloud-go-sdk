// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransferDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeTransferDomainsRequest
	GetDomainName() *string
	SetFromUserId(v int64) *DescribeTransferDomainsRequest
	GetFromUserId() *int64
	SetLang(v string) *DescribeTransferDomainsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeTransferDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeTransferDomainsRequest
	GetPageSize() *int64
	SetTargetUserId(v int64) *DescribeTransferDomainsRequest
	GetTargetUserId() *int64
	SetTransferType(v string) *DescribeTransferDomainsRequest
	GetTransferType() *string
}

type DescribeTransferDomainsRequest struct {
	// The domain name. Use this parameter to query the transfer records of a specific domain name.
	//
	// example:
	//
	// example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the source account. Use this parameter to query the list of domain names transferred from this account to the current account.
	//
	// example:
	//
	// 1*******
	FromUserId *int64 `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// The language of the response.
	//
	// Valid values:
	//
	// zh: Chinese
	//
	// en: English
	//
	// The default value is zh.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the destination account. Use this parameter to query the list of domain names transferred from the current account to this account.
	//
	// example:
	//
	// 1*******
	TargetUserId *int64 `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	// The type of transfer. Valid values:
	//
	// - IN: The list of domain names transferred to the current account.
	//
	// - OUT: The list of domain names transferred from the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// IN
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s DescribeTransferDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeTransferDomainsRequest) GetFromUserId() *int64 {
	return s.FromUserId
}

func (s *DescribeTransferDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTransferDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeTransferDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTransferDomainsRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *DescribeTransferDomainsRequest) GetTransferType() *string {
	return s.TransferType
}

func (s *DescribeTransferDomainsRequest) SetDomainName(v string) *DescribeTransferDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetFromUserId(v int64) *DescribeTransferDomainsRequest {
	s.FromUserId = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetLang(v string) *DescribeTransferDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetPageNumber(v int64) *DescribeTransferDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetPageSize(v int64) *DescribeTransferDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetTargetUserId(v int64) *DescribeTransferDomainsRequest {
	s.TargetUserId = &v
	return s
}

func (s *DescribeTransferDomainsRequest) SetTransferType(v string) *DescribeTransferDomainsRequest {
	s.TransferType = &v
	return s
}

func (s *DescribeTransferDomainsRequest) Validate() error {
	return dara.Validate(s)
}
