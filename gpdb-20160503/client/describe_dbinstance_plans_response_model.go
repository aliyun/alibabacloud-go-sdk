// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancePlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancePlansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancePlansResponseBody) *DescribeDBInstancePlansResponse
	GetBody() *DescribeDBInstancePlansResponseBody
}

type DescribeDBInstancePlansResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancePlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancePlansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancePlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancePlansResponse) GetBody() *DescribeDBInstancePlansResponseBody {
	return s.Body
}

func (s *DescribeDBInstancePlansResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePlansResponse) SetStatusCode(v int32) *DescribeDBInstancePlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePlansResponse) SetBody(v *DescribeDBInstancePlansResponseBody) *DescribeDBInstancePlansResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancePlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
