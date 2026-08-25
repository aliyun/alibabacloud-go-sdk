// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDelegateAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDelegateAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDelegateAccountResponse
	GetStatusCode() *int32
	SetBody(v *DisableDelegateAccountResponseBody) *DisableDelegateAccountResponse
	GetBody() *DisableDelegateAccountResponseBody
}

type DisableDelegateAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDelegateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDelegateAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDelegateAccountResponse) GoString() string {
	return s.String()
}

func (s *DisableDelegateAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDelegateAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDelegateAccountResponse) GetBody() *DisableDelegateAccountResponseBody {
	return s.Body
}

func (s *DisableDelegateAccountResponse) SetHeaders(v map[string]*string) *DisableDelegateAccountResponse {
	s.Headers = v
	return s
}

func (s *DisableDelegateAccountResponse) SetStatusCode(v int32) *DisableDelegateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDelegateAccountResponse) SetBody(v *DisableDelegateAccountResponseBody) *DisableDelegateAccountResponse {
	s.Body = v
	return s
}

func (s *DisableDelegateAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
