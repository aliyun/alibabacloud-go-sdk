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
	SetCreateFromFlowCode(v string) *CreateChatFlowRequest
	GetCreateFromFlowCode() *string
	SetCreateFromFlowVersion(v string) *CreateChatFlowRequest
	GetCreateFromFlowVersion() *string
	SetFlowTriggerType(v string) *CreateChatFlowRequest
	GetFlowTriggerType() *string
	SetLifeCycleExtendData(v map[string]*string) *CreateChatFlowRequest
	GetLifeCycleExtendData() map[string]*string
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
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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
	LifeCycleExtendData map[string]*string `json:"LifeCycleExtendData,omitempty" xml:"LifeCycleExtendData,omitempty"`
	OwnerId             *int64             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *CreateChatFlowRequest) GetCreateFromFlowCode() *string {
	return s.CreateFromFlowCode
}

func (s *CreateChatFlowRequest) GetCreateFromFlowVersion() *string {
	return s.CreateFromFlowVersion
}

func (s *CreateChatFlowRequest) GetFlowTriggerType() *string {
	return s.FlowTriggerType
}

func (s *CreateChatFlowRequest) GetLifeCycleExtendData() map[string]*string {
	return s.LifeCycleExtendData
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

func (s *CreateChatFlowRequest) SetCreateFromFlowCode(v string) *CreateChatFlowRequest {
	s.CreateFromFlowCode = &v
	return s
}

func (s *CreateChatFlowRequest) SetCreateFromFlowVersion(v string) *CreateChatFlowRequest {
	s.CreateFromFlowVersion = &v
	return s
}

func (s *CreateChatFlowRequest) SetFlowTriggerType(v string) *CreateChatFlowRequest {
	s.FlowTriggerType = &v
	return s
}

func (s *CreateChatFlowRequest) SetLifeCycleExtendData(v map[string]*string) *CreateChatFlowRequest {
	s.LifeCycleExtendData = v
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
