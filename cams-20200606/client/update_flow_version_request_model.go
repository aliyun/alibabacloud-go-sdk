// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *UpdateFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *UpdateFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *UpdateFlowVersionRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *UpdateFlowVersionRequest
	GetFlowVersion() *string
	SetFlowViewModel(v string) *UpdateFlowVersionRequest
	GetFlowViewModel() *string
	SetOwnerId(v int64) *UpdateFlowVersionRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateFlowVersionRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateFlowVersionRequest
	GetResourceOwnerId() *int64
	SetType(v string) *UpdateFlowVersionRequest
	GetType() *string
}

type UpdateFlowVersionRequest struct {
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
	// The flow code. You can view the flow code on the [flow editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. You can click a flow name on the [flow editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page to go to the flow editor canvas page and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The DSL data of the flow version. This is a JSON-formatted data string. You can orchestrate flow components on the flow editor canvas in advance, save the flow, and then click **Settings*	- > **Export*	- in the upper-right corner of the canvas orchestration page to export a JSON-formatted data file for viewing.
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
	// The version remarks.
	//
	// example:
	//
	// Fix WhatsApp message sending error
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The save type.
	//
	// example:
	//
	// Sample value
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *UpdateFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *UpdateFlowVersionRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *UpdateFlowVersionRequest) GetFlowViewModel() *string {
	return s.FlowViewModel
}

func (s *UpdateFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateFlowVersionRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateFlowVersionRequest) GetType() *string {
	return s.Type
}

func (s *UpdateFlowVersionRequest) SetBizCode(v string) *UpdateFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetBizExtend(v map[string]interface{}) *UpdateFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *UpdateFlowVersionRequest) SetFlowCode(v string) *UpdateFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetFlowVersion(v string) *UpdateFlowVersionRequest {
	s.FlowVersion = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetFlowViewModel(v string) *UpdateFlowVersionRequest {
	s.FlowViewModel = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetOwnerId(v int64) *UpdateFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetRemark(v string) *UpdateFlowVersionRequest {
	s.Remark = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetResourceOwnerAccount(v string) *UpdateFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetResourceOwnerId(v int64) *UpdateFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateFlowVersionRequest) SetType(v string) *UpdateFlowVersionRequest {
	s.Type = &v
	return s
}

func (s *UpdateFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
