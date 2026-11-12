// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *InstallCloudAssistantRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *InstallCloudAssistantRequest
	GetRegionId() *string
}

type InstallCloudAssistantRequest struct {
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAssistantRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *InstallCloudAssistantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallCloudAssistantRequest) SetInstanceIds(v []*string) *InstallCloudAssistantRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallCloudAssistantRequest) SetRegionId(v string) *InstallCloudAssistantRequest {
	s.RegionId = &v
	return s
}

func (s *InstallCloudAssistantRequest) Validate() error {
	return dara.Validate(s)
}
