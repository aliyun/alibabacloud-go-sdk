// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridCloudGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *DeleteHybridCloudGroupRequest
	GetClusterId() *int64
	SetGroupId(v int64) *DeleteHybridCloudGroupRequest
	GetGroupId() *int64
	SetInstanceId(v string) *DeleteHybridCloudGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteHybridCloudGroupRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteHybridCloudGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteHybridCloudGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-*****tm005
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteHybridCloudGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridCloudGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridCloudGroupRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *DeleteHybridCloudGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteHybridCloudGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHybridCloudGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHybridCloudGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteHybridCloudGroupRequest) SetClusterId(v int64) *DeleteHybridCloudGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHybridCloudGroupRequest) SetGroupId(v int64) *DeleteHybridCloudGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteHybridCloudGroupRequest) SetInstanceId(v string) *DeleteHybridCloudGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHybridCloudGroupRequest) SetRegionId(v string) *DeleteHybridCloudGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHybridCloudGroupRequest) SetResourceManagerResourceGroupId(v string) *DeleteHybridCloudGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteHybridCloudGroupRequest) Validate() error {
	return dara.Validate(s)
}
