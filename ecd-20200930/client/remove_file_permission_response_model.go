// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFilePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveFilePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveFilePermissionResponse
	GetStatusCode() *int32
	SetBody(v *RemoveFilePermissionResponseBody) *RemoveFilePermissionResponse
	GetBody() *RemoveFilePermissionResponseBody
}

type RemoveFilePermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFilePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFilePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveFilePermissionResponse) GoString() string {
	return s.String()
}

func (s *RemoveFilePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveFilePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveFilePermissionResponse) GetBody() *RemoveFilePermissionResponseBody {
	return s.Body
}

func (s *RemoveFilePermissionResponse) SetHeaders(v map[string]*string) *RemoveFilePermissionResponse {
	s.Headers = v
	return s
}

func (s *RemoveFilePermissionResponse) SetStatusCode(v int32) *RemoveFilePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFilePermissionResponse) SetBody(v *RemoveFilePermissionResponseBody) *RemoveFilePermissionResponse {
	s.Body = v
	return s
}

func (s *RemoveFilePermissionResponse) Validate() error {
	return dara.Validate(s)
}
