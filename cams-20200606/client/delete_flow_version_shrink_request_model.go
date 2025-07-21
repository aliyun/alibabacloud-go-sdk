// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DeleteFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *DeleteFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *DeleteFlowVersionShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *DeleteFlowVersionShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *DeleteFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
}

type DeleteFlowVersionShrinkRequest struct {
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
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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

func (s DeleteFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DeleteFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *DeleteFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *DeleteFlowVersionShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *DeleteFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFlowVersionShrinkRequest) SetBizCode(v string) *DeleteFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) SetBizExtendShrink(v string) *DeleteFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) SetFlowCode(v string) *DeleteFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) SetFlowVersion(v string) *DeleteFlowVersionShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) SetOwnerId(v int64) *DeleteFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *DeleteFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *DeleteFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
