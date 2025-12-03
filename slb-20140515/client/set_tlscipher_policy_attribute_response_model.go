// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTLSCipherPolicyAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTLSCipherPolicyAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTLSCipherPolicyAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetTLSCipherPolicyAttributeResponseBody) *SetTLSCipherPolicyAttributeResponse
	GetBody() *SetTLSCipherPolicyAttributeResponseBody
}

type SetTLSCipherPolicyAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTLSCipherPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTLSCipherPolicyAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTLSCipherPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetTLSCipherPolicyAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTLSCipherPolicyAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTLSCipherPolicyAttributeResponse) GetBody() *SetTLSCipherPolicyAttributeResponseBody {
	return s.Body
}

func (s *SetTLSCipherPolicyAttributeResponse) SetHeaders(v map[string]*string) *SetTLSCipherPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponse) SetStatusCode(v int32) *SetTLSCipherPolicyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponse) SetBody(v *SetTLSCipherPolicyAttributeResponseBody) *SetTLSCipherPolicyAttributeResponse {
	s.Body = v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
