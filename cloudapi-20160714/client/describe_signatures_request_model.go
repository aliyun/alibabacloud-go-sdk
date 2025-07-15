// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSignaturesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSignaturesRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeSignaturesRequest
	GetSecurityToken() *string
	SetSignatureId(v string) *DescribeSignaturesRequest
	GetSignatureId() *string
	SetSignatureName(v string) *DescribeSignaturesRequest
	GetSignatureName() *string
}

type DescribeSignaturesRequest struct {
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
	// The IDs of the keys to query.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The names of the keys to query.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s DescribeSignaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSignaturesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSignaturesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSignaturesRequest) GetSignatureId() *string {
	return s.SignatureId
}

func (s *DescribeSignaturesRequest) GetSignatureName() *string {
	return s.SignatureName
}

func (s *DescribeSignaturesRequest) SetPageNumber(v int32) *DescribeSignaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSignaturesRequest) SetPageSize(v int32) *DescribeSignaturesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSignaturesRequest) SetSecurityToken(v string) *DescribeSignaturesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSignaturesRequest) SetSignatureId(v string) *DescribeSignaturesRequest {
	s.SignatureId = &v
	return s
}

func (s *DescribeSignaturesRequest) SetSignatureName(v string) *DescribeSignaturesRequest {
	s.SignatureName = &v
	return s
}

func (s *DescribeSignaturesRequest) Validate() error {
	return dara.Validate(s)
}
