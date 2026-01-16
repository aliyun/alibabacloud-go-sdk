// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterNodePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterNodePoolShrinkRequest
	GetClusterId() *string
	SetKubernetesConfigShrink(v string) *CreateClusterNodePoolShrinkRequest
	GetKubernetesConfigShrink() *string
	SetNodepoolInfoShrink(v string) *CreateClusterNodePoolShrinkRequest
	GetNodepoolInfoShrink() *string
	SetScalingGroupShrink(v string) *CreateClusterNodePoolShrinkRequest
	GetScalingGroupShrink() *string
}

type CreateClusterNodePoolShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId              *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	KubernetesConfigShrink *string `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty"`
	// This parameter is required.
	NodepoolInfoShrink *string `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty"`
	// This parameter is required.
	ScalingGroupShrink *string `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty"`
}

func (s CreateClusterNodePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterNodePoolShrinkRequest) GetKubernetesConfigShrink() *string {
	return s.KubernetesConfigShrink
}

func (s *CreateClusterNodePoolShrinkRequest) GetNodepoolInfoShrink() *string {
	return s.NodepoolInfoShrink
}

func (s *CreateClusterNodePoolShrinkRequest) GetScalingGroupShrink() *string {
	return s.ScalingGroupShrink
}

func (s *CreateClusterNodePoolShrinkRequest) SetClusterId(v string) *CreateClusterNodePoolShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterNodePoolShrinkRequest) SetKubernetesConfigShrink(v string) *CreateClusterNodePoolShrinkRequest {
	s.KubernetesConfigShrink = &v
	return s
}

func (s *CreateClusterNodePoolShrinkRequest) SetNodepoolInfoShrink(v string) *CreateClusterNodePoolShrinkRequest {
	s.NodepoolInfoShrink = &v
	return s
}

func (s *CreateClusterNodePoolShrinkRequest) SetScalingGroupShrink(v string) *CreateClusterNodePoolShrinkRequest {
	s.ScalingGroupShrink = &v
	return s
}

func (s *CreateClusterNodePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
