// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAITrafficAnalysisStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAITrafficAnalysisStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAITrafficAnalysisStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAITrafficAnalysisStatusResponseBody) *DescribeAITrafficAnalysisStatusResponse
	GetBody() *DescribeAITrafficAnalysisStatusResponseBody
}

type DescribeAITrafficAnalysisStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAITrafficAnalysisStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAITrafficAnalysisStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAITrafficAnalysisStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAITrafficAnalysisStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAITrafficAnalysisStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAITrafficAnalysisStatusResponse) GetBody() *DescribeAITrafficAnalysisStatusResponseBody {
	return s.Body
}

func (s *DescribeAITrafficAnalysisStatusResponse) SetHeaders(v map[string]*string) *DescribeAITrafficAnalysisStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAITrafficAnalysisStatusResponse) SetStatusCode(v int32) *DescribeAITrafficAnalysisStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAITrafficAnalysisStatusResponse) SetBody(v *DescribeAITrafficAnalysisStatusResponseBody) *DescribeAITrafficAnalysisStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAITrafficAnalysisStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
