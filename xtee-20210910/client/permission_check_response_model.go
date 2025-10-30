// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PermissionCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PermissionCheckResponse
	GetStatusCode() *int32
	SetBody(v *PermissionCheckResponseBody) *PermissionCheckResponse
	GetBody() *PermissionCheckResponseBody
}

type PermissionCheckResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PermissionCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PermissionCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s PermissionCheckResponse) GoString() string {
	return s.String()
}

func (s *PermissionCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PermissionCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PermissionCheckResponse) GetBody() *PermissionCheckResponseBody {
	return s.Body
}

func (s *PermissionCheckResponse) SetHeaders(v map[string]*string) *PermissionCheckResponse {
	s.Headers = v
	return s
}

func (s *PermissionCheckResponse) SetStatusCode(v int32) *PermissionCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *PermissionCheckResponse) SetBody(v *PermissionCheckResponseBody) *PermissionCheckResponse {
	s.Body = v
	return s
}

func (s *PermissionCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
