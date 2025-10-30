// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDiagnosisSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceDiagnosisSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceDiagnosisSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceDiagnosisSummaryResponseBody) *DescribeDBInstanceDiagnosisSummaryResponse
	GetBody() *DescribeDBInstanceDiagnosisSummaryResponseBody
}

type DescribeDBInstanceDiagnosisSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceDiagnosisSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) GetBody() *DescribeDBInstanceDiagnosisSummaryResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetStatusCode(v int32) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) SetBody(v *DescribeDBInstanceDiagnosisSummaryResponseBody) *DescribeDBInstanceDiagnosisSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
