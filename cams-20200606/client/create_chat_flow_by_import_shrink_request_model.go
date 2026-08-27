// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowByImportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateChatFlowByImportShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *CreateChatFlowByImportShrinkRequest
	GetBizExtendShrink() *string
	SetFlowViewModel(v string) *CreateChatFlowByImportShrinkRequest
	GetFlowViewModel() *string
	SetOwnerId(v int64) *CreateChatFlowByImportShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateChatFlowByImportShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateChatFlowByImportShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatFlowByImportShrinkRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *CreateChatFlowByImportShrinkRequest
	GetTitle() *string
}

type CreateChatFlowByImportShrinkRequest struct {
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
	// The imported flow DSL data, which is a JSON-formatted string. You can arrange flow components on the canvas in the flow orchestration console in advance, save the flow, and then click **Settings*	- > **Export*	- in the upper-right corner of the canvas to export a JSON data file for viewing.
	//
	// example:
	//
	// {
	//
	//   "schema": {
	//
	//     "namespace": "External",
	//
	//     "version": "1.0",
	//
	//     "copyright": "Alibaba Cloud"
	//
	//   },
	//
	//   "editor": "H4sIAAAAAAAAA+1YbU/c***********************",
	//
	//   "flow": {
	//
	//     "triggerType": "TriggeredByWhatsApp"
	//
	//   }
	//
	// }
	FlowViewModel *string `json:"FlowViewModel,omitempty" xml:"FlowViewModel,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The flow remarks.
	//
	// example:
	//
	// Trigger Subscription.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The flow title.
	//
	// example:
	//
	// WhatsApp Trigger Subscription.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateChatFlowByImportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowByImportShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatFlowByImportShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateChatFlowByImportShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *CreateChatFlowByImportShrinkRequest) GetFlowViewModel() *string {
	return s.FlowViewModel
}

func (s *CreateChatFlowByImportShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatFlowByImportShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateChatFlowByImportShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatFlowByImportShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatFlowByImportShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChatFlowByImportShrinkRequest) SetBizCode(v string) *CreateChatFlowByImportShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetBizExtendShrink(v string) *CreateChatFlowByImportShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetFlowViewModel(v string) *CreateChatFlowByImportShrinkRequest {
	s.FlowViewModel = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetOwnerId(v int64) *CreateChatFlowByImportShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetRemark(v string) *CreateChatFlowByImportShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetResourceOwnerAccount(v string) *CreateChatFlowByImportShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetResourceOwnerId(v int64) *CreateChatFlowByImportShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetTitle(v string) *CreateChatFlowByImportShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
