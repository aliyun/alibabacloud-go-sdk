// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupPlanListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupPlanListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupPlanListResponseBody) *DescribeBackupPlanListResponse
	GetBody() *DescribeBackupPlanListResponseBody
}

type DescribeBackupPlanListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPlanListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPlanListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupPlanListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupPlanListResponse) GetBody() *DescribeBackupPlanListResponseBody {
	return s.Body
}

func (s *DescribeBackupPlanListResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanListResponse) SetStatusCode(v int32) *DescribeBackupPlanListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPlanListResponse) SetBody(v *DescribeBackupPlanListResponseBody) *DescribeBackupPlanListResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupPlanListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
