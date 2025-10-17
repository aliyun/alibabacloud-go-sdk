// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialRecognitionIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialRecognitionIntlResponse
	GetStatusCode() *int32
	SetBody(v *CredentialRecognitionIntlResponseBody) *CredentialRecognitionIntlResponse
	GetBody() *CredentialRecognitionIntlResponseBody
}

type CredentialRecognitionIntlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialRecognitionIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialRecognitionIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlResponse) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialRecognitionIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialRecognitionIntlResponse) GetBody() *CredentialRecognitionIntlResponseBody {
	return s.Body
}

func (s *CredentialRecognitionIntlResponse) SetHeaders(v map[string]*string) *CredentialRecognitionIntlResponse {
	s.Headers = v
	return s
}

func (s *CredentialRecognitionIntlResponse) SetStatusCode(v int32) *CredentialRecognitionIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *CredentialRecognitionIntlResponse) SetBody(v *CredentialRecognitionIntlResponseBody) *CredentialRecognitionIntlResponse {
	s.Body = v
	return s
}

func (s *CredentialRecognitionIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
