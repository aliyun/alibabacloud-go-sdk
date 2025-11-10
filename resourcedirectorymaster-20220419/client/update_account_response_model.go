// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccountResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccountResponseBody) *UpdateAccountResponse
	GetBody() *UpdateAccountResponseBody
}

type UpdateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccountResponse) GetBody() *UpdateAccountResponseBody {
	return s.Body
}

func (s *UpdateAccountResponse) SetHeaders(v map[string]*string) *UpdateAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountResponse) SetStatusCode(v int32) *UpdateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountResponse) SetBody(v *UpdateAccountResponseBody) *UpdateAccountResponse {
	s.Body = v
	return s
}

func (s *UpdateAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
