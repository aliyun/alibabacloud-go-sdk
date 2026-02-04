// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceCLSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceCLSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceCLSResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceCLSResponseBody) *DescribeDBInstanceCLSResponse
	GetBody() *DescribeDBInstanceCLSResponseBody
}

type DescribeDBInstanceCLSResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceCLSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceCLSResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceCLSResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceCLSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceCLSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceCLSResponse) GetBody() *DescribeDBInstanceCLSResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceCLSResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceCLSResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceCLSResponse) SetStatusCode(v int32) *DescribeDBInstanceCLSResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceCLSResponse) SetBody(v *DescribeDBInstanceCLSResponseBody) *DescribeDBInstanceCLSResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceCLSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
