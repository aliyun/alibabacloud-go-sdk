// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTestFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateTestFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateTestFileResponse
	GetStatusCode() *int32
	SetBody(v *ValidateTestFileResponseBody) *ValidateTestFileResponse
	GetBody() *ValidateTestFileResponseBody
}

type ValidateTestFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateTestFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateTestFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateTestFileResponse) GoString() string {
	return s.String()
}

func (s *ValidateTestFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateTestFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateTestFileResponse) GetBody() *ValidateTestFileResponseBody {
	return s.Body
}

func (s *ValidateTestFileResponse) SetHeaders(v map[string]*string) *ValidateTestFileResponse {
	s.Headers = v
	return s
}

func (s *ValidateTestFileResponse) SetStatusCode(v int32) *ValidateTestFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateTestFileResponse) SetBody(v *ValidateTestFileResponseBody) *ValidateTestFileResponse {
	s.Body = v
	return s
}

func (s *ValidateTestFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
