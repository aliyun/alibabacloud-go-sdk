// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewaySecretDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGatewaySecretDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGatewaySecretDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGatewaySecretDetailsResponseBody) *DescribeGatewaySecretDetailsResponse
	GetBody() *DescribeGatewaySecretDetailsResponseBody
}

type DescribeGatewaySecretDetailsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGatewaySecretDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGatewaySecretDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewaySecretDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGatewaySecretDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGatewaySecretDetailsResponse) GetBody() *DescribeGatewaySecretDetailsResponseBody {
	return s.Body
}

func (s *DescribeGatewaySecretDetailsResponse) SetHeaders(v map[string]*string) *DescribeGatewaySecretDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewaySecretDetailsResponse) SetStatusCode(v int32) *DescribeGatewaySecretDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponse) SetBody(v *DescribeGatewaySecretDetailsResponseBody) *DescribeGatewaySecretDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeGatewaySecretDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
