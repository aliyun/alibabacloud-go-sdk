// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForResourceServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUsersForResourceServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUsersForResourceServerResponse
	GetStatusCode() *int32
	SetBody(v *ListUsersForResourceServerResponseBody) *ListUsersForResourceServerResponse
	GetBody() *ListUsersForResourceServerResponseBody
}

type ListUsersForResourceServerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForResourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForResourceServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForResourceServerResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForResourceServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUsersForResourceServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUsersForResourceServerResponse) GetBody() *ListUsersForResourceServerResponseBody {
	return s.Body
}

func (s *ListUsersForResourceServerResponse) SetHeaders(v map[string]*string) *ListUsersForResourceServerResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForResourceServerResponse) SetStatusCode(v int32) *ListUsersForResourceServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForResourceServerResponse) SetBody(v *ListUsersForResourceServerResponseBody) *ListUsersForResourceServerResponse {
	s.Body = v
	return s
}

func (s *ListUsersForResourceServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
