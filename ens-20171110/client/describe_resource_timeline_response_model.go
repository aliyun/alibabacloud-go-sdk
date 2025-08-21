// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceTimelineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceTimelineResponseBody) *DescribeResourceTimelineResponse
	GetBody() *DescribeResourceTimelineResponseBody
}

type DescribeResourceTimelineResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceTimelineResponse) GetBody() *DescribeResourceTimelineResponseBody {
	return s.Body
}

func (s *DescribeResourceTimelineResponse) SetHeaders(v map[string]*string) *DescribeResourceTimelineResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceTimelineResponse) SetStatusCode(v int32) *DescribeResourceTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceTimelineResponse) SetBody(v *DescribeResourceTimelineResponseBody) *DescribeResourceTimelineResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceTimelineResponse) Validate() error {
	return dara.Validate(s)
}
