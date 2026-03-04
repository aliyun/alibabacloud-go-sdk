// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeProductionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListYikeProductionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListYikeProductionsRequest
	GetNextToken() *string
	SetPageNo(v int32) *ListYikeProductionsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListYikeProductionsRequest
	GetPageSize() *int32
}

type ListYikeProductionsRequest struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListYikeProductionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListYikeProductionsRequest) GoString() string {
	return s.String()
}

func (s *ListYikeProductionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListYikeProductionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListYikeProductionsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListYikeProductionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListYikeProductionsRequest) SetMaxResults(v int32) *ListYikeProductionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListYikeProductionsRequest) SetNextToken(v string) *ListYikeProductionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListYikeProductionsRequest) SetPageNo(v int32) *ListYikeProductionsRequest {
	s.PageNo = &v
	return s
}

func (s *ListYikeProductionsRequest) SetPageSize(v int32) *ListYikeProductionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListYikeProductionsRequest) Validate() error {
	return dara.Validate(s)
}
