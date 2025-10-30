// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogDeliveryConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogDeliveryConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogDeliveryConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogDeliveryConfigsResponseBody) *DescribeLogDeliveryConfigsResponse
	GetBody() *DescribeLogDeliveryConfigsResponseBody
}

type DescribeLogDeliveryConfigsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogDeliveryConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogDeliveryConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogDeliveryConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogDeliveryConfigsResponse) GetBody() *DescribeLogDeliveryConfigsResponseBody {
	return s.Body
}

func (s *DescribeLogDeliveryConfigsResponse) SetHeaders(v map[string]*string) *DescribeLogDeliveryConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogDeliveryConfigsResponse) SetStatusCode(v int32) *DescribeLogDeliveryConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogDeliveryConfigsResponse) SetBody(v *DescribeLogDeliveryConfigsResponseBody) *DescribeLogDeliveryConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeLogDeliveryConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
