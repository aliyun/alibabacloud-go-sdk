// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCorpNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *ListCorpNumbersRequest
	GetNumber() *string
	SetPageNumber(v int32) *ListCorpNumbersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCorpNumbersRequest
	GetPageSize() *int32
	SetSipTrunkId(v string) *ListCorpNumbersRequest
	GetSipTrunkId() *string
	SetTenantId(v string) *ListCorpNumbersRequest
	GetTenantId() *string
}

type ListCorpNumbersRequest struct {
	// The bound number.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The SIP trunk ID.
	//
	// example:
	//
	// 159d4b8361804ebda21242216a8981c2
	SipTrunkId *string `json:"SipTrunkId,omitempty" xml:"SipTrunkId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListCorpNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCorpNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListCorpNumbersRequest) GetNumber() *string {
	return s.Number
}

func (s *ListCorpNumbersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCorpNumbersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCorpNumbersRequest) GetSipTrunkId() *string {
	return s.SipTrunkId
}

func (s *ListCorpNumbersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCorpNumbersRequest) SetNumber(v string) *ListCorpNumbersRequest {
	s.Number = &v
	return s
}

func (s *ListCorpNumbersRequest) SetPageNumber(v int32) *ListCorpNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCorpNumbersRequest) SetPageSize(v int32) *ListCorpNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListCorpNumbersRequest) SetSipTrunkId(v string) *ListCorpNumbersRequest {
	s.SipTrunkId = &v
	return s
}

func (s *ListCorpNumbersRequest) SetTenantId(v string) *ListCorpNumbersRequest {
	s.TenantId = &v
	return s
}

func (s *ListCorpNumbersRequest) Validate() error {
	return dara.Validate(s)
}
