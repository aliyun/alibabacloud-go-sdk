// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OnlineSkillRequestBody) *OnlineSkillRequest
	GetBody() *OnlineSkillRequestBody
}

type OnlineSkillRequest struct {
	// The request body.
	Body *OnlineSkillRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s OnlineSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineSkillRequest) GoString() string {
	return s.String()
}

func (s *OnlineSkillRequest) GetBody() *OnlineSkillRequestBody {
	return s.Body
}

func (s *OnlineSkillRequest) SetBody(v *OnlineSkillRequestBody) *OnlineSkillRequest {
	s.Body = v
	return s
}

func (s *OnlineSkillRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnlineSkillRequestBody struct {
	// The operation scope. Valid values:
	//
	// - skill: the entire Skill.
	//
	// - version: a specified version.
	//
	// example:
	//
	// version
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The Skill version number.
	//
	// example:
	//
	// 0.0.1
	SkillVersion *string `json:"skillVersion,omitempty" xml:"skillVersion,omitempty"`
}

func (s OnlineSkillRequestBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineSkillRequestBody) GoString() string {
	return s.String()
}

func (s *OnlineSkillRequestBody) GetScope() *string {
	return s.Scope
}

func (s *OnlineSkillRequestBody) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *OnlineSkillRequestBody) SetScope(v string) *OnlineSkillRequestBody {
	s.Scope = &v
	return s
}

func (s *OnlineSkillRequestBody) SetSkillVersion(v string) *OnlineSkillRequestBody {
	s.SkillVersion = &v
	return s
}

func (s *OnlineSkillRequestBody) Validate() error {
	return dara.Validate(s)
}
