// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmartAccessGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmartAccessGatewayAttributeResponseBody) *DescribeSmartAccessGatewayAttributeResponse
	GetBody() *DescribeSmartAccessGatewayAttributeResponseBody
}

type DescribeSmartAccessGatewayAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartAccessGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartAccessGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmartAccessGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmartAccessGatewayAttributeResponse) GetBody() *DescribeSmartAccessGatewayAttributeResponseBody {
	return s.Body
}

func (s *DescribeSmartAccessGatewayAttributeResponse) SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponse) SetStatusCode(v int32) *DescribeSmartAccessGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponse) SetBody(v *DescribeSmartAccessGatewayAttributeResponseBody) *DescribeSmartAccessGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
