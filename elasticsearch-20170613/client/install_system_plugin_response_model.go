// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSystemPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallSystemPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallSystemPluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallSystemPluginResponseBody) *InstallSystemPluginResponse
	GetBody() *InstallSystemPluginResponseBody
}

type InstallSystemPluginResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallSystemPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallSystemPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallSystemPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallSystemPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallSystemPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallSystemPluginResponse) GetBody() *InstallSystemPluginResponseBody {
	return s.Body
}

func (s *InstallSystemPluginResponse) SetHeaders(v map[string]*string) *InstallSystemPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallSystemPluginResponse) SetStatusCode(v int32) *InstallSystemPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallSystemPluginResponse) SetBody(v *InstallSystemPluginResponseBody) *InstallSystemPluginResponse {
	s.Body = v
	return s
}

func (s *InstallSystemPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
