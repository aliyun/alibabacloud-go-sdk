// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileAddPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileAddPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileAddPermissionResponse
	GetStatusCode() *int32
}

type FileAddPermissionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s FileAddPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s FileAddPermissionResponse) GoString() string {
	return s.String()
}

func (s *FileAddPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileAddPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileAddPermissionResponse) SetHeaders(v map[string]*string) *FileAddPermissionResponse {
	s.Headers = v
	return s
}

func (s *FileAddPermissionResponse) SetStatusCode(v int32) *FileAddPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *FileAddPermissionResponse) Validate() error {
	return dara.Validate(s)
}
