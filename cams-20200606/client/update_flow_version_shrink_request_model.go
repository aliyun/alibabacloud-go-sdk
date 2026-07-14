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
	// The tenant code. Default value: ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The extended business information. The default value is an empty collection.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The code of the flow. View the flow code in the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder).
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The version of the flow. In the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder), click the flow name to open the canvas and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The DSL data of the flow version, in JSON format. To obtain this data, orchestrate the components on the canvas in the Flow Editor. After you save the flow, click **Settings*	- > **Export*	- in the upper-right corner of the canvas to export the flow as a JSON data file.
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
	// The remarks for the version.
	//
	// example:
	//
	// 修复发送WhatsApp消息错误
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
