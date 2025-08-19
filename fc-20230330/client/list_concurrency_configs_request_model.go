// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConcurrencyConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ListConcurrencyConfigsRequest
	GetFunctionName() *string
	SetLimit(v int32) *ListConcurrencyConfigsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListConcurrencyConfigsRequest
	GetNextToken() *string
}

type ListConcurrencyConfigsRequest struct {
	// The function name. If you leave this parameter empty, the concurrency configurations of all functions are returned.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The maximum number of entries returned.
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

func (s ListConcurrencyConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConcurrencyConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConcurrencyConfigsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListConcurrencyConfigsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListConcurrencyConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConcurrencyConfigsRequest) SetFunctionName(v string) *ListConcurrencyConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListConcurrencyConfigsRequest) SetLimit(v int32) *ListConcurrencyConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListConcurrencyConfigsRequest) SetNextToken(v string) *ListConcurrencyConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConcurrencyConfigsRequest) Validate() error {
	return dara.Validate(s)
}
