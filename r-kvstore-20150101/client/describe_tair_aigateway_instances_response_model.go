// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairAIGatewayInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairAIGatewayInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairAIGatewayInstancesResponseBody) *DescribeTairAIGatewayInstancesResponse
	GetBody() *DescribeTairAIGatewayInstancesResponseBody
}

type DescribeTairAIGatewayInstancesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairAIGatewayInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairAIGatewayInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairAIGatewayInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairAIGatewayInstancesResponse) GetBody() *DescribeTairAIGatewayInstancesResponseBody {
	return s.Body
}

func (s *DescribeTairAIGatewayInstancesResponse) SetHeaders(v map[string]*string) *DescribeTairAIGatewayInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponse) SetStatusCode(v int32) *DescribeTairAIGatewayInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponse) SetBody(v *DescribeTairAIGatewayInstancesResponseBody) *DescribeTairAIGatewayInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
