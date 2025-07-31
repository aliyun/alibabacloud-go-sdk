// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddableRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddableRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListAddableRolesResponseBody) *ListAddableRolesResponse
	GetBody() *ListAddableRolesResponseBody
}

type ListAddableRolesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddableRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddableRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddableRolesResponse) GoString() string {
	return s.String()
}

func (s *ListAddableRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddableRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddableRolesResponse) GetBody() *ListAddableRolesResponseBody {
	return s.Body
}

func (s *ListAddableRolesResponse) SetHeaders(v map[string]*string) *ListAddableRolesResponse {
	s.Headers = v
	return s
}

func (s *ListAddableRolesResponse) SetStatusCode(v int32) *ListAddableRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddableRolesResponse) SetBody(v *ListAddableRolesResponseBody) *ListAddableRolesResponse {
	s.Body = v
	return s
}

func (s *ListAddableRolesResponse) Validate() error {
	return dara.Validate(s)
}
