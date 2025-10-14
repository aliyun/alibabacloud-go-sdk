// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceViaEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceViaEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceViaEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceViaEndpointResponseBody) *DescribeDBInstanceViaEndpointResponse
	GetBody() *DescribeDBInstanceViaEndpointResponseBody
}

type DescribeDBInstanceViaEndpointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceViaEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceViaEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceViaEndpointResponse) GetBody() *DescribeDBInstanceViaEndpointResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceViaEndpointResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceViaEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponse) SetStatusCode(v int32) *DescribeDBInstanceViaEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponse) SetBody(v *DescribeDBInstanceViaEndpointResponseBody) *DescribeDBInstanceViaEndpointResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
