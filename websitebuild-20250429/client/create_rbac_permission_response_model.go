// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRbacPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRbacPermissionResponse
	GetStatusCode() *int32
	SetBody(v *CreateRbacPermissionResponseBody) *CreateRbacPermissionResponse
	GetBody() *CreateRbacPermissionResponseBody
}

type CreateRbacPermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRbacPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRbacPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacPermissionResponse) GoString() string {
	return s.String()
}

func (s *CreateRbacPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRbacPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRbacPermissionResponse) GetBody() *CreateRbacPermissionResponseBody {
	return s.Body
}

func (s *CreateRbacPermissionResponse) SetHeaders(v map[string]*string) *CreateRbacPermissionResponse {
	s.Headers = v
	return s
}

func (s *CreateRbacPermissionResponse) SetStatusCode(v int32) *CreateRbacPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRbacPermissionResponse) SetBody(v *CreateRbacPermissionResponseBody) *CreateRbacPermissionResponse {
	s.Body = v
	return s
}

func (s *CreateRbacPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
