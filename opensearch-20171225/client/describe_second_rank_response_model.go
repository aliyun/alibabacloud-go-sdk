// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecondRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecondRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecondRankResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecondRankResponseBody) *DescribeSecondRankResponse
	GetBody() *DescribeSecondRankResponseBody
}

type DescribeSecondRankResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecondRankResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondRankResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecondRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecondRankResponse) GetBody() *DescribeSecondRankResponseBody {
	return s.Body
}

func (s *DescribeSecondRankResponse) SetHeaders(v map[string]*string) *DescribeSecondRankResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecondRankResponse) SetStatusCode(v int32) *DescribeSecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecondRankResponse) SetBody(v *DescribeSecondRankResponseBody) *DescribeSecondRankResponse {
	s.Body = v
	return s
}

func (s *DescribeSecondRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
