// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatGatewayStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNatGatewayStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNatGatewayStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetNatGatewayStatusResponseBody) *GetNatGatewayStatusResponse
	GetBody() *GetNatGatewayStatusResponseBody
}

type GetNatGatewayStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNatGatewayStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNatGatewayStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayStatusResponse) GoString() string {
	return s.String()
}

func (s *GetNatGatewayStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNatGatewayStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNatGatewayStatusResponse) GetBody() *GetNatGatewayStatusResponseBody {
	return s.Body
}

func (s *GetNatGatewayStatusResponse) SetHeaders(v map[string]*string) *GetNatGatewayStatusResponse {
	s.Headers = v
	return s
}

func (s *GetNatGatewayStatusResponse) SetStatusCode(v int32) *GetNatGatewayStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNatGatewayStatusResponse) SetBody(v *GetNatGatewayStatusResponseBody) *GetNatGatewayStatusResponse {
	s.Body = v
	return s
}

func (s *GetNatGatewayStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
