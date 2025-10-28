// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogHistogramAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowLogHistogramAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowLogHistogramAsyncResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowLogHistogramAsyncResponseBody) *DescribeSlowLogHistogramAsyncResponse
	GetBody() *DescribeSlowLogHistogramAsyncResponseBody
}

type DescribeSlowLogHistogramAsyncResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogHistogramAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowLogHistogramAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowLogHistogramAsyncResponse) GetBody() *DescribeSlowLogHistogramAsyncResponseBody {
	return s.Body
}

func (s *DescribeSlowLogHistogramAsyncResponse) SetHeaders(v map[string]*string) *DescribeSlowLogHistogramAsyncResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponse) SetStatusCode(v int32) *DescribeSlowLogHistogramAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponse) SetBody(v *DescribeSlowLogHistogramAsyncResponseBody) *DescribeSlowLogHistogramAsyncResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
