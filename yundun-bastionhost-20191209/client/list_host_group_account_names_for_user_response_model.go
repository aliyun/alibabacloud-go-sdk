// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupAccountNamesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostGroupAccountNamesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostGroupAccountNamesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListHostGroupAccountNamesForUserResponseBody) *ListHostGroupAccountNamesForUserResponse
	GetBody() *ListHostGroupAccountNamesForUserResponseBody
}

type ListHostGroupAccountNamesForUserResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupAccountNamesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostGroupAccountNamesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostGroupAccountNamesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostGroupAccountNamesForUserResponse) GetBody() *ListHostGroupAccountNamesForUserResponseBody {
	return s.Body
}

func (s *ListHostGroupAccountNamesForUserResponse) SetHeaders(v map[string]*string) *ListHostGroupAccountNamesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponse) SetStatusCode(v int32) *ListHostGroupAccountNamesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponse) SetBody(v *ListHostGroupAccountNamesForUserResponseBody) *ListHostGroupAccountNamesForUserResponse {
	s.Body = v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponse) Validate() error {
	return dara.Validate(s)
}
