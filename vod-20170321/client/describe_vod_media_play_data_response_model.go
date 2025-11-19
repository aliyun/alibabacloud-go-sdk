// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodMediaPlayDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodMediaPlayDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodMediaPlayDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodMediaPlayDataResponseBody) *DescribeVodMediaPlayDataResponse
	GetBody() *DescribeVodMediaPlayDataResponseBody
}

type DescribeVodMediaPlayDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodMediaPlayDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodMediaPlayDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMediaPlayDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodMediaPlayDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodMediaPlayDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodMediaPlayDataResponse) GetBody() *DescribeVodMediaPlayDataResponseBody {
	return s.Body
}

func (s *DescribeVodMediaPlayDataResponse) SetHeaders(v map[string]*string) *DescribeVodMediaPlayDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodMediaPlayDataResponse) SetStatusCode(v int32) *DescribeVodMediaPlayDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponse) SetBody(v *DescribeVodMediaPlayDataResponseBody) *DescribeVodMediaPlayDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodMediaPlayDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
