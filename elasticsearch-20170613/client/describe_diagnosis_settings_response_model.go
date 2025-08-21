// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosisSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosisSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosisSettingsResponseBody) *DescribeDiagnosisSettingsResponse
	GetBody() *DescribeDiagnosisSettingsResponseBody
}

type DescribeDiagnosisSettingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosisSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosisSettingsResponse) GetBody() *DescribeDiagnosisSettingsResponseBody {
	return s.Body
}

func (s *DescribeDiagnosisSettingsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSettingsResponse) SetStatusCode(v int32) *DescribeDiagnosisSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponse) SetBody(v *DescribeDiagnosisSettingsResponseBody) *DescribeDiagnosisSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosisSettingsResponse) Validate() error {
	return dara.Validate(s)
}
