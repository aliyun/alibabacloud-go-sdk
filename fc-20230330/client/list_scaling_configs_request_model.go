// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScalingConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ListScalingConfigsRequest
	GetFunctionName() *string
	SetLimit(v int32) *ListScalingConfigsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListScalingConfigsRequest
	GetNextToken() *string
}

type ListScalingConfigsRequest struct {
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListScalingConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScalingConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListScalingConfigsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListScalingConfigsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListScalingConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScalingConfigsRequest) SetFunctionName(v string) *ListScalingConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListScalingConfigsRequest) SetLimit(v int32) *ListScalingConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListScalingConfigsRequest) SetNextToken(v string) *ListScalingConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListScalingConfigsRequest) Validate() error {
	return dara.Validate(s)
}
