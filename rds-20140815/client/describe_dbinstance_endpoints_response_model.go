// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceEndpointsResponseBody) *DescribeDBInstanceEndpointsResponse
	GetBody() *DescribeDBInstanceEndpointsResponseBody
}

type DescribeDBInstanceEndpointsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceEndpointsResponse) GetBody() *DescribeDBInstanceEndpointsResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceEndpointsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponse) SetStatusCode(v int32) *DescribeDBInstanceEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponse) SetBody(v *DescribeDBInstanceEndpointsResponseBody) *DescribeDBInstanceEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
