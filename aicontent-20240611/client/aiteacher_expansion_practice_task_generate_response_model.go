// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAITeacherExpansionPracticeTaskGenerateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AITeacherExpansionPracticeTaskGenerateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponse
	GetStatusCode() *int32
	SetBody(v *AITeacherExpansionPracticeTaskGenerateResponseBody) *AITeacherExpansionPracticeTaskGenerateResponse
	GetBody() *AITeacherExpansionPracticeTaskGenerateResponseBody
}

type AITeacherExpansionPracticeTaskGenerateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AITeacherExpansionPracticeTaskGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponse) String() string {
	return dara.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponse) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) GetBody() *AITeacherExpansionPracticeTaskGenerateResponseBody {
	return s.Body
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetHeaders(v map[string]*string) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.Headers = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetBody(v *AITeacherExpansionPracticeTaskGenerateResponseBody) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.Body = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
