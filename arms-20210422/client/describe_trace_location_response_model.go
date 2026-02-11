// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTraceLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTraceLocationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTraceLocationResponseBody) *DescribeTraceLocationResponse
	GetBody() *DescribeTraceLocationResponseBody
}

type DescribeTraceLocationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTraceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTraceLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLocationResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTraceLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTraceLocationResponse) GetBody() *DescribeTraceLocationResponseBody {
	return s.Body
}

func (s *DescribeTraceLocationResponse) SetHeaders(v map[string]*string) *DescribeTraceLocationResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceLocationResponse) SetStatusCode(v int32) *DescribeTraceLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceLocationResponse) SetBody(v *DescribeTraceLocationResponseBody) *DescribeTraceLocationResponse {
	s.Body = v
	return s
}

func (s *DescribeTraceLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
