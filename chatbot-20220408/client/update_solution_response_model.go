// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSolutionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSolutionResponseBody) *UpdateSolutionResponse
	GetBody() *UpdateSolutionResponseBody
}

type UpdateSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSolutionResponse) GetBody() *UpdateSolutionResponseBody {
	return s.Body
}

func (s *UpdateSolutionResponse) SetHeaders(v map[string]*string) *UpdateSolutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSolutionResponse) SetStatusCode(v int32) *UpdateSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSolutionResponse) SetBody(v *UpdateSolutionResponseBody) *UpdateSolutionResponse {
	s.Body = v
	return s
}

func (s *UpdateSolutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
