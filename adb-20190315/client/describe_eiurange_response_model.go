// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEIURangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEIURangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEIURangeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEIURangeResponseBody) *DescribeEIURangeResponse
	GetBody() *DescribeEIURangeResponseBody
}

type DescribeEIURangeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEIURangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEIURangeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEIURangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEIURangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEIURangeResponse) GetBody() *DescribeEIURangeResponseBody {
	return s.Body
}

func (s *DescribeEIURangeResponse) SetHeaders(v map[string]*string) *DescribeEIURangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEIURangeResponse) SetStatusCode(v int32) *DescribeEIURangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEIURangeResponse) SetBody(v *DescribeEIURangeResponseBody) *DescribeEIURangeResponse {
	s.Body = v
	return s
}

func (s *DescribeEIURangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
