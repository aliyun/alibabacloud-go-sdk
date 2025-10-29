// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallPluginResponse
	GetStatusCode() *int32
	SetBody(v *UninstallPluginResponseBody) *UninstallPluginResponse
	GetBody() *UninstallPluginResponseBody
}

type UninstallPluginResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallPluginResponse) GetBody() *UninstallPluginResponseBody {
	return s.Body
}

func (s *UninstallPluginResponse) SetHeaders(v map[string]*string) *UninstallPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallPluginResponse) SetStatusCode(v int32) *UninstallPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallPluginResponse) SetBody(v *UninstallPluginResponseBody) *UninstallPluginResponse {
	s.Body = v
	return s
}

func (s *UninstallPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
