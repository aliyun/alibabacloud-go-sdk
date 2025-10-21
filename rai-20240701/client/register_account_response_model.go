// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterAccountResponse
	GetStatusCode() *int32
	SetBody(v *RegisterAccountResponseBody) *RegisterAccountResponse
	GetBody() *RegisterAccountResponseBody
}

type RegisterAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterAccountResponse) GoString() string {
	return s.String()
}

func (s *RegisterAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterAccountResponse) GetBody() *RegisterAccountResponseBody {
	return s.Body
}

func (s *RegisterAccountResponse) SetHeaders(v map[string]*string) *RegisterAccountResponse {
	s.Headers = v
	return s
}

func (s *RegisterAccountResponse) SetStatusCode(v int32) *RegisterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterAccountResponse) SetBody(v *RegisterAccountResponseBody) *RegisterAccountResponse {
	s.Body = v
	return s
}

func (s *RegisterAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
