// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceInfoDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTraceInfoDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTraceInfoDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTraceInfoDetailResponseBody) *DescribeTraceInfoDetailResponse
	GetBody() *DescribeTraceInfoDetailResponseBody
}

type DescribeTraceInfoDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTraceInfoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTraceInfoDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTraceInfoDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTraceInfoDetailResponse) GetBody() *DescribeTraceInfoDetailResponseBody {
	return s.Body
}

func (s *DescribeTraceInfoDetailResponse) SetHeaders(v map[string]*string) *DescribeTraceInfoDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceInfoDetailResponse) SetStatusCode(v int32) *DescribeTraceInfoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceInfoDetailResponse) SetBody(v *DescribeTraceInfoDetailResponseBody) *DescribeTraceInfoDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeTraceInfoDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
