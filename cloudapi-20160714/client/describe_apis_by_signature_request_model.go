// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisBySignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeApisBySignatureRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisBySignatureRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApisBySignatureRequest
	GetSecurityToken() *string
	SetSignatureId(v string) *DescribeApisBySignatureRequest
	GetSignatureId() *string
}

type DescribeApisBySignatureRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the signature key.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
}

func (s DescribeApisBySignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisBySignatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisBySignatureRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisBySignatureRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisBySignatureRequest) GetSignatureId() *string {
	return s.SignatureId
}

func (s *DescribeApisBySignatureRequest) SetPageNumber(v int32) *DescribeApisBySignatureRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisBySignatureRequest) SetPageSize(v int32) *DescribeApisBySignatureRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisBySignatureRequest) SetSecurityToken(v string) *DescribeApisBySignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisBySignatureRequest) SetSignatureId(v string) *DescribeApisBySignatureRequest {
	s.SignatureId = &v
	return s
}

func (s *DescribeApisBySignatureRequest) Validate() error {
	return dara.Validate(s)
}
