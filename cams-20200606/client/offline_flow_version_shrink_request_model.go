// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *OfflineFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *OfflineFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *OfflineFlowVersionShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *OfflineFlowVersionShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *OfflineFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *OfflineFlowVersionShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *OfflineFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OfflineFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
}

type OfflineFlowVersionShrinkRequest struct {
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
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Flow remarks
	//
	// example:
	//
	// We don\\"t need this old version.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OfflineFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OfflineFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *OfflineFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *OfflineFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *OfflineFlowVersionShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *OfflineFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OfflineFlowVersionShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *OfflineFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OfflineFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OfflineFlowVersionShrinkRequest) SetBizCode(v string) *OfflineFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetBizExtendShrink(v string) *OfflineFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetFlowCode(v string) *OfflineFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetFlowVersion(v string) *OfflineFlowVersionShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetOwnerId(v int64) *OfflineFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetRemark(v string) *OfflineFlowVersionShrinkRequest {
	s.Remark = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *OfflineFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *OfflineFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OfflineFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
