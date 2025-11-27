// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceConnectivityResponseBody) *DescribeDBInstanceConnectivityResponse
	GetBody() *DescribeDBInstanceConnectivityResponseBody
}

type DescribeDBInstanceConnectivityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConnectivityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceConnectivityResponse) GetBody() *DescribeDBInstanceConnectivityResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceConnectivityResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceConnectivityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceConnectivityResponse) SetStatusCode(v int32) *DescribeDBInstanceConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceConnectivityResponse) SetBody(v *DescribeDBInstanceConnectivityResponseBody) *DescribeDBInstanceConnectivityResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
