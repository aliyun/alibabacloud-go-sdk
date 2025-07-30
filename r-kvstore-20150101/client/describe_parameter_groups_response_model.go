// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterGroupsResponseBody) *DescribeParameterGroupsResponse
	GetBody() *DescribeParameterGroupsResponseBody
}

type DescribeParameterGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterGroupsResponse) GetBody() *DescribeParameterGroupsResponseBody {
	return s.Body
}

func (s *DescribeParameterGroupsResponse) SetHeaders(v map[string]*string) *DescribeParameterGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterGroupsResponse) SetStatusCode(v int32) *DescribeParameterGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterGroupsResponse) SetBody(v *DescribeParameterGroupsResponseBody) *DescribeParameterGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterGroupsResponse) Validate() error {
	return dara.Validate(s)
}
