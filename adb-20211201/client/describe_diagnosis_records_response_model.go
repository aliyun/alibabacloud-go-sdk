// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosisRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse
	GetBody() *DescribeDiagnosisRecordsResponseBody
}

type DescribeDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosisRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosisRecordsResponse) GetBody() *DescribeDiagnosisRecordsResponseBody {
	return s.Body
}

func (s *DescribeDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetStatusCode(v int32) *DescribeDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) Validate() error {
	return dara.Validate(s)
}
