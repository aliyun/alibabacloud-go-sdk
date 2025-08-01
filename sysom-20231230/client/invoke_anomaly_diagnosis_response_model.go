// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAnomalyDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeAnomalyDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeAnomalyDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *InvokeAnomalyDiagnosisResponseBody) *InvokeAnomalyDiagnosisResponse
	GetBody() *InvokeAnomalyDiagnosisResponseBody
}

type InvokeAnomalyDiagnosisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeAnomalyDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeAnomalyDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeAnomalyDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *InvokeAnomalyDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeAnomalyDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeAnomalyDiagnosisResponse) GetBody() *InvokeAnomalyDiagnosisResponseBody {
	return s.Body
}

func (s *InvokeAnomalyDiagnosisResponse) SetHeaders(v map[string]*string) *InvokeAnomalyDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *InvokeAnomalyDiagnosisResponse) SetStatusCode(v int32) *InvokeAnomalyDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponse) SetBody(v *InvokeAnomalyDiagnosisResponseBody) *InvokeAnomalyDiagnosisResponse {
	s.Body = v
	return s
}

func (s *InvokeAnomalyDiagnosisResponse) Validate() error {
	return dara.Validate(s)
}
