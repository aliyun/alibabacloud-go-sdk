// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DescribeResolverRulesRequest
	GetEndpointId() *string
	SetKeyword(v string) *DescribeResolverRulesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeResolverRulesRequest
	GetLang() *string
	SetNeedDetailAttributes(v bool) *DescribeResolverRulesRequest
	GetNeedDetailAttributes() *bool
	SetPageNumber(v int32) *DescribeResolverRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResolverRulesRequest
	GetPageSize() *int32
}

type DescribeResolverRulesRequest struct {
	// The outbound endpoint ID.
	//
	// example:
	//
	// hr****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The keyword of the forwarding rule name. Fuzzy search is supported. The value is not case-sensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to return virtual private clouds (VPCs) associated with the forwarding rule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	NeedDetailAttributes *bool `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeResolverRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeResolverRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeResolverRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResolverRulesRequest) GetNeedDetailAttributes() *bool {
	return s.NeedDetailAttributes
}

func (s *DescribeResolverRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResolverRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResolverRulesRequest) SetEndpointId(v string) *DescribeResolverRulesRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetKeyword(v string) *DescribeResolverRulesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetLang(v string) *DescribeResolverRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetNeedDetailAttributes(v bool) *DescribeResolverRulesRequest {
	s.NeedDetailAttributes = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetPageNumber(v int32) *DescribeResolverRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetPageSize(v int32) *DescribeResolverRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverRulesRequest) Validate() error {
	return dara.Validate(s)
}
