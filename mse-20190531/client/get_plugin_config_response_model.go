// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPluginConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPluginConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPluginConfigResponseBody) *GetPluginConfigResponse
	GetBody() *GetPluginConfigResponseBody
}

type GetPluginConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPluginConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPluginConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPluginConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPluginConfigResponse) GetBody() *GetPluginConfigResponseBody {
	return s.Body
}

func (s *GetPluginConfigResponse) SetHeaders(v map[string]*string) *GetPluginConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPluginConfigResponse) SetStatusCode(v int32) *GetPluginConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginConfigResponse) SetBody(v *GetPluginConfigResponseBody) *GetPluginConfigResponse {
	s.Body = v
	return s
}

func (s *GetPluginConfigResponse) Validate() error {
	return dara.Validate(s)
}
