// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyPolicyResponseBody) *GetAccessKeyPolicyResponse
	GetBody() *GetAccessKeyPolicyResponseBody
}

type GetAccessKeyPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyPolicyResponse) GetBody() *GetAccessKeyPolicyResponseBody {
	return s.Body
}

func (s *GetAccessKeyPolicyResponse) SetHeaders(v map[string]*string) *GetAccessKeyPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyPolicyResponse) SetStatusCode(v int32) *GetAccessKeyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyPolicyResponse) SetBody(v *GetAccessKeyPolicyResponseBody) *GetAccessKeyPolicyResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
