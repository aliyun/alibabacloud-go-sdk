// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedraftSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RedraftSkillVersionRequestBody) *RedraftSkillVersionRequest
	GetBody() *RedraftSkillVersionRequestBody
}

type RedraftSkillVersionRequest struct {
	// The request body.
	Body *RedraftSkillVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s RedraftSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s RedraftSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *RedraftSkillVersionRequest) GetBody() *RedraftSkillVersionRequestBody {
	return s.Body
}

func (s *RedraftSkillVersionRequest) SetBody(v *RedraftSkillVersionRequestBody) *RedraftSkillVersionRequest {
	s.Body = v
	return s
}

func (s *RedraftSkillVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RedraftSkillVersionRequestBody struct {
}

func (s RedraftSkillVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s RedraftSkillVersionRequestBody) GoString() string {
	return s.String()
}

func (s *RedraftSkillVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
