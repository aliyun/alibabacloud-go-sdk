// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryPaidResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryPaidResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponse
	GetStatusCode() *int32
	SetBody(v *AliyunConsoleOpenApiQueryPaidResourceResponseBody) *AliyunConsoleOpenApiQueryPaidResourceResponse
	GetBody() *AliyunConsoleOpenApiQueryPaidResourceResponseBody
}

type AliyunConsoleOpenApiQueryPaidResourceResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryPaidResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryPaidResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryPaidResourceResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) GetBody() *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	return s.Body
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryPaidResourceResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) SetBody(v *AliyunConsoleOpenApiQueryPaidResourceResponseBody) *AliyunConsoleOpenApiQueryPaidResourceResponse {
	s.Body = v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
