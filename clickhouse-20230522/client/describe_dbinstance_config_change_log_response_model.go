// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigChangeLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceConfigChangeLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceConfigChangeLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceConfigChangeLogResponseBody) *DescribeDBInstanceConfigChangeLogResponse
	GetBody() *DescribeDBInstanceConfigChangeLogResponseBody
}

type DescribeDBInstanceConfigChangeLogResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceConfigChangeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceConfigChangeLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigChangeLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigChangeLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceConfigChangeLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceConfigChangeLogResponse) GetBody() *DescribeDBInstanceConfigChangeLogResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceConfigChangeLogResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceConfigChangeLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponse) SetStatusCode(v int32) *DescribeDBInstanceConfigChangeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponse) SetBody(v *DescribeDBInstanceConfigChangeLogResponseBody) *DescribeDBInstanceConfigChangeLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
