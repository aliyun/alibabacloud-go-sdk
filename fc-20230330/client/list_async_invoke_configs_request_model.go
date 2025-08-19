// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncInvokeConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ListAsyncInvokeConfigsRequest
	GetFunctionName() *string
	SetLimit(v int32) *ListAsyncInvokeConfigsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListAsyncInvokeConfigsRequest
	GetNextToken() *string
}

type ListAsyncInvokeConfigsRequest struct {
	// The function name. If you do not configure this parameter, the asynchronous invocation configurations of all functions are displayed.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The maximum number of entries to be returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The paging information. This parameter specifies the start point of the query.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAsyncInvokeConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncInvokeConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncInvokeConfigsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListAsyncInvokeConfigsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAsyncInvokeConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAsyncInvokeConfigsRequest) SetFunctionName(v string) *ListAsyncInvokeConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListAsyncInvokeConfigsRequest) SetLimit(v int32) *ListAsyncInvokeConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListAsyncInvokeConfigsRequest) SetNextToken(v string) *ListAsyncInvokeConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAsyncInvokeConfigsRequest) Validate() error {
	return dara.Validate(s)
}
