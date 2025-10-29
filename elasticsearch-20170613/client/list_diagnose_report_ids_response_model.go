// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseReportIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnoseReportIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnoseReportIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnoseReportIdsResponseBody) *ListDiagnoseReportIdsResponse
	GetBody() *ListDiagnoseReportIdsResponseBody
}

type ListDiagnoseReportIdsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnoseReportIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnoseReportIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportIdsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnoseReportIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnoseReportIdsResponse) GetBody() *ListDiagnoseReportIdsResponseBody {
	return s.Body
}

func (s *ListDiagnoseReportIdsResponse) SetHeaders(v map[string]*string) *ListDiagnoseReportIdsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportIdsResponse) SetStatusCode(v int32) *ListDiagnoseReportIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnoseReportIdsResponse) SetBody(v *ListDiagnoseReportIdsResponseBody) *ListDiagnoseReportIdsResponse {
	s.Body = v
	return s
}

func (s *ListDiagnoseReportIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
