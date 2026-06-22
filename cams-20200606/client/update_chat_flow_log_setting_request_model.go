// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowLogSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowCode(v string) *UpdateChatFlowLogSettingRequest
	GetFlowCode() *string
	SetId(v int64) *UpdateChatFlowLogSettingRequest
	GetId() *int64
	SetOwnerId(v int64) *UpdateChatFlowLogSettingRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateChatFlowLogSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateChatFlowLogSettingRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *UpdateChatFlowLogSettingRequest
	GetStatus() *string
}

type UpdateChatFlowLogSettingRequest struct {
	// example:
	//
	// 示例值示例值示例值
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 74
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateChatFlowLogSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowLogSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowLogSettingRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *UpdateChatFlowLogSettingRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateChatFlowLogSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateChatFlowLogSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateChatFlowLogSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateChatFlowLogSettingRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateChatFlowLogSettingRequest) SetFlowCode(v string) *UpdateChatFlowLogSettingRequest {
	s.FlowCode = &v
	return s
}

func (s *UpdateChatFlowLogSettingRequest) SetId(v int64) *UpdateChatFlowLogSettingRequest {
	s.Id = &v
	return s
}

func (s *UpdateChatFlowLogSettingRequest) SetOwnerId(v int64) *UpdateChatFlowLogSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateChatFlowLogSettingRequest) SetResourceOwnerAccount(v string) *UpdateChatFlowLogSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateChatFlowLogSettingRequest) SetResourceOwnerId(v int64) *UpdateChatFlowLogSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateChatFlowLogSettingRequest) SetStatus(v string) *UpdateChatFlowLogSettingRequest {
	s.Status = &v
	return s
}

func (s *UpdateChatFlowLogSettingRequest) Validate() error {
	return dara.Validate(s)
}
