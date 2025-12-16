// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserAnalyzerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserAnalyzerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserAnalyzerResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserAnalyzerResponseBody) *CreateUserAnalyzerResponse
	GetBody() *CreateUserAnalyzerResponseBody
}

type CreateUserAnalyzerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserAnalyzerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserAnalyzerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserAnalyzerResponse) GetBody() *CreateUserAnalyzerResponseBody {
	return s.Body
}

func (s *CreateUserAnalyzerResponse) SetHeaders(v map[string]*string) *CreateUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *CreateUserAnalyzerResponse) SetStatusCode(v int32) *CreateUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserAnalyzerResponse) SetBody(v *CreateUserAnalyzerResponseBody) *CreateUserAnalyzerResponse {
	s.Body = v
	return s
}

func (s *CreateUserAnalyzerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
