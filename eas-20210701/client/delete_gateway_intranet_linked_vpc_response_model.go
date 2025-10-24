// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayIntranetLinkedVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayIntranetLinkedVpcResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayIntranetLinkedVpcResponseBody) *DeleteGatewayIntranetLinkedVpcResponse
	GetBody() *DeleteGatewayIntranetLinkedVpcResponseBody
}

type DeleteGatewayIntranetLinkedVpcResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayIntranetLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) GetBody() *DeleteGatewayIntranetLinkedVpcResponseBody {
	return s.Body
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) SetHeaders(v map[string]*string) *DeleteGatewayIntranetLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) SetStatusCode(v int32) *DeleteGatewayIntranetLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) SetBody(v *DeleteGatewayIntranetLinkedVpcResponseBody) *DeleteGatewayIntranetLinkedVpcResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
