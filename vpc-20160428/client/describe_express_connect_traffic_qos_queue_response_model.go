// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectTrafficQosQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectTrafficQosQueueResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectTrafficQosQueueResponseBody) *DescribeExpressConnectTrafficQosQueueResponse
	GetBody() *DescribeExpressConnectTrafficQosQueueResponseBody
}

type DescribeExpressConnectTrafficQosQueueResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectTrafficQosQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) GetBody() *DescribeExpressConnectTrafficQosQueueResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) SetStatusCode(v int32) *DescribeExpressConnectTrafficQosQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) SetBody(v *DescribeExpressConnectTrafficQosQueueResponseBody) *DescribeExpressConnectTrafficQosQueueResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
