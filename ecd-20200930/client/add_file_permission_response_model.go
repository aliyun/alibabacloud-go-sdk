// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFilePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFilePermissionResponse
	GetStatusCode() *int32
	SetBody(v *AddFilePermissionResponseBody) *AddFilePermissionResponse
	GetBody() *AddFilePermissionResponseBody
}

type AddFilePermissionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFilePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFilePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFilePermissionResponse) GoString() string {
	return s.String()
}

func (s *AddFilePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFilePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFilePermissionResponse) GetBody() *AddFilePermissionResponseBody {
	return s.Body
}

func (s *AddFilePermissionResponse) SetHeaders(v map[string]*string) *AddFilePermissionResponse {
	s.Headers = v
	return s
}

func (s *AddFilePermissionResponse) SetStatusCode(v int32) *AddFilePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFilePermissionResponse) SetBody(v *AddFilePermissionResponseBody) *AddFilePermissionResponse {
	s.Body = v
	return s
}

func (s *AddFilePermissionResponse) Validate() error {
	return dara.Validate(s)
}
