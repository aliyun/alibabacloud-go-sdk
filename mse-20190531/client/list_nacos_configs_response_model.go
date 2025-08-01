// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNacosConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNacosConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListNacosConfigsResponseBody) *ListNacosConfigsResponse
	GetBody() *ListNacosConfigsResponseBody
}

type ListNacosConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNacosConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNacosConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNacosConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNacosConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNacosConfigsResponse) GetBody() *ListNacosConfigsResponseBody {
	return s.Body
}

func (s *ListNacosConfigsResponse) SetHeaders(v map[string]*string) *ListNacosConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListNacosConfigsResponse) SetStatusCode(v int32) *ListNacosConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNacosConfigsResponse) SetBody(v *ListNacosConfigsResponseBody) *ListNacosConfigsResponse {
	s.Body = v
	return s
}

func (s *ListNacosConfigsResponse) Validate() error {
	return dara.Validate(s)
}
