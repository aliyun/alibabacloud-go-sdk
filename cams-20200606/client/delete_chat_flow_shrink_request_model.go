// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DeleteChatFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *DeleteChatFlowShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *DeleteChatFlowShrinkRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *DeleteChatFlowShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatFlowShrinkRequest
	GetResourceOwnerId() *int64
}

type DeleteChatFlowShrinkRequest struct {
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
	// Process code.
	//
	// example:
	//
	// 示例值
	FlowCode             *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteChatFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DeleteChatFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *DeleteChatFlowShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *DeleteChatFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatFlowShrinkRequest) SetBizCode(v string) *DeleteChatFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *DeleteChatFlowShrinkRequest) SetBizExtendShrink(v string) *DeleteChatFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *DeleteChatFlowShrinkRequest) SetFlowCode(v string) *DeleteChatFlowShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *DeleteChatFlowShrinkRequest) SetOwnerId(v int64) *DeleteChatFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatFlowShrinkRequest) SetResourceOwnerAccount(v string) *DeleteChatFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatFlowShrinkRequest) SetResourceOwnerId(v int64) *DeleteChatFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
