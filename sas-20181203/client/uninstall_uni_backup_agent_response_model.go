// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallUniBackupAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallUniBackupAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallUniBackupAgentResponse
	GetStatusCode() *int32
	SetBody(v *UninstallUniBackupAgentResponseBody) *UninstallUniBackupAgentResponse
	GetBody() *UninstallUniBackupAgentResponseBody
}

type UninstallUniBackupAgentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallUniBackupAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallUniBackupAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallUniBackupAgentResponse) GoString() string {
	return s.String()
}

func (s *UninstallUniBackupAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallUniBackupAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallUniBackupAgentResponse) GetBody() *UninstallUniBackupAgentResponseBody {
	return s.Body
}

func (s *UninstallUniBackupAgentResponse) SetHeaders(v map[string]*string) *UninstallUniBackupAgentResponse {
	s.Headers = v
	return s
}

func (s *UninstallUniBackupAgentResponse) SetStatusCode(v int32) *UninstallUniBackupAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallUniBackupAgentResponse) SetBody(v *UninstallUniBackupAgentResponseBody) *UninstallUniBackupAgentResponse {
	s.Body = v
	return s
}

func (s *UninstallUniBackupAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
