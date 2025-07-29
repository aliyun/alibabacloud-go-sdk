// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CleanUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CleanUserPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *CleanUserPermissionsResponseBody) *CleanUserPermissionsResponse
	GetBody() *CleanUserPermissionsResponseBody
}

type CleanUserPermissionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CleanUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CleanUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CleanUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CleanUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CleanUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CleanUserPermissionsResponse) GetBody() *CleanUserPermissionsResponseBody {
	return s.Body
}

func (s *CleanUserPermissionsResponse) SetHeaders(v map[string]*string) *CleanUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CleanUserPermissionsResponse) SetStatusCode(v int32) *CleanUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CleanUserPermissionsResponse) SetBody(v *CleanUserPermissionsResponseBody) *CleanUserPermissionsResponse {
	s.Body = v
	return s
}

func (s *CleanUserPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
