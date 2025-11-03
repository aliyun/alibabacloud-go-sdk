// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNatGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNatGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetNatGatewayAttributeResponseBody) *GetNatGatewayAttributeResponse
	GetBody() *GetNatGatewayAttributeResponseBody
}

type GetNatGatewayAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNatGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNatGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNatGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNatGatewayAttributeResponse) GetBody() *GetNatGatewayAttributeResponseBody {
	return s.Body
}

func (s *GetNatGatewayAttributeResponse) SetHeaders(v map[string]*string) *GetNatGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetNatGatewayAttributeResponse) SetStatusCode(v int32) *GetNatGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNatGatewayAttributeResponse) SetBody(v *GetNatGatewayAttributeResponseBody) *GetNatGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *GetNatGatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
