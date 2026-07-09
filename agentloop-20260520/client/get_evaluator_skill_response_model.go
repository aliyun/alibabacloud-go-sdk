// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluatorSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEvaluatorSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEvaluatorSkillResponse
	GetStatusCode() *int32
	SetBody(v *GetEvaluatorSkillResponseBody) *GetEvaluatorSkillResponse
	GetBody() *GetEvaluatorSkillResponseBody
}

type GetEvaluatorSkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEvaluatorSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEvaluatorSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorSkillResponse) GoString() string {
	return s.String()
}

func (s *GetEvaluatorSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEvaluatorSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEvaluatorSkillResponse) GetBody() *GetEvaluatorSkillResponseBody {
	return s.Body
}

func (s *GetEvaluatorSkillResponse) SetHeaders(v map[string]*string) *GetEvaluatorSkillResponse {
	s.Headers = v
	return s
}

func (s *GetEvaluatorSkillResponse) SetStatusCode(v int32) *GetEvaluatorSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEvaluatorSkillResponse) SetBody(v *GetEvaluatorSkillResponseBody) *GetEvaluatorSkillResponse {
	s.Body = v
	return s
}

func (s *GetEvaluatorSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
