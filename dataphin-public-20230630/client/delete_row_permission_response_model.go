// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRowPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRowPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRowPermissionResponseBody) *DeleteRowPermissionResponse
	GetBody() *DeleteRowPermissionResponseBody
}

type DeleteRowPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRowPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRowPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowPermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRowPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRowPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRowPermissionResponse) GetBody() *DeleteRowPermissionResponseBody {
	return s.Body
}

func (s *DeleteRowPermissionResponse) SetHeaders(v map[string]*string) *DeleteRowPermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRowPermissionResponse) SetStatusCode(v int32) *DeleteRowPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRowPermissionResponse) SetBody(v *DeleteRowPermissionResponseBody) *DeleteRowPermissionResponse {
	s.Body = v
	return s
}

func (s *DeleteRowPermissionResponse) Validate() error {
	return dara.Validate(s)
}
