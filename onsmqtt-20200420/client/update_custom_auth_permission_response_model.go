// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAuthPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomAuthPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomAuthPermissionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomAuthPermissionResponseBody) *UpdateCustomAuthPermissionResponse
	GetBody() *UpdateCustomAuthPermissionResponseBody
}

type UpdateCustomAuthPermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomAuthPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomAuthPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAuthPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomAuthPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomAuthPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomAuthPermissionResponse) GetBody() *UpdateCustomAuthPermissionResponseBody {
	return s.Body
}

func (s *UpdateCustomAuthPermissionResponse) SetHeaders(v map[string]*string) *UpdateCustomAuthPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomAuthPermissionResponse) SetStatusCode(v int32) *UpdateCustomAuthPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomAuthPermissionResponse) SetBody(v *UpdateCustomAuthPermissionResponseBody) *UpdateCustomAuthPermissionResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomAuthPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
