// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForcePublishSkillVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ForcePublishSkillVersionRequestBody) *ForcePublishSkillVersionRequest
	GetBody() *ForcePublishSkillVersionRequestBody
}

type ForcePublishSkillVersionRequest struct {
	// The request body.
	Body *ForcePublishSkillVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ForcePublishSkillVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ForcePublishSkillVersionRequest) GoString() string {
	return s.String()
}

func (s *ForcePublishSkillVersionRequest) GetBody() *ForcePublishSkillVersionRequestBody {
	return s.Body
}

func (s *ForcePublishSkillVersionRequest) SetBody(v *ForcePublishSkillVersionRequestBody) *ForcePublishSkillVersionRequest {
	s.Body = v
	return s
}

func (s *ForcePublishSkillVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ForcePublishSkillVersionRequestBody struct {
	// Specifies whether to update the latest label.
	//
	// example:
	//
	// true
	UpdateLatestLabel *bool `json:"updateLatestLabel,omitempty" xml:"updateLatestLabel,omitempty"`
}

func (s ForcePublishSkillVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ForcePublishSkillVersionRequestBody) GoString() string {
	return s.String()
}

func (s *ForcePublishSkillVersionRequestBody) GetUpdateLatestLabel() *bool {
	return s.UpdateLatestLabel
}

func (s *ForcePublishSkillVersionRequestBody) SetUpdateLatestLabel(v bool) *ForcePublishSkillVersionRequestBody {
	s.UpdateLatestLabel = &v
	return s
}

func (s *ForcePublishSkillVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
