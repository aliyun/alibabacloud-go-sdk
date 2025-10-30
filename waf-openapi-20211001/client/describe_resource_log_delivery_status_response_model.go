// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogDeliveryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceLogDeliveryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceLogDeliveryStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceLogDeliveryStatusResponseBody) *DescribeResourceLogDeliveryStatusResponse
	GetBody() *DescribeResourceLogDeliveryStatusResponseBody
}

type DescribeResourceLogDeliveryStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLogDeliveryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLogDeliveryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogDeliveryStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogDeliveryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceLogDeliveryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceLogDeliveryStatusResponse) GetBody() *DescribeResourceLogDeliveryStatusResponseBody {
	return s.Body
}

func (s *DescribeResourceLogDeliveryStatusResponse) SetHeaders(v map[string]*string) *DescribeResourceLogDeliveryStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponse) SetStatusCode(v int32) *DescribeResourceLogDeliveryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponse) SetBody(v *DescribeResourceLogDeliveryStatusResponseBody) *DescribeResourceLogDeliveryStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceLogDeliveryStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
