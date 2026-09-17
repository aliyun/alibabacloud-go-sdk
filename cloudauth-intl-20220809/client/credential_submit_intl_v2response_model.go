// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialSubmitIntlV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialSubmitIntlV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialSubmitIntlV2Response
	GetStatusCode() *int32
	SetBody(v *CredentialSubmitIntlV2ResponseBody) *CredentialSubmitIntlV2Response
	GetBody() *CredentialSubmitIntlV2ResponseBody
}

type CredentialSubmitIntlV2Response struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialSubmitIntlV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialSubmitIntlV2Response) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlV2Response) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialSubmitIntlV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialSubmitIntlV2Response) GetBody() *CredentialSubmitIntlV2ResponseBody {
	return s.Body
}

func (s *CredentialSubmitIntlV2Response) SetHeaders(v map[string]*string) *CredentialSubmitIntlV2Response {
	s.Headers = v
	return s
}

func (s *CredentialSubmitIntlV2Response) SetStatusCode(v int32) *CredentialSubmitIntlV2Response {
	s.StatusCode = &v
	return s
}

func (s *CredentialSubmitIntlV2Response) SetBody(v *CredentialSubmitIntlV2ResponseBody) *CredentialSubmitIntlV2Response {
	s.Body = v
	return s
}

func (s *CredentialSubmitIntlV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
