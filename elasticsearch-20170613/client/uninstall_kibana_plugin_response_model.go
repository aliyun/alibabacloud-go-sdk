// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallKibanaPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallKibanaPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallKibanaPluginResponse
	GetStatusCode() *int32
	SetBody(v *UninstallKibanaPluginResponseBody) *UninstallKibanaPluginResponse
	GetBody() *UninstallKibanaPluginResponseBody
}

type UninstallKibanaPluginResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallKibanaPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallKibanaPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallKibanaPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallKibanaPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallKibanaPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallKibanaPluginResponse) GetBody() *UninstallKibanaPluginResponseBody {
	return s.Body
}

func (s *UninstallKibanaPluginResponse) SetHeaders(v map[string]*string) *UninstallKibanaPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallKibanaPluginResponse) SetStatusCode(v int32) *UninstallKibanaPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallKibanaPluginResponse) SetBody(v *UninstallKibanaPluginResponseBody) *UninstallKibanaPluginResponse {
	s.Body = v
	return s
}

func (s *UninstallKibanaPluginResponse) Validate() error {
	return dara.Validate(s)
}
