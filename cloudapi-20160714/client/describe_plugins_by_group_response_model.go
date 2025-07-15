// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsByGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginsByGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginsByGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginsByGroupResponseBody) *DescribePluginsByGroupResponse
	GetBody() *DescribePluginsByGroupResponseBody
}

type DescribePluginsByGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginsByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginsByGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginsByGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginsByGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginsByGroupResponse) GetBody() *DescribePluginsByGroupResponseBody {
	return s.Body
}

func (s *DescribePluginsByGroupResponse) SetHeaders(v map[string]*string) *DescribePluginsByGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginsByGroupResponse) SetStatusCode(v int32) *DescribePluginsByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginsByGroupResponse) SetBody(v *DescribePluginsByGroupResponseBody) *DescribePluginsByGroupResponse {
	s.Body = v
	return s
}

func (s *DescribePluginsByGroupResponse) Validate() error {
	return dara.Validate(s)
}
