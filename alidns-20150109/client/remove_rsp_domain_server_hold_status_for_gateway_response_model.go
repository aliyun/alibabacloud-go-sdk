// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRspDomainServerHoldStatusForGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRspDomainServerHoldStatusForGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRspDomainServerHoldStatusForGatewayResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRspDomainServerHoldStatusForGatewayResponseBody) *RemoveRspDomainServerHoldStatusForGatewayResponse
	GetBody() *RemoveRspDomainServerHoldStatusForGatewayResponseBody
}

type RemoveRspDomainServerHoldStatusForGatewayResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRspDomainServerHoldStatusForGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRspDomainServerHoldStatusForGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayResponse) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) GetBody() *RemoveRspDomainServerHoldStatusForGatewayResponseBody {
	return s.Body
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) SetHeaders(v map[string]*string) *RemoveRspDomainServerHoldStatusForGatewayResponse {
	s.Headers = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) SetStatusCode(v int32) *RemoveRspDomainServerHoldStatusForGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) SetBody(v *RemoveRspDomainServerHoldStatusForGatewayResponseBody) *RemoveRspDomainServerHoldStatusForGatewayResponse {
	s.Body = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
