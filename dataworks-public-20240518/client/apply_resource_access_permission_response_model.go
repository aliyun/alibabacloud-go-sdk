// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourceAccessPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyResourceAccessPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyResourceAccessPermissionResponse
	GetStatusCode() *int32
	SetBody(v *ApplyResourceAccessPermissionResponseBody) *ApplyResourceAccessPermissionResponse
	GetBody() *ApplyResourceAccessPermissionResponseBody
}

type ApplyResourceAccessPermissionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyResourceAccessPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyResourceAccessPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionResponse) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyResourceAccessPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyResourceAccessPermissionResponse) GetBody() *ApplyResourceAccessPermissionResponseBody {
	return s.Body
}

func (s *ApplyResourceAccessPermissionResponse) SetHeaders(v map[string]*string) *ApplyResourceAccessPermissionResponse {
	s.Headers = v
	return s
}

func (s *ApplyResourceAccessPermissionResponse) SetStatusCode(v int32) *ApplyResourceAccessPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyResourceAccessPermissionResponse) SetBody(v *ApplyResourceAccessPermissionResponseBody) *ApplyResourceAccessPermissionResponse {
	s.Body = v
	return s
}

func (s *ApplyResourceAccessPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
