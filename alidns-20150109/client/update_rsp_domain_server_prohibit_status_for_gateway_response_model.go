// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerProhibitStatusForGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRspDomainServerProhibitStatusForGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRspDomainServerProhibitStatusForGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) *UpdateRspDomainServerProhibitStatusForGatewayResponse
	GetBody() *UpdateRspDomainServerProhibitStatusForGatewayResponseBody
}

type UpdateRspDomainServerProhibitStatusForGatewayResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRspDomainServerProhibitStatusForGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) GetBody() *UpdateRspDomainServerProhibitStatusForGatewayResponseBody {
	return s.Body
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) SetHeaders(v map[string]*string) *UpdateRspDomainServerProhibitStatusForGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) SetStatusCode(v int32) *UpdateRspDomainServerProhibitStatusForGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) SetBody(v *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) *UpdateRspDomainServerProhibitStatusForGatewayResponse {
	s.Body = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
