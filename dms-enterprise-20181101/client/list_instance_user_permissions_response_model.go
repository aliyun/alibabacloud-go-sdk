// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceUserPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceUserPermissionsResponseBody) *ListInstanceUserPermissionsResponse
	GetBody() *ListInstanceUserPermissionsResponseBody
}

type ListInstanceUserPermissionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceUserPermissionsResponse) GetBody() *ListInstanceUserPermissionsResponseBody {
	return s.Body
}

func (s *ListInstanceUserPermissionsResponse) SetHeaders(v map[string]*string) *ListInstanceUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceUserPermissionsResponse) SetStatusCode(v int32) *ListInstanceUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceUserPermissionsResponse) SetBody(v *ListInstanceUserPermissionsResponseBody) *ListInstanceUserPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceUserPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
