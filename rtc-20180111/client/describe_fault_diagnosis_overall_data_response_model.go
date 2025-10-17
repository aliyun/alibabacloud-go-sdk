// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisOverallDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaultDiagnosisOverallDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaultDiagnosisOverallDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaultDiagnosisOverallDataResponseBody) *DescribeFaultDiagnosisOverallDataResponse
	GetBody() *DescribeFaultDiagnosisOverallDataResponseBody
}

type DescribeFaultDiagnosisOverallDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaultDiagnosisOverallDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaultDiagnosisOverallDataResponse) GetBody() *DescribeFaultDiagnosisOverallDataResponseBody {
	return s.Body
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetBody(v *DescribeFaultDiagnosisOverallDataResponseBody) *DescribeFaultDiagnosisOverallDataResponse {
	s.Body = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
