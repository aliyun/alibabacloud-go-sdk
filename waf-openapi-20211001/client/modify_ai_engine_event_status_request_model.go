// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiEngineEventStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyAiEngineEventStatusRequest
	GetComment() *string
	SetEventId(v string) *ModifyAiEngineEventStatusRequest
	GetEventId() *string
	SetInstanceId(v string) *ModifyAiEngineEventStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyAiEngineEventStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyAiEngineEventStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetTargetStatus(v string) *ModifyAiEngineEventStatusRequest
	GetTargetStatus() *string
}

type ModifyAiEngineEventStatusRequest struct {
	// The comment. The value can be up to 512 characters in length.
	//
	// example:
	//
	// testComment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVT-20261025109
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_des*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-de***
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The target status of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// handling
	TargetStatus *string `json:"TargetStatus,omitempty" xml:"TargetStatus,omitempty"`
}

func (s ModifyAiEngineEventStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiEngineEventStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAiEngineEventStatusRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyAiEngineEventStatusRequest) GetEventId() *string {
	return s.EventId
}

func (s *ModifyAiEngineEventStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAiEngineEventStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAiEngineEventStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyAiEngineEventStatusRequest) GetTargetStatus() *string {
	return s.TargetStatus
}

func (s *ModifyAiEngineEventStatusRequest) SetComment(v string) *ModifyAiEngineEventStatusRequest {
	s.Comment = &v
	return s
}

func (s *ModifyAiEngineEventStatusRequest) SetEventId(v string) *ModifyAiEngineEventStatusRequest {
	s.EventId = &v
	return s
}

func (s *ModifyAiEngineEventStatusRequest) SetInstanceId(v string) *ModifyAiEngineEventStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAiEngineEventStatusRequest) SetRegionId(v string) *ModifyAiEngineEventStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAiEngineEventStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyAiEngineEventStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyAiEngineEventStatusRequest) SetTargetStatus(v string) *ModifyAiEngineEventStatusRequest {
	s.TargetStatus = &v
	return s
}

func (s *ModifyAiEngineEventStatusRequest) Validate() error {
	return dara.Validate(s)
}
