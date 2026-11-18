// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairAIGatewayParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairAIGatewayParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairAIGatewayParametersResponseBody) *DescribeTairAIGatewayParametersResponse
	GetBody() *DescribeTairAIGatewayParametersResponseBody
}

type DescribeTairAIGatewayParametersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairAIGatewayParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairAIGatewayParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairAIGatewayParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairAIGatewayParametersResponse) GetBody() *DescribeTairAIGatewayParametersResponseBody {
	return s.Body
}

func (s *DescribeTairAIGatewayParametersResponse) SetHeaders(v map[string]*string) *DescribeTairAIGatewayParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairAIGatewayParametersResponse) SetStatusCode(v int32) *DescribeTairAIGatewayParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairAIGatewayParametersResponse) SetBody(v *DescribeTairAIGatewayParametersResponseBody) *DescribeTairAIGatewayParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeTairAIGatewayParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
