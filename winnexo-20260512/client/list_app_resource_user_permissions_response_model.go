// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourceUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppResourceUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppResourceUserPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppResourceUserPermissionsResponseBody) *ListAppResourceUserPermissionsResponse
	GetBody() *ListAppResourceUserPermissionsResponseBody
}

type ListAppResourceUserPermissionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppResourceUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppResourceUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourceUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppResourceUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppResourceUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppResourceUserPermissionsResponse) GetBody() *ListAppResourceUserPermissionsResponseBody {
	return s.Body
}

func (s *ListAppResourceUserPermissionsResponse) SetHeaders(v map[string]*string) *ListAppResourceUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppResourceUserPermissionsResponse) SetStatusCode(v int32) *ListAppResourceUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppResourceUserPermissionsResponse) SetBody(v *ListAppResourceUserPermissionsResponseBody) *ListAppResourceUserPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListAppResourceUserPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
