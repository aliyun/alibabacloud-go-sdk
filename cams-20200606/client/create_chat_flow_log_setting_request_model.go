// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowLogSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowCode(v string) *CreateChatFlowLogSettingRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *CreateChatFlowLogSettingRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateChatFlowLogSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatFlowLogSettingRequest
	GetResourceOwnerId() *int64
}

type CreateChatFlowLogSettingRequest struct {
	// Process code.
	//
	// example:
	//
	// 示例值
	FlowCode             *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateChatFlowLogSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowLogSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateChatFlowLogSettingRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *CreateChatFlowLogSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatFlowLogSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatFlowLogSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatFlowLogSettingRequest) SetFlowCode(v string) *CreateChatFlowLogSettingRequest {
	s.FlowCode = &v
	return s
}

func (s *CreateChatFlowLogSettingRequest) SetOwnerId(v int64) *CreateChatFlowLogSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatFlowLogSettingRequest) SetResourceOwnerAccount(v string) *CreateChatFlowLogSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatFlowLogSettingRequest) SetResourceOwnerId(v int64) *CreateChatFlowLogSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatFlowLogSettingRequest) Validate() error {
	return dara.Validate(s)
}
