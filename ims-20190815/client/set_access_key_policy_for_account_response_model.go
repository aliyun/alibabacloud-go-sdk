// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessKeyPolicyForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAccessKeyPolicyForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAccessKeyPolicyForAccountResponse
	GetStatusCode() *int32
	SetBody(v *SetAccessKeyPolicyForAccountResponseBody) *SetAccessKeyPolicyForAccountResponse
	GetBody() *SetAccessKeyPolicyForAccountResponseBody
}

type SetAccessKeyPolicyForAccountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessKeyPolicyForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessKeyPolicyForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAccessKeyPolicyForAccountResponse) GoString() string {
	return s.String()
}

func (s *SetAccessKeyPolicyForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAccessKeyPolicyForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAccessKeyPolicyForAccountResponse) GetBody() *SetAccessKeyPolicyForAccountResponseBody {
	return s.Body
}

func (s *SetAccessKeyPolicyForAccountResponse) SetHeaders(v map[string]*string) *SetAccessKeyPolicyForAccountResponse {
	s.Headers = v
	return s
}

func (s *SetAccessKeyPolicyForAccountResponse) SetStatusCode(v int32) *SetAccessKeyPolicyForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessKeyPolicyForAccountResponse) SetBody(v *SetAccessKeyPolicyForAccountResponseBody) *SetAccessKeyPolicyForAccountResponse {
	s.Body = v
	return s
}

func (s *SetAccessKeyPolicyForAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
