// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyCommandConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandConfigShrink(v string) *SetPolicyCommandConfigShrinkRequest
	GetCommandConfigShrink() *string
	SetInstanceId(v string) *SetPolicyCommandConfigShrinkRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyCommandConfigShrinkRequest
	GetPolicyId() *string
	SetProjectId(v int64) *SetPolicyCommandConfigShrinkRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyCommandConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyCommandConfigShrinkRequest struct {
	// The command control settings.
	//
	// > This parameter applies only to Linux hosts.
	//
	// This parameter is required.
	CommandConfigShrink *string `json:"CommandConfig,omitempty" xml:"CommandConfig,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyCommandConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigShrinkRequest) GetCommandConfigShrink() *string {
	return s.CommandConfigShrink
}

func (s *SetPolicyCommandConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyCommandConfigShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyCommandConfigShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SetPolicyCommandConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyCommandConfigShrinkRequest) SetCommandConfigShrink(v string) *SetPolicyCommandConfigShrinkRequest {
	s.CommandConfigShrink = &v
	return s
}

func (s *SetPolicyCommandConfigShrinkRequest) SetInstanceId(v string) *SetPolicyCommandConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyCommandConfigShrinkRequest) SetPolicyId(v string) *SetPolicyCommandConfigShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyCommandConfigShrinkRequest) SetProjectId(v int64) *SetPolicyCommandConfigShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyCommandConfigShrinkRequest) SetRegionId(v string) *SetPolicyCommandConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyCommandConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
