// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSpecInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceSpecInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceSpecInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceSpecInfoResponseBody) *DescribeDBInstanceSpecInfoResponse
	GetBody() *DescribeDBInstanceSpecInfoResponseBody
}

type DescribeDBInstanceSpecInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceSpecInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSpecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSpecInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceSpecInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceSpecInfoResponse) GetBody() *DescribeDBInstanceSpecInfoResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceSpecInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSpecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponse) SetStatusCode(v int32) *DescribeDBInstanceSpecInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponse) SetBody(v *DescribeDBInstanceSpecInfoResponseBody) *DescribeDBInstanceSpecInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
