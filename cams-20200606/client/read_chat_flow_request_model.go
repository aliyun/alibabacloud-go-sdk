// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadChatFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ReadChatFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *ReadChatFlowRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *ReadChatFlowRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *ReadChatFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReadChatFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReadChatFlowRequest
	GetResourceOwnerId() *int64
}

type ReadChatFlowRequest struct {
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
	// 示例值
	FlowCode             *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReadChatFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadChatFlowRequest) GoString() string {
	return s.String()
}

func (s *ReadChatFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ReadChatFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ReadChatFlowRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ReadChatFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReadChatFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReadChatFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReadChatFlowRequest) SetBizCode(v string) *ReadChatFlowRequest {
	s.BizCode = &v
	return s
}

func (s *ReadChatFlowRequest) SetBizExtend(v map[string]interface{}) *ReadChatFlowRequest {
	s.BizExtend = v
	return s
}

func (s *ReadChatFlowRequest) SetFlowCode(v string) *ReadChatFlowRequest {
	s.FlowCode = &v
	return s
}

func (s *ReadChatFlowRequest) SetOwnerId(v int64) *ReadChatFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *ReadChatFlowRequest) SetResourceOwnerAccount(v string) *ReadChatFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReadChatFlowRequest) SetResourceOwnerId(v int64) *ReadChatFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReadChatFlowRequest) Validate() error {
	return dara.Validate(s)
}
