// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *UpdateChatFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *UpdateChatFlowRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *UpdateChatFlowRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *UpdateChatFlowRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateChatFlowRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateChatFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateChatFlowRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *UpdateChatFlowRequest
	GetTitle() *string
}

type UpdateChatFlowRequest struct {
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

func (s UpdateChatFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateChatFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *UpdateChatFlowRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *UpdateChatFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateChatFlowRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateChatFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateChatFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateChatFlowRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateChatFlowRequest) SetBizCode(v string) *UpdateChatFlowRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateChatFlowRequest) SetBizExtend(v map[string]interface{}) *UpdateChatFlowRequest {
	s.BizExtend = v
	return s
}

func (s *UpdateChatFlowRequest) SetFlowCode(v string) *UpdateChatFlowRequest {
	s.FlowCode = &v
	return s
}

func (s *UpdateChatFlowRequest) SetOwnerId(v int64) *UpdateChatFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateChatFlowRequest) SetRemark(v string) *UpdateChatFlowRequest {
	s.Remark = &v
	return s
}

func (s *UpdateChatFlowRequest) SetResourceOwnerAccount(v string) *UpdateChatFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateChatFlowRequest) SetResourceOwnerId(v int64) *UpdateChatFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateChatFlowRequest) SetTitle(v string) *UpdateChatFlowRequest {
	s.Title = &v
	return s
}

func (s *UpdateChatFlowRequest) Validate() error {
	return dara.Validate(s)
}
