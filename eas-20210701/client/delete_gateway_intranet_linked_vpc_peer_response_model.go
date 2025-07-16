// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcPeerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayIntranetLinkedVpcPeerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayIntranetLinkedVpcPeerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayIntranetLinkedVpcPeerResponseBody) *DeleteGatewayIntranetLinkedVpcPeerResponse
	GetBody() *DeleteGatewayIntranetLinkedVpcPeerResponseBody
}

type DeleteGatewayIntranetLinkedVpcPeerResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayIntranetLinkedVpcPeerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcPeerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcPeerResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) GetBody() *DeleteGatewayIntranetLinkedVpcPeerResponseBody {
	return s.Body
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) SetHeaders(v map[string]*string) *DeleteGatewayIntranetLinkedVpcPeerResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) SetStatusCode(v int32) *DeleteGatewayIntranetLinkedVpcPeerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) SetBody(v *DeleteGatewayIntranetLinkedVpcPeerResponseBody) *DeleteGatewayIntranetLinkedVpcPeerResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponse) Validate() error {
	return dara.Validate(s)
}
