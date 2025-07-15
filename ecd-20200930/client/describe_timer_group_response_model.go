// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTimerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTimerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTimerGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTimerGroupResponseBody) *DescribeTimerGroupResponse
	GetBody() *DescribeTimerGroupResponseBody
}

type DescribeTimerGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTimerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTimerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTimerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTimerGroupResponse) GetBody() *DescribeTimerGroupResponseBody {
	return s.Body
}

func (s *DescribeTimerGroupResponse) SetHeaders(v map[string]*string) *DescribeTimerGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeTimerGroupResponse) SetStatusCode(v int32) *DescribeTimerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTimerGroupResponse) SetBody(v *DescribeTimerGroupResponseBody) *DescribeTimerGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeTimerGroupResponse) Validate() error {
	return dara.Validate(s)
}
