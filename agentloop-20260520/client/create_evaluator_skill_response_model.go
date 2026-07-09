// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluatorSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEvaluatorSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEvaluatorSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateEvaluatorSkillResponseBody) *CreateEvaluatorSkillResponse
	GetBody() *CreateEvaluatorSkillResponseBody
}

type CreateEvaluatorSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEvaluatorSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEvaluatorSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluatorSkillResponse) GoString() string {
	return s.String()
}

func (s *CreateEvaluatorSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEvaluatorSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEvaluatorSkillResponse) GetBody() *CreateEvaluatorSkillResponseBody {
	return s.Body
}

func (s *CreateEvaluatorSkillResponse) SetHeaders(v map[string]*string) *CreateEvaluatorSkillResponse {
	s.Headers = v
	return s
}

func (s *CreateEvaluatorSkillResponse) SetStatusCode(v int32) *CreateEvaluatorSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEvaluatorSkillResponse) SetBody(v *CreateEvaluatorSkillResponseBody) *CreateEvaluatorSkillResponse {
	s.Body = v
	return s
}

func (s *CreateEvaluatorSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
