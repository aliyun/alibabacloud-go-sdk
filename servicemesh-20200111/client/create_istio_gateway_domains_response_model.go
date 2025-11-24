// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIstioGatewayDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIstioGatewayDomainsResponse
	GetStatusCode() *int32
	SetBody(v *CreateIstioGatewayDomainsResponseBody) *CreateIstioGatewayDomainsResponse
	GetBody() *CreateIstioGatewayDomainsResponseBody
}

type CreateIstioGatewayDomainsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIstioGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIstioGatewayDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIstioGatewayDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIstioGatewayDomainsResponse) GetBody() *CreateIstioGatewayDomainsResponseBody {
	return s.Body
}

func (s *CreateIstioGatewayDomainsResponse) SetHeaders(v map[string]*string) *CreateIstioGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *CreateIstioGatewayDomainsResponse) SetStatusCode(v int32) *CreateIstioGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIstioGatewayDomainsResponse) SetBody(v *CreateIstioGatewayDomainsResponseBody) *CreateIstioGatewayDomainsResponse {
	s.Body = v
	return s
}

func (s *CreateIstioGatewayDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
