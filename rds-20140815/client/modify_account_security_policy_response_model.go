// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountSecurityPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountSecurityPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountSecurityPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountSecurityPolicyResponseBody) *ModifyAccountSecurityPolicyResponse
	GetBody() *ModifyAccountSecurityPolicyResponseBody
}

type ModifyAccountSecurityPolicyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountSecurityPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountSecurityPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountSecurityPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountSecurityPolicyResponse) GetBody() *ModifyAccountSecurityPolicyResponseBody {
	return s.Body
}

func (s *ModifyAccountSecurityPolicyResponse) SetHeaders(v map[string]*string) *ModifyAccountSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountSecurityPolicyResponse) SetStatusCode(v int32) *ModifyAccountSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountSecurityPolicyResponse) SetBody(v *ModifyAccountSecurityPolicyResponseBody) *ModifyAccountSecurityPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountSecurityPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
