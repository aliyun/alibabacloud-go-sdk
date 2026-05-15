// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *QueryAiVoiceAgentDetailRequest
	GetAgentId() *string
	SetOwnerId(v int64) *QueryAiVoiceAgentDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryAiVoiceAgentDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAiVoiceAgentDetailRequest
	GetResourceOwnerId() *int64
}

type QueryAiVoiceAgentDetailRequest struct {
	// example:
	//
	// 123123********
	AgentId              *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryAiVoiceAgentDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *QueryAiVoiceAgentDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAiVoiceAgentDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAiVoiceAgentDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAiVoiceAgentDetailRequest) SetAgentId(v string) *QueryAiVoiceAgentDetailRequest {
	s.AgentId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailRequest) SetOwnerId(v int64) *QueryAiVoiceAgentDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailRequest) SetResourceOwnerAccount(v string) *QueryAiVoiceAgentDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAiVoiceAgentDetailRequest) SetResourceOwnerId(v int64) *QueryAiVoiceAgentDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAiVoiceAgentDetailRequest) Validate() error {
	return dara.Validate(s)
}
