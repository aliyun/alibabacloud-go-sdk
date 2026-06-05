// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteAgentRequest
	GetName() *string
}

type DeleteAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) GetName() *string {
	return s.Name
}

func (s *DeleteAgentRequest) SetName(v string) *DeleteAgentRequest {
	s.Name = &v
	return s
}

func (s *DeleteAgentRequest) Validate() error {
	return dara.Validate(s)
}
