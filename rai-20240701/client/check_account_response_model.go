// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccountResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccountResponseBody) *CheckAccountResponse
	GetBody() *CheckAccountResponseBody
}

type CheckAccountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccountResponse) GetBody() *CheckAccountResponseBody {
	return s.Body
}

func (s *CheckAccountResponse) SetHeaders(v map[string]*string) *CheckAccountResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountResponse) SetStatusCode(v int32) *CheckAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountResponse) SetBody(v *CheckAccountResponseBody) *CheckAccountResponse {
	s.Body = v
	return s
}

func (s *CheckAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
