// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialRecognitionIntlV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialRecognitionIntlV2Response
	GetStatusCode() *int32
	SetBody(v *CredentialRecognitionIntlV2ResponseBody) *CredentialRecognitionIntlV2Response
	GetBody() *CredentialRecognitionIntlV2ResponseBody
}

type CredentialRecognitionIntlV2Response struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialRecognitionIntlV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialRecognitionIntlV2Response) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlV2Response) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialRecognitionIntlV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialRecognitionIntlV2Response) GetBody() *CredentialRecognitionIntlV2ResponseBody {
	return s.Body
}

func (s *CredentialRecognitionIntlV2Response) SetHeaders(v map[string]*string) *CredentialRecognitionIntlV2Response {
	s.Headers = v
	return s
}

func (s *CredentialRecognitionIntlV2Response) SetStatusCode(v int32) *CredentialRecognitionIntlV2Response {
	s.StatusCode = &v
	return s
}

func (s *CredentialRecognitionIntlV2Response) SetBody(v *CredentialRecognitionIntlV2ResponseBody) *CredentialRecognitionIntlV2Response {
	s.Body = v
	return s
}

func (s *CredentialRecognitionIntlV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
