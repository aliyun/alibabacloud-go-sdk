// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *OnlineFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *OnlineFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *OnlineFlowVersionRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *OnlineFlowVersionRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *OnlineFlowVersionRequest
	GetOwnerId() *int64
	SetRemark(v string) *OnlineFlowVersionRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *OnlineFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OnlineFlowVersionRequest
	GetResourceOwnerId() *int64
}

type OnlineFlowVersionRequest struct {
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
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Remark
	//
	// example:
	//
	// Let\\"s go online.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OnlineFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *OnlineFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *OnlineFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *OnlineFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *OnlineFlowVersionRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *OnlineFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OnlineFlowVersionRequest) GetRemark() *string {
	return s.Remark
}

func (s *OnlineFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OnlineFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OnlineFlowVersionRequest) SetBizCode(v string) *OnlineFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *OnlineFlowVersionRequest) SetBizExtend(v map[string]interface{}) *OnlineFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *OnlineFlowVersionRequest) SetFlowCode(v string) *OnlineFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *OnlineFlowVersionRequest) SetFlowVersion(v string) *OnlineFlowVersionRequest {
	s.FlowVersion = &v
	return s
}

func (s *OnlineFlowVersionRequest) SetOwnerId(v int64) *OnlineFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *OnlineFlowVersionRequest) SetRemark(v string) *OnlineFlowVersionRequest {
	s.Remark = &v
	return s
}

func (s *OnlineFlowVersionRequest) SetResourceOwnerAccount(v string) *OnlineFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OnlineFlowVersionRequest) SetResourceOwnerId(v int64) *OnlineFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OnlineFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
