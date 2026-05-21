// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListNodeTypesRequest
	GetAcceleratorType() *string
	SetGPUType(v string) *ListNodeTypesRequest
	GetGPUType() *string
	SetNodeTypes(v string) *ListNodeTypesRequest
	GetNodeTypes() *string
	SetQuotaId(v string) *ListNodeTypesRequest
	GetQuotaId() *string
	SetResourceGroupIds(v string) *ListNodeTypesRequest
	GetResourceGroupIds() *string
}

type ListNodeTypesRequest struct {
	// example:
	//
	// GPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// example:
	//
	// A10
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// ecs.c6.large,ecs.gn7i-c32g1.8xlarge
	NodeTypes *string `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty"`
	// example:
	//
	// quota123
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// rg123,rg456
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
}

func (s ListNodeTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeTypesRequest) GoString() string {
	return s.String()
}

func (s *ListNodeTypesRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListNodeTypesRequest) GetGPUType() *string {
	return s.GPUType
}

func (s *ListNodeTypesRequest) GetNodeTypes() *string {
	return s.NodeTypes
}

func (s *ListNodeTypesRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListNodeTypesRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *ListNodeTypesRequest) SetAcceleratorType(v string) *ListNodeTypesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListNodeTypesRequest) SetGPUType(v string) *ListNodeTypesRequest {
	s.GPUType = &v
	return s
}

func (s *ListNodeTypesRequest) SetNodeTypes(v string) *ListNodeTypesRequest {
	s.NodeTypes = &v
	return s
}

func (s *ListNodeTypesRequest) SetQuotaId(v string) *ListNodeTypesRequest {
	s.QuotaId = &v
	return s
}

func (s *ListNodeTypesRequest) SetResourceGroupIds(v string) *ListNodeTypesRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *ListNodeTypesRequest) Validate() error {
	return dara.Validate(s)
}
