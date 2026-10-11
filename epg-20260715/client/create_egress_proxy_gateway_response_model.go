// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEgressProxyGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEgressProxyGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateEgressProxyGatewayResponseBody) *CreateEgressProxyGatewayResponse
	GetBody() *CreateEgressProxyGatewayResponseBody
}

type CreateEgressProxyGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEgressProxyGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEgressProxyGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEgressProxyGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEgressProxyGatewayResponse) GetBody() *CreateEgressProxyGatewayResponseBody {
	return s.Body
}

func (s *CreateEgressProxyGatewayResponse) SetHeaders(v map[string]*string) *CreateEgressProxyGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateEgressProxyGatewayResponse) SetStatusCode(v int32) *CreateEgressProxyGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEgressProxyGatewayResponse) SetBody(v *CreateEgressProxyGatewayResponseBody) *CreateEgressProxyGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateEgressProxyGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
