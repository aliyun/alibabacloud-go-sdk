// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowByImportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateChatFlowByImportRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *CreateChatFlowByImportRequest
	GetBizExtend() map[string]interface{}
	SetFlowViewModel(v string) *CreateChatFlowByImportRequest
	GetFlowViewModel() *string
	SetOwnerId(v int64) *CreateChatFlowByImportRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateChatFlowByImportRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateChatFlowByImportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatFlowByImportRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *CreateChatFlowByImportRequest
	GetTitle() *string
}

type CreateChatFlowByImportRequest struct {
	// The business tenant code. The default value is ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The business extension information. The default value is an empty collection.
	//
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The flow DSL data to import. This is a block of data in JSON format. To obtain this data, arrange the components on the canvas in the Flow Editor, save the flow, and then click **Settings*	- > **Export*	- in the upper-right corner of the canvas. The flow is exported as a JSON data file.
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
	// The remarks for the flow.
	//
	// example:
	//
	// 触发订阅
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The title of the flow.
	//
	// example:
	//
	// WhatsApp触发订阅
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateChatFlowByImportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowByImportRequest) GoString() string {
	return s.String()
}

func (s *CreateChatFlowByImportRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateChatFlowByImportRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *CreateChatFlowByImportRequest) GetFlowViewModel() *string {
	return s.FlowViewModel
}

func (s *CreateChatFlowByImportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatFlowByImportRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateChatFlowByImportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatFlowByImportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatFlowByImportRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChatFlowByImportRequest) SetBizCode(v string) *CreateChatFlowByImportRequest {
	s.BizCode = &v
	return s
}

func (s *CreateChatFlowByImportRequest) SetBizExtend(v map[string]interface{}) *CreateChatFlowByImportRequest {
	s.BizExtend = v
	return s
}

func (s *CreateChatFlowByImportRequest) SetFlowViewModel(v string) *CreateChatFlowByImportRequest {
	s.FlowViewModel = &v
	return s
}

func (s *CreateChatFlowByImportRequest) SetOwnerId(v int64) *CreateChatFlowByImportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatFlowByImportRequest) SetRemark(v string) *CreateChatFlowByImportRequest {
	s.Remark = &v
	return s
}

func (s *CreateChatFlowByImportRequest) SetResourceOwnerAccount(v string) *CreateChatFlowByImportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatFlowByImportRequest) SetResourceOwnerId(v int64) *CreateChatFlowByImportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatFlowByImportRequest) SetTitle(v string) *CreateChatFlowByImportRequest {
	s.Title = &v
	return s
}

func (s *CreateChatFlowByImportRequest) Validate() error {
	return dara.Validate(s)
}
