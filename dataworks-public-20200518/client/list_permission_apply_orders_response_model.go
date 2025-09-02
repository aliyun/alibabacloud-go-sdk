// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionApplyOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionApplyOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionApplyOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ListPermissionApplyOrdersResponseBody) *ListPermissionApplyOrdersResponse
	GetBody() *ListPermissionApplyOrdersResponseBody
}

type ListPermissionApplyOrdersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionApplyOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionApplyOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionApplyOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionApplyOrdersResponse) GetBody() *ListPermissionApplyOrdersResponseBody {
	return s.Body
}

func (s *ListPermissionApplyOrdersResponse) SetHeaders(v map[string]*string) *ListPermissionApplyOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionApplyOrdersResponse) SetStatusCode(v int32) *ListPermissionApplyOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionApplyOrdersResponse) SetBody(v *ListPermissionApplyOrdersResponseBody) *ListPermissionApplyOrdersResponse {
	s.Body = v
	return s
}

func (s *ListPermissionApplyOrdersResponse) Validate() error {
	return dara.Validate(s)
}
