// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetLinkedVpcPeerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayIntranetLinkedVpcPeerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayIntranetLinkedVpcPeerResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayIntranetLinkedVpcPeerResponseBody) *ListGatewayIntranetLinkedVpcPeerResponse
	GetBody() *ListGatewayIntranetLinkedVpcPeerResponseBody
}

type ListGatewayIntranetLinkedVpcPeerResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayIntranetLinkedVpcPeerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcPeerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcPeerResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) GetBody() *ListGatewayIntranetLinkedVpcPeerResponseBody {
	return s.Body
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) SetHeaders(v map[string]*string) *ListGatewayIntranetLinkedVpcPeerResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) SetStatusCode(v int32) *ListGatewayIntranetLinkedVpcPeerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) SetBody(v *ListGatewayIntranetLinkedVpcPeerResponseBody) *ListGatewayIntranetLinkedVpcPeerResponse {
	s.Body = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerResponse) Validate() error {
	return dara.Validate(s)
}
