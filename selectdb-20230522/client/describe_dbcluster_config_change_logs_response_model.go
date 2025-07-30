// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigChangeLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterConfigChangeLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterConfigChangeLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterConfigChangeLogsResponseBody) *DescribeDBClusterConfigChangeLogsResponse
	GetBody() *DescribeDBClusterConfigChangeLogsResponseBody
}

type DescribeDBClusterConfigChangeLogsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConfigChangeLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigChangeLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterConfigChangeLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterConfigChangeLogsResponse) GetBody() *DescribeDBClusterConfigChangeLogsResponseBody {
	return s.Body
}

func (s *DescribeDBClusterConfigChangeLogsResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponse) SetStatusCode(v int32) *DescribeDBClusterConfigChangeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponse) SetBody(v *DescribeDBClusterConfigChangeLogsResponseBody) *DescribeDBClusterConfigChangeLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterConfigChangeLogsResponse) Validate() error {
	return dara.Validate(s)
}
