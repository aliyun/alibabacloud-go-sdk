// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetLinkedVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayIntranetLinkedVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayIntranetLinkedVpcResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayIntranetLinkedVpcResponseBody) *ListGatewayIntranetLinkedVpcResponse
	GetBody() *ListGatewayIntranetLinkedVpcResponseBody
}

type ListGatewayIntranetLinkedVpcResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayIntranetLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayIntranetLinkedVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayIntranetLinkedVpcResponse) GetBody() *ListGatewayIntranetLinkedVpcResponseBody {
	return s.Body
}

func (s *ListGatewayIntranetLinkedVpcResponse) SetHeaders(v map[string]*string) *ListGatewayIntranetLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponse) SetStatusCode(v int32) *ListGatewayIntranetLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponse) SetBody(v *ListGatewayIntranetLinkedVpcResponseBody) *ListGatewayIntranetLinkedVpcResponse {
	s.Body = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponse) Validate() error {
	return dara.Validate(s)
}
