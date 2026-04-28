// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileListPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileListPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileListPermissionResponse
	GetStatusCode() *int32
	SetBody(v []*FilePermissionMember) *FileListPermissionResponse
	GetBody() []*FilePermissionMember
}

type FileListPermissionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*FilePermissionMember `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s FileListPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s FileListPermissionResponse) GoString() string {
	return s.String()
}

func (s *FileListPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileListPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileListPermissionResponse) GetBody() []*FilePermissionMember {
	return s.Body
}

func (s *FileListPermissionResponse) SetHeaders(v map[string]*string) *FileListPermissionResponse {
	s.Headers = v
	return s
}

func (s *FileListPermissionResponse) SetStatusCode(v int32) *FileListPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *FileListPermissionResponse) SetBody(v []*FilePermissionMember) *FileListPermissionResponse {
	s.Body = v
	return s
}

func (s *FileListPermissionResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
