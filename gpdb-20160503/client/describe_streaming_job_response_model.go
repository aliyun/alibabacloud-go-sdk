// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamingJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamingJobResponseBody) *DescribeStreamingJobResponse
	GetBody() *DescribeStreamingJobResponseBody
}

type DescribeStreamingJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamingJobResponse) GetBody() *DescribeStreamingJobResponseBody {
	return s.Body
}

func (s *DescribeStreamingJobResponse) SetHeaders(v map[string]*string) *DescribeStreamingJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamingJobResponse) SetStatusCode(v int32) *DescribeStreamingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamingJobResponse) SetBody(v *DescribeStreamingJobResponseBody) *DescribeStreamingJobResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
