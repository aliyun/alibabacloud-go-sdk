// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ReadChatFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *ReadChatFlowShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *ReadChatFlowShrinkRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *ReadChatFlowShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReadChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReadChatFlowShrinkRequest
	GetResourceOwnerId() *int64
}

type ReadChatFlowShrinkRequest struct {
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
	// The flow code. View the flow code in the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder).
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode             *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReadChatFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadChatFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ReadChatFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *ReadChatFlowShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ReadChatFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReadChatFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReadChatFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReadChatFlowShrinkRequest) SetBizCode(v string) *ReadChatFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ReadChatFlowShrinkRequest) SetBizExtendShrink(v string) *ReadChatFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *ReadChatFlowShrinkRequest) SetFlowCode(v string) *ReadChatFlowShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *ReadChatFlowShrinkRequest) SetOwnerId(v int64) *ReadChatFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ReadChatFlowShrinkRequest) SetResourceOwnerAccount(v string) *ReadChatFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReadChatFlowShrinkRequest) SetResourceOwnerId(v int64) *ReadChatFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReadChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
