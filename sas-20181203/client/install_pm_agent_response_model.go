// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPmAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallPmAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallPmAgentResponse
	GetStatusCode() *int32
	SetBody(v *InstallPmAgentResponseBody) *InstallPmAgentResponse
	GetBody() *InstallPmAgentResponseBody
}

type InstallPmAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallPmAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallPmAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallPmAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallPmAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallPmAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallPmAgentResponse) GetBody() *InstallPmAgentResponseBody {
	return s.Body
}

func (s *InstallPmAgentResponse) SetHeaders(v map[string]*string) *InstallPmAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallPmAgentResponse) SetStatusCode(v int32) *InstallPmAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallPmAgentResponse) SetBody(v *InstallPmAgentResponseBody) *InstallPmAgentResponse {
	s.Body = v
	return s
}

func (s *InstallPmAgentResponse) Validate() error {
	return dara.Validate(s)
}
