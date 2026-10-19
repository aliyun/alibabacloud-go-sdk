// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResourcePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppResourcePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppResourcePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppResourcePermissionsResponseBody) *DeleteAppResourcePermissionsResponse
	GetBody() *DeleteAppResourcePermissionsResponseBody
}

type DeleteAppResourcePermissionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResourcePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResourcePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppResourcePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppResourcePermissionsResponse) GetBody() *DeleteAppResourcePermissionsResponseBody {
	return s.Body
}

func (s *DeleteAppResourcePermissionsResponse) SetHeaders(v map[string]*string) *DeleteAppResourcePermissionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResourcePermissionsResponse) SetStatusCode(v int32) *DeleteAppResourcePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponse) SetBody(v *DeleteAppResourcePermissionsResponseBody) *DeleteAppResourcePermissionsResponse {
	s.Body = v
	return s
}

func (s *DeleteAppResourcePermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
