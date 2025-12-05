// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKeyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKeyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetKeyPolicyResponseBody) *GetKeyPolicyResponse
	GetBody() *GetKeyPolicyResponseBody
}

type GetKeyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKeyPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetKeyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKeyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKeyPolicyResponse) GetBody() *GetKeyPolicyResponseBody {
	return s.Body
}

func (s *GetKeyPolicyResponse) SetHeaders(v map[string]*string) *GetKeyPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetKeyPolicyResponse) SetStatusCode(v int32) *GetKeyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeyPolicyResponse) SetBody(v *GetKeyPolicyResponseBody) *GetKeyPolicyResponse {
	s.Body = v
	return s
}

func (s *GetKeyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
