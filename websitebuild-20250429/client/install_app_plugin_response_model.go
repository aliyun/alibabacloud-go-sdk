// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAppPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAppPluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallAppPluginResponseBody) *InstallAppPluginResponse
	GetBody() *InstallAppPluginResponseBody
}

type InstallAppPluginResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAppPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAppPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallAppPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAppPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAppPluginResponse) GetBody() *InstallAppPluginResponseBody {
	return s.Body
}

func (s *InstallAppPluginResponse) SetHeaders(v map[string]*string) *InstallAppPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallAppPluginResponse) SetStatusCode(v int32) *InstallAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAppPluginResponse) SetBody(v *InstallAppPluginResponseBody) *InstallAppPluginResponse {
	s.Body = v
	return s
}

func (s *InstallAppPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
