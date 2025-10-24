// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendConfig(v string) *UpdatePipelineRequest
	GetExtendConfig() *string
	SetName(v string) *UpdatePipelineRequest
	GetName() *string
	SetNotifyConfig(v string) *UpdatePipelineRequest
	GetNotifyConfig() *string
	SetOwnerAccount(v string) *UpdatePipelineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdatePipelineRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *UpdatePipelineRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *UpdatePipelineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdatePipelineRequest
	GetResourceOwnerId() *int64
	SetRole(v string) *UpdatePipelineRequest
	GetRole() *string
	SetState(v string) *UpdatePipelineRequest
	GetState() *string
}

type UpdatePipelineRequest struct {
	ExtendConfig *string `json:"ExtendConfig,omitempty" xml:"ExtendConfig,omitempty"`
	// The new name of the MPS queue. The value can contain letters, digits, and special characters such as hyphens (-) and can be up to 128 bytes in size. The value cannot start with a special character.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-pipeline-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Message Service (MNS) configuration, such as the information about the MNS queue or topic. For more information, see the "NotifyConfig" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// example:
	//
	// {"Topic":"example-topic-****"}
	NotifyConfig *string `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue that you want to update. To view the MPS queue ID, log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role that is assigned to the current RAM user. To obtain the role, you can log on to the **Resource Access Management (RAM) console*	- and choose **Identities*	- > **Roles*	- in the left-side navigation pane.
	//
	// example:
	//
	// AliyunMTSDefaultRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The new state of the MPS queue.
	//
	// 	- **Active**: The MPS queue is active. Jobs in the MPS queue can be scheduled and run by MPS.
	//
	// 	- **Paused**: The MPS queue is paused. Jobs in the MPS queue cannot be scheduled or run by MPS, and all jobs remain in the Submitted state. Jobs that are running will not be affected.
	//
	// This parameter is required.
	//
	// example:
	//
	// Paused
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequest) GetExtendConfig() *string {
	return s.ExtendConfig
}

func (s *UpdatePipelineRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePipelineRequest) GetNotifyConfig() *string {
	return s.NotifyConfig
}

func (s *UpdatePipelineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdatePipelineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdatePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *UpdatePipelineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdatePipelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdatePipelineRequest) GetRole() *string {
	return s.Role
}

func (s *UpdatePipelineRequest) GetState() *string {
	return s.State
}

func (s *UpdatePipelineRequest) SetExtendConfig(v string) *UpdatePipelineRequest {
	s.ExtendConfig = &v
	return s
}

func (s *UpdatePipelineRequest) SetName(v string) *UpdatePipelineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePipelineRequest) SetNotifyConfig(v string) *UpdatePipelineRequest {
	s.NotifyConfig = &v
	return s
}

func (s *UpdatePipelineRequest) SetOwnerAccount(v string) *UpdatePipelineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdatePipelineRequest) SetOwnerId(v int64) *UpdatePipelineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePipelineRequest) SetPipelineId(v string) *UpdatePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineRequest) SetResourceOwnerAccount(v string) *UpdatePipelineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePipelineRequest) SetResourceOwnerId(v int64) *UpdatePipelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePipelineRequest) SetRole(v string) *UpdatePipelineRequest {
	s.Role = &v
	return s
}

func (s *UpdatePipelineRequest) SetState(v string) *UpdatePipelineRequest {
	s.State = &v
	return s
}

func (s *UpdatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
