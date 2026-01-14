// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCancelCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentCancelCallShrinkRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentCancelCallShrinkRequest
	GetAgentTag() *string
	SetNumbersShrink(v string) *AgentCancelCallShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *AgentCancelCallShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AgentCancelCallShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentCancelCallShrinkRequest
	GetResourceOwnerId() *int64
	SetTagsShrink(v string) *AgentCancelCallShrinkRequest
	GetTagsShrink() *string
}

type AgentCancelCallShrinkRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 64
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 号码列表
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AgentCancelCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentCancelCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgentCancelCallShrinkRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentCancelCallShrinkRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentCancelCallShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *AgentCancelCallShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentCancelCallShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentCancelCallShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentCancelCallShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AgentCancelCallShrinkRequest) SetAgentId(v int64) *AgentCancelCallShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetAgentTag(v string) *AgentCancelCallShrinkRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetNumbersShrink(v string) *AgentCancelCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetOwnerId(v int64) *AgentCancelCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetResourceOwnerAccount(v string) *AgentCancelCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetResourceOwnerId(v int64) *AgentCancelCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetTagsShrink(v string) *AgentCancelCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
