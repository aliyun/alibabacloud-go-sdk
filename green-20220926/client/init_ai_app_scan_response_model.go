// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAiAppScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitAiAppScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitAiAppScanResponse
	GetStatusCode() *int32
	SetBody(v *InitAiAppScanResponseBody) *InitAiAppScanResponse
	GetBody() *InitAiAppScanResponseBody
}

type InitAiAppScanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitAiAppScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitAiAppScanResponse) String() string {
	return dara.Prettify(s)
}

func (s InitAiAppScanResponse) GoString() string {
	return s.String()
}

func (s *InitAiAppScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitAiAppScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitAiAppScanResponse) GetBody() *InitAiAppScanResponseBody {
	return s.Body
}

func (s *InitAiAppScanResponse) SetHeaders(v map[string]*string) *InitAiAppScanResponse {
	s.Headers = v
	return s
}

func (s *InitAiAppScanResponse) SetStatusCode(v int32) *InitAiAppScanResponse {
	s.StatusCode = &v
	return s
}

func (s *InitAiAppScanResponse) SetBody(v *InitAiAppScanResponseBody) *InitAiAppScanResponse {
	s.Body = v
	return s
}

func (s *InitAiAppScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
