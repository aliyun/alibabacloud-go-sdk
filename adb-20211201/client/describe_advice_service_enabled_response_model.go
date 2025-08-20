// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdviceServiceEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdviceServiceEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdviceServiceEnabledResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdviceServiceEnabledResponseBody) *DescribeAdviceServiceEnabledResponse
	GetBody() *DescribeAdviceServiceEnabledResponseBody
}

type DescribeAdviceServiceEnabledResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdviceServiceEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdviceServiceEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdviceServiceEnabledResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdviceServiceEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdviceServiceEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdviceServiceEnabledResponse) GetBody() *DescribeAdviceServiceEnabledResponseBody {
	return s.Body
}

func (s *DescribeAdviceServiceEnabledResponse) SetHeaders(v map[string]*string) *DescribeAdviceServiceEnabledResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdviceServiceEnabledResponse) SetStatusCode(v int32) *DescribeAdviceServiceEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponse) SetBody(v *DescribeAdviceServiceEnabledResponseBody) *DescribeAdviceServiceEnabledResponse {
	s.Body = v
	return s
}

func (s *DescribeAdviceServiceEnabledResponse) Validate() error {
	return dara.Validate(s)
}
