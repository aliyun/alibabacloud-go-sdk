// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePerspectiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeletePerspectiveRequest
	GetAgentKey() *string
	SetPerspectiveId(v string) *DeletePerspectiveRequest
	GetPerspectiveId() *string
}

type DeletePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DeletePerspectiveRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeletePerspectiveRequest) GetPerspectiveId() *string {
	return s.PerspectiveId
}

func (s *DeletePerspectiveRequest) SetAgentKey(v string) *DeletePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *DeletePerspectiveRequest) SetPerspectiveId(v string) *DeletePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

func (s *DeletePerspectiveRequest) Validate() error {
	return dara.Validate(s)
}
