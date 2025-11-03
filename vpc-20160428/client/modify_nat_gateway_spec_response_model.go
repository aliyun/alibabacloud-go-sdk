// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewaySpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatGatewaySpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatGatewaySpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatGatewaySpecResponseBody) *ModifyNatGatewaySpecResponse
	GetBody() *ModifyNatGatewaySpecResponseBody
}

type ModifyNatGatewaySpecResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatGatewaySpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatGatewaySpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewaySpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewaySpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatGatewaySpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatGatewaySpecResponse) GetBody() *ModifyNatGatewaySpecResponseBody {
	return s.Body
}

func (s *ModifyNatGatewaySpecResponse) SetHeaders(v map[string]*string) *ModifyNatGatewaySpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatGatewaySpecResponse) SetStatusCode(v int32) *ModifyNatGatewaySpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatGatewaySpecResponse) SetBody(v *ModifyNatGatewaySpecResponseBody) *ModifyNatGatewaySpecResponse {
	s.Body = v
	return s
}

func (s *ModifyNatGatewaySpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
