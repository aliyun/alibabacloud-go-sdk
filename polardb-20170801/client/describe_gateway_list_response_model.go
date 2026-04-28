// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGatewayListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGatewayListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGatewayListResponseBody) *DescribeGatewayListResponse
	GetBody() *DescribeGatewayListResponseBody
}

type DescribeGatewayListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGatewayListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGatewayListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGatewayListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGatewayListResponse) GetBody() *DescribeGatewayListResponseBody {
	return s.Body
}

func (s *DescribeGatewayListResponse) SetHeaders(v map[string]*string) *DescribeGatewayListResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayListResponse) SetStatusCode(v int32) *DescribeGatewayListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayListResponse) SetBody(v *DescribeGatewayListResponseBody) *DescribeGatewayListResponse {
	s.Body = v
	return s
}

func (s *DescribeGatewayListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
