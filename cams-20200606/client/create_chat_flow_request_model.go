// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateChatFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *CreateChatFlowRequest
	GetBizExtend() map[string]interface{}
	SetFlowTriggerType(v string) *CreateChatFlowRequest
	GetFlowTriggerType() *string
	SetOwnerId(v int64) *CreateChatFlowRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateChatFlowRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateChatFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatFlowRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *CreateChatFlowRequest
	GetTitle() *string
}

type CreateChatFlowRequest struct {
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
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// Flow trigger type
	//
	// example:
	//
	// TriggeredByWhatsApp
	FlowTriggerType *string `json:"FlowTriggerType,omitempty" xml:"FlowTriggerType,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Flow remarks
	//
	// example:
	//
	// ChatFlow for WhatsApp Customer Service Auto-Reply.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Flow title
	//
	// example:
	//
	// Auto Reply
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateChatFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateChatFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateChatFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *CreateChatFlowRequest) GetFlowTriggerType() *string {
	return s.FlowTriggerType
}

func (s *CreateChatFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatFlowRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateChatFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatFlowRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChatFlowRequest) SetBizCode(v string) *CreateChatFlowRequest {
	s.BizCode = &v
	return s
}

func (s *CreateChatFlowRequest) SetBizExtend(v map[string]interface{}) *CreateChatFlowRequest {
	s.BizExtend = v
	return s
}

func (s *CreateChatFlowRequest) SetFlowTriggerType(v string) *CreateChatFlowRequest {
	s.FlowTriggerType = &v
	return s
}

func (s *CreateChatFlowRequest) SetOwnerId(v int64) *CreateChatFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatFlowRequest) SetRemark(v string) *CreateChatFlowRequest {
	s.Remark = &v
	return s
}

func (s *CreateChatFlowRequest) SetResourceOwnerAccount(v string) *CreateChatFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatFlowRequest) SetResourceOwnerId(v int64) *CreateChatFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatFlowRequest) SetTitle(v string) *CreateChatFlowRequest {
	s.Title = &v
	return s
}

func (s *CreateChatFlowRequest) Validate() error {
	return dara.Validate(s)
}
