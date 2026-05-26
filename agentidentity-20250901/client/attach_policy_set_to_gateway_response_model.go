// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicySetToGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPolicySetToGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPolicySetToGatewayResponse
	GetStatusCode() *int32
	SetBody(v *AttachPolicySetToGatewayResponseBody) *AttachPolicySetToGatewayResponse
	GetBody() *AttachPolicySetToGatewayResponseBody
}

type AttachPolicySetToGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicySetToGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicySetToGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicySetToGatewayResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicySetToGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPolicySetToGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPolicySetToGatewayResponse) GetBody() *AttachPolicySetToGatewayResponseBody {
	return s.Body
}

func (s *AttachPolicySetToGatewayResponse) SetHeaders(v map[string]*string) *AttachPolicySetToGatewayResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicySetToGatewayResponse) SetStatusCode(v int32) *AttachPolicySetToGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicySetToGatewayResponse) SetBody(v *AttachPolicySetToGatewayResponseBody) *AttachPolicySetToGatewayResponse {
	s.Body = v
	return s
}

func (s *AttachPolicySetToGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
