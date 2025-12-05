// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostAccountsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostAccountsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListHostAccountsForUserResponseBody) *ListHostAccountsForUserResponse
	GetBody() *ListHostAccountsForUserResponseBody
}

type ListHostAccountsForUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostAccountsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostAccountsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostAccountsForUserResponse) GetBody() *ListHostAccountsForUserResponseBody {
	return s.Body
}

func (s *ListHostAccountsForUserResponse) SetHeaders(v map[string]*string) *ListHostAccountsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsForUserResponse) SetStatusCode(v int32) *ListHostAccountsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsForUserResponse) SetBody(v *ListHostAccountsForUserResponseBody) *ListHostAccountsForUserResponse {
	s.Body = v
	return s
}

func (s *ListHostAccountsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
