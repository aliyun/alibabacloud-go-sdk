// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSwimmingLaneResponseBody) *DescribeSwimmingLaneResponse
	GetBody() *DescribeSwimmingLaneResponseBody
}

type DescribeSwimmingLaneResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSwimmingLaneResponse) GetBody() *DescribeSwimmingLaneResponseBody {
	return s.Body
}

func (s *DescribeSwimmingLaneResponse) SetHeaders(v map[string]*string) *DescribeSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *DescribeSwimmingLaneResponse) SetStatusCode(v int32) *DescribeSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSwimmingLaneResponse) SetBody(v *DescribeSwimmingLaneResponseBody) *DescribeSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *DescribeSwimmingLaneResponse) Validate() error {
	return dara.Validate(s)
}
