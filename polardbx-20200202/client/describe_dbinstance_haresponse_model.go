// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceHAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceHAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceHAResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceHAResponseBody) *DescribeDBInstanceHAResponse
	GetBody() *DescribeDBInstanceHAResponseBody
}

type DescribeDBInstanceHAResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceHAResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceHAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceHAResponse) GetBody() *DescribeDBInstanceHAResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceHAResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceHAResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceHAResponse) SetStatusCode(v int32) *DescribeDBInstanceHAResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceHAResponse) SetBody(v *DescribeDBInstanceHAResponseBody) *DescribeDBInstanceHAResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceHAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
