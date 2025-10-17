// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobStepResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobStepResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobStepResponse
	GetStatusCode() *int32
	SetBody(v *ListJobStepResponseBody) *ListJobStepResponse
	GetBody() *ListJobStepResponseBody
}

type ListJobStepResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobStepResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobStepResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepResponse) GoString() string {
	return s.String()
}

func (s *ListJobStepResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobStepResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobStepResponse) GetBody() *ListJobStepResponseBody {
	return s.Body
}

func (s *ListJobStepResponse) SetHeaders(v map[string]*string) *ListJobStepResponse {
	s.Headers = v
	return s
}

func (s *ListJobStepResponse) SetStatusCode(v int32) *ListJobStepResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobStepResponse) SetBody(v *ListJobStepResponseBody) *ListJobStepResponse {
	s.Body = v
	return s
}

func (s *ListJobStepResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
