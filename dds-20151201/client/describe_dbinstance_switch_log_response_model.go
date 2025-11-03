// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSwitchLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceSwitchLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceSwitchLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceSwitchLogResponseBody) *DescribeDBInstanceSwitchLogResponse
	GetBody() *DescribeDBInstanceSwitchLogResponseBody
}

type DescribeDBInstanceSwitchLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceSwitchLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceSwitchLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceSwitchLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceSwitchLogResponse) GetBody() *DescribeDBInstanceSwitchLogResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceSwitchLogResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceSwitchLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponse) SetStatusCode(v int32) *DescribeDBInstanceSwitchLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponse) SetBody(v *DescribeDBInstanceSwitchLogResponseBody) *DescribeDBInstanceSwitchLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
