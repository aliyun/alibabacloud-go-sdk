// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupByResourceGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupList(v []*string) *CreateMonitorGroupByResourceGroupIdRequest
	GetContactGroupList() []*string
	SetEnableInstallAgent(v bool) *CreateMonitorGroupByResourceGroupIdRequest
	GetEnableInstallAgent() *bool
	SetEnableSubscribeEvent(v bool) *CreateMonitorGroupByResourceGroupIdRequest
	GetEnableSubscribeEvent() *bool
	SetRegionId(v string) *CreateMonitorGroupByResourceGroupIdRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateMonitorGroupByResourceGroupIdRequest
	GetResourceGroupId() *string
	SetResourceGroupName(v string) *CreateMonitorGroupByResourceGroupIdRequest
	GetResourceGroupName() *string
}

type CreateMonitorGroupByResourceGroupIdRequest struct {
	// The alert contact groups. The alert notifications of the application group are sent to the alert contacts that belong to the specified alert contact groups.
	//
	// An alert contact group can contain one or more alert contacts. For information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html). For information about how to obtain alert contact groups, see [DescribeContactGroupList](https://help.aliyun.com/document_detail/114922.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroupList []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	// Specifies whether the CloudMonitor agent is automatically installed for the application group. CloudMonitor determines whether to automatically install the CloudMonitor agent for the hosts in an application group based on the value of this parameter. Valid values:
	//
	// 	- true: The CloudMonitor agent is automatically installed.
	//
	// 	- false (default): The CloudMonitor agent is not automatically installed.
	//
	// example:
	//
	// true
	EnableInstallAgent *bool `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	// Specifies whether the application group automatically subscribes to event notifications. If events whose severity level is critical or warning occur on resources in an application group, CloudMonitor sends alert notifications. Valid values:
	//
	// 	- true: The application group automatically subscribes to event notifications.
	//
	// 	- false (default): The application group does not automatically subscribe to event notifications.
	//
	// example:
	//
	// true
	EnableSubscribeEvent *bool `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	// The ID of the region where the resource group resides.
	//
	// For information about how to obtain the ID of the region where a resource group resides, see [GetResourceGroup](https://help.aliyun.com/document_detail/158866.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For information about how to obtain the ID of a resource group, see [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the resource group.
	//
	// For information about how to obtain the name of a resource group, see [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudMonitor
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CreateMonitorGroupByResourceGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupByResourceGroupIdRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) GetContactGroupList() []*string {
	return s.ContactGroupList
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) GetEnableInstallAgent() *bool {
	return s.EnableInstallAgent
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) GetEnableSubscribeEvent() *bool {
	return s.EnableSubscribeEvent
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetContactGroupList(v []*string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.ContactGroupList = v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetEnableInstallAgent(v bool) *CreateMonitorGroupByResourceGroupIdRequest {
	s.EnableInstallAgent = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetEnableSubscribeEvent(v bool) *CreateMonitorGroupByResourceGroupIdRequest {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetRegionId(v string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetResourceGroupId(v string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetResourceGroupName(v string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
