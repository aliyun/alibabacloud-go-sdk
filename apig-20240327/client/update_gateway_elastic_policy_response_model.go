// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayElasticPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayElasticPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayElasticPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayElasticPolicyResponseBody) *UpdateGatewayElasticPolicyResponse
	GetBody() *UpdateGatewayElasticPolicyResponseBody
}

type UpdateGatewayElasticPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayElasticPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayElasticPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayElasticPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayElasticPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayElasticPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayElasticPolicyResponse) GetBody() *UpdateGatewayElasticPolicyResponseBody {
	return s.Body
}

func (s *UpdateGatewayElasticPolicyResponse) SetHeaders(v map[string]*string) *UpdateGatewayElasticPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayElasticPolicyResponse) SetStatusCode(v int32) *UpdateGatewayElasticPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayElasticPolicyResponse) SetBody(v *UpdateGatewayElasticPolicyResponseBody) *UpdateGatewayElasticPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayElasticPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
