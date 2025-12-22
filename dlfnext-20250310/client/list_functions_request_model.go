// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionNamePattern(v string) *ListFunctionsRequest
	GetFunctionNamePattern() *string
	SetMaxResults(v int32) *ListFunctionsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListFunctionsRequest
	GetPageToken() *string
}

type ListFunctionsRequest struct {
	// example:
	//
	// function%
	FunctionNamePattern *string `json:"functionNamePattern,omitempty" xml:"functionNamePattern,omitempty"`
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) GetFunctionNamePattern() *string {
	return s.FunctionNamePattern
}

func (s *ListFunctionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFunctionsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListFunctionsRequest) SetFunctionNamePattern(v string) *ListFunctionsRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListFunctionsRequest) SetMaxResults(v int32) *ListFunctionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFunctionsRequest) SetPageToken(v string) *ListFunctionsRequest {
	s.PageToken = &v
	return s
}

func (s *ListFunctionsRequest) Validate() error {
	return dara.Validate(s)
}
