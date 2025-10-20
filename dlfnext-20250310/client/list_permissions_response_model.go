// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse
	GetBody() *ListPermissionsResponseBody
}

type ListPermissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionsResponse) GetBody() *ListPermissionsResponseBody {
	return s.Body
}

func (s *ListPermissionsResponse) SetHeaders(v map[string]*string) *ListPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionsResponse) SetStatusCode(v int32) *ListPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
