// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowLogSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowCode(v string) *ReadChatFlowLogSettingRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *ReadChatFlowLogSettingRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReadChatFlowLogSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReadChatFlowLogSettingRequest
	GetResourceOwnerId() *int64
}

type ReadChatFlowLogSettingRequest struct {
	// Process code.
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf****
	FlowCode             *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReadChatFlowLogSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowLogSettingRequest) GoString() string {
	return s.String()
}

func (s *ReadChatFlowLogSettingRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ReadChatFlowLogSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReadChatFlowLogSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReadChatFlowLogSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReadChatFlowLogSettingRequest) SetFlowCode(v string) *ReadChatFlowLogSettingRequest {
	s.FlowCode = &v
	return s
}

func (s *ReadChatFlowLogSettingRequest) SetOwnerId(v int64) *ReadChatFlowLogSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *ReadChatFlowLogSettingRequest) SetResourceOwnerAccount(v string) *ReadChatFlowLogSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReadChatFlowLogSettingRequest) SetResourceOwnerId(v int64) *ReadChatFlowLogSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReadChatFlowLogSettingRequest) Validate() error {
	return dara.Validate(s)
}
