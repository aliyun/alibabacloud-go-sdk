// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupResponseBody) *DescribeGroupResponse
	GetBody() *DescribeGroupResponseBody
}

type DescribeGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupResponse) GetBody() *DescribeGroupResponseBody {
	return s.Body
}

func (s *DescribeGroupResponse) SetHeaders(v map[string]*string) *DescribeGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupResponse) SetStatusCode(v int32) *DescribeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupResponse) SetBody(v *DescribeGroupResponseBody) *DescribeGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupResponse) Validate() error {
	return dara.Validate(s)
}
