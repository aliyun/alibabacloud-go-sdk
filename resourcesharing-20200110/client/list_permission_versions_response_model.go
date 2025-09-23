// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListPermissionVersionsResponseBody) *ListPermissionVersionsResponse
	GetBody() *ListPermissionVersionsResponseBody
}

type ListPermissionVersionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionVersionsResponse) GetBody() *ListPermissionVersionsResponseBody {
	return s.Body
}

func (s *ListPermissionVersionsResponse) SetHeaders(v map[string]*string) *ListPermissionVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionVersionsResponse) SetStatusCode(v int32) *ListPermissionVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionVersionsResponse) SetBody(v *ListPermissionVersionsResponseBody) *ListPermissionVersionsResponse {
	s.Body = v
	return s
}

func (s *ListPermissionVersionsResponse) Validate() error {
	return dara.Validate(s)
}
