// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNodePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterNodePoolShrinkRequest
	GetClusterId() *string
	SetKubernetesConfigShrink(v string) *ModifyClusterNodePoolShrinkRequest
	GetKubernetesConfigShrink() *string
	SetNodepoolInfoShrink(v string) *ModifyClusterNodePoolShrinkRequest
	GetNodepoolInfoShrink() *string
	SetScalingGroupShrink(v string) *ModifyClusterNodePoolShrinkRequest
	GetScalingGroupShrink() *string
}

type ModifyClusterNodePoolShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId              *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	KubernetesConfigShrink *string `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty"`
	// This parameter is required.
	NodepoolInfoShrink *string `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty"`
	ScalingGroupShrink *string `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty"`
}

func (s ModifyClusterNodePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterNodePoolShrinkRequest) GetKubernetesConfigShrink() *string {
	return s.KubernetesConfigShrink
}

func (s *ModifyClusterNodePoolShrinkRequest) GetNodepoolInfoShrink() *string {
	return s.NodepoolInfoShrink
}

func (s *ModifyClusterNodePoolShrinkRequest) GetScalingGroupShrink() *string {
	return s.ScalingGroupShrink
}

func (s *ModifyClusterNodePoolShrinkRequest) SetClusterId(v string) *ModifyClusterNodePoolShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterNodePoolShrinkRequest) SetKubernetesConfigShrink(v string) *ModifyClusterNodePoolShrinkRequest {
	s.KubernetesConfigShrink = &v
	return s
}

func (s *ModifyClusterNodePoolShrinkRequest) SetNodepoolInfoShrink(v string) *ModifyClusterNodePoolShrinkRequest {
	s.NodepoolInfoShrink = &v
	return s
}

func (s *ModifyClusterNodePoolShrinkRequest) SetScalingGroupShrink(v string) *ModifyClusterNodePoolShrinkRequest {
	s.ScalingGroupShrink = &v
	return s
}

func (s *ModifyClusterNodePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
