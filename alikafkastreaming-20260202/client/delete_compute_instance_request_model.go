// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteComputeInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteComputeInstanceRequest
	GetRegionId() *string
	SetResourceType(v string) *DeleteComputeInstanceRequest
	GetResourceType() *string
}

type DeleteComputeInstanceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteComputeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteComputeInstanceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteComputeInstanceRequest) SetInstanceId(v string) *DeleteComputeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteComputeInstanceRequest) SetRegionId(v string) *DeleteComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteComputeInstanceRequest) SetResourceType(v string) *DeleteComputeInstanceRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
