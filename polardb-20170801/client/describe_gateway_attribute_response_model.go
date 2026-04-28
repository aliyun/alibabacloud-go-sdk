// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGatewayAttributeResponseBody) *DescribeGatewayAttributeResponse
	GetBody() *DescribeGatewayAttributeResponseBody
}

type DescribeGatewayAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGatewayAttributeResponse) GetBody() *DescribeGatewayAttributeResponseBody {
	return s.Body
}

func (s *DescribeGatewayAttributeResponse) SetHeaders(v map[string]*string) *DescribeGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayAttributeResponse) SetStatusCode(v int32) *DescribeGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayAttributeResponse) SetBody(v *DescribeGatewayAttributeResponseBody) *DescribeGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
