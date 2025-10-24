// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateTagRequest
	GetAgentKey() *string
	SetClientToken(v string) *CreateTagRequest
	GetClientToken() *string
	SetGroupId(v int64) *CreateTagRequest
	GetGroupId() *int64
	SetTagName(v string) *CreateTagRequest
	GetTagName() *string
}

type CreateTagRequest struct {
	// example:
	//
	// xxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4132342
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标签组1
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateTagRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTagRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *CreateTagRequest) SetAgentKey(v string) *CreateTagRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateTagRequest) SetClientToken(v string) *CreateTagRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTagRequest) SetGroupId(v int64) *CreateTagRequest {
	s.GroupId = &v
	return s
}

func (s *CreateTagRequest) SetTagName(v string) *CreateTagRequest {
	s.TagName = &v
	return s
}

func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
