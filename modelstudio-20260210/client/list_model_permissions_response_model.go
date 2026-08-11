// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelPermissionsResponseBody) *ListModelPermissionsResponse
	GetBody() *ListModelPermissionsResponseBody
}

type ListModelPermissionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListModelPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelPermissionsResponse) GetBody() *ListModelPermissionsResponseBody {
	return s.Body
}

func (s *ListModelPermissionsResponse) SetHeaders(v map[string]*string) *ListModelPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListModelPermissionsResponse) SetStatusCode(v int32) *ListModelPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelPermissionsResponse) SetBody(v *ListModelPermissionsResponseBody) *ListModelPermissionsResponse {
	s.Body = v
	return s
}

func (s *ListModelPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
