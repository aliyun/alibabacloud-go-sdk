// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenSourcePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpenSourcePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpenSourcePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListOpenSourcePermissionsResponseBody) *ListOpenSourcePermissionsResponse
	GetBody() *ListOpenSourcePermissionsResponseBody
}

type ListOpenSourcePermissionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpenSourcePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpenSourcePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourcePermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListOpenSourcePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpenSourcePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpenSourcePermissionsResponse) GetBody() *ListOpenSourcePermissionsResponseBody {
	return s.Body
}

func (s *ListOpenSourcePermissionsResponse) SetHeaders(v map[string]*string) *ListOpenSourcePermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListOpenSourcePermissionsResponse) SetStatusCode(v int32) *ListOpenSourcePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenSourcePermissionsResponse) SetBody(v *ListOpenSourcePermissionsResponseBody) *ListOpenSourcePermissionsResponse {
	s.Body = v
	return s
}

func (s *ListOpenSourcePermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
