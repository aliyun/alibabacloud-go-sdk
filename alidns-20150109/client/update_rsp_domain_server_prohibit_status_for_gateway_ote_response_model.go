// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerProhibitStatusForGatewayOteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRspDomainServerProhibitStatusForGatewayOteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) *UpdateRspDomainServerProhibitStatusForGatewayOteResponse
	GetBody() *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody
}

type UpdateRspDomainServerProhibitStatusForGatewayOteResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponse) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) GetBody() *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody {
	return s.Body
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) SetHeaders(v map[string]*string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponse {
	s.Headers = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) SetStatusCode(v int32) *UpdateRspDomainServerProhibitStatusForGatewayOteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) SetBody(v *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) *UpdateRspDomainServerProhibitStatusForGatewayOteResponse {
	s.Body = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
