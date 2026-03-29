// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataDiagnosisStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDataDiagnosisStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDataDiagnosisStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *QueryDataDiagnosisStatisticsResponseBody) *QueryDataDiagnosisStatisticsResponse
	GetBody() *QueryDataDiagnosisStatisticsResponseBody
}

type QueryDataDiagnosisStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataDiagnosisStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataDiagnosisStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataDiagnosisStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QueryDataDiagnosisStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDataDiagnosisStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDataDiagnosisStatisticsResponse) GetBody() *QueryDataDiagnosisStatisticsResponseBody {
	return s.Body
}

func (s *QueryDataDiagnosisStatisticsResponse) SetHeaders(v map[string]*string) *QueryDataDiagnosisStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponse) SetStatusCode(v int32) *QueryDataDiagnosisStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponse) SetBody(v *QueryDataDiagnosisStatisticsResponseBody) *QueryDataDiagnosisStatisticsResponse {
	s.Body = v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
