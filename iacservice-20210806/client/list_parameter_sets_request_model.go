// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListParameterSetsRequest
	GetKeyword() *string
	SetKmsKeyId(v string) *ListParameterSetsRequest
	GetKmsKeyId() *string
	SetPageNumber(v int32) *ListParameterSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListParameterSetsRequest
	GetPageSize() *int32
}

type ListParameterSetsRequest struct {
	// The keyword for exact match search. The search is case-insensitive.
	//
	// example:
	//
	// vpc
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 21a90f5d-a469-4ac4-a8ea-f6e1e7470e6f
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListParameterSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetsRequest) GoString() string {
	return s.String()
}

func (s *ListParameterSetsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListParameterSetsRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *ListParameterSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParameterSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParameterSetsRequest) SetKeyword(v string) *ListParameterSetsRequest {
	s.Keyword = &v
	return s
}

func (s *ListParameterSetsRequest) SetKmsKeyId(v string) *ListParameterSetsRequest {
	s.KmsKeyId = &v
	return s
}

func (s *ListParameterSetsRequest) SetPageNumber(v int32) *ListParameterSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParameterSetsRequest) SetPageSize(v int32) *ListParameterSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListParameterSetsRequest) Validate() error {
	return dara.Validate(s)
}
