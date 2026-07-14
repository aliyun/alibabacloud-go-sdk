// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *CreateFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *CreateFlowVersionShrinkRequest
	GetFlowCode() *string
	SetFlowVersionCopyFrom(v string) *CreateFlowVersionShrinkRequest
	GetFlowVersionCopyFrom() *string
	SetOwnerId(v int64) *CreateFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateFlowVersionShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
}

type CreateFlowVersionShrinkRequest struct {
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
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The code of the flow. You can view the flow code in the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder).
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The version of the flow to copy. To view the historical versions of a flow, go to the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) and click a flow name to open the orchestration canvas.
	//
	// example:
	//
	// 1
	FlowVersionCopyFrom *string `json:"FlowVersionCopyFrom,omitempty" xml:"FlowVersionCopyFrom,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The remarks for the version.
	//
	// example:
	//
	// 修复发送WhatsApp消息错误
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *CreateFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *CreateFlowVersionShrinkRequest) GetFlowVersionCopyFrom() *string {
	return s.FlowVersionCopyFrom
}

func (s *CreateFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowVersionShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowVersionShrinkRequest) SetBizCode(v string) *CreateFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetBizExtendShrink(v string) *CreateFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetFlowCode(v string) *CreateFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetFlowVersionCopyFrom(v string) *CreateFlowVersionShrinkRequest {
	s.FlowVersionCopyFrom = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetOwnerId(v int64) *CreateFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetRemark(v string) *CreateFlowVersionShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *CreateFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *CreateFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
