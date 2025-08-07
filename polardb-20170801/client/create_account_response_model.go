// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccountResponseBody) *CreateAccountResponse
	GetBody() *CreateAccountResponseBody
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccountResponse) GetBody() *CreateAccountResponseBody {
	return s.Body
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

func (s *CreateAccountResponse) Validate() error {
	return dara.Validate(s)
}
