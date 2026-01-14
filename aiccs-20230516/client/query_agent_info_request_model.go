// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgentInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *QueryAgentInfoRequest
	GetAgentId() *int64
	SetAgentTag(v string) *QueryAgentInfoRequest
	GetAgentTag() *string
	SetOwnerId(v int64) *QueryAgentInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryAgentInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAgentInfoRequest
	GetResourceOwnerId() *int64
}

type QueryAgentInfoRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// "1"
	AgentTag             *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryAgentInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAgentInfoRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *QueryAgentInfoRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *QueryAgentInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAgentInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAgentInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAgentInfoRequest) SetAgentId(v int64) *QueryAgentInfoRequest {
	s.AgentId = &v
	return s
}

func (s *QueryAgentInfoRequest) SetAgentTag(v string) *QueryAgentInfoRequest {
	s.AgentTag = &v
	return s
}

func (s *QueryAgentInfoRequest) SetOwnerId(v int64) *QueryAgentInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAgentInfoRequest) SetResourceOwnerAccount(v string) *QueryAgentInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAgentInfoRequest) SetResourceOwnerId(v int64) *QueryAgentInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAgentInfoRequest) Validate() error {
	return dara.Validate(s)
}
