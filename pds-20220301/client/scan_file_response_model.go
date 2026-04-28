// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanFileResponse
	GetStatusCode() *int32
	SetBody(v *ScanFileResponseBody) *ScanFileResponse
	GetBody() *ScanFileResponseBody
}

type ScanFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ScanFileResponse) GoString() string {
	return s.String()
}

func (s *ScanFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanFileResponse) GetBody() *ScanFileResponseBody {
	return s.Body
}

func (s *ScanFileResponse) SetHeaders(v map[string]*string) *ScanFileResponse {
	s.Headers = v
	return s
}

func (s *ScanFileResponse) SetStatusCode(v int32) *ScanFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanFileResponse) SetBody(v *ScanFileResponseBody) *ScanFileResponse {
	s.Body = v
	return s
}

func (s *ScanFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
