// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *ResumeAgentTaskResponseBody) *ResumeAgentTaskResponse
	GetBody() *ResumeAgentTaskResponseBody
}

type ResumeAgentTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeAgentTaskResponse) GetBody() *ResumeAgentTaskResponseBody {
	return s.Body
}

func (s *ResumeAgentTaskResponse) SetHeaders(v map[string]*string) *ResumeAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeAgentTaskResponse) SetStatusCode(v int32) *ResumeAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeAgentTaskResponse) SetBody(v *ResumeAgentTaskResponseBody) *ResumeAgentTaskResponse {
	s.Body = v
	return s
}

func (s *ResumeAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
