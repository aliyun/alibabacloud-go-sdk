// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetAgentRequest
	GetName() *string
}

type GetAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetName() *string {
	return s.Name
}

func (s *GetAgentRequest) SetName(v string) *GetAgentRequest {
	s.Name = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
