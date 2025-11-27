// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceNetInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceNetInfoResponseBody) *DescribeDBInstanceNetInfoResponse
	GetBody() *DescribeDBInstanceNetInfoResponseBody
}

type DescribeDBInstanceNetInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceNetInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceNetInfoResponse) GetBody() *DescribeDBInstanceNetInfoResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetStatusCode(v int32) *DescribeDBInstanceNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetBody(v *DescribeDBInstanceNetInfoResponseBody) *DescribeDBInstanceNetInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
