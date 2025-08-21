// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallLogstashSystemPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallLogstashSystemPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallLogstashSystemPluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallLogstashSystemPluginResponseBody) *InstallLogstashSystemPluginResponse
	GetBody() *InstallLogstashSystemPluginResponseBody
}

type InstallLogstashSystemPluginResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallLogstashSystemPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallLogstashSystemPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallLogstashSystemPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallLogstashSystemPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallLogstashSystemPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallLogstashSystemPluginResponse) GetBody() *InstallLogstashSystemPluginResponseBody {
	return s.Body
}

func (s *InstallLogstashSystemPluginResponse) SetHeaders(v map[string]*string) *InstallLogstashSystemPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallLogstashSystemPluginResponse) SetStatusCode(v int32) *InstallLogstashSystemPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallLogstashSystemPluginResponse) SetBody(v *InstallLogstashSystemPluginResponseBody) *InstallLogstashSystemPluginResponse {
	s.Body = v
	return s
}

func (s *InstallLogstashSystemPluginResponse) Validate() error {
	return dara.Validate(s)
}
