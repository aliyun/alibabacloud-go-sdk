// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultiPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultiPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultiPriceResponseBody) *DescribeMultiPriceResponse
	GetBody() *DescribeMultiPriceResponseBody
}

type DescribeMultiPriceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultiPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultiPriceResponse) GetBody() *DescribeMultiPriceResponseBody {
	return s.Body
}

func (s *DescribeMultiPriceResponse) SetHeaders(v map[string]*string) *DescribeMultiPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiPriceResponse) SetStatusCode(v int32) *DescribeMultiPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiPriceResponse) SetBody(v *DescribeMultiPriceResponseBody) *DescribeMultiPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeMultiPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
