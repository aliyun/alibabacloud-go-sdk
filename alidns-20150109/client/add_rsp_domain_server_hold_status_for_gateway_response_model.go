// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRspDomainServerHoldStatusForGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRspDomainServerHoldStatusForGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRspDomainServerHoldStatusForGatewayResponse
	GetStatusCode() *int32
	SetBody(v *AddRspDomainServerHoldStatusForGatewayResponseBody) *AddRspDomainServerHoldStatusForGatewayResponse
	GetBody() *AddRspDomainServerHoldStatusForGatewayResponseBody
}

type AddRspDomainServerHoldStatusForGatewayResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRspDomainServerHoldStatusForGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayResponse) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) GetBody() *AddRspDomainServerHoldStatusForGatewayResponseBody {
	return s.Body
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) SetHeaders(v map[string]*string) *AddRspDomainServerHoldStatusForGatewayResponse {
	s.Headers = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) SetStatusCode(v int32) *AddRspDomainServerHoldStatusForGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) SetBody(v *AddRspDomainServerHoldStatusForGatewayResponseBody) *AddRspDomainServerHoldStatusForGatewayResponse {
	s.Body = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
