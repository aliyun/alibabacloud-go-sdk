// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallBackupClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallBackupClientResponse
	GetStatusCode() *int32
	SetBody(v *InstallBackupClientResponseBody) *InstallBackupClientResponse
	GetBody() *InstallBackupClientResponseBody
}

type InstallBackupClientResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallBackupClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallBackupClientResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientResponse) GoString() string {
	return s.String()
}

func (s *InstallBackupClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallBackupClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallBackupClientResponse) GetBody() *InstallBackupClientResponseBody {
	return s.Body
}

func (s *InstallBackupClientResponse) SetHeaders(v map[string]*string) *InstallBackupClientResponse {
	s.Headers = v
	return s
}

func (s *InstallBackupClientResponse) SetStatusCode(v int32) *InstallBackupClientResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallBackupClientResponse) SetBody(v *InstallBackupClientResponseBody) *InstallBackupClientResponse {
	s.Body = v
	return s
}

func (s *InstallBackupClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
