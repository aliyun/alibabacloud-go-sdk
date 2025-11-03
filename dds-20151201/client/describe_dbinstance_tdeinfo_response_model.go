// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDEInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceTDEInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceTDEInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceTDEInfoResponseBody) *DescribeDBInstanceTDEInfoResponse
	GetBody() *DescribeDBInstanceTDEInfoResponseBody
}

type DescribeDBInstanceTDEInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceTDEInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceTDEInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceTDEInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceTDEInfoResponse) GetBody() *DescribeDBInstanceTDEInfoResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceTDEInfoResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceTDEInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponse) SetStatusCode(v int32) *DescribeDBInstanceTDEInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponse) SetBody(v *DescribeDBInstanceTDEInfoResponseBody) *DescribeDBInstanceTDEInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
