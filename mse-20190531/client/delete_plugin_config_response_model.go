// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePluginConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePluginConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeletePluginConfigResponseBody) *DeletePluginConfigResponse
	GetBody() *DeletePluginConfigResponseBody
}

type DeletePluginConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePluginConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginConfigResponse) GoString() string {
	return s.String()
}

func (s *DeletePluginConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePluginConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePluginConfigResponse) GetBody() *DeletePluginConfigResponseBody {
	return s.Body
}

func (s *DeletePluginConfigResponse) SetHeaders(v map[string]*string) *DeletePluginConfigResponse {
	s.Headers = v
	return s
}

func (s *DeletePluginConfigResponse) SetStatusCode(v int32) *DeletePluginConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePluginConfigResponse) SetBody(v *DeletePluginConfigResponseBody) *DeletePluginConfigResponse {
	s.Body = v
	return s
}

func (s *DeletePluginConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
