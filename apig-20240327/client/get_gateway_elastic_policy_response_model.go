// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayElasticPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayElasticPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayElasticPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayElasticPolicyResponseBody) *GetGatewayElasticPolicyResponse
	GetBody() *GetGatewayElasticPolicyResponseBody
}

type GetGatewayElasticPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayElasticPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayElasticPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayElasticPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayElasticPolicyResponse) GetBody() *GetGatewayElasticPolicyResponseBody {
	return s.Body
}

func (s *GetGatewayElasticPolicyResponse) SetHeaders(v map[string]*string) *GetGatewayElasticPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayElasticPolicyResponse) SetStatusCode(v int32) *GetGatewayElasticPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayElasticPolicyResponse) SetBody(v *GetGatewayElasticPolicyResponseBody) *GetGatewayElasticPolicyResponse {
	s.Body = v
	return s
}

func (s *GetGatewayElasticPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
