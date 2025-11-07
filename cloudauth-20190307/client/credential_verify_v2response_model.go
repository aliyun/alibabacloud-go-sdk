// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialVerifyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialVerifyV2Response
	GetStatusCode() *int32
	SetBody(v *CredentialVerifyV2ResponseBody) *CredentialVerifyV2Response
	GetBody() *CredentialVerifyV2ResponseBody
}

type CredentialVerifyV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialVerifyV2Response) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2Response) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialVerifyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialVerifyV2Response) GetBody() *CredentialVerifyV2ResponseBody {
	return s.Body
}

func (s *CredentialVerifyV2Response) SetHeaders(v map[string]*string) *CredentialVerifyV2Response {
	s.Headers = v
	return s
}

func (s *CredentialVerifyV2Response) SetStatusCode(v int32) *CredentialVerifyV2Response {
	s.StatusCode = &v
	return s
}

func (s *CredentialVerifyV2Response) SetBody(v *CredentialVerifyV2ResponseBody) *CredentialVerifyV2Response {
	s.Body = v
	return s
}

func (s *CredentialVerifyV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
