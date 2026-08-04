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
	// The accelerator type used to filter the results. Valid values: `CPU` and `GPU`.
	//
	// example:
	//
	// GPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The GPU type used to filter the results. Fuzzy matching is supported.
	//
	// example:
	//
	// A10
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// A comma-separated list of node types. If omitted, all node types are returned.
	//
	// example:
	//
	// ecs.c6.large,ecs.gn7i-c32g1.8xlarge
	NodeTypes *string `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty"`
	// The quota ID. You must specify either this parameter or `ResourceGroupIds`.
	//
	// example:
	//
	// quota123
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// A comma-separated list of resource group IDs. You must specify either this parameter or `QuotaId`. The `UserId` for the specified `ResourceGroupId` must match the `UserId` of the requesting user. All specified `resource group` instances must be of the same type and associated with the same Virtual Private Cloud (VPC).
	//
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
