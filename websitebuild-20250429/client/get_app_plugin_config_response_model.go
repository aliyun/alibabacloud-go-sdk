// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppPluginConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppPluginConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAppPluginConfigResponseBody) *GetAppPluginConfigResponse
	GetBody() *GetAppPluginConfigResponseBody
}

type GetAppPluginConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppPluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppPluginConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppPluginConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppPluginConfigResponse) GetBody() *GetAppPluginConfigResponseBody {
	return s.Body
}

func (s *GetAppPluginConfigResponse) SetHeaders(v map[string]*string) *GetAppPluginConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAppPluginConfigResponse) SetStatusCode(v int32) *GetAppPluginConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppPluginConfigResponse) SetBody(v *GetAppPluginConfigResponseBody) *GetAppPluginConfigResponse {
	s.Body = v
	return s
}

func (s *GetAppPluginConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
