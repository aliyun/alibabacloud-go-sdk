// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
	GetStatusCode() *int32
	SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
	GetBody() *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) GetBody() *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	return s.Body
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.Body = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
