// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPermissionResponse
	GetStatusCode() *int32
	SetBody(v *GetPermissionResponseBody) *GetPermissionResponse
	GetBody() *GetPermissionResponseBody
}

type GetPermissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPermissionResponse) GetBody() *GetPermissionResponseBody {
	return s.Body
}

func (s *GetPermissionResponse) SetHeaders(v map[string]*string) *GetPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionResponse) SetStatusCode(v int32) *GetPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionResponse) SetBody(v *GetPermissionResponseBody) *GetPermissionResponse {
	s.Body = v
	return s
}

func (s *GetPermissionResponse) Validate() error {
	return dara.Validate(s)
}
