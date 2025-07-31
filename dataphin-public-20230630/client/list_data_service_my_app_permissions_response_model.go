// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyAppPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceMyAppPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceMyAppPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceMyAppPermissionsResponseBody) *ListDataServiceMyAppPermissionsResponse
	GetBody() *ListDataServiceMyAppPermissionsResponseBody
}

type ListDataServiceMyAppPermissionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceMyAppPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceMyAppPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceMyAppPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceMyAppPermissionsResponse) GetBody() *ListDataServiceMyAppPermissionsResponseBody {
	return s.Body
}

func (s *ListDataServiceMyAppPermissionsResponse) SetHeaders(v map[string]*string) *ListDataServiceMyAppPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponse) SetStatusCode(v int32) *ListDataServiceMyAppPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponse) SetBody(v *ListDataServiceMyAppPermissionsResponseBody) *ListDataServiceMyAppPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
