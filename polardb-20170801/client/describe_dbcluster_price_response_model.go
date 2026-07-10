// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterPriceResponseBody) *DescribeDBClusterPriceResponse
	GetBody() *DescribeDBClusterPriceResponseBody
}

type DescribeDBClusterPriceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterPriceResponse) GetBody() *DescribeDBClusterPriceResponseBody {
	return s.Body
}

func (s *DescribeDBClusterPriceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPriceResponse) SetStatusCode(v int32) *DescribeDBClusterPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterPriceResponse) SetBody(v *DescribeDBClusterPriceResponseBody) *DescribeDBClusterPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
