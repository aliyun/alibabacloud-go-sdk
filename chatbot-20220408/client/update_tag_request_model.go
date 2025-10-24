// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateTagRequest
	GetAgentKey() *string
	SetClientToken(v string) *UpdateTagRequest
	GetClientToken() *string
	SetGroupId(v int64) *UpdateTagRequest
	GetGroupId() *int64
	SetId(v int64) *UpdateTagRequest
	GetId() *int64
	SetTagName(v string) *UpdateTagRequest
	GetTagName() *string
}

type UpdateTagRequest struct {
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
	// 455745
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 839232
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标签组1
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s UpdateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateTagRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateTagRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTagRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateTagRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *UpdateTagRequest) SetAgentKey(v string) *UpdateTagRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateTagRequest) SetClientToken(v string) *UpdateTagRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTagRequest) SetGroupId(v int64) *UpdateTagRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateTagRequest) SetId(v int64) *UpdateTagRequest {
	s.Id = &v
	return s
}

func (s *UpdateTagRequest) SetTagName(v string) *UpdateTagRequest {
	s.TagName = &v
	return s
}

func (s *UpdateTagRequest) Validate() error {
	return dara.Validate(s)
}
