// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayIntranetLinkedVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayIntranetLinkedVpcResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayIntranetLinkedVpcResponseBody) *CreateGatewayIntranetLinkedVpcResponse
	GetBody() *CreateGatewayIntranetLinkedVpcResponseBody
}

type CreateGatewayIntranetLinkedVpcResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayIntranetLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayIntranetLinkedVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayIntranetLinkedVpcResponse) GetBody() *CreateGatewayIntranetLinkedVpcResponseBody {
	return s.Body
}

func (s *CreateGatewayIntranetLinkedVpcResponse) SetHeaders(v map[string]*string) *CreateGatewayIntranetLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponse) SetStatusCode(v int32) *CreateGatewayIntranetLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponse) SetBody(v *CreateGatewayIntranetLinkedVpcResponseBody) *CreateGatewayIntranetLinkedVpcResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponse) Validate() error {
	return dara.Validate(s)
}
