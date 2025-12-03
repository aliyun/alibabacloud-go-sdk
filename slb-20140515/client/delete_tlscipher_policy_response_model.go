// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTLSCipherPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTLSCipherPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTLSCipherPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTLSCipherPolicyResponseBody) *DeleteTLSCipherPolicyResponse
	GetBody() *DeleteTLSCipherPolicyResponseBody
}

type DeleteTLSCipherPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTLSCipherPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTLSCipherPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTLSCipherPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteTLSCipherPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTLSCipherPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTLSCipherPolicyResponse) GetBody() *DeleteTLSCipherPolicyResponseBody {
	return s.Body
}

func (s *DeleteTLSCipherPolicyResponse) SetHeaders(v map[string]*string) *DeleteTLSCipherPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteTLSCipherPolicyResponse) SetStatusCode(v int32) *DeleteTLSCipherPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTLSCipherPolicyResponse) SetBody(v *DeleteTLSCipherPolicyResponseBody) *DeleteTLSCipherPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteTLSCipherPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
