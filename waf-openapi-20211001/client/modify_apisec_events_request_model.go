// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyApisecEventsRequest
	GetClusterId() *string
	SetEventIds(v []*string) *ModifyApisecEventsRequest
	GetEventIds() []*string
	SetEventScope(v string) *ModifyApisecEventsRequest
	GetEventScope() *string
	SetInstanceId(v string) *ModifyApisecEventsRequest
	GetInstanceId() *string
	SetNote(v string) *ModifyApisecEventsRequest
	GetNote() *string
	SetRegionId(v string) *ModifyApisecEventsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecEventsRequest
	GetResourceManagerResourceGroupId() *string
	SetUserStatus(v string) *ModifyApisecEventsRequest
	GetUserStatus() *string
}

type ModifyApisecEventsRequest struct {
	// The hybrid cloud cluster ID.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) to obtain hybrid cloud cluster information.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of API security event IDs.
	//
	// This parameter is required.
	EventIds []*string `json:"EventIds,omitempty" xml:"EventIds,omitempty" type:"Repeated"`
	// The dimension of the security event.
	//
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// already confirmed.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The event status. Valid values:
	//
	// - **toBeConfirmed**: To be confirmed.
	//
	// - **confirmed**: Confirmed.
	//
	// - **ignored**: Ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// confirmed
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s ModifyApisecEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecEventsRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecEventsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyApisecEventsRequest) GetEventIds() []*string {
	return s.EventIds
}

func (s *ModifyApisecEventsRequest) GetEventScope() *string {
	return s.EventScope
}

func (s *ModifyApisecEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecEventsRequest) GetNote() *string {
	return s.Note
}

func (s *ModifyApisecEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecEventsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecEventsRequest) GetUserStatus() *string {
	return s.UserStatus
}

func (s *ModifyApisecEventsRequest) SetClusterId(v string) *ModifyApisecEventsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyApisecEventsRequest) SetEventIds(v []*string) *ModifyApisecEventsRequest {
	s.EventIds = v
	return s
}

func (s *ModifyApisecEventsRequest) SetEventScope(v string) *ModifyApisecEventsRequest {
	s.EventScope = &v
	return s
}

func (s *ModifyApisecEventsRequest) SetInstanceId(v string) *ModifyApisecEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecEventsRequest) SetNote(v string) *ModifyApisecEventsRequest {
	s.Note = &v
	return s
}

func (s *ModifyApisecEventsRequest) SetRegionId(v string) *ModifyApisecEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecEventsRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecEventsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecEventsRequest) SetUserStatus(v string) *ModifyApisecEventsRequest {
	s.UserStatus = &v
	return s
}

func (s *ModifyApisecEventsRequest) Validate() error {
	return dara.Validate(s)
}
