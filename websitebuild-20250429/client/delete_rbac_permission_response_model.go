// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRbacPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRbacPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRbacPermissionResponseBody) *DeleteRbacPermissionResponse
	GetBody() *DeleteRbacPermissionResponseBody
}

type DeleteRbacPermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRbacPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRbacPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacPermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRbacPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRbacPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRbacPermissionResponse) GetBody() *DeleteRbacPermissionResponseBody {
	return s.Body
}

func (s *DeleteRbacPermissionResponse) SetHeaders(v map[string]*string) *DeleteRbacPermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRbacPermissionResponse) SetStatusCode(v int32) *DeleteRbacPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRbacPermissionResponse) SetBody(v *DeleteRbacPermissionResponseBody) *DeleteRbacPermissionResponse {
	s.Body = v
	return s
}

func (s *DeleteRbacPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
