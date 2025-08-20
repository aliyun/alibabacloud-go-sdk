// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkSQLDiagnosisAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkSQLDiagnosisAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkSQLDiagnosisAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkSQLDiagnosisAttributeResponseBody) *DescribeSparkSQLDiagnosisAttributeResponse
	GetBody() *DescribeSparkSQLDiagnosisAttributeResponseBody
}

type DescribeSparkSQLDiagnosisAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkSQLDiagnosisAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkSQLDiagnosisAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) GetBody() *DescribeSparkSQLDiagnosisAttributeResponseBody {
	return s.Body
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) SetHeaders(v map[string]*string) *DescribeSparkSQLDiagnosisAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) SetStatusCode(v int32) *DescribeSparkSQLDiagnosisAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) SetBody(v *DescribeSparkSQLDiagnosisAttributeResponseBody) *DescribeSparkSQLDiagnosisAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponse) Validate() error {
	return dara.Validate(s)
}
