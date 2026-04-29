// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallPolarClawPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallPolarClawPluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallPolarClawPluginResponseBody) *InstallPolarClawPluginResponse
	GetBody() *InstallPolarClawPluginResponseBody
}

type InstallPolarClawPluginResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallPolarClawPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallPolarClawPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallPolarClawPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallPolarClawPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallPolarClawPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallPolarClawPluginResponse) GetBody() *InstallPolarClawPluginResponseBody {
	return s.Body
}

func (s *InstallPolarClawPluginResponse) SetHeaders(v map[string]*string) *InstallPolarClawPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallPolarClawPluginResponse) SetStatusCode(v int32) *InstallPolarClawPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallPolarClawPluginResponse) SetBody(v *InstallPolarClawPluginResponseBody) *InstallPolarClawPluginResponse {
	s.Body = v
	return s
}

func (s *InstallPolarClawPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
