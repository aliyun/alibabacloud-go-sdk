// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *UpdateChatFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *UpdateChatFlowShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *UpdateChatFlowShrinkRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *UpdateChatFlowShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateChatFlowShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateChatFlowShrinkRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *UpdateChatFlowShrinkRequest
	GetTitle() *string
}

type UpdateChatFlowShrinkRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information, default is “{}”.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// Process code.
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf8c70
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Process remarks
	//
	// example:
	//
	// This is Customer Service WhatsApp Auto-Reply Flow
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Process title
	//
	// example:
	//
	// Customer Service WhatsApp Auto-Reply Flow
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateChatFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateChatFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *UpdateChatFlowShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *UpdateChatFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateChatFlowShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateChatFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateChatFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateChatFlowShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateChatFlowShrinkRequest) SetBizCode(v string) *UpdateChatFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetBizExtendShrink(v string) *UpdateChatFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetFlowCode(v string) *UpdateChatFlowShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetOwnerId(v int64) *UpdateChatFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetRemark(v string) *UpdateChatFlowShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetResourceOwnerAccount(v string) *UpdateChatFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetResourceOwnerId(v int64) *UpdateChatFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) SetTitle(v string) *UpdateChatFlowShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
