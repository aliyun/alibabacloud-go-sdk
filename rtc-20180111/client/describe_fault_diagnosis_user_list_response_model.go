// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaultDiagnosisUserListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaultDiagnosisUserListResponseBody) *DescribeFaultDiagnosisUserListResponse
	GetBody() *DescribeFaultDiagnosisUserListResponseBody
}

type DescribeFaultDiagnosisUserListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaultDiagnosisUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaultDiagnosisUserListResponse) GetBody() *DescribeFaultDiagnosisUserListResponseBody {
	return s.Body
}

func (s *DescribeFaultDiagnosisUserListResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) SetBody(v *DescribeFaultDiagnosisUserListResponseBody) *DescribeFaultDiagnosisUserListResponse {
	s.Body = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
