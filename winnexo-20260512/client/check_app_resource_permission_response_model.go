// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAppResourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAppResourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *CheckAppResourcePermissionResponseBody) *CheckAppResourcePermissionResponse
	GetBody() *CheckAppResourcePermissionResponseBody
}

type CheckAppResourcePermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAppResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAppResourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAppResourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAppResourcePermissionResponse) GetBody() *CheckAppResourcePermissionResponseBody {
	return s.Body
}

func (s *CheckAppResourcePermissionResponse) SetHeaders(v map[string]*string) *CheckAppResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *CheckAppResourcePermissionResponse) SetStatusCode(v int32) *CheckAppResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAppResourcePermissionResponse) SetBody(v *CheckAppResourcePermissionResponseBody) *CheckAppResourcePermissionResponse {
	s.Body = v
	return s
}

func (s *CheckAppResourcePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
