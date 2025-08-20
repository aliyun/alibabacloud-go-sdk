// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDiagnosisRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadDiagnosisRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse
	GetBody() *DownloadDiagnosisRecordsResponseBody
}

type DownloadDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadDiagnosisRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadDiagnosisRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadDiagnosisRecordsResponse) GetBody() *DownloadDiagnosisRecordsResponseBody {
	return s.Body
}

func (s *DownloadDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetStatusCode(v int32) *DownloadDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse {
	s.Body = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) Validate() error {
	return dara.Validate(s)
}
