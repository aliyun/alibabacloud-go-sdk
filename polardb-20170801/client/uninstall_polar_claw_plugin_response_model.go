// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPolarClawPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallPolarClawPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallPolarClawPluginResponse
	GetStatusCode() *int32
	SetBody(v *UninstallPolarClawPluginResponseBody) *UninstallPolarClawPluginResponse
	GetBody() *UninstallPolarClawPluginResponseBody
}

type UninstallPolarClawPluginResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallPolarClawPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallPolarClawPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallPolarClawPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallPolarClawPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallPolarClawPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallPolarClawPluginResponse) GetBody() *UninstallPolarClawPluginResponseBody {
	return s.Body
}

func (s *UninstallPolarClawPluginResponse) SetHeaders(v map[string]*string) *UninstallPolarClawPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallPolarClawPluginResponse) SetStatusCode(v int32) *UninstallPolarClawPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallPolarClawPluginResponse) SetBody(v *UninstallPolarClawPluginResponseBody) *UninstallPolarClawPluginResponse {
	s.Body = v
	return s
}

func (s *UninstallPolarClawPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
