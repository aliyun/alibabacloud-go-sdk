// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskStepResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskStepResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskStepResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskStepResponseBody) *ListTaskStepResponse
	GetBody() *ListTaskStepResponseBody
}

type ListTaskStepResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskStepResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskStepResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskStepResponse) GoString() string {
	return s.String()
}

func (s *ListTaskStepResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskStepResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskStepResponse) GetBody() *ListTaskStepResponseBody {
	return s.Body
}

func (s *ListTaskStepResponse) SetHeaders(v map[string]*string) *ListTaskStepResponse {
	s.Headers = v
	return s
}

func (s *ListTaskStepResponse) SetStatusCode(v int32) *ListTaskStepResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskStepResponse) SetBody(v *ListTaskStepResponseBody) *ListTaskStepResponse {
	s.Body = v
	return s
}

func (s *ListTaskStepResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
