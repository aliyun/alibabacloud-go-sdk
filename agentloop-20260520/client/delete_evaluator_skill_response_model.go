// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluatorSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEvaluatorSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEvaluatorSkillResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEvaluatorSkillResponseBody) *DeleteEvaluatorSkillResponse
	GetBody() *DeleteEvaluatorSkillResponseBody
}

type DeleteEvaluatorSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEvaluatorSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEvaluatorSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluatorSkillResponse) GoString() string {
	return s.String()
}

func (s *DeleteEvaluatorSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEvaluatorSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEvaluatorSkillResponse) GetBody() *DeleteEvaluatorSkillResponseBody {
	return s.Body
}

func (s *DeleteEvaluatorSkillResponse) SetHeaders(v map[string]*string) *DeleteEvaluatorSkillResponse {
	s.Headers = v
	return s
}

func (s *DeleteEvaluatorSkillResponse) SetStatusCode(v int32) *DeleteEvaluatorSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEvaluatorSkillResponse) SetBody(v *DeleteEvaluatorSkillResponseBody) *DeleteEvaluatorSkillResponse {
	s.Body = v
	return s
}

func (s *DeleteEvaluatorSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
