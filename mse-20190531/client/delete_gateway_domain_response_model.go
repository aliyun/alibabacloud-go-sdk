// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayDomainResponseBody) *DeleteGatewayDomainResponse
	GetBody() *DeleteGatewayDomainResponseBody
}

type DeleteGatewayDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayDomainResponse) GetBody() *DeleteGatewayDomainResponseBody {
	return s.Body
}

func (s *DeleteGatewayDomainResponse) SetHeaders(v map[string]*string) *DeleteGatewayDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayDomainResponse) SetStatusCode(v int32) *DeleteGatewayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayDomainResponse) SetBody(v *DeleteGatewayDomainResponseBody) *DeleteGatewayDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayDomainResponse) Validate() error {
	return dara.Validate(s)
}
