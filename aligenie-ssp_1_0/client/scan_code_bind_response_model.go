// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanCodeBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanCodeBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanCodeBindResponse
	GetStatusCode() *int32
	SetBody(v *ScanCodeBindResponseBody) *ScanCodeBindResponse
	GetBody() *ScanCodeBindResponseBody
}

type ScanCodeBindResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanCodeBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanCodeBindResponse) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindResponse) GoString() string {
	return s.String()
}

func (s *ScanCodeBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanCodeBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanCodeBindResponse) GetBody() *ScanCodeBindResponseBody {
	return s.Body
}

func (s *ScanCodeBindResponse) SetHeaders(v map[string]*string) *ScanCodeBindResponse {
	s.Headers = v
	return s
}

func (s *ScanCodeBindResponse) SetStatusCode(v int32) *ScanCodeBindResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanCodeBindResponse) SetBody(v *ScanCodeBindResponseBody) *ScanCodeBindResponse {
	s.Body = v
	return s
}

func (s *ScanCodeBindResponse) Validate() error {
	return dara.Validate(s)
}
