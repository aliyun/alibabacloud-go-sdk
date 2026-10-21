// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyPolicyForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyPolicyForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyPolicyForAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyPolicyForAccountResponseBody) *GetAccessKeyPolicyForAccountResponse
	GetBody() *GetAccessKeyPolicyForAccountResponseBody
}

type GetAccessKeyPolicyForAccountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyPolicyForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyPolicyForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyPolicyForAccountResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyPolicyForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyPolicyForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyPolicyForAccountResponse) GetBody() *GetAccessKeyPolicyForAccountResponseBody {
	return s.Body
}

func (s *GetAccessKeyPolicyForAccountResponse) SetHeaders(v map[string]*string) *GetAccessKeyPolicyForAccountResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyPolicyForAccountResponse) SetStatusCode(v int32) *GetAccessKeyPolicyForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyPolicyForAccountResponse) SetBody(v *GetAccessKeyPolicyForAccountResponseBody) *GetAccessKeyPolicyForAccountResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyPolicyForAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
