// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRealtimeDeliveryAccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRealtimeDeliveryAccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRealtimeDeliveryAccResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRealtimeDeliveryAccResponseBody) *DescribeRealtimeDeliveryAccResponse
	GetBody() *DescribeRealtimeDeliveryAccResponseBody
}

type DescribeRealtimeDeliveryAccResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRealtimeDeliveryAccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRealtimeDeliveryAccResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponse) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRealtimeDeliveryAccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRealtimeDeliveryAccResponse) GetBody() *DescribeRealtimeDeliveryAccResponseBody {
	return s.Body
}

func (s *DescribeRealtimeDeliveryAccResponse) SetHeaders(v map[string]*string) *DescribeRealtimeDeliveryAccResponse {
	s.Headers = v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponse) SetStatusCode(v int32) *DescribeRealtimeDeliveryAccResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponse) SetBody(v *DescribeRealtimeDeliveryAccResponseBody) *DescribeRealtimeDeliveryAccResponse {
	s.Body = v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponse) Validate() error {
	return dara.Validate(s)
}
