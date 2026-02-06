// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobGroupResponseBody) *DescribeJobGroupResponse
	GetBody() *DescribeJobGroupResponseBody
}

type DescribeJobGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobGroupResponse) GetBody() *DescribeJobGroupResponseBody {
	return s.Body
}

func (s *DescribeJobGroupResponse) SetHeaders(v map[string]*string) *DescribeJobGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobGroupResponse) SetStatusCode(v int32) *DescribeJobGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobGroupResponse) SetBody(v *DescribeJobGroupResponseBody) *DescribeJobGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeJobGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
