// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserPoolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolsRequest
	GetNextToken() *string
}

type ListUserPoolsRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdFVzZXJQb29sczo6MTA=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPoolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolsRequest) SetMaxResults(v int32) *ListUserPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolsRequest) SetNextToken(v string) *ListUserPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolsRequest) Validate() error {
	return dara.Validate(s)
}
