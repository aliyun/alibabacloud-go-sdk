// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterConnectivityResponseBody) *DescribeDBClusterConnectivityResponse
	GetBody() *DescribeDBClusterConnectivityResponseBody
}

type DescribeDBClusterConnectivityResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConnectivityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterConnectivityResponse) GetBody() *DescribeDBClusterConnectivityResponseBody {
	return s.Body
}

func (s *DescribeDBClusterConnectivityResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConnectivityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConnectivityResponse) SetStatusCode(v int32) *DescribeDBClusterConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConnectivityResponse) SetBody(v *DescribeDBClusterConnectivityResponseBody) *DescribeDBClusterConnectivityResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
