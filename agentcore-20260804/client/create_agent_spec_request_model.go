// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAgentSpecRequestBody) *CreateAgentSpecRequest
	GetBody() *CreateAgentSpecRequestBody
}

type CreateAgentSpecRequest struct {
	// The request body.
	Body *CreateAgentSpecRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateAgentSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecRequest) GetBody() *CreateAgentSpecRequestBody {
	return s.Body
}

func (s *CreateAgentSpecRequest) SetBody(v *CreateAgentSpecRequestBody) *CreateAgentSpecRequest {
	s.Body = v
	return s
}

func (s *CreateAgentSpecRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSpecRequestBody struct {
	// The unique name of the AgentSpec.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-worker
	AgentSpecName *string `json:"agentSpecName,omitempty" xml:"agentSpecName,omitempty"`
	// The draft version number. If not specified, the default value is 0.0.1.
	//
	// example:
	//
	// 0.0.1
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s CreateAgentSpecRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecRequestBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecRequestBody) GetAgentSpecName() *string {
	return s.AgentSpecName
}

func (s *CreateAgentSpecRequestBody) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateAgentSpecRequestBody) SetAgentSpecName(v string) *CreateAgentSpecRequestBody {
	s.AgentSpecName = &v
	return s
}

func (s *CreateAgentSpecRequestBody) SetTargetVersion(v string) *CreateAgentSpecRequestBody {
	s.TargetVersion = &v
	return s
}

func (s *CreateAgentSpecRequestBody) Validate() error {
	return dara.Validate(s)
}
