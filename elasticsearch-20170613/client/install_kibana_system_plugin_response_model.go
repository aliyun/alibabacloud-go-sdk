// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallKibanaSystemPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallKibanaSystemPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallKibanaSystemPluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallKibanaSystemPluginResponseBody) *InstallKibanaSystemPluginResponse
	GetBody() *InstallKibanaSystemPluginResponseBody
}

type InstallKibanaSystemPluginResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallKibanaSystemPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallKibanaSystemPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallKibanaSystemPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallKibanaSystemPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallKibanaSystemPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallKibanaSystemPluginResponse) GetBody() *InstallKibanaSystemPluginResponseBody {
	return s.Body
}

func (s *InstallKibanaSystemPluginResponse) SetHeaders(v map[string]*string) *InstallKibanaSystemPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallKibanaSystemPluginResponse) SetStatusCode(v int32) *InstallKibanaSystemPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallKibanaSystemPluginResponse) SetBody(v *InstallKibanaSystemPluginResponseBody) *InstallKibanaSystemPluginResponse {
	s.Body = v
	return s
}

func (s *InstallKibanaSystemPluginResponse) Validate() error {
	return dara.Validate(s)
}
