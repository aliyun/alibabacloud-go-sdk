// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudMonitorAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallCloudMonitorAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallCloudMonitorAgentResponse
	GetStatusCode() *int32
	SetBody(v *InstallCloudMonitorAgentResponseBody) *InstallCloudMonitorAgentResponse
	GetBody() *InstallCloudMonitorAgentResponseBody
}

type InstallCloudMonitorAgentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCloudMonitorAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCloudMonitorAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudMonitorAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallCloudMonitorAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallCloudMonitorAgentResponse) GetBody() *InstallCloudMonitorAgentResponseBody {
	return s.Body
}

func (s *InstallCloudMonitorAgentResponse) SetHeaders(v map[string]*string) *InstallCloudMonitorAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudMonitorAgentResponse) SetStatusCode(v int32) *InstallCloudMonitorAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudMonitorAgentResponse) SetBody(v *InstallCloudMonitorAgentResponseBody) *InstallCloudMonitorAgentResponse {
	s.Body = v
	return s
}

func (s *InstallCloudMonitorAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
