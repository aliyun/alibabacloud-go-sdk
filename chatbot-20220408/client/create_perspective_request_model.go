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
	// The key of the business space. If this parameter is not specified, the default business space is accessed. You can obtain the key from the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The description of the perspective.
	//
	// example:
	//
	// A mobile perspective for the shopping app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the perspective. The name can be up to 50 characters in length.
	//
	// example:
	//
	// Mobile perspective
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
