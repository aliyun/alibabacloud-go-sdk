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
}

type UpdateFlowVersionRequest struct {
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

func (s *UpdateFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
