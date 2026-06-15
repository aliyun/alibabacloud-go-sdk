// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPaidType(v int64) *CreateComputeInstanceRequest
	GetPaidType() *int64
	SetRegionId(v string) *CreateComputeInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateComputeInstanceRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *CreateComputeInstanceRequest
	GetResourceType() *string
}

type CreateComputeInstanceRequest struct {
	// This parameter is required.
	PaidType *int64 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// This parameter is required.
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeInstanceRequest) GetPaidType() *int64 {
	return s.PaidType
}

func (s *CreateComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateComputeInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateComputeInstanceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateComputeInstanceRequest) SetPaidType(v int64) *CreateComputeInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreateComputeInstanceRequest) SetRegionId(v string) *CreateComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateComputeInstanceRequest) SetResourceGroupId(v string) *CreateComputeInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateComputeInstanceRequest) SetResourceType(v string) *CreateComputeInstanceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
