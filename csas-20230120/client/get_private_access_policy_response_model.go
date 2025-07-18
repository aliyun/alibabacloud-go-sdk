// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetPrivateAccessPolicyResponseBody) *GetPrivateAccessPolicyResponse
	GetBody() *GetPrivateAccessPolicyResponseBody
}

type GetPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrivateAccessPolicyResponse) GetBody() *GetPrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *GetPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *GetPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateAccessPolicyResponse) SetStatusCode(v int32) *GetPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateAccessPolicyResponse) SetBody(v *GetPrivateAccessPolicyResponseBody) *GetPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *GetPrivateAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
