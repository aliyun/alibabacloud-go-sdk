// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRbacPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRbacPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRbacPermissionsResponseBody) *ListRbacPermissionsResponse
	GetBody() *ListRbacPermissionsResponseBody
}

type ListRbacPermissionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRbacPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRbacPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRbacPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListRbacPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRbacPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRbacPermissionsResponse) GetBody() *ListRbacPermissionsResponseBody {
	return s.Body
}

func (s *ListRbacPermissionsResponse) SetHeaders(v map[string]*string) *ListRbacPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListRbacPermissionsResponse) SetStatusCode(v int32) *ListRbacPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRbacPermissionsResponse) SetBody(v *ListRbacPermissionsResponseBody) *ListRbacPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListRbacPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
