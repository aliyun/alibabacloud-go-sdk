// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationRolesResponseBody) *ListApplicationRolesResponse
	GetBody() *ListApplicationRolesResponseBody
}

type ListApplicationRolesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRolesResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationRolesResponse) GetBody() *ListApplicationRolesResponseBody {
	return s.Body
}

func (s *ListApplicationRolesResponse) SetHeaders(v map[string]*string) *ListApplicationRolesResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationRolesResponse) SetStatusCode(v int32) *ListApplicationRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationRolesResponse) SetBody(v *ListApplicationRolesResponseBody) *ListApplicationRolesResponse {
	s.Body = v
	return s
}

func (s *ListApplicationRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
