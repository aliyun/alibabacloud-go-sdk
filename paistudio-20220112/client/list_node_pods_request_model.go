// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePodsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGPUIndexes(v string) *ListNodePodsRequest
	GetGPUIndexes() *string
	SetOversoldTypes(v string) *ListNodePodsRequest
	GetOversoldTypes() *string
	SetResourceGroupId(v string) *ListNodePodsRequest
	GetResourceGroupId() *string
}

type ListNodePodsRequest struct {
	// The GPU index number.
	//
	// example:
	//
	// 1
	GPUIndexes *string `json:"GPUIndexes,omitempty" xml:"GPUIndexes,omitempty"`
	// The resource type used by the pod.
	//
	// example:
	//
	// ForceQuotaOversold
	OversoldTypes *string `json:"OversoldTypes,omitempty" xml:"OversoldTypes,omitempty"`
	// The ID of the resource group to which the node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// lingj19q90jp66nq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodePodsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodePodsRequest) GoString() string {
	return s.String()
}

func (s *ListNodePodsRequest) GetGPUIndexes() *string {
	return s.GPUIndexes
}

func (s *ListNodePodsRequest) GetOversoldTypes() *string {
	return s.OversoldTypes
}

func (s *ListNodePodsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNodePodsRequest) SetGPUIndexes(v string) *ListNodePodsRequest {
	s.GPUIndexes = &v
	return s
}

func (s *ListNodePodsRequest) SetOversoldTypes(v string) *ListNodePodsRequest {
	s.OversoldTypes = &v
	return s
}

func (s *ListNodePodsRequest) SetResourceGroupId(v string) *ListNodePodsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNodePodsRequest) Validate() error {
	return dara.Validate(s)
}
