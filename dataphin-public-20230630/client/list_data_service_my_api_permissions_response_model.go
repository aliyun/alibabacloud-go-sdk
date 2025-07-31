// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyApiPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceMyApiPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceMyApiPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceMyApiPermissionsResponseBody) *ListDataServiceMyApiPermissionsResponse
	GetBody() *ListDataServiceMyApiPermissionsResponseBody
}

type ListDataServiceMyApiPermissionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceMyApiPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceMyApiPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceMyApiPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceMyApiPermissionsResponse) GetBody() *ListDataServiceMyApiPermissionsResponseBody {
	return s.Body
}

func (s *ListDataServiceMyApiPermissionsResponse) SetHeaders(v map[string]*string) *ListDataServiceMyApiPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponse) SetStatusCode(v int32) *ListDataServiceMyApiPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponse) SetBody(v *ListDataServiceMyApiPermissionsResponseBody) *ListDataServiceMyApiPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
