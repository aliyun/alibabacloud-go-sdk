// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePermissionResponse
	GetStatusCode() *int32
	SetBody(v *DeletePermissionResponseBody) *DeletePermissionResponse
	GetBody() *DeletePermissionResponseBody
}

type DeletePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionResponse) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePermissionResponse) GetBody() *DeletePermissionResponseBody {
	return s.Body
}

func (s *DeletePermissionResponse) SetHeaders(v map[string]*string) *DeletePermissionResponse {
	s.Headers = v
	return s
}

func (s *DeletePermissionResponse) SetStatusCode(v int32) *DeletePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePermissionResponse) SetBody(v *DeletePermissionResponseBody) *DeletePermissionResponse {
	s.Body = v
	return s
}

func (s *DeletePermissionResponse) Validate() error {
	return dara.Validate(s)
}
