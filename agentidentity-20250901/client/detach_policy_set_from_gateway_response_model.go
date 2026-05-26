// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicySetFromGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPolicySetFromGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPolicySetFromGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DetachPolicySetFromGatewayResponseBody) *DetachPolicySetFromGatewayResponse
	GetBody() *DetachPolicySetFromGatewayResponseBody
}

type DetachPolicySetFromGatewayResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicySetFromGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicySetFromGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicySetFromGatewayResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicySetFromGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPolicySetFromGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPolicySetFromGatewayResponse) GetBody() *DetachPolicySetFromGatewayResponseBody {
	return s.Body
}

func (s *DetachPolicySetFromGatewayResponse) SetHeaders(v map[string]*string) *DetachPolicySetFromGatewayResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicySetFromGatewayResponse) SetStatusCode(v int32) *DetachPolicySetFromGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicySetFromGatewayResponse) SetBody(v *DetachPolicySetFromGatewayResponseBody) *DetachPolicySetFromGatewayResponse {
	s.Body = v
	return s
}

func (s *DetachPolicySetFromGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
