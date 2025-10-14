// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScaleOutMigrateTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScaleOutMigrateTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScaleOutMigrateTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScaleOutMigrateTaskListResponseBody) *DescribeScaleOutMigrateTaskListResponse
	GetBody() *DescribeScaleOutMigrateTaskListResponseBody
}

type DescribeScaleOutMigrateTaskListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScaleOutMigrateTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScaleOutMigrateTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScaleOutMigrateTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScaleOutMigrateTaskListResponse) GetBody() *DescribeScaleOutMigrateTaskListResponseBody {
	return s.Body
}

func (s *DescribeScaleOutMigrateTaskListResponse) SetHeaders(v map[string]*string) *DescribeScaleOutMigrateTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponse) SetStatusCode(v int32) *DescribeScaleOutMigrateTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponse) SetBody(v *DescribeScaleOutMigrateTaskListResponseBody) *DescribeScaleOutMigrateTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
