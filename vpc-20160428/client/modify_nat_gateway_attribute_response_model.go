// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatGatewayAttributeResponseBody) *ModifyNatGatewayAttributeResponse
	GetBody() *ModifyNatGatewayAttributeResponseBody
}

type ModifyNatGatewayAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatGatewayAttributeResponse) GetBody() *ModifyNatGatewayAttributeResponseBody {
	return s.Body
}

func (s *ModifyNatGatewayAttributeResponse) SetHeaders(v map[string]*string) *ModifyNatGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatGatewayAttributeResponse) SetStatusCode(v int32) *ModifyNatGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatGatewayAttributeResponse) SetBody(v *ModifyNatGatewayAttributeResponseBody) *ModifyNatGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNatGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
