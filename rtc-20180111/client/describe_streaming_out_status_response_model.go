// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingOutStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamingOutStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamingOutStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamingOutStatusResponseBody) *DescribeStreamingOutStatusResponse
	GetBody() *DescribeStreamingOutStatusResponseBody
}

type DescribeStreamingOutStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamingOutStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamingOutStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingOutStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamingOutStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamingOutStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamingOutStatusResponse) GetBody() *DescribeStreamingOutStatusResponseBody {
	return s.Body
}

func (s *DescribeStreamingOutStatusResponse) SetHeaders(v map[string]*string) *DescribeStreamingOutStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamingOutStatusResponse) SetStatusCode(v int32) *DescribeStreamingOutStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamingOutStatusResponse) SetBody(v *DescribeStreamingOutStatusResponseBody) *DescribeStreamingOutStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamingOutStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
