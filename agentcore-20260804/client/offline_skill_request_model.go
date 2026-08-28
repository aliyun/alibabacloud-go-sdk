// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OfflineSkillRequestBody) *OfflineSkillRequest
	GetBody() *OfflineSkillRequestBody
}

type OfflineSkillRequest struct {
	// The request body.
	Body *OfflineSkillRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s OfflineSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineSkillRequest) GoString() string {
	return s.String()
}

func (s *OfflineSkillRequest) GetBody() *OfflineSkillRequestBody {
	return s.Body
}

func (s *OfflineSkillRequest) SetBody(v *OfflineSkillRequestBody) *OfflineSkillRequest {
	s.Body = v
	return s
}

func (s *OfflineSkillRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OfflineSkillRequestBody struct {
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

func (s OfflineSkillRequestBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineSkillRequestBody) GoString() string {
	return s.String()
}

func (s *OfflineSkillRequestBody) GetScope() *string {
	return s.Scope
}

func (s *OfflineSkillRequestBody) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *OfflineSkillRequestBody) SetScope(v string) *OfflineSkillRequestBody {
	s.Scope = &v
	return s
}

func (s *OfflineSkillRequestBody) SetSkillVersion(v string) *OfflineSkillRequestBody {
	s.SkillVersion = &v
	return s
}

func (s *OfflineSkillRequestBody) Validate() error {
	return dara.Validate(s)
}
