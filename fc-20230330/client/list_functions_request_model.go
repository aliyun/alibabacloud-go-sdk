// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListFunctionsRequest
	GetDescription() *string
	SetFcVersion(v string) *ListFunctionsRequest
	GetFcVersion() *string
	SetFunctionName(v string) *ListFunctionsRequest
	GetFunctionName() *string
	SetGpuType(v string) *ListFunctionsRequest
	GetGpuType() *string
	SetLimit(v int32) *ListFunctionsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListFunctionsRequest
	GetNextToken() *string
	SetPrefix(v string) *ListFunctionsRequest
	GetPrefix() *string
	SetResourceGroupId(v string) *ListFunctionsRequest
	GetResourceGroupId() *string
	SetRuntime(v string) *ListFunctionsRequest
	GetRuntime() *string
	SetTags(v []*Tag) *ListFunctionsRequest
	GetTags() []*Tag
}

type ListFunctionsRequest struct {
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
	Tags []*Tag `json:"tags" xml:"tags" type:"Repeated"`
}

func (s ListFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListFunctionsRequest) GetFcVersion() *string {
	return s.FcVersion
}

func (s *ListFunctionsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListFunctionsRequest) GetGpuType() *string {
	return s.GpuType
}

func (s *ListFunctionsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFunctionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFunctionsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListFunctionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFunctionsRequest) GetRuntime() *string {
	return s.Runtime
}

func (s *ListFunctionsRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *ListFunctionsRequest) SetDescription(v string) *ListFunctionsRequest {
	s.Description = &v
	return s
}

func (s *ListFunctionsRequest) SetFcVersion(v string) *ListFunctionsRequest {
	s.FcVersion = &v
	return s
}

func (s *ListFunctionsRequest) SetFunctionName(v string) *ListFunctionsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionsRequest) SetGpuType(v string) *ListFunctionsRequest {
	s.GpuType = &v
	return s
}

func (s *ListFunctionsRequest) SetLimit(v int32) *ListFunctionsRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionsRequest) SetNextToken(v string) *ListFunctionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFunctionsRequest) SetPrefix(v string) *ListFunctionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListFunctionsRequest) SetResourceGroupId(v string) *ListFunctionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFunctionsRequest) SetRuntime(v string) *ListFunctionsRequest {
	s.Runtime = &v
	return s
}

func (s *ListFunctionsRequest) SetTags(v []*Tag) *ListFunctionsRequest {
	s.Tags = v
	return s
}

func (s *ListFunctionsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
