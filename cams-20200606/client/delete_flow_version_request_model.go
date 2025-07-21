// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DeleteFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *DeleteFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *DeleteFlowVersionRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *DeleteFlowVersionRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *DeleteFlowVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFlowVersionRequest
	GetResourceOwnerId() *int64
}

type DeleteFlowVersionRequest struct {
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
}

func (s DeleteFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DeleteFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *DeleteFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *DeleteFlowVersionRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *DeleteFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFlowVersionRequest) SetBizCode(v string) *DeleteFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *DeleteFlowVersionRequest) SetBizExtend(v map[string]interface{}) *DeleteFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *DeleteFlowVersionRequest) SetFlowCode(v string) *DeleteFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *DeleteFlowVersionRequest) SetFlowVersion(v string) *DeleteFlowVersionRequest {
	s.FlowVersion = &v
	return s
}

func (s *DeleteFlowVersionRequest) SetOwnerId(v int64) *DeleteFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowVersionRequest) SetResourceOwnerAccount(v string) *DeleteFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFlowVersionRequest) SetResourceOwnerId(v int64) *DeleteFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
