// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNacosGrayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryNacosGrayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryNacosGrayConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryNacosGrayConfigResponseBody) *QueryNacosGrayConfigResponse
	GetBody() *QueryNacosGrayConfigResponseBody
}

type QueryNacosGrayConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNacosGrayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNacosGrayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryNacosGrayConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryNacosGrayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryNacosGrayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryNacosGrayConfigResponse) GetBody() *QueryNacosGrayConfigResponseBody {
	return s.Body
}

func (s *QueryNacosGrayConfigResponse) SetHeaders(v map[string]*string) *QueryNacosGrayConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryNacosGrayConfigResponse) SetStatusCode(v int32) *QueryNacosGrayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNacosGrayConfigResponse) SetBody(v *QueryNacosGrayConfigResponseBody) *QueryNacosGrayConfigResponse {
	s.Body = v
	return s
}

func (s *QueryNacosGrayConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
