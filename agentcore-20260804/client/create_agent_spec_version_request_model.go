// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAgentSpecVersionRequestBody) *CreateAgentSpecVersionRequest
	GetBody() *CreateAgentSpecVersionRequestBody
}

type CreateAgentSpecVersionRequest struct {
	// The request body.
	Body *CreateAgentSpecVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateAgentSpecVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecVersionRequest) GetBody() *CreateAgentSpecVersionRequestBody {
	return s.Body
}

func (s *CreateAgentSpecVersionRequest) SetBody(v *CreateAgentSpecVersionRequestBody) *CreateAgentSpecVersionRequest {
	s.Body = v
	return s
}

func (s *CreateAgentSpecVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSpecVersionRequestBody struct {
	// The existing version on which to base the draft.
	//
	// example:
	//
	// 1.0.0
	BasedOnVersion *string `json:"basedOnVersion,omitempty" xml:"basedOnVersion,omitempty"`
	// The version number for the draft. If not specified, the version number is automatically incremented.
	//
	// example:
	//
	// 2.0.0
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s CreateAgentSpecVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecVersionRequestBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecVersionRequestBody) GetBasedOnVersion() *string {
	return s.BasedOnVersion
}

func (s *CreateAgentSpecVersionRequestBody) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateAgentSpecVersionRequestBody) SetBasedOnVersion(v string) *CreateAgentSpecVersionRequestBody {
	s.BasedOnVersion = &v
	return s
}

func (s *CreateAgentSpecVersionRequestBody) SetTargetVersion(v string) *CreateAgentSpecVersionRequestBody {
	s.TargetVersion = &v
	return s
}

func (s *CreateAgentSpecVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
