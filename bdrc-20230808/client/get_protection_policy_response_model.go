// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtectionPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProtectionPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProtectionPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetProtectionPolicyResponseBody) *GetProtectionPolicyResponse
	GetBody() *GetProtectionPolicyResponseBody
}

type GetProtectionPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProtectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProtectionPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProtectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetProtectionPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProtectionPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProtectionPolicyResponse) GetBody() *GetProtectionPolicyResponseBody {
	return s.Body
}

func (s *GetProtectionPolicyResponse) SetHeaders(v map[string]*string) *GetProtectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetProtectionPolicyResponse) SetStatusCode(v int32) *GetProtectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProtectionPolicyResponse) SetBody(v *GetProtectionPolicyResponseBody) *GetProtectionPolicyResponse {
	s.Body = v
	return s
}

func (s *GetProtectionPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
