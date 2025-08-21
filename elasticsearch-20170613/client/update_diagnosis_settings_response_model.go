// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiagnosisSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDiagnosisSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDiagnosisSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDiagnosisSettingsResponseBody) *UpdateDiagnosisSettingsResponse
	GetBody() *UpdateDiagnosisSettingsResponseBody
}

type UpdateDiagnosisSettingsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDiagnosisSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDiagnosisSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiagnosisSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateDiagnosisSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDiagnosisSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDiagnosisSettingsResponse) GetBody() *UpdateDiagnosisSettingsResponseBody {
	return s.Body
}

func (s *UpdateDiagnosisSettingsResponse) SetHeaders(v map[string]*string) *UpdateDiagnosisSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateDiagnosisSettingsResponse) SetStatusCode(v int32) *UpdateDiagnosisSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDiagnosisSettingsResponse) SetBody(v *UpdateDiagnosisSettingsResponseBody) *UpdateDiagnosisSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateDiagnosisSettingsResponse) Validate() error {
	return dara.Validate(s)
}
