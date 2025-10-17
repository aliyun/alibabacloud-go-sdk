// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSolutionResponse
	GetStatusCode() *int32
	SetBody(v *CreateSolutionResponseBody) *CreateSolutionResponse
	GetBody() *CreateSolutionResponseBody
}

type CreateSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSolutionResponse) GoString() string {
	return s.String()
}

func (s *CreateSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSolutionResponse) GetBody() *CreateSolutionResponseBody {
	return s.Body
}

func (s *CreateSolutionResponse) SetHeaders(v map[string]*string) *CreateSolutionResponse {
	s.Headers = v
	return s
}

func (s *CreateSolutionResponse) SetStatusCode(v int32) *CreateSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSolutionResponse) SetBody(v *CreateSolutionResponseBody) *CreateSolutionResponse {
	s.Body = v
	return s
}

func (s *CreateSolutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
