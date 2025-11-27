// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketingActivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMarketingActivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMarketingActivityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMarketingActivityResponseBody) *DescribeMarketingActivityResponse
	GetBody() *DescribeMarketingActivityResponseBody
}

type DescribeMarketingActivityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMarketingActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMarketingActivityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketingActivityResponse) GoString() string {
	return s.String()
}

func (s *DescribeMarketingActivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMarketingActivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMarketingActivityResponse) GetBody() *DescribeMarketingActivityResponseBody {
	return s.Body
}

func (s *DescribeMarketingActivityResponse) SetHeaders(v map[string]*string) *DescribeMarketingActivityResponse {
	s.Headers = v
	return s
}

func (s *DescribeMarketingActivityResponse) SetStatusCode(v int32) *DescribeMarketingActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMarketingActivityResponse) SetBody(v *DescribeMarketingActivityResponseBody) *DescribeMarketingActivityResponse {
	s.Body = v
	return s
}

func (s *DescribeMarketingActivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
