// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAppDiagnosisInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkAppDiagnosisInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkAppDiagnosisInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkAppDiagnosisInfoResponseBody) *DescribeSparkAppDiagnosisInfoResponse
	GetBody() *DescribeSparkAppDiagnosisInfoResponseBody
}

type DescribeSparkAppDiagnosisInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkAppDiagnosisInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkAppDiagnosisInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAppDiagnosisInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkAppDiagnosisInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkAppDiagnosisInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkAppDiagnosisInfoResponse) GetBody() *DescribeSparkAppDiagnosisInfoResponseBody {
	return s.Body
}

func (s *DescribeSparkAppDiagnosisInfoResponse) SetHeaders(v map[string]*string) *DescribeSparkAppDiagnosisInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponse) SetStatusCode(v int32) *DescribeSparkAppDiagnosisInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponse) SetBody(v *DescribeSparkAppDiagnosisInfoResponseBody) *DescribeSparkAppDiagnosisInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponse) Validate() error {
	return dara.Validate(s)
}
