// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesForCloneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancesForCloneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancesForCloneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancesForCloneResponseBody) *DescribeDBInstancesForCloneResponse
	GetBody() *DescribeDBInstancesForCloneResponseBody
}

type DescribeDBInstancesForCloneResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancesForCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancesForCloneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancesForCloneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancesForCloneResponse) GetBody() *DescribeDBInstancesForCloneResponseBody {
	return s.Body
}

func (s *DescribeDBInstancesForCloneResponse) SetHeaders(v map[string]*string) *DescribeDBInstancesForCloneResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancesForCloneResponse) SetStatusCode(v int32) *DescribeDBInstancesForCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponse) SetBody(v *DescribeDBInstancesForCloneResponseBody) *DescribeDBInstancesForCloneResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancesForCloneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
