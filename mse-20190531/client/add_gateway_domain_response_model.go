// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayDomainResponseBody) *AddGatewayDomainResponse
	GetBody() *AddGatewayDomainResponseBody
}

type AddGatewayDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayDomainResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayDomainResponse) GetBody() *AddGatewayDomainResponseBody {
	return s.Body
}

func (s *AddGatewayDomainResponse) SetHeaders(v map[string]*string) *AddGatewayDomainResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayDomainResponse) SetStatusCode(v int32) *AddGatewayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayDomainResponse) SetBody(v *AddGatewayDomainResponseBody) *AddGatewayDomainResponse {
	s.Body = v
	return s
}

func (s *AddGatewayDomainResponse) Validate() error {
	return dara.Validate(s)
}
