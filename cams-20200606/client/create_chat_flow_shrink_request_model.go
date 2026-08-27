// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateChatFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *CreateChatFlowShrinkRequest
	GetBizExtendShrink() *string
	SetCreateFromFlowCode(v string) *CreateChatFlowShrinkRequest
	GetCreateFromFlowCode() *string
	SetCreateFromFlowVersion(v string) *CreateChatFlowShrinkRequest
	GetCreateFromFlowVersion() *string
	SetFlowTriggerType(v string) *CreateChatFlowShrinkRequest
	GetFlowTriggerType() *string
	SetLifeCycleExtendDataShrink(v string) *CreateChatFlowShrinkRequest
	GetLifeCycleExtendDataShrink() *string
	SetOwnerId(v int64) *CreateChatFlowShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateChatFlowShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatFlowShrinkRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *CreateChatFlowShrinkRequest
	GetTitle() *string
}

type CreateChatFlowShrinkRequest struct {
	// The business tenant code. Default value: ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The business extension information. Default value: an empty collection.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The source flowCode for creation.
	//
	// example:
	//
	// 示例值
	CreateFromFlowCode *string `json:"CreateFromFlowCode,omitempty" xml:"CreateFromFlowCode,omitempty"`
	// The source flowVersion for creation.
	//
	// example:
	//
	// 示例值示例值示例值
	CreateFromFlowVersion *string `json:"CreateFromFlowVersion,omitempty" xml:"CreateFromFlowVersion,omitempty"`
	// The flow trigger type. Valid values:
	//
	//  - TriggeredManually
	//
	// - TriggeredByWhatsApp
	//
	// - TriggeredByMessenger
	//
	// - TriggeredByInstagram
	//
	// - TriggeredByViber
	//
	// example:
	//
	// TriggeredByWhatsApp
	FlowTriggerType *string `json:"FlowTriggerType,omitempty" xml:"FlowTriggerType,omitempty"`
	// The lifecycle extension input parameters.
	LifeCycleExtendDataShrink *string `json:"LifeCycleExtendData,omitempty" xml:"LifeCycleExtendData,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The flow remarks.
	//
	// example:
	//
	// Send verification template triggered by API
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The flow title.
	//
	// example:
	//
	// WhatsApp auto-reply
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateChatFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateChatFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *CreateChatFlowShrinkRequest) GetCreateFromFlowCode() *string {
	return s.CreateFromFlowCode
}

func (s *CreateChatFlowShrinkRequest) GetCreateFromFlowVersion() *string {
	return s.CreateFromFlowVersion
}

func (s *CreateChatFlowShrinkRequest) GetFlowTriggerType() *string {
	return s.FlowTriggerType
}

func (s *CreateChatFlowShrinkRequest) GetLifeCycleExtendDataShrink() *string {
	return s.LifeCycleExtendDataShrink
}

func (s *CreateChatFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatFlowShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateChatFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatFlowShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChatFlowShrinkRequest) SetBizCode(v string) *CreateChatFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetBizExtendShrink(v string) *CreateChatFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetCreateFromFlowCode(v string) *CreateChatFlowShrinkRequest {
	s.CreateFromFlowCode = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetCreateFromFlowVersion(v string) *CreateChatFlowShrinkRequest {
	s.CreateFromFlowVersion = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetFlowTriggerType(v string) *CreateChatFlowShrinkRequest {
	s.FlowTriggerType = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetLifeCycleExtendDataShrink(v string) *CreateChatFlowShrinkRequest {
	s.LifeCycleExtendDataShrink = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetOwnerId(v int64) *CreateChatFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetRemark(v string) *CreateChatFlowShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetResourceOwnerAccount(v string) *CreateChatFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetResourceOwnerId(v int64) *CreateChatFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) SetTitle(v string) *CreateChatFlowShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
