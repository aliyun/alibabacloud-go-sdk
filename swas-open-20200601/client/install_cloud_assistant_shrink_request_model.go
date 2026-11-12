// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAssistantShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *InstallCloudAssistantShrinkRequest
	GetInstanceIdsShrink() *string
	SetRegionId(v string) *InstallCloudAssistantShrinkRequest
	GetRegionId() *string
}

type InstallCloudAssistantShrinkRequest struct {
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudAssistantShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *InstallCloudAssistantShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallCloudAssistantShrinkRequest) SetInstanceIdsShrink(v string) *InstallCloudAssistantShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *InstallCloudAssistantShrinkRequest) SetRegionId(v string) *InstallCloudAssistantShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *InstallCloudAssistantShrinkRequest) Validate() error {
	return dara.Validate(s)
}
