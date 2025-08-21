// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallUserPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallUserPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallUserPluginsResponse
	GetStatusCode() *int32
	SetBody(v *InstallUserPluginsResponseBody) *InstallUserPluginsResponse
	GetBody() *InstallUserPluginsResponseBody
}

type InstallUserPluginsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallUserPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallUserPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallUserPluginsResponse) GoString() string {
	return s.String()
}

func (s *InstallUserPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallUserPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallUserPluginsResponse) GetBody() *InstallUserPluginsResponseBody {
	return s.Body
}

func (s *InstallUserPluginsResponse) SetHeaders(v map[string]*string) *InstallUserPluginsResponse {
	s.Headers = v
	return s
}

func (s *InstallUserPluginsResponse) SetStatusCode(v int32) *InstallUserPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallUserPluginsResponse) SetBody(v *InstallUserPluginsResponseBody) *InstallUserPluginsResponse {
	s.Body = v
	return s
}

func (s *InstallUserPluginsResponse) Validate() error {
	return dara.Validate(s)
}
