// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkSQLDiagnosisListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkSQLDiagnosisListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkSQLDiagnosisListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkSQLDiagnosisListResponseBody) *DescribeSparkSQLDiagnosisListResponse
	GetBody() *DescribeSparkSQLDiagnosisListResponseBody
}

type DescribeSparkSQLDiagnosisListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkSQLDiagnosisListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkSQLDiagnosisListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkSQLDiagnosisListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkSQLDiagnosisListResponse) GetBody() *DescribeSparkSQLDiagnosisListResponseBody {
	return s.Body
}

func (s *DescribeSparkSQLDiagnosisListResponse) SetHeaders(v map[string]*string) *DescribeSparkSQLDiagnosisListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponse) SetStatusCode(v int32) *DescribeSparkSQLDiagnosisListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponse) SetBody(v *DescribeSparkSQLDiagnosisListResponseBody) *DescribeSparkSQLDiagnosisListResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkSQLDiagnosisListResponse) Validate() error {
	return dara.Validate(s)
}
