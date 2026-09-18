// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopScannerTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopScannerTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopScannerTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopScannerTaskResponseBody) *StopScannerTaskResponse
	GetBody() *StopScannerTaskResponseBody
}

type StopScannerTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopScannerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopScannerTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopScannerTaskResponse) GoString() string {
	return s.String()
}

func (s *StopScannerTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopScannerTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopScannerTaskResponse) GetBody() *StopScannerTaskResponseBody {
	return s.Body
}

func (s *StopScannerTaskResponse) SetHeaders(v map[string]*string) *StopScannerTaskResponse {
	s.Headers = v
	return s
}

func (s *StopScannerTaskResponse) SetStatusCode(v int32) *StopScannerTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopScannerTaskResponse) SetBody(v *StopScannerTaskResponseBody) *StopScannerTaskResponse {
	s.Body = v
	return s
}

func (s *StopScannerTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
