// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceReplicationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceReplicationResponseBody) *DescribeDBInstanceReplicationResponse
	GetBody() *DescribeDBInstanceReplicationResponseBody
}

type DescribeDBInstanceReplicationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceReplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceReplicationResponse) GetBody() *DescribeDBInstanceReplicationResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceReplicationResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceReplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceReplicationResponse) SetStatusCode(v int32) *DescribeDBInstanceReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponse) SetBody(v *DescribeDBInstanceReplicationResponseBody) *DescribeDBInstanceReplicationResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
