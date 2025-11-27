// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesByExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancesByExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancesByExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancesByExpireTimeResponseBody) *DescribeDBInstancesByExpireTimeResponse
	GetBody() *DescribeDBInstancesByExpireTimeResponseBody
}

type DescribeDBInstancesByExpireTimeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesByExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesByExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancesByExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancesByExpireTimeResponse) GetBody() *DescribeDBInstancesByExpireTimeResponseBody {
	return s.Body
}

func (s *DescribeDBInstancesByExpireTimeResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesByExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponse) SetStatusCode(v int32) *DescribeDBInstancesByExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponse) SetBody(v *DescribeDBInstancesByExpireTimeResponseBody) *DescribeDBInstancesByExpireTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
