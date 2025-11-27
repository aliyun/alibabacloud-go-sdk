// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallRCCloudAssistantShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *InstallRCCloudAssistantShrinkRequest
	GetInstanceIdsShrink() *string
	SetRegionId(v string) *InstallRCCloudAssistantShrinkRequest
	GetRegionId() *string
}

type InstallRCCloudAssistantShrinkRequest struct {
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallRCCloudAssistantShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallRCCloudAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallRCCloudAssistantShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *InstallRCCloudAssistantShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallRCCloudAssistantShrinkRequest) SetInstanceIdsShrink(v string) *InstallRCCloudAssistantShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *InstallRCCloudAssistantShrinkRequest) SetRegionId(v string) *InstallRCCloudAssistantShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *InstallRCCloudAssistantShrinkRequest) Validate() error {
	return dara.Validate(s)
}
