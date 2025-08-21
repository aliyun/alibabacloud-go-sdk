// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallLogstashPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallLogstashPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallLogstashPluginResponse
	GetStatusCode() *int32
	SetBody(v *UninstallLogstashPluginResponseBody) *UninstallLogstashPluginResponse
	GetBody() *UninstallLogstashPluginResponseBody
}

type UninstallLogstashPluginResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallLogstashPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallLogstashPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallLogstashPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallLogstashPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallLogstashPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallLogstashPluginResponse) GetBody() *UninstallLogstashPluginResponseBody {
	return s.Body
}

func (s *UninstallLogstashPluginResponse) SetHeaders(v map[string]*string) *UninstallLogstashPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallLogstashPluginResponse) SetStatusCode(v int32) *UninstallLogstashPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallLogstashPluginResponse) SetBody(v *UninstallLogstashPluginResponseBody) *UninstallLogstashPluginResponse {
	s.Body = v
	return s
}

func (s *UninstallLogstashPluginResponse) Validate() error {
	return dara.Validate(s)
}
