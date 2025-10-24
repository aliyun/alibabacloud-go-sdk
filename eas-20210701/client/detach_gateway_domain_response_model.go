// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGatewayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachGatewayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachGatewayDomainResponse
	GetStatusCode() *int32
	SetBody(v *DetachGatewayDomainResponseBody) *DetachGatewayDomainResponse
	GetBody() *DetachGatewayDomainResponseBody
}

type DetachGatewayDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachGatewayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachGatewayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachGatewayDomainResponse) GoString() string {
	return s.String()
}

func (s *DetachGatewayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachGatewayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachGatewayDomainResponse) GetBody() *DetachGatewayDomainResponseBody {
	return s.Body
}

func (s *DetachGatewayDomainResponse) SetHeaders(v map[string]*string) *DetachGatewayDomainResponse {
	s.Headers = v
	return s
}

func (s *DetachGatewayDomainResponse) SetStatusCode(v int32) *DetachGatewayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachGatewayDomainResponse) SetBody(v *DetachGatewayDomainResponseBody) *DetachGatewayDomainResponse {
	s.Body = v
	return s
}

func (s *DetachGatewayDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
