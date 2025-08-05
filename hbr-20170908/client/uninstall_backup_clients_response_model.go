// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallBackupClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallBackupClientsResponse
	GetStatusCode() *int32
	SetBody(v *UninstallBackupClientsResponseBody) *UninstallBackupClientsResponse
	GetBody() *UninstallBackupClientsResponseBody
}

type UninstallBackupClientsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallBackupClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallBackupClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallBackupClientsResponse) GetBody() *UninstallBackupClientsResponseBody {
	return s.Body
}

func (s *UninstallBackupClientsResponse) SetHeaders(v map[string]*string) *UninstallBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *UninstallBackupClientsResponse) SetStatusCode(v int32) *UninstallBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallBackupClientsResponse) SetBody(v *UninstallBackupClientsResponseBody) *UninstallBackupClientsResponse {
	s.Body = v
	return s
}

func (s *UninstallBackupClientsResponse) Validate() error {
	return dara.Validate(s)
}
