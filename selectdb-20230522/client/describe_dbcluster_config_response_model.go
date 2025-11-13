// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterConfigResponseBody) *DescribeDBClusterConfigResponse
	GetBody() *DescribeDBClusterConfigResponseBody
}

type DescribeDBClusterConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterConfigResponse) GetBody() *DescribeDBClusterConfigResponseBody {
	return s.Body
}

func (s *DescribeDBClusterConfigResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigResponse) SetStatusCode(v int32) *DescribeDBClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigResponse) SetBody(v *DescribeDBClusterConfigResponseBody) *DescribeDBClusterConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
