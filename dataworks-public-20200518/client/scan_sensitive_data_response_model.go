// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanSensitiveDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanSensitiveDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanSensitiveDataResponse
	GetStatusCode() *int32
	SetBody(v *ScanSensitiveDataResponseBody) *ScanSensitiveDataResponse
	GetBody() *ScanSensitiveDataResponseBody
}

type ScanSensitiveDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanSensitiveDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanSensitiveDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ScanSensitiveDataResponse) GoString() string {
	return s.String()
}

func (s *ScanSensitiveDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanSensitiveDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanSensitiveDataResponse) GetBody() *ScanSensitiveDataResponseBody {
	return s.Body
}

func (s *ScanSensitiveDataResponse) SetHeaders(v map[string]*string) *ScanSensitiveDataResponse {
	s.Headers = v
	return s
}

func (s *ScanSensitiveDataResponse) SetStatusCode(v int32) *ScanSensitiveDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanSensitiveDataResponse) SetBody(v *ScanSensitiveDataResponseBody) *ScanSensitiveDataResponse {
	s.Body = v
	return s
}

func (s *ScanSensitiveDataResponse) Validate() error {
	return dara.Validate(s)
}
