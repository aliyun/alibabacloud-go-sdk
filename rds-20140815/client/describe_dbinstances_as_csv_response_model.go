// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesAsCsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancesAsCsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancesAsCsvResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancesAsCsvResponseBody) *DescribeDBInstancesAsCsvResponse
	GetBody() *DescribeDBInstancesAsCsvResponseBody
}

type DescribeDBInstancesAsCsvResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesAsCsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesAsCsvResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesAsCsvResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesAsCsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancesAsCsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancesAsCsvResponse) GetBody() *DescribeDBInstancesAsCsvResponseBody {
	return s.Body
}

func (s *DescribeDBInstancesAsCsvResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesAsCsvResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesAsCsvResponse) SetStatusCode(v int32) *DescribeDBInstancesAsCsvResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponse) SetBody(v *DescribeDBInstancesAsCsvResponseBody) *DescribeDBInstancesAsCsvResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancesAsCsvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
