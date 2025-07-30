// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterGroupResponseBody) *DescribeParameterGroupResponse
	GetBody() *DescribeParameterGroupResponseBody
}

type DescribeParameterGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterGroupResponse) GetBody() *DescribeParameterGroupResponseBody {
	return s.Body
}

func (s *DescribeParameterGroupResponse) SetHeaders(v map[string]*string) *DescribeParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterGroupResponse) SetStatusCode(v int32) *DescribeParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterGroupResponse) SetBody(v *DescribeParameterGroupResponseBody) *DescribeParameterGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterGroupResponse) Validate() error {
	return dara.Validate(s)
}
