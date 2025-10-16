// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAnalyzeOpenStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsAnalyzeOpenStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsAnalyzeOpenStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsAnalyzeOpenStatusResponseBody) *DescribeSlsAnalyzeOpenStatusResponse
	GetBody() *DescribeSlsAnalyzeOpenStatusResponseBody
}

type DescribeSlsAnalyzeOpenStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsAnalyzeOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsAnalyzeOpenStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAnalyzeOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) GetBody() *DescribeSlsAnalyzeOpenStatusResponseBody {
	return s.Body
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsAnalyzeOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) SetStatusCode(v int32) *DescribeSlsAnalyzeOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) SetBody(v *DescribeSlsAnalyzeOpenStatusResponseBody) *DescribeSlsAnalyzeOpenStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsAnalyzeOpenStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
