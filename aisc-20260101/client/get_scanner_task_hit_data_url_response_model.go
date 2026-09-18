// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScannerTaskHitDataUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScannerTaskHitDataUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScannerTaskHitDataUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetScannerTaskHitDataUrlResponseBody) *GetScannerTaskHitDataUrlResponse
	GetBody() *GetScannerTaskHitDataUrlResponseBody
}

type GetScannerTaskHitDataUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScannerTaskHitDataUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScannerTaskHitDataUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskHitDataUrlResponse) GoString() string {
	return s.String()
}

func (s *GetScannerTaskHitDataUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScannerTaskHitDataUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScannerTaskHitDataUrlResponse) GetBody() *GetScannerTaskHitDataUrlResponseBody {
	return s.Body
}

func (s *GetScannerTaskHitDataUrlResponse) SetHeaders(v map[string]*string) *GetScannerTaskHitDataUrlResponse {
	s.Headers = v
	return s
}

func (s *GetScannerTaskHitDataUrlResponse) SetStatusCode(v int32) *GetScannerTaskHitDataUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScannerTaskHitDataUrlResponse) SetBody(v *GetScannerTaskHitDataUrlResponseBody) *GetScannerTaskHitDataUrlResponse {
	s.Body = v
	return s
}

func (s *GetScannerTaskHitDataUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
