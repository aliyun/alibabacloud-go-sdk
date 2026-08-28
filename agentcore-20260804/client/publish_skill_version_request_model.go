// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PublishSkillVersionRequestBody) *PublishSkillVersionRequest
	GetBody() *PublishSkillVersionRequestBody
}

type PublishSkillVersionRequest struct {
	// The request body.
	Body *PublishSkillVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PublishSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishSkillVersionRequest) GetBody() *PublishSkillVersionRequestBody {
	return s.Body
}

func (s *PublishSkillVersionRequest) SetBody(v *PublishSkillVersionRequestBody) *PublishSkillVersionRequest {
	s.Body = v
	return s
}

func (s *PublishSkillVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishSkillVersionRequestBody struct {
	// Specifies whether to update the latest label.
	//
	// example:
	//
	// true
	UpdateLatestLabel *bool `json:"updateLatestLabel,omitempty" xml:"updateLatestLabel,omitempty"`
}

func (s PublishSkillVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s PublishSkillVersionRequestBody) GoString() string {
	return s.String()
}

func (s *PublishSkillVersionRequestBody) GetUpdateLatestLabel() *bool {
	return s.UpdateLatestLabel
}

func (s *PublishSkillVersionRequestBody) SetUpdateLatestLabel(v bool) *PublishSkillVersionRequestBody {
	s.UpdateLatestLabel = &v
	return s
}

func (s *PublishSkillVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
