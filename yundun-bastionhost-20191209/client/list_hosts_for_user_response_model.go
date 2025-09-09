// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListHostsForUserResponseBody) *ListHostsForUserResponse
	GetBody() *ListHostsForUserResponseBody
}

type ListHostsForUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostsForUserResponse) GetBody() *ListHostsForUserResponseBody {
	return s.Body
}

func (s *ListHostsForUserResponse) SetHeaders(v map[string]*string) *ListHostsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostsForUserResponse) SetStatusCode(v int32) *ListHostsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostsForUserResponse) SetBody(v *ListHostsForUserResponseBody) *ListHostsForUserResponse {
	s.Body = v
	return s
}

func (s *ListHostsForUserResponse) Validate() error {
	return dara.Validate(s)
}
