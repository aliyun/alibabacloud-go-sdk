// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallAppPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallAppPluginResponse
	GetStatusCode() *int32
	SetBody(v *UninstallAppPluginResponseBody) *UninstallAppPluginResponse
	GetBody() *UninstallAppPluginResponseBody
}

type UninstallAppPluginResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAppPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallAppPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallAppPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallAppPluginResponse) GetBody() *UninstallAppPluginResponseBody {
	return s.Body
}

func (s *UninstallAppPluginResponse) SetHeaders(v map[string]*string) *UninstallAppPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallAppPluginResponse) SetStatusCode(v int32) *UninstallAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAppPluginResponse) SetBody(v *UninstallAppPluginResponseBody) *UninstallAppPluginResponse {
	s.Body = v
	return s
}

func (s *UninstallAppPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
