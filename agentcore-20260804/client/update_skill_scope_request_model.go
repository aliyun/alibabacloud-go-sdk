// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateSkillScopeRequestBody) *UpdateSkillScopeRequest
	GetBody() *UpdateSkillScopeRequestBody
}

type UpdateSkillScopeRequest struct {
	// The request body.
	Body *UpdateSkillScopeRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateSkillScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeRequest) GetBody() *UpdateSkillScopeRequestBody {
	return s.Body
}

func (s *UpdateSkillScopeRequest) SetBody(v *UpdateSkillScopeRequestBody) *UpdateSkillScopeRequest {
	s.Body = v
	return s
}

func (s *UpdateSkillScopeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSkillScopeRequestBody struct {
	// The visibility scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s UpdateSkillScopeRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeRequestBody) GetScope() *string {
	return s.Scope
}

func (s *UpdateSkillScopeRequestBody) SetScope(v string) *UpdateSkillScopeRequestBody {
	s.Scope = &v
	return s
}

func (s *UpdateSkillScopeRequestBody) Validate() error {
	return dara.Validate(s)
}
