// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiMarketAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiMarketAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiMarketAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiMarketAttributesResponseBody) *DescribeApiMarketAttributesResponse
	GetBody() *DescribeApiMarketAttributesResponseBody
}

type DescribeApiMarketAttributesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiMarketAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiMarketAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMarketAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiMarketAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiMarketAttributesResponse) GetBody() *DescribeApiMarketAttributesResponseBody {
	return s.Body
}

func (s *DescribeApiMarketAttributesResponse) SetHeaders(v map[string]*string) *DescribeApiMarketAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiMarketAttributesResponse) SetStatusCode(v int32) *DescribeApiMarketAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiMarketAttributesResponse) SetBody(v *DescribeApiMarketAttributesResponseBody) *DescribeApiMarketAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeApiMarketAttributesResponse) Validate() error {
	return dara.Validate(s)
}
