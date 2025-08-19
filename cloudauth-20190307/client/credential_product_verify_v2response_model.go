// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialProductVerifyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialProductVerifyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialProductVerifyV2Response
	GetStatusCode() *int32
	SetBody(v *CredentialProductVerifyV2ResponseBody) *CredentialProductVerifyV2Response
	GetBody() *CredentialProductVerifyV2ResponseBody
}

type CredentialProductVerifyV2Response struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialProductVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialProductVerifyV2Response) String() string {
	return dara.Prettify(s)
}

func (s CredentialProductVerifyV2Response) GoString() string {
	return s.String()
}

func (s *CredentialProductVerifyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialProductVerifyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialProductVerifyV2Response) GetBody() *CredentialProductVerifyV2ResponseBody {
	return s.Body
}

func (s *CredentialProductVerifyV2Response) SetHeaders(v map[string]*string) *CredentialProductVerifyV2Response {
	s.Headers = v
	return s
}

func (s *CredentialProductVerifyV2Response) SetStatusCode(v int32) *CredentialProductVerifyV2Response {
	s.StatusCode = &v
	return s
}

func (s *CredentialProductVerifyV2Response) SetBody(v *CredentialProductVerifyV2ResponseBody) *CredentialProductVerifyV2Response {
	s.Body = v
	return s
}

func (s *CredentialProductVerifyV2Response) Validate() error {
	return dara.Validate(s)
}
