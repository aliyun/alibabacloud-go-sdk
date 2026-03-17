// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmartAccessGatewayHaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmartAccessGatewayHaResponseBody) *DescribeSmartAccessGatewayHaResponse
	GetBody() *DescribeSmartAccessGatewayHaResponseBody
}

type DescribeSmartAccessGatewayHaResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartAccessGatewayHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartAccessGatewayHaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayHaResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmartAccessGatewayHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmartAccessGatewayHaResponse) GetBody() *DescribeSmartAccessGatewayHaResponseBody {
	return s.Body
}

func (s *DescribeSmartAccessGatewayHaResponse) SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayHaResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponse) SetStatusCode(v int32) *DescribeSmartAccessGatewayHaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponse) SetBody(v *DescribeSmartAccessGatewayHaResponseBody) *DescribeSmartAccessGatewayHaResponse {
	s.Body = v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
