// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListFunctionsShrinkRequest
	GetDescription() *string
	SetFcVersion(v string) *ListFunctionsShrinkRequest
	GetFcVersion() *string
	SetFunctionName(v string) *ListFunctionsShrinkRequest
	GetFunctionName() *string
	SetGpuType(v string) *ListFunctionsShrinkRequest
	GetGpuType() *string
	SetLimit(v int32) *ListFunctionsShrinkRequest
	GetLimit() *int32
	SetNextToken(v string) *ListFunctionsShrinkRequest
	GetNextToken() *string
	SetPrefix(v string) *ListFunctionsShrinkRequest
	GetPrefix() *string
	SetResourceGroupId(v string) *ListFunctionsShrinkRequest
	GetResourceGroupId() *string
	SetRuntime(v string) *ListFunctionsShrinkRequest
	GetRuntime() *string
	SetTagsShrink(v string) *ListFunctionsShrinkRequest
	GetTagsShrink() *string
}

type ListFunctionsShrinkRequest struct {
	// The description of the functions to retrieve.
	//
	// example:
	//
	// test_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The version of Function Compute to which the functions belong.
	//
	// 	- v3: Only lists functions of Function Compute 3.0.
	//
	// 	- v2: Only lists functions of Function Compute 2.0.
	//
	// By default, this parameter is left empty and functions in both Function Compute 3.0 and Function Compute 2.0 are listed.
	//
	// example:
	//
	// v3
	FcVersion    *string `json:"fcVersion,omitempty" xml:"fcVersion,omitempty"`
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The GPU type of the functions to retrieve.
	//
	// example:
	//
	// fc.gpu.tesla.1
	GpuType *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
	// The number of functions to return. The minimum value is 1 and the maximum value is 100.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The prefix of the function name.
	//
	// example:
	//
	// my-func
	Prefix          *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The runtime of the functions to retrieve.
	//
	// example:
	//
	// python3.10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The tag of the functions to retrieve.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListFunctionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ListFunctionsShrinkRequest) GetFcVersion() *string {
	return s.FcVersion
}

func (s *ListFunctionsShrinkRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListFunctionsShrinkRequest) GetGpuType() *string {
	return s.GpuType
}

func (s *ListFunctionsShrinkRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFunctionsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFunctionsShrinkRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListFunctionsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFunctionsShrinkRequest) GetRuntime() *string {
	return s.Runtime
}

func (s *ListFunctionsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListFunctionsShrinkRequest) SetDescription(v string) *ListFunctionsShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetFcVersion(v string) *ListFunctionsShrinkRequest {
	s.FcVersion = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetFunctionName(v string) *ListFunctionsShrinkRequest {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetGpuType(v string) *ListFunctionsShrinkRequest {
	s.GpuType = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetLimit(v int32) *ListFunctionsShrinkRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetNextToken(v string) *ListFunctionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetPrefix(v string) *ListFunctionsShrinkRequest {
	s.Prefix = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetResourceGroupId(v string) *ListFunctionsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetRuntime(v string) *ListFunctionsShrinkRequest {
	s.Runtime = &v
	return s
}

func (s *ListFunctionsShrinkRequest) SetTagsShrink(v string) *ListFunctionsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListFunctionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
