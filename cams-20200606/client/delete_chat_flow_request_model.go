// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DeleteChatFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *DeleteChatFlowRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *DeleteChatFlowRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *DeleteChatFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatFlowRequest
	GetResourceOwnerId() *int64
}

type DeleteChatFlowRequest struct {
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

func (s DeleteChatFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DeleteChatFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *DeleteChatFlowRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *DeleteChatFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatFlowRequest) SetBizCode(v string) *DeleteChatFlowRequest {
	s.BizCode = &v
	return s
}

func (s *DeleteChatFlowRequest) SetBizExtend(v map[string]interface{}) *DeleteChatFlowRequest {
	s.BizExtend = v
	return s
}

func (s *DeleteChatFlowRequest) SetFlowCode(v string) *DeleteChatFlowRequest {
	s.FlowCode = &v
	return s
}

func (s *DeleteChatFlowRequest) SetOwnerId(v int64) *DeleteChatFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatFlowRequest) SetResourceOwnerAccount(v string) *DeleteChatFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatFlowRequest) SetResourceOwnerId(v int64) *DeleteChatFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatFlowRequest) Validate() error {
	return dara.Validate(s)
}
