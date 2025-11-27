// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallRCCloudAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *InstallRCCloudAssistantRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *InstallRCCloudAssistantRequest
	GetRegionId() *string
}

type InstallRCCloudAssistantRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallRCCloudAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallRCCloudAssistantRequest) GoString() string {
	return s.String()
}

func (s *InstallRCCloudAssistantRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *InstallRCCloudAssistantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallRCCloudAssistantRequest) SetInstanceIds(v []*string) *InstallRCCloudAssistantRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallRCCloudAssistantRequest) SetRegionId(v string) *InstallRCCloudAssistantRequest {
	s.RegionId = &v
	return s
}

func (s *InstallRCCloudAssistantRequest) Validate() error {
	return dara.Validate(s)
}
