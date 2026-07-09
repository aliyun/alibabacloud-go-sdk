// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluatorSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEvaluatorSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEvaluatorSkillResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEvaluatorSkillResponseBody) *UpdateEvaluatorSkillResponse
	GetBody() *UpdateEvaluatorSkillResponseBody
}

type UpdateEvaluatorSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEvaluatorSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEvaluatorSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorSkillResponse) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEvaluatorSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEvaluatorSkillResponse) GetBody() *UpdateEvaluatorSkillResponseBody {
	return s.Body
}

func (s *UpdateEvaluatorSkillResponse) SetHeaders(v map[string]*string) *UpdateEvaluatorSkillResponse {
	s.Headers = v
	return s
}

func (s *UpdateEvaluatorSkillResponse) SetStatusCode(v int32) *UpdateEvaluatorSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEvaluatorSkillResponse) SetBody(v *UpdateEvaluatorSkillResponseBody) *UpdateEvaluatorSkillResponse {
	s.Body = v
	return s
}

func (s *UpdateEvaluatorSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
