// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogDeliveryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogDeliveryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogDeliveryConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogDeliveryConfigResponseBody) *DescribeLogDeliveryConfigResponse
	GetBody() *DescribeLogDeliveryConfigResponseBody
}

type DescribeLogDeliveryConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogDeliveryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogDeliveryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogDeliveryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogDeliveryConfigResponse) GetBody() *DescribeLogDeliveryConfigResponseBody {
	return s.Body
}

func (s *DescribeLogDeliveryConfigResponse) SetHeaders(v map[string]*string) *DescribeLogDeliveryConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogDeliveryConfigResponse) SetStatusCode(v int32) *DescribeLogDeliveryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogDeliveryConfigResponse) SetBody(v *DescribeLogDeliveryConfigResponseBody) *DescribeLogDeliveryConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLogDeliveryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
