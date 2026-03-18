// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAITeacherSyncPracticeTaskGenerateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AITeacherSyncPracticeTaskGenerateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponse
	GetStatusCode() *int32
	SetBody(v *AITeacherSyncPracticeTaskGenerateResponseBody) *AITeacherSyncPracticeTaskGenerateResponse
	GetBody() *AITeacherSyncPracticeTaskGenerateResponseBody
}

type AITeacherSyncPracticeTaskGenerateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AITeacherSyncPracticeTaskGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponse) String() string {
	return dara.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponse) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) GetBody() *AITeacherSyncPracticeTaskGenerateResponseBody {
	return s.Body
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetHeaders(v map[string]*string) *AITeacherSyncPracticeTaskGenerateResponse {
	s.Headers = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetBody(v *AITeacherSyncPracticeTaskGenerateResponseBody) *AITeacherSyncPracticeTaskGenerateResponse {
	s.Body = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
