// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventTrendsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEventTrendsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEventTrendsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEventTrendsResponseBody) *DescribeAiEngineEventTrendsResponse
	GetBody() *DescribeAiEngineEventTrendsResponseBody
}

type DescribeAiEngineEventTrendsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEventTrendsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEventTrendsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventTrendsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventTrendsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEventTrendsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEventTrendsResponse) GetBody() *DescribeAiEngineEventTrendsResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEventTrendsResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEventTrendsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEventTrendsResponse) SetStatusCode(v int32) *DescribeAiEngineEventTrendsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponse) SetBody(v *DescribeAiEngineEventTrendsResponseBody) *DescribeAiEngineEventTrendsResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEventTrendsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
