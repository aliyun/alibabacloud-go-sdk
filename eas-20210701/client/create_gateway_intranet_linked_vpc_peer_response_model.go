// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcPeerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayIntranetLinkedVpcPeerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayIntranetLinkedVpcPeerResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayIntranetLinkedVpcPeerResponseBody) *CreateGatewayIntranetLinkedVpcPeerResponse
	GetBody() *CreateGatewayIntranetLinkedVpcPeerResponseBody
}

type CreateGatewayIntranetLinkedVpcPeerResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayIntranetLinkedVpcPeerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcPeerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcPeerResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) GetBody() *CreateGatewayIntranetLinkedVpcPeerResponseBody {
	return s.Body
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) SetHeaders(v map[string]*string) *CreateGatewayIntranetLinkedVpcPeerResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) SetStatusCode(v int32) *CreateGatewayIntranetLinkedVpcPeerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) SetBody(v *CreateGatewayIntranetLinkedVpcPeerResponseBody) *CreateGatewayIntranetLinkedVpcPeerResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
