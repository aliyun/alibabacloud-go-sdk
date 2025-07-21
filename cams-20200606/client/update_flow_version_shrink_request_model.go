// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *UpdateFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *UpdateFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *UpdateFlowVersionShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *UpdateFlowVersionShrinkRequest
	GetFlowVersion() *string
	SetFlowViewModel(v string) *UpdateFlowVersionShrinkRequest
	GetFlowViewModel() *string
	SetOwnerId(v int64) *UpdateFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateFlowVersionShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
}

type UpdateFlowVersionShrinkRequest struct {
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
	// Flow code.
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf****
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// Flow version
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// DSL data of the flow version
	//
	// example:
	//
	// {}
	FlowViewModel *string `json:"FlowViewModel,omitempty" xml:"FlowViewModel,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Version remarks
	//
	// example:
	//
	// Fix Send WhatsApp Message Error
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *UpdateFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *UpdateFlowVersionShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *UpdateFlowVersionShrinkRequest) GetFlowViewModel() *string {
	return s.FlowViewModel
}

func (s *UpdateFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateFlowVersionShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateFlowVersionShrinkRequest) SetBizCode(v string) *UpdateFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetBizExtendShrink(v string) *UpdateFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetFlowCode(v string) *UpdateFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetFlowVersion(v string) *UpdateFlowVersionShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetFlowViewModel(v string) *UpdateFlowVersionShrinkRequest {
	s.FlowViewModel = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetOwnerId(v int64) *UpdateFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetRemark(v string) *UpdateFlowVersionShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *UpdateFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *UpdateFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
