// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRowPermissionByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRowPermissionByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *ListRowPermissionByUserIdResponseBody) *ListRowPermissionByUserIdResponse
	GetBody() *ListRowPermissionByUserIdResponseBody
}

type ListRowPermissionByUserIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRowPermissionByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRowPermissionByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRowPermissionByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRowPermissionByUserIdResponse) GetBody() *ListRowPermissionByUserIdResponseBody {
	return s.Body
}

func (s *ListRowPermissionByUserIdResponse) SetHeaders(v map[string]*string) *ListRowPermissionByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListRowPermissionByUserIdResponse) SetStatusCode(v int32) *ListRowPermissionByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRowPermissionByUserIdResponse) SetBody(v *ListRowPermissionByUserIdResponseBody) *ListRowPermissionByUserIdResponse {
	s.Body = v
	return s
}

func (s *ListRowPermissionByUserIdResponse) Validate() error {
	return dara.Validate(s)
}
