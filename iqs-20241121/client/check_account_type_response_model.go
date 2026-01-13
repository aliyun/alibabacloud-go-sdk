// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountTypeResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountTypeResponseBody) *CheckAccountTypeResponse
	GetBody() *CheckAccountTypeResponseBody
}

type CheckAccountTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountTypeResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountTypeResponse) GetBody() *CheckAccountTypeResponseBody {
	return s.Body
}

func (s *CheckAccountTypeResponse) SetHeaders(v map[string]*string) *CheckAccountTypeResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountTypeResponse) SetStatusCode(v int32) *CheckAccountTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountTypeResponse) SetBody(v *CheckAccountTypeResponseBody) *CheckAccountTypeResponse {
	s.Body = v
	return s
}

func (s *CheckAccountTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
