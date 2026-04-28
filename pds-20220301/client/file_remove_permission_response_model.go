// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileRemovePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileRemovePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileRemovePermissionResponse
	GetStatusCode() *int32
}

type FileRemovePermissionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s FileRemovePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s FileRemovePermissionResponse) GoString() string {
	return s.String()
}

func (s *FileRemovePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileRemovePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileRemovePermissionResponse) SetHeaders(v map[string]*string) *FileRemovePermissionResponse {
	s.Headers = v
	return s
}

func (s *FileRemovePermissionResponse) SetStatusCode(v int32) *FileRemovePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *FileRemovePermissionResponse) Validate() error {
	return dara.Validate(s)
}
