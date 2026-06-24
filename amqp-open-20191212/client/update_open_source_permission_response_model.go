// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpenSourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOpenSourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOpenSourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOpenSourcePermissionResponseBody) *UpdateOpenSourcePermissionResponse
	GetBody() *UpdateOpenSourcePermissionResponseBody
}

type UpdateOpenSourcePermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOpenSourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOpenSourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpenSourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOpenSourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOpenSourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOpenSourcePermissionResponse) GetBody() *UpdateOpenSourcePermissionResponseBody {
	return s.Body
}

func (s *UpdateOpenSourcePermissionResponse) SetHeaders(v map[string]*string) *UpdateOpenSourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOpenSourcePermissionResponse) SetStatusCode(v int32) *UpdateOpenSourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOpenSourcePermissionResponse) SetBody(v *UpdateOpenSourcePermissionResponseBody) *UpdateOpenSourcePermissionResponse {
	s.Body = v
	return s
}

func (s *UpdateOpenSourcePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
