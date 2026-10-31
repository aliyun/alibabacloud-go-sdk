// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyResourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyResourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *ApplyResourcePermissionResponseBody) *ApplyResourcePermissionResponse
	GetBody() *ApplyResourcePermissionResponseBody
}

type ApplyResourcePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyResourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyResourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyResourcePermissionResponse) GetBody() *ApplyResourcePermissionResponseBody {
	return s.Body
}

func (s *ApplyResourcePermissionResponse) SetHeaders(v map[string]*string) *ApplyResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *ApplyResourcePermissionResponse) SetStatusCode(v int32) *ApplyResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyResourcePermissionResponse) SetBody(v *ApplyResourcePermissionResponseBody) *ApplyResourcePermissionResponse {
	s.Body = v
	return s
}

func (s *ApplyResourcePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
