// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayApikeyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGatewayApikeyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGatewayApikeyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGatewayApikeyListResponseBody) *DescribeGatewayApikeyListResponse
	GetBody() *DescribeGatewayApikeyListResponseBody
}

type DescribeGatewayApikeyListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGatewayApikeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGatewayApikeyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayApikeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayApikeyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGatewayApikeyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGatewayApikeyListResponse) GetBody() *DescribeGatewayApikeyListResponseBody {
	return s.Body
}

func (s *DescribeGatewayApikeyListResponse) SetHeaders(v map[string]*string) *DescribeGatewayApikeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayApikeyListResponse) SetStatusCode(v int32) *DescribeGatewayApikeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayApikeyListResponse) SetBody(v *DescribeGatewayApikeyListResponseBody) *DescribeGatewayApikeyListResponse {
	s.Body = v
	return s
}

func (s *DescribeGatewayApikeyListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
