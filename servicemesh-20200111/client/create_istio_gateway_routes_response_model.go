// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIstioGatewayRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIstioGatewayRoutesResponse
	GetStatusCode() *int32
	SetBody(v *CreateIstioGatewayRoutesResponseBody) *CreateIstioGatewayRoutesResponse
	GetBody() *CreateIstioGatewayRoutesResponseBody
}

type CreateIstioGatewayRoutesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIstioGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIstioGatewayRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIstioGatewayRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIstioGatewayRoutesResponse) GetBody() *CreateIstioGatewayRoutesResponseBody {
	return s.Body
}

func (s *CreateIstioGatewayRoutesResponse) SetHeaders(v map[string]*string) *CreateIstioGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *CreateIstioGatewayRoutesResponse) SetStatusCode(v int32) *CreateIstioGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIstioGatewayRoutesResponse) SetBody(v *CreateIstioGatewayRoutesResponseBody) *CreateIstioGatewayRoutesResponse {
	s.Body = v
	return s
}

func (s *CreateIstioGatewayRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
