// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserTenantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserTenantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserTenantsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserTenantsResponseBody) *ListUserTenantsResponse
	GetBody() *ListUserTenantsResponseBody
}

type ListUserTenantsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserTenantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserTenantsResponse) GoString() string {
	return s.String()
}

func (s *ListUserTenantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserTenantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserTenantsResponse) GetBody() *ListUserTenantsResponseBody {
	return s.Body
}

func (s *ListUserTenantsResponse) SetHeaders(v map[string]*string) *ListUserTenantsResponse {
	s.Headers = v
	return s
}

func (s *ListUserTenantsResponse) SetStatusCode(v int32) *ListUserTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserTenantsResponse) SetBody(v *ListUserTenantsResponseBody) *ListUserTenantsResponse {
	s.Body = v
	return s
}

func (s *ListUserTenantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
