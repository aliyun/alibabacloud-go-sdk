// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCancelCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentCancelCallRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentCancelCallRequest
	GetAgentTag() *string
	SetNumbers(v []*string) *AgentCancelCallRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *AgentCancelCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AgentCancelCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentCancelCallRequest
	GetResourceOwnerId() *int64
	SetTags(v []*string) *AgentCancelCallRequest
	GetTags() []*string
}

type AgentCancelCallRequest struct {
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
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AgentCancelCallRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentCancelCallRequest) GoString() string {
	return s.String()
}

func (s *AgentCancelCallRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentCancelCallRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentCancelCallRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *AgentCancelCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentCancelCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentCancelCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentCancelCallRequest) GetTags() []*string {
	return s.Tags
}

func (s *AgentCancelCallRequest) SetAgentId(v int64) *AgentCancelCallRequest {
	s.AgentId = &v
	return s
}

func (s *AgentCancelCallRequest) SetAgentTag(v string) *AgentCancelCallRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentCancelCallRequest) SetNumbers(v []*string) *AgentCancelCallRequest {
	s.Numbers = v
	return s
}

func (s *AgentCancelCallRequest) SetOwnerId(v int64) *AgentCancelCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentCancelCallRequest) SetResourceOwnerAccount(v string) *AgentCancelCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentCancelCallRequest) SetResourceOwnerId(v int64) *AgentCancelCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentCancelCallRequest) SetTags(v []*string) *AgentCancelCallRequest {
	s.Tags = v
	return s
}

func (s *AgentCancelCallRequest) Validate() error {
	return dara.Validate(s)
}
