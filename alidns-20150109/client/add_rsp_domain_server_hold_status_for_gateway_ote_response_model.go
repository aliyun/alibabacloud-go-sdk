// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRspDomainServerHoldStatusForGatewayOteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRspDomainServerHoldStatusForGatewayOteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRspDomainServerHoldStatusForGatewayOteResponse
	GetStatusCode() *int32
	SetBody(v *AddRspDomainServerHoldStatusForGatewayOteResponseBody) *AddRspDomainServerHoldStatusForGatewayOteResponse
	GetBody() *AddRspDomainServerHoldStatusForGatewayOteResponseBody
}

type AddRspDomainServerHoldStatusForGatewayOteResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRspDomainServerHoldStatusForGatewayOteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponse) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) GetBody() *AddRspDomainServerHoldStatusForGatewayOteResponseBody {
	return s.Body
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) SetHeaders(v map[string]*string) *AddRspDomainServerHoldStatusForGatewayOteResponse {
	s.Headers = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) SetStatusCode(v int32) *AddRspDomainServerHoldStatusForGatewayOteResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) SetBody(v *AddRspDomainServerHoldStatusForGatewayOteResponseBody) *AddRspDomainServerHoldStatusForGatewayOteResponse {
	s.Body = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
