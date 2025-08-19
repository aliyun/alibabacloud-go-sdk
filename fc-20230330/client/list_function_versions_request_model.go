// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListFunctionVersionsRequest
	GetDirection() *string
	SetLimit(v int32) *ListFunctionVersionsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListFunctionVersionsRequest
	GetNextToken() *string
}

type ListFunctionVersionsRequest struct {
	// The sorting mode of function versions. Valid values: BACKWARD and FORWARD.
	//
	// example:
	//
	// BACKWARD
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The number of function versions that are returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionVersionsRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListFunctionVersionsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFunctionVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFunctionVersionsRequest) SetDirection(v string) *ListFunctionVersionsRequest {
	s.Direction = &v
	return s
}

func (s *ListFunctionVersionsRequest) SetLimit(v int32) *ListFunctionVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionVersionsRequest) SetNextToken(v string) *ListFunctionVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFunctionVersionsRequest) Validate() error {
	return dara.Validate(s)
}
