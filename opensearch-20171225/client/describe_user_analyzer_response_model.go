// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAnalyzerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAnalyzerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAnalyzerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAnalyzerResponseBody) *DescribeUserAnalyzerResponse
	GetBody() *DescribeUserAnalyzerResponseBody
}

type DescribeUserAnalyzerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAnalyzerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAnalyzerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAnalyzerResponse) GetBody() *DescribeUserAnalyzerResponseBody {
	return s.Body
}

func (s *DescribeUserAnalyzerResponse) SetHeaders(v map[string]*string) *DescribeUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAnalyzerResponse) SetStatusCode(v int32) *DescribeUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAnalyzerResponse) SetBody(v *DescribeUserAnalyzerResponseBody) *DescribeUserAnalyzerResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAnalyzerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
