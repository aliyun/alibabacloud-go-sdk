// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceEndpointResponseBody) *DescribeDBInstanceEndpointResponse
	GetBody() *DescribeDBInstanceEndpointResponseBody
}

type DescribeDBInstanceEndpointResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceEndpointResponse) GetBody() *DescribeDBInstanceEndpointResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceEndpointResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceEndpointResponse) SetStatusCode(v int32) *DescribeDBInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponse) SetBody(v *DescribeDBInstanceEndpointResponseBody) *DescribeDBInstanceEndpointResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
