// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateTagGroupRequest
	GetAgentKey() *string
	SetClientToken(v string) *CreateTagGroupRequest
	GetClientToken() *string
	SetGroupName(v string) *CreateTagGroupRequest
	GetGroupName() *string
}

type CreateTagGroupRequest struct {
	// example:
	//
	// xxxxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标签组1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTagGroupRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateTagGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTagGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateTagGroupRequest) SetAgentKey(v string) *CreateTagGroupRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateTagGroupRequest) SetClientToken(v string) *CreateTagGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTagGroupRequest) SetGroupName(v string) *CreateTagGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
