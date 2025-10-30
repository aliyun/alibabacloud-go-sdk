// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRatioLineChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsRatioLineChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsRatioLineChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsRatioLineChartResponseBody) *DescribeTagsRatioLineChartResponse
	GetBody() *DescribeTagsRatioLineChartResponseBody
}

type DescribeTagsRatioLineChartResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsRatioLineChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsRatioLineChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRatioLineChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsRatioLineChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsRatioLineChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsRatioLineChartResponse) GetBody() *DescribeTagsRatioLineChartResponseBody {
	return s.Body
}

func (s *DescribeTagsRatioLineChartResponse) SetHeaders(v map[string]*string) *DescribeTagsRatioLineChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsRatioLineChartResponse) SetStatusCode(v int32) *DescribeTagsRatioLineChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsRatioLineChartResponse) SetBody(v *DescribeTagsRatioLineChartResponseBody) *DescribeTagsRatioLineChartResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsRatioLineChartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
