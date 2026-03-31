// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *DescribeDefenseSceneConfigRequest
	GetConfigKey() *string
	SetDefenseScene(v string) *DescribeDefenseSceneConfigRequest
	GetDefenseScene() *string
	SetInstanceId(v string) *DescribeDefenseSceneConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseSceneConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseSceneConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseSceneConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// autoEnableStatus
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// apisec
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqtm005
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

func (s DescribeDefenseSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseSceneConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeDefenseSceneConfigRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseSceneConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseSceneConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseSceneConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseSceneConfigRequest) SetConfigKey(v string) *DescribeDefenseSceneConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeDefenseSceneConfigRequest) SetDefenseScene(v string) *DescribeDefenseSceneConfigRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseSceneConfigRequest) SetInstanceId(v string) *DescribeDefenseSceneConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseSceneConfigRequest) SetRegionId(v string) *DescribeDefenseSceneConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseSceneConfigRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseSceneConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
