// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasPermissionGroupInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSaasPermissionGroupInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSaasPermissionGroupInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListSaasPermissionGroupInfosResponseBody) *ListSaasPermissionGroupInfosResponse
	GetBody() *ListSaasPermissionGroupInfosResponseBody
}

type ListSaasPermissionGroupInfosResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSaasPermissionGroupInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponse) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSaasPermissionGroupInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSaasPermissionGroupInfosResponse) GetBody() *ListSaasPermissionGroupInfosResponseBody {
	return s.Body
}

func (s *ListSaasPermissionGroupInfosResponse) SetHeaders(v map[string]*string) *ListSaasPermissionGroupInfosResponse {
	s.Headers = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponse) SetStatusCode(v int32) *ListSaasPermissionGroupInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponse) SetBody(v *ListSaasPermissionGroupInfosResponseBody) *ListSaasPermissionGroupInfosResponse {
	s.Body = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponse) Validate() error {
	return dara.Validate(s)
}
