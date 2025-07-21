// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *CreateFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *CreateFlowVersionRequest
	GetFlowCode() *string
	SetFlowVersionCopyFrom(v string) *CreateFlowVersionRequest
	GetFlowVersionCopyFrom() *string
	SetOwnerId(v int64) *CreateFlowVersionRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateFlowVersionRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowVersionRequest
	GetResourceOwnerId() *int64
}

type CreateFlowVersionRequest struct {
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
	// The flow version to be copied.
	//
	// example:
	//
	// 1
	FlowVersionCopyFrom *string `json:"FlowVersionCopyFrom,omitempty" xml:"FlowVersionCopyFrom,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Version remarks.
	//
	// example:
	//
	// Fix Send WhatsApp Message Error
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *CreateFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *CreateFlowVersionRequest) GetFlowVersionCopyFrom() *string {
	return s.FlowVersionCopyFrom
}

func (s *CreateFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowVersionRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowVersionRequest) SetBizCode(v string) *CreateFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *CreateFlowVersionRequest) SetBizExtend(v map[string]interface{}) *CreateFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *CreateFlowVersionRequest) SetFlowCode(v string) *CreateFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *CreateFlowVersionRequest) SetFlowVersionCopyFrom(v string) *CreateFlowVersionRequest {
	s.FlowVersionCopyFrom = &v
	return s
}

func (s *CreateFlowVersionRequest) SetOwnerId(v int64) *CreateFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowVersionRequest) SetRemark(v string) *CreateFlowVersionRequest {
	s.Remark = &v
	return s
}

func (s *CreateFlowVersionRequest) SetResourceOwnerAccount(v string) *CreateFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowVersionRequest) SetResourceOwnerId(v int64) *CreateFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
