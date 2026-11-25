// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAppJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIAppJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIAppJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIAppJobResponseBody) *SubmitAIAppJobResponse
	GetBody() *SubmitAIAppJobResponseBody
}

type SubmitAIAppJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIAppJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIAppJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAppJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIAppJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIAppJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIAppJobResponse) GetBody() *SubmitAIAppJobResponseBody {
	return s.Body
}

func (s *SubmitAIAppJobResponse) SetHeaders(v map[string]*string) *SubmitAIAppJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIAppJobResponse) SetStatusCode(v int32) *SubmitAIAppJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIAppJobResponse) SetBody(v *SubmitAIAppJobResponseBody) *SubmitAIAppJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIAppJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
