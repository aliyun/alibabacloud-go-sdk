// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDatabaseAccountsFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDatabaseAccountsFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDatabaseAccountsFromUserResponse
	GetStatusCode() *int32
	SetBody(v *DetachDatabaseAccountsFromUserResponseBody) *DetachDatabaseAccountsFromUserResponse
	GetBody() *DetachDatabaseAccountsFromUserResponseBody
}

type DetachDatabaseAccountsFromUserResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDatabaseAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDatabaseAccountsFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDatabaseAccountsFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDatabaseAccountsFromUserResponse) GetBody() *DetachDatabaseAccountsFromUserResponseBody {
	return s.Body
}

func (s *DetachDatabaseAccountsFromUserResponse) SetHeaders(v map[string]*string) *DetachDatabaseAccountsFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponse) SetStatusCode(v int32) *DetachDatabaseAccountsFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponse) SetBody(v *DetachDatabaseAccountsFromUserResponseBody) *DetachDatabaseAccountsFromUserResponse {
	s.Body = v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
