// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPendingCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPendingCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPendingCertificateResponse
	GetStatusCode() *int32
	SetBody(v *CancelPendingCertificateResponseBody) *CancelPendingCertificateResponse
	GetBody() *CancelPendingCertificateResponseBody
}

type CancelPendingCertificateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPendingCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPendingCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPendingCertificateResponse) GoString() string {
	return s.String()
}

func (s *CancelPendingCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPendingCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPendingCertificateResponse) GetBody() *CancelPendingCertificateResponseBody {
	return s.Body
}

func (s *CancelPendingCertificateResponse) SetHeaders(v map[string]*string) *CancelPendingCertificateResponse {
	s.Headers = v
	return s
}

func (s *CancelPendingCertificateResponse) SetStatusCode(v int32) *CancelPendingCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPendingCertificateResponse) SetBody(v *CancelPendingCertificateResponseBody) *CancelPendingCertificateResponse {
	s.Body = v
	return s
}

func (s *CancelPendingCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
