// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteTagRequest
	GetAgentKey() *string
	SetClientToken(v string) *DeleteTagRequest
	GetClientToken() *string
	SetGroupId(v int64) *DeleteTagRequest
	GetGroupId() *int64
	SetId(v int64) *DeleteTagRequest
	GetId() *int64
}

type DeleteTagRequest struct {
	// example:
	//
	// xxxx
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 43435
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 839232
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteTagRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTagRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteTagRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteTagRequest) SetAgentKey(v string) *DeleteTagRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteTagRequest) SetClientToken(v string) *DeleteTagRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTagRequest) SetGroupId(v int64) *DeleteTagRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteTagRequest) SetId(v int64) *DeleteTagRequest {
	s.Id = &v
	return s
}

func (s *DeleteTagRequest) Validate() error {
	return dara.Validate(s)
}
