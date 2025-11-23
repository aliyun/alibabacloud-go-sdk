// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPermissionsResponseBody) *ListUserPermissionsResponse
	GetBody() *ListUserPermissionsResponseBody
}

type ListUserPermissionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPermissionsResponse) GetBody() *ListUserPermissionsResponseBody {
	return s.Body
}

func (s *ListUserPermissionsResponse) SetHeaders(v map[string]*string) *ListUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserPermissionsResponse) SetStatusCode(v int32) *ListUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPermissionsResponse) SetBody(v *ListUserPermissionsResponseBody) *ListUserPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListUserPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
