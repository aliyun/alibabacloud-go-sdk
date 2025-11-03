// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreDBInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreDBInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreDBInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreDBInstanceListResponseBody) *DescribeRestoreDBInstanceListResponse
	GetBody() *DescribeRestoreDBInstanceListResponseBody
}

type DescribeRestoreDBInstanceListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreDBInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreDBInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreDBInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreDBInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreDBInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreDBInstanceListResponse) GetBody() *DescribeRestoreDBInstanceListResponseBody {
	return s.Body
}

func (s *DescribeRestoreDBInstanceListResponse) SetHeaders(v map[string]*string) *DescribeRestoreDBInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreDBInstanceListResponse) SetStatusCode(v int32) *DescribeRestoreDBInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreDBInstanceListResponse) SetBody(v *DescribeRestoreDBInstanceListResponseBody) *DescribeRestoreDBInstanceListResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreDBInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
