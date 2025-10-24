// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGatewayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachGatewayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachGatewayDomainResponse
	GetStatusCode() *int32
	SetBody(v *AttachGatewayDomainResponseBody) *AttachGatewayDomainResponse
	GetBody() *AttachGatewayDomainResponseBody
}

type AttachGatewayDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachGatewayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachGatewayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachGatewayDomainResponse) GoString() string {
	return s.String()
}

func (s *AttachGatewayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachGatewayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachGatewayDomainResponse) GetBody() *AttachGatewayDomainResponseBody {
	return s.Body
}

func (s *AttachGatewayDomainResponse) SetHeaders(v map[string]*string) *AttachGatewayDomainResponse {
	s.Headers = v
	return s
}

func (s *AttachGatewayDomainResponse) SetStatusCode(v int32) *AttachGatewayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachGatewayDomainResponse) SetBody(v *AttachGatewayDomainResponseBody) *AttachGatewayDomainResponse {
	s.Body = v
	return s
}

func (s *AttachGatewayDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
