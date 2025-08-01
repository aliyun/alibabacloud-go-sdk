// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAgentResponse
	GetStatusCode() *int32
	SetBody(v *InstallAgentResponseBody) *InstallAgentResponse
	GetBody() *InstallAgentResponseBody
}

type InstallAgentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAgentResponse) GetBody() *InstallAgentResponseBody {
	return s.Body
}

func (s *InstallAgentResponse) SetHeaders(v map[string]*string) *InstallAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallAgentResponse) SetStatusCode(v int32) *InstallAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAgentResponse) SetBody(v *InstallAgentResponseBody) *InstallAgentResponse {
	s.Body = v
	return s
}

func (s *InstallAgentResponse) Validate() error {
	return dara.Validate(s)
}
