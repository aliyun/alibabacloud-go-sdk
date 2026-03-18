// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
	GetStatusCode() *int32
	SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
	GetBody() *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) GetBody() *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	return s.Body
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.Body = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
