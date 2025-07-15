// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectTrafficQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectTrafficQosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectTrafficQosResponseBody) *DescribeExpressConnectTrafficQosResponse
	GetBody() *DescribeExpressConnectTrafficQosResponseBody
}

type DescribeExpressConnectTrafficQosResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectTrafficQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectTrafficQosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectTrafficQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectTrafficQosResponse) GetBody() *DescribeExpressConnectTrafficQosResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectTrafficQosResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectTrafficQosResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponse) SetStatusCode(v int32) *DescribeExpressConnectTrafficQosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponse) SetBody(v *DescribeExpressConnectTrafficQosResponseBody) *DescribeExpressConnectTrafficQosResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectTrafficQosResponse) Validate() error {
	return dara.Validate(s)
}
