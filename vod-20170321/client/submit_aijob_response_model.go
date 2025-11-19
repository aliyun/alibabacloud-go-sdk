// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIJobResponseBody) *SubmitAIJobResponse
	GetBody() *SubmitAIJobResponseBody
}

type SubmitAIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIJobResponse) GetBody() *SubmitAIJobResponseBody {
	return s.Body
}

func (s *SubmitAIJobResponse) SetHeaders(v map[string]*string) *SubmitAIJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIJobResponse) SetStatusCode(v int32) *SubmitAIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIJobResponse) SetBody(v *SubmitAIJobResponseBody) *SubmitAIJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
