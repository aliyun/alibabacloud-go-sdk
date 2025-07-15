// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginGroupsResponseBody) *DescribePluginGroupsResponse
	GetBody() *DescribePluginGroupsResponseBody
}

type DescribePluginGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginGroupsResponse) GetBody() *DescribePluginGroupsResponseBody {
	return s.Body
}

func (s *DescribePluginGroupsResponse) SetHeaders(v map[string]*string) *DescribePluginGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginGroupsResponse) SetStatusCode(v int32) *DescribePluginGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginGroupsResponse) SetBody(v *DescribePluginGroupsResponseBody) *DescribePluginGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribePluginGroupsResponse) Validate() error {
	return dara.Validate(s)
}
