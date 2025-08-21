// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePerspectiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdatePerspectiveRequest
	GetAgentKey() *string
	SetName(v string) *UpdatePerspectiveRequest
	GetName() *string
	SetPerspectiveId(v string) *UpdatePerspectiveRequest
	GetPerspectiveId() *string
}

type UpdatePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 客户端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s UpdatePerspectiveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdatePerspectiveRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePerspectiveRequest) GetPerspectiveId() *string {
	return s.PerspectiveId
}

func (s *UpdatePerspectiveRequest) SetAgentKey(v string) *UpdatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetName(v string) *UpdatePerspectiveRequest {
	s.Name = &v
	return s
}

func (s *UpdatePerspectiveRequest) SetPerspectiveId(v string) *UpdatePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

func (s *UpdatePerspectiveRequest) Validate() error {
	return dara.Validate(s)
}
