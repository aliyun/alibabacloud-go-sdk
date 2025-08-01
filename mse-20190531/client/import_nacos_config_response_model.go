// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNacosConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportNacosConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportNacosConfigResponse
	GetStatusCode() *int32
	SetBody(v *ImportNacosConfigResponseBody) *ImportNacosConfigResponse
	GetBody() *ImportNacosConfigResponseBody
}

type ImportNacosConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportNacosConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportNacosConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportNacosConfigResponse) GetBody() *ImportNacosConfigResponseBody {
	return s.Body
}

func (s *ImportNacosConfigResponse) SetHeaders(v map[string]*string) *ImportNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *ImportNacosConfigResponse) SetStatusCode(v int32) *ImportNacosConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportNacosConfigResponse) SetBody(v *ImportNacosConfigResponseBody) *ImportNacosConfigResponse {
	s.Body = v
	return s
}

func (s *ImportNacosConfigResponse) Validate() error {
	return dara.Validate(s)
}
