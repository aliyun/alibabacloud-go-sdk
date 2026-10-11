// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEgressProxyGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEgressProxyGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetEgressProxyGatewayAttributeResponseBody) *GetEgressProxyGatewayAttributeResponse
	GetBody() *GetEgressProxyGatewayAttributeResponseBody
}

type GetEgressProxyGatewayAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEgressProxyGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEgressProxyGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEgressProxyGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEgressProxyGatewayAttributeResponse) GetBody() *GetEgressProxyGatewayAttributeResponseBody {
	return s.Body
}

func (s *GetEgressProxyGatewayAttributeResponse) SetHeaders(v map[string]*string) *GetEgressProxyGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponse) SetStatusCode(v int32) *GetEgressProxyGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponse) SetBody(v *GetEgressProxyGatewayAttributeResponseBody) *GetEgressProxyGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
