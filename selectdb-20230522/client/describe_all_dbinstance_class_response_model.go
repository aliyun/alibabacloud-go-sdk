// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDBInstanceClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllDBInstanceClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllDBInstanceClassResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllDBInstanceClassResponseBody) *DescribeAllDBInstanceClassResponse
	GetBody() *DescribeAllDBInstanceClassResponseBody
}

type DescribeAllDBInstanceClassResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllDBInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllDBInstanceClassResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDBInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDBInstanceClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllDBInstanceClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllDBInstanceClassResponse) GetBody() *DescribeAllDBInstanceClassResponseBody {
	return s.Body
}

func (s *DescribeAllDBInstanceClassResponse) SetHeaders(v map[string]*string) *DescribeAllDBInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDBInstanceClassResponse) SetStatusCode(v int32) *DescribeAllDBInstanceClassResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponse) SetBody(v *DescribeAllDBInstanceClassResponseBody) *DescribeAllDBInstanceClassResponse {
	s.Body = v
	return s
}

func (s *DescribeAllDBInstanceClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
