// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisFactorDistributionStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaultDiagnosisFactorDistributionStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaultDiagnosisFactorDistributionStatResponseBody) *DescribeFaultDiagnosisFactorDistributionStatResponse
	GetBody() *DescribeFaultDiagnosisFactorDistributionStatResponseBody
}

type DescribeFaultDiagnosisFactorDistributionStatResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisFactorDistributionStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) GetBody() *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	return s.Body
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetBody(v *DescribeFaultDiagnosisFactorDistributionStatResponseBody) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.Body = v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
