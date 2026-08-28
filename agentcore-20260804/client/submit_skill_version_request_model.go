// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SubmitSkillVersionRequestBody) *SubmitSkillVersionRequest
	GetBody() *SubmitSkillVersionRequestBody
}

type SubmitSkillVersionRequest struct {
	// The request body.
	Body *SubmitSkillVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SubmitSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionRequest) GetBody() *SubmitSkillVersionRequestBody {
	return s.Body
}

func (s *SubmitSkillVersionRequest) SetBody(v *SubmitSkillVersionRequestBody) *SubmitSkillVersionRequest {
	s.Body = v
	return s
}

func (s *SubmitSkillVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSkillVersionRequestBody struct {
}

func (s SubmitSkillVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionRequestBody) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
