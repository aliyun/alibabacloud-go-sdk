// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallUniBackupAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallUniBackupAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallUniBackupAgentResponse
	GetStatusCode() *int32
	SetBody(v *InstallUniBackupAgentResponseBody) *InstallUniBackupAgentResponse
	GetBody() *InstallUniBackupAgentResponseBody
}

type InstallUniBackupAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallUniBackupAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallUniBackupAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallUniBackupAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallUniBackupAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallUniBackupAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallUniBackupAgentResponse) GetBody() *InstallUniBackupAgentResponseBody {
	return s.Body
}

func (s *InstallUniBackupAgentResponse) SetHeaders(v map[string]*string) *InstallUniBackupAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallUniBackupAgentResponse) SetStatusCode(v int32) *InstallUniBackupAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallUniBackupAgentResponse) SetBody(v *InstallUniBackupAgentResponseBody) *InstallUniBackupAgentResponse {
	s.Body = v
	return s
}

func (s *InstallUniBackupAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
