// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *OnlineFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *OnlineFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *OnlineFlowVersionShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *OnlineFlowVersionShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *OnlineFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *OnlineFlowVersionShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *OnlineFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OnlineFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
}

type OnlineFlowVersionShrinkRequest struct {
	// The business tenant code. The default value is ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information. The default value is an empty object.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The flow code. View the flow code in the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder).
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf****
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. In the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder), click the flow name to open the orchestration canvas and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The remarks for the flow. You can view the remarks in the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder).
	//
	// example:
	//
	// 通过API触发下发验证模板
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OnlineFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OnlineFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *OnlineFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *OnlineFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *OnlineFlowVersionShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *OnlineFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OnlineFlowVersionShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *OnlineFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OnlineFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OnlineFlowVersionShrinkRequest) SetBizCode(v string) *OnlineFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetBizExtendShrink(v string) *OnlineFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetFlowCode(v string) *OnlineFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetFlowVersion(v string) *OnlineFlowVersionShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetOwnerId(v int64) *OnlineFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetRemark(v string) *OnlineFlowVersionShrinkRequest {
	s.Remark = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *OnlineFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *OnlineFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OnlineFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
