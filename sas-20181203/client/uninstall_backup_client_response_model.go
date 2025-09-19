// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallBackupClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallBackupClientResponse
	GetStatusCode() *int32
	SetBody(v *UninstallBackupClientResponseBody) *UninstallBackupClientResponse
	GetBody() *UninstallBackupClientResponseBody
}

type UninstallBackupClientResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallBackupClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallBackupClientResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientResponse) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallBackupClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallBackupClientResponse) GetBody() *UninstallBackupClientResponseBody {
	return s.Body
}

func (s *UninstallBackupClientResponse) SetHeaders(v map[string]*string) *UninstallBackupClientResponse {
	s.Headers = v
	return s
}

func (s *UninstallBackupClientResponse) SetStatusCode(v int32) *UninstallBackupClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallBackupClientResponse) SetBody(v *UninstallBackupClientResponseBody) *UninstallBackupClientResponse {
	s.Body = v
	return s
}

func (s *UninstallBackupClientResponse) Validate() error {
	return dara.Validate(s)
}
