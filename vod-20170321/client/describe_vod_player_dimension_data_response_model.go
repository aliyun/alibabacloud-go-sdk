// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerDimensionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodPlayerDimensionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodPlayerDimensionDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodPlayerDimensionDataResponseBody) *DescribeVodPlayerDimensionDataResponse
	GetBody() *DescribeVodPlayerDimensionDataResponseBody
}

type DescribeVodPlayerDimensionDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodPlayerDimensionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodPlayerDimensionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerDimensionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerDimensionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodPlayerDimensionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodPlayerDimensionDataResponse) GetBody() *DescribeVodPlayerDimensionDataResponseBody {
	return s.Body
}

func (s *DescribeVodPlayerDimensionDataResponse) SetHeaders(v map[string]*string) *DescribeVodPlayerDimensionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodPlayerDimensionDataResponse) SetStatusCode(v int32) *DescribeVodPlayerDimensionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataResponse) SetBody(v *DescribeVodPlayerDimensionDataResponseBody) *DescribeVodPlayerDimensionDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodPlayerDimensionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
