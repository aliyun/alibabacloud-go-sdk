// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScannerTaskResultHtmlUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScannerTaskResultHtmlUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScannerTaskResultHtmlUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetScannerTaskResultHtmlUrlResponseBody) *GetScannerTaskResultHtmlUrlResponse
	GetBody() *GetScannerTaskResultHtmlUrlResponseBody
}

type GetScannerTaskResultHtmlUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScannerTaskResultHtmlUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScannerTaskResultHtmlUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskResultHtmlUrlResponse) GoString() string {
	return s.String()
}

func (s *GetScannerTaskResultHtmlUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScannerTaskResultHtmlUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScannerTaskResultHtmlUrlResponse) GetBody() *GetScannerTaskResultHtmlUrlResponseBody {
	return s.Body
}

func (s *GetScannerTaskResultHtmlUrlResponse) SetHeaders(v map[string]*string) *GetScannerTaskResultHtmlUrlResponse {
	s.Headers = v
	return s
}

func (s *GetScannerTaskResultHtmlUrlResponse) SetStatusCode(v int32) *GetScannerTaskResultHtmlUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScannerTaskResultHtmlUrlResponse) SetBody(v *GetScannerTaskResultHtmlUrlResponseBody) *GetScannerTaskResultHtmlUrlResponse {
	s.Body = v
	return s
}

func (s *GetScannerTaskResultHtmlUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
