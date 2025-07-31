// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePermissionsResponseBody) *ListResourcePermissionsResponse
	GetBody() *ListResourcePermissionsResponseBody
}

type ListResourcePermissionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePermissionsResponse) GetBody() *ListResourcePermissionsResponseBody {
	return s.Body
}

func (s *ListResourcePermissionsResponse) SetHeaders(v map[string]*string) *ListResourcePermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePermissionsResponse) SetStatusCode(v int32) *ListResourcePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePermissionsResponse) SetBody(v *ListResourcePermissionsResponseBody) *ListResourcePermissionsResponse {
	s.Body = v
	return s
}

func (s *ListResourcePermissionsResponse) Validate() error {
	return dara.Validate(s)
}
