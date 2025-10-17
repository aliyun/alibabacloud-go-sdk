// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisUserDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaultDiagnosisUserDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaultDiagnosisUserDetailResponseBody) *DescribeFaultDiagnosisUserDetailResponse
	GetBody() *DescribeFaultDiagnosisUserDetailResponseBody
}

type DescribeFaultDiagnosisUserDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaultDiagnosisUserDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaultDiagnosisUserDetailResponse) GetBody() *DescribeFaultDiagnosisUserDetailResponseBody {
	return s.Body
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisUserDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetBody(v *DescribeFaultDiagnosisUserDetailResponseBody) *DescribeFaultDiagnosisUserDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
