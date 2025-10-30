// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamingDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamingDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamingDataSourceResponseBody) *DescribeStreamingDataSourceResponse
	GetBody() *DescribeStreamingDataSourceResponseBody
}

type DescribeStreamingDataSourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamingDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamingDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamingDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamingDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamingDataSourceResponse) GetBody() *DescribeStreamingDataSourceResponseBody {
	return s.Body
}

func (s *DescribeStreamingDataSourceResponse) SetHeaders(v map[string]*string) *DescribeStreamingDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamingDataSourceResponse) SetStatusCode(v int32) *DescribeStreamingDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamingDataSourceResponse) SetBody(v *DescribeStreamingDataSourceResponseBody) *DescribeStreamingDataSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamingDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
