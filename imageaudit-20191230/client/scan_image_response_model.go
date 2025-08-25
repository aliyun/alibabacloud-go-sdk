// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanImageResponse
	GetStatusCode() *int32
	SetBody(v *ScanImageResponseBody) *ScanImageResponse
	GetBody() *ScanImageResponseBody
}

type ScanImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ScanImageResponse) GoString() string {
	return s.String()
}

func (s *ScanImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanImageResponse) GetBody() *ScanImageResponseBody {
	return s.Body
}

func (s *ScanImageResponse) SetHeaders(v map[string]*string) *ScanImageResponse {
	s.Headers = v
	return s
}

func (s *ScanImageResponse) SetStatusCode(v int32) *ScanImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanImageResponse) SetBody(v *ScanImageResponseBody) *ScanImageResponse {
	s.Body = v
	return s
}

func (s *ScanImageResponse) Validate() error {
	return dara.Validate(s)
}
