// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupPlanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupPlanConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupPlanConfigResponseBody) *DescribeBackupPlanConfigResponse
	GetBody() *DescribeBackupPlanConfigResponseBody
}

type DescribeBackupPlanConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPlanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPlanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupPlanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupPlanConfigResponse) GetBody() *DescribeBackupPlanConfigResponseBody {
	return s.Body
}

func (s *DescribeBackupPlanConfigResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanConfigResponse) SetStatusCode(v int32) *DescribeBackupPlanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPlanConfigResponse) SetBody(v *DescribeBackupPlanConfigResponseBody) *DescribeBackupPlanConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupPlanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
