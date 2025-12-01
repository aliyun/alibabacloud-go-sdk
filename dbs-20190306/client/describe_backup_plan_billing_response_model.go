// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanBillingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupPlanBillingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupPlanBillingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupPlanBillingResponseBody) *DescribeBackupPlanBillingResponse
	GetBody() *DescribeBackupPlanBillingResponseBody
}

type DescribeBackupPlanBillingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPlanBillingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPlanBillingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanBillingResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupPlanBillingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupPlanBillingResponse) GetBody() *DescribeBackupPlanBillingResponseBody {
	return s.Body
}

func (s *DescribeBackupPlanBillingResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanBillingResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanBillingResponse) SetStatusCode(v int32) *DescribeBackupPlanBillingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPlanBillingResponse) SetBody(v *DescribeBackupPlanBillingResponseBody) *DescribeBackupPlanBillingResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupPlanBillingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
