// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRspDomainServerHoldStatusForGatewayOteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRspDomainServerHoldStatusForGatewayOteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRspDomainServerHoldStatusForGatewayOteResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) *RemoveRspDomainServerHoldStatusForGatewayOteResponse
	GetBody() *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody
}

type RemoveRspDomainServerHoldStatusForGatewayOteResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponse) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) GetBody() *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody {
	return s.Body
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) SetHeaders(v map[string]*string) *RemoveRspDomainServerHoldStatusForGatewayOteResponse {
	s.Headers = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) SetStatusCode(v int32) *RemoveRspDomainServerHoldStatusForGatewayOteResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) SetBody(v *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) *RemoveRspDomainServerHoldStatusForGatewayOteResponse {
	s.Body = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
