// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataPermissionsResponseBody) *ListDataPermissionsResponse
	GetBody() *ListDataPermissionsResponseBody
}

type ListDataPermissionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataPermissionsResponse) GetBody() *ListDataPermissionsResponseBody {
	return s.Body
}

func (s *ListDataPermissionsResponse) SetHeaders(v map[string]*string) *ListDataPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataPermissionsResponse) SetStatusCode(v int32) *ListDataPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataPermissionsResponse) SetBody(v *ListDataPermissionsResponseBody) *ListDataPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListDataPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
