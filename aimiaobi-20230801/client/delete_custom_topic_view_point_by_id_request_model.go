// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTopicViewPointByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteCustomTopicViewPointByIdRequest
	GetAgentKey() *string
	SetCustomViewPointId(v string) *DeleteCustomTopicViewPointByIdRequest
	GetCustomViewPointId() *string
}

type DeleteCustomTopicViewPointByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dfd73894e6a94fd79fe7ffbe865796fb
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
}

func (s DeleteCustomTopicViewPointByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTopicViewPointByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicViewPointByIdRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteCustomTopicViewPointByIdRequest) GetCustomViewPointId() *string {
	return s.CustomViewPointId
}

func (s *DeleteCustomTopicViewPointByIdRequest) SetAgentKey(v string) *DeleteCustomTopicViewPointByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdRequest) SetCustomViewPointId(v string) *DeleteCustomTopicViewPointByIdRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdRequest) Validate() error {
	return dara.Validate(s)
}
