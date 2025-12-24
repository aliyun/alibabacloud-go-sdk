// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPythonFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyPythonFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyPythonFileResponse
	GetStatusCode() *int32
	SetBody(v *VerifyPythonFileResponseBody) *VerifyPythonFileResponse
	GetBody() *VerifyPythonFileResponseBody
}

type VerifyPythonFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPythonFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPythonFileResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyPythonFileResponse) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyPythonFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyPythonFileResponse) GetBody() *VerifyPythonFileResponseBody {
	return s.Body
}

func (s *VerifyPythonFileResponse) SetHeaders(v map[string]*string) *VerifyPythonFileResponse {
	s.Headers = v
	return s
}

func (s *VerifyPythonFileResponse) SetStatusCode(v int32) *VerifyPythonFileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPythonFileResponse) SetBody(v *VerifyPythonFileResponseBody) *VerifyPythonFileResponse {
	s.Body = v
	return s
}

func (s *VerifyPythonFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
