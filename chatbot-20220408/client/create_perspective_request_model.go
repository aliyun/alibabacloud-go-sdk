// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerspectiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreatePerspectiveRequest
	GetAgentKey() *string
	SetDescription(v string) *CreatePerspectiveRequest
	GetDescription() *string
	SetName(v string) *CreatePerspectiveRequest
	GetName() *string
}

type CreatePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 用于购物APP的移动端视角
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 移动端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePerspectiveRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreatePerspectiveRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePerspectiveRequest) GetName() *string {
	return s.Name
}

func (s *CreatePerspectiveRequest) SetAgentKey(v string) *CreatePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePerspectiveRequest) SetDescription(v string) *CreatePerspectiveRequest {
	s.Description = &v
	return s
}

func (s *CreatePerspectiveRequest) SetName(v string) *CreatePerspectiveRequest {
	s.Name = &v
	return s
}

func (s *CreatePerspectiveRequest) Validate() error {
	return dara.Validate(s)
}
