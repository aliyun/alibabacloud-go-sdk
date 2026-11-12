// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateSkillLabelsRequestBody) *UpdateSkillLabelsRequest
	GetBody() *UpdateSkillLabelsRequestBody
}

type UpdateSkillLabelsRequest struct {
	// The request body.
	Body *UpdateSkillLabelsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateSkillLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillLabelsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillLabelsRequest) GetBody() *UpdateSkillLabelsRequestBody {
	return s.Body
}

func (s *UpdateSkillLabelsRequest) SetBody(v *UpdateSkillLabelsRequestBody) *UpdateSkillLabelsRequest {
	s.Body = v
	return s
}

func (s *UpdateSkillLabelsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSkillLabelsRequestBody struct {
	// The version label mapping JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"latest":"0.0.2","stable":"0.0.1"}
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
}

func (s UpdateSkillLabelsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillLabelsRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillLabelsRequestBody) GetLabels() *string {
	return s.Labels
}

func (s *UpdateSkillLabelsRequestBody) SetLabels(v string) *UpdateSkillLabelsRequestBody {
	s.Labels = &v
	return s
}

func (s *UpdateSkillLabelsRequestBody) Validate() error {
	return dara.Validate(s)
}
