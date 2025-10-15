// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccountsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationAccountsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationAccountsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationAccountsForUserResponseBody) *ListApplicationAccountsForUserResponse
	GetBody() *ListApplicationAccountsForUserResponseBody
}

type ListApplicationAccountsForUserResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationAccountsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationAccountsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationAccountsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationAccountsForUserResponse) GetBody() *ListApplicationAccountsForUserResponseBody {
	return s.Body
}

func (s *ListApplicationAccountsForUserResponse) SetHeaders(v map[string]*string) *ListApplicationAccountsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationAccountsForUserResponse) SetStatusCode(v int32) *ListApplicationAccountsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationAccountsForUserResponse) SetBody(v *ListApplicationAccountsForUserResponseBody) *ListApplicationAccountsForUserResponse {
	s.Body = v
	return s
}

func (s *ListApplicationAccountsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
