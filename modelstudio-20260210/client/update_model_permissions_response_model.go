// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelPermissionsResponseBody) *UpdateModelPermissionsResponse
	GetBody() *UpdateModelPermissionsResponseBody
}

type UpdateModelPermissionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelPermissionsResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelPermissionsResponse) GetBody() *UpdateModelPermissionsResponseBody {
	return s.Body
}

func (s *UpdateModelPermissionsResponse) SetHeaders(v map[string]*string) *UpdateModelPermissionsResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelPermissionsResponse) SetStatusCode(v int32) *UpdateModelPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelPermissionsResponse) SetBody(v *UpdateModelPermissionsResponseBody) *UpdateModelPermissionsResponse {
	s.Body = v
	return s
}

func (s *UpdateModelPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
