// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIstioGatewayDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIstioGatewayDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIstioGatewayDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIstioGatewayDomainsResponseBody) *DeleteIstioGatewayDomainsResponse
	GetBody() *DeleteIstioGatewayDomainsResponseBody
}

type DeleteIstioGatewayDomainsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIstioGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIstioGatewayDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIstioGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *DeleteIstioGatewayDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIstioGatewayDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIstioGatewayDomainsResponse) GetBody() *DeleteIstioGatewayDomainsResponseBody {
	return s.Body
}

func (s *DeleteIstioGatewayDomainsResponse) SetHeaders(v map[string]*string) *DeleteIstioGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *DeleteIstioGatewayDomainsResponse) SetStatusCode(v int32) *DeleteIstioGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIstioGatewayDomainsResponse) SetBody(v *DeleteIstioGatewayDomainsResponseBody) *DeleteIstioGatewayDomainsResponse {
	s.Body = v
	return s
}

func (s *DeleteIstioGatewayDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
