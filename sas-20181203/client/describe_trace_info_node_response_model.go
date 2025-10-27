// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceInfoNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTraceInfoNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTraceInfoNodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTraceInfoNodeResponseBody) *DescribeTraceInfoNodeResponse
	GetBody() *DescribeTraceInfoNodeResponseBody
}

type DescribeTraceInfoNodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTraceInfoNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTraceInfoNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTraceInfoNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTraceInfoNodeResponse) GetBody() *DescribeTraceInfoNodeResponseBody {
	return s.Body
}

func (s *DescribeTraceInfoNodeResponse) SetHeaders(v map[string]*string) *DescribeTraceInfoNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceInfoNodeResponse) SetStatusCode(v int32) *DescribeTraceInfoNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceInfoNodeResponse) SetBody(v *DescribeTraceInfoNodeResponseBody) *DescribeTraceInfoNodeResponse {
	s.Body = v
	return s
}

func (s *DescribeTraceInfoNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
