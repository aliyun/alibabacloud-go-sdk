// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNacosConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNacosConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetNacosConfigResponseBody) *GetNacosConfigResponse
	GetBody() *GetNacosConfigResponseBody
}

type GetNacosConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNacosConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNacosConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNacosConfigResponse) GetBody() *GetNacosConfigResponseBody {
	return s.Body
}

func (s *GetNacosConfigResponse) SetHeaders(v map[string]*string) *GetNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNacosConfigResponse) SetStatusCode(v int32) *GetNacosConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNacosConfigResponse) SetBody(v *GetNacosConfigResponseBody) *GetNacosConfigResponse {
	s.Body = v
	return s
}

func (s *GetNacosConfigResponse) Validate() error {
	return dara.Validate(s)
}
