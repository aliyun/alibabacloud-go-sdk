// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModificationPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModificationPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModificationPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModificationPriceResponseBody) *DescribeModificationPriceResponse
	GetBody() *DescribeModificationPriceResponseBody
}

type DescribeModificationPriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModificationPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModificationPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModificationPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModificationPriceResponse) GetBody() *DescribeModificationPriceResponseBody {
	return s.Body
}

func (s *DescribeModificationPriceResponse) SetHeaders(v map[string]*string) *DescribeModificationPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeModificationPriceResponse) SetStatusCode(v int32) *DescribeModificationPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModificationPriceResponse) SetBody(v *DescribeModificationPriceResponseBody) *DescribeModificationPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeModificationPriceResponse) Validate() error {
	return dara.Validate(s)
}
