// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanTextResponse
	GetStatusCode() *int32
	SetBody(v *ScanTextResponseBody) *ScanTextResponse
	GetBody() *ScanTextResponseBody
}

type ScanTextResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanTextResponse) String() string {
	return dara.Prettify(s)
}

func (s ScanTextResponse) GoString() string {
	return s.String()
}

func (s *ScanTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanTextResponse) GetBody() *ScanTextResponseBody {
	return s.Body
}

func (s *ScanTextResponse) SetHeaders(v map[string]*string) *ScanTextResponse {
	s.Headers = v
	return s
}

func (s *ScanTextResponse) SetStatusCode(v int32) *ScanTextResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanTextResponse) SetBody(v *ScanTextResponseBody) *ScanTextResponse {
	s.Body = v
	return s
}

func (s *ScanTextResponse) Validate() error {
	return dara.Validate(s)
}
