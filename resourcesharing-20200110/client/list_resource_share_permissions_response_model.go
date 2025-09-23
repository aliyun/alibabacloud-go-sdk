// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceSharePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceSharePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceSharePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceSharePermissionsResponseBody) *ListResourceSharePermissionsResponse
	GetBody() *ListResourceSharePermissionsResponseBody
}

type ListResourceSharePermissionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceSharePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceSharePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharePermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceSharePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceSharePermissionsResponse) GetBody() *ListResourceSharePermissionsResponseBody {
	return s.Body
}

func (s *ListResourceSharePermissionsResponse) SetHeaders(v map[string]*string) *ListResourceSharePermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceSharePermissionsResponse) SetStatusCode(v int32) *ListResourceSharePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceSharePermissionsResponse) SetBody(v *ListResourceSharePermissionsResponseBody) *ListResourceSharePermissionsResponse {
	s.Body = v
	return s
}

func (s *ListResourceSharePermissionsResponse) Validate() error {
	return dara.Validate(s)
}
