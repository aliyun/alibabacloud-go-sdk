// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteTagGroupRequest
	GetAgentKey() *string
	SetClientToken(v string) *DeleteTagGroupRequest
	GetClientToken() *string
	SetId(v int64) *DeleteTagGroupRequest
	GetId() *int64
}

type DeleteTagGroupRequest struct {
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
	// 839232
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagGroupRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteTagGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTagGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteTagGroupRequest) SetAgentKey(v string) *DeleteTagGroupRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteTagGroupRequest) SetClientToken(v string) *DeleteTagGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTagGroupRequest) SetId(v int64) *DeleteTagGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
