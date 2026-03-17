// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmartAccessGatewayVersionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmartAccessGatewayVersionsResponseBody) *DescribeSmartAccessGatewayVersionsResponse
	GetBody() *DescribeSmartAccessGatewayVersionsResponseBody
}

type DescribeSmartAccessGatewayVersionsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartAccessGatewayVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartAccessGatewayVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmartAccessGatewayVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmartAccessGatewayVersionsResponse) GetBody() *DescribeSmartAccessGatewayVersionsResponseBody {
	return s.Body
}

func (s *DescribeSmartAccessGatewayVersionsResponse) SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponse) SetStatusCode(v int32) *DescribeSmartAccessGatewayVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponse) SetBody(v *DescribeSmartAccessGatewayVersionsResponseBody) *DescribeSmartAccessGatewayVersionsResponse {
	s.Body = v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
