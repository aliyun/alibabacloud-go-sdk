// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyImageSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyImageSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyImageSignatureResponse
	GetStatusCode() *int32
	SetBody(v *VerifyImageSignatureResponseBody) *VerifyImageSignatureResponse
	GetBody() *VerifyImageSignatureResponseBody
}

type VerifyImageSignatureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyImageSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyImageSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureResponse) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyImageSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyImageSignatureResponse) GetBody() *VerifyImageSignatureResponseBody {
	return s.Body
}

func (s *VerifyImageSignatureResponse) SetHeaders(v map[string]*string) *VerifyImageSignatureResponse {
	s.Headers = v
	return s
}

func (s *VerifyImageSignatureResponse) SetStatusCode(v int32) *VerifyImageSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyImageSignatureResponse) SetBody(v *VerifyImageSignatureResponseBody) *VerifyImageSignatureResponse {
	s.Body = v
	return s
}

func (s *VerifyImageSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
