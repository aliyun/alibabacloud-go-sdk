// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentWithTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAgentWithTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAgentWithTypeResponse
	GetStatusCode() *int32
	SetBody(v *InstallAgentWithTypeResponseBody) *InstallAgentWithTypeResponse
	GetBody() *InstallAgentWithTypeResponseBody
}

type InstallAgentWithTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAgentWithTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAgentWithTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentWithTypeResponse) GoString() string {
	return s.String()
}

func (s *InstallAgentWithTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAgentWithTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAgentWithTypeResponse) GetBody() *InstallAgentWithTypeResponseBody {
	return s.Body
}

func (s *InstallAgentWithTypeResponse) SetHeaders(v map[string]*string) *InstallAgentWithTypeResponse {
	s.Headers = v
	return s
}

func (s *InstallAgentWithTypeResponse) SetStatusCode(v int32) *InstallAgentWithTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAgentWithTypeResponse) SetBody(v *InstallAgentWithTypeResponseBody) *InstallAgentWithTypeResponse {
	s.Body = v
	return s
}

func (s *InstallAgentWithTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
