// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMGatewayMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLLMGatewayMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLLMGatewayMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ListLLMGatewayMetricsResponseBody) *ListLLMGatewayMetricsResponse
	GetBody() *ListLLMGatewayMetricsResponseBody
}

type ListLLMGatewayMetricsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLLMGatewayMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLLMGatewayMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLLMGatewayMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListLLMGatewayMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLLMGatewayMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLLMGatewayMetricsResponse) GetBody() *ListLLMGatewayMetricsResponseBody {
	return s.Body
}

func (s *ListLLMGatewayMetricsResponse) SetHeaders(v map[string]*string) *ListLLMGatewayMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListLLMGatewayMetricsResponse) SetStatusCode(v int32) *ListLLMGatewayMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLLMGatewayMetricsResponse) SetBody(v *ListLLMGatewayMetricsResponseBody) *ListLLMGatewayMetricsResponse {
	s.Body = v
	return s
}

func (s *ListLLMGatewayMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
