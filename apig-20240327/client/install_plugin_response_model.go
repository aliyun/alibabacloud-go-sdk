// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallPluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallPluginResponseBody) *InstallPluginResponse
	GetBody() *InstallPluginResponseBody
}

type InstallPluginResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallPluginResponse) GetBody() *InstallPluginResponseBody {
	return s.Body
}

func (s *InstallPluginResponse) SetHeaders(v map[string]*string) *InstallPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallPluginResponse) SetStatusCode(v int32) *InstallPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallPluginResponse) SetBody(v *InstallPluginResponseBody) *InstallPluginResponse {
	s.Body = v
	return s
}

func (s *InstallPluginResponse) Validate() error {
	return dara.Validate(s)
}
