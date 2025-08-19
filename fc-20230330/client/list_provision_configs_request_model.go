// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ListProvisionConfigsRequest
	GetFunctionName() *string
	SetLimit(v int32) *ListProvisionConfigsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListProvisionConfigsRequest
	GetNextToken() *string
}

type ListProvisionConfigsRequest struct {
	// The name of the function. If this parameter is not specified, the provisioned configurations of all functions are listed.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// Number of provisioned configurations to return.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListProvisionConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListProvisionConfigsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListProvisionConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProvisionConfigsRequest) SetFunctionName(v string) *ListProvisionConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetLimit(v int32) *ListProvisionConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetNextToken(v string) *ListProvisionConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProvisionConfigsRequest) Validate() error {
	return dara.Validate(s)
}
