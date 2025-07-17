// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagToFlinkClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddTagToFlinkClusterRequest
	GetClusterId() *string
	SetFlinkWorkSpaceId(v string) *AddTagToFlinkClusterRequest
	GetFlinkWorkSpaceId() *string
	SetFlinkWorkSpaceName(v string) *AddTagToFlinkClusterRequest
	GetFlinkWorkSpaceName() *string
	SetRegionId(v string) *AddTagToFlinkClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AddTagToFlinkClusterRequest
	GetResourceGroupId() *string
	SetTargetUserId(v string) *AddTagToFlinkClusterRequest
	GetTargetUserId() *string
}

type AddTagToFlinkClusterRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5defa51f******c92bd2ef5fb093269
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Flink workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ab18f***93744d
	FlinkWorkSpaceId *string `json:"FlinkWorkSpaceId,omitempty" xml:"FlinkWorkSpaceId,omitempty"`
	// The name of the Flink workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// flink-demo
	FlinkWorkSpaceName *string `json:"FlinkWorkSpaceName,omitempty" xml:"FlinkWorkSpaceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Alibaba Cloud account to which the Flink workspace belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 198608******7619
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s AddTagToFlinkClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagToFlinkClusterRequest) GoString() string {
	return s.String()
}

func (s *AddTagToFlinkClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddTagToFlinkClusterRequest) GetFlinkWorkSpaceId() *string {
	return s.FlinkWorkSpaceId
}

func (s *AddTagToFlinkClusterRequest) GetFlinkWorkSpaceName() *string {
	return s.FlinkWorkSpaceName
}

func (s *AddTagToFlinkClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagToFlinkClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddTagToFlinkClusterRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *AddTagToFlinkClusterRequest) SetClusterId(v string) *AddTagToFlinkClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *AddTagToFlinkClusterRequest) SetFlinkWorkSpaceId(v string) *AddTagToFlinkClusterRequest {
	s.FlinkWorkSpaceId = &v
	return s
}

func (s *AddTagToFlinkClusterRequest) SetFlinkWorkSpaceName(v string) *AddTagToFlinkClusterRequest {
	s.FlinkWorkSpaceName = &v
	return s
}

func (s *AddTagToFlinkClusterRequest) SetRegionId(v string) *AddTagToFlinkClusterRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagToFlinkClusterRequest) SetResourceGroupId(v string) *AddTagToFlinkClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddTagToFlinkClusterRequest) SetTargetUserId(v string) *AddTagToFlinkClusterRequest {
	s.TargetUserId = &v
	return s
}

func (s *AddTagToFlinkClusterRequest) Validate() error {
	return dara.Validate(s)
}
