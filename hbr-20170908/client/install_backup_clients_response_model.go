// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallBackupClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallBackupClientsResponse
	GetStatusCode() *int32
	SetBody(v *InstallBackupClientsResponseBody) *InstallBackupClientsResponse
	GetBody() *InstallBackupClientsResponseBody
}

type InstallBackupClientsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallBackupClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallBackupClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallBackupClientsResponse) GetBody() *InstallBackupClientsResponseBody {
	return s.Body
}

func (s *InstallBackupClientsResponse) SetHeaders(v map[string]*string) *InstallBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *InstallBackupClientsResponse) SetStatusCode(v int32) *InstallBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallBackupClientsResponse) SetBody(v *InstallBackupClientsResponseBody) *InstallBackupClientsResponse {
	s.Body = v
	return s
}

func (s *InstallBackupClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
