// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpv4GatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpv4GatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpv4GatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetIpv4GatewayAttributeResponseBody) *GetIpv4GatewayAttributeResponse
	GetBody() *GetIpv4GatewayAttributeResponseBody
}

type GetIpv4GatewayAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpv4GatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpv4GatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpv4GatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetIpv4GatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpv4GatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpv4GatewayAttributeResponse) GetBody() *GetIpv4GatewayAttributeResponseBody {
	return s.Body
}

func (s *GetIpv4GatewayAttributeResponse) SetHeaders(v map[string]*string) *GetIpv4GatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetIpv4GatewayAttributeResponse) SetStatusCode(v int32) *GetIpv4GatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponse) SetBody(v *GetIpv4GatewayAttributeResponseBody) *GetIpv4GatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *GetIpv4GatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
