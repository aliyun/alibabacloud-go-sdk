// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateTagGroupRequest
	GetAgentKey() *string
	SetClientToken(v string) *UpdateTagGroupRequest
	GetClientToken() *string
	SetGroupName(v string) *UpdateTagGroupRequest
	GetGroupName() *string
	SetId(v int64) *UpdateTagGroupRequest
	GetId() *int64
}

type UpdateTagGroupRequest struct {
	// example:
	//
	// xxx
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
	// This parameter is required.
	//
	// example:
	//
	// 839232
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateTagGroupRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateTagGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTagGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateTagGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateTagGroupRequest) SetAgentKey(v string) *UpdateTagGroupRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateTagGroupRequest) SetClientToken(v string) *UpdateTagGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTagGroupRequest) SetGroupName(v string) *UpdateTagGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateTagGroupRequest) SetId(v int64) *UpdateTagGroupRequest {
	s.Id = &v
	return s
}

func (s *UpdateTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
