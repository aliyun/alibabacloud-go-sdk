// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ReadFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *ReadFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *ReadFlowVersionRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *ReadFlowVersionRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *ReadFlowVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReadFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReadFlowVersionRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ReadFlowVersionRequest
	GetStatus() *string
}

type ReadFlowVersionRequest struct {
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
	FlowVersion          *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Flow version status.
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReadFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *ReadFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ReadFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ReadFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ReadFlowVersionRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *ReadFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReadFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReadFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReadFlowVersionRequest) GetStatus() *string {
	return s.Status
}

func (s *ReadFlowVersionRequest) SetBizCode(v string) *ReadFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *ReadFlowVersionRequest) SetBizExtend(v map[string]interface{}) *ReadFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *ReadFlowVersionRequest) SetFlowCode(v string) *ReadFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *ReadFlowVersionRequest) SetFlowVersion(v string) *ReadFlowVersionRequest {
	s.FlowVersion = &v
	return s
}

func (s *ReadFlowVersionRequest) SetOwnerId(v int64) *ReadFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReadFlowVersionRequest) SetResourceOwnerAccount(v string) *ReadFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReadFlowVersionRequest) SetResourceOwnerId(v int64) *ReadFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReadFlowVersionRequest) SetStatus(v string) *ReadFlowVersionRequest {
	s.Status = &v
	return s
}

func (s *ReadFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
