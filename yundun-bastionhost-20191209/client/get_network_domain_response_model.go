// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkDomainResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkDomainResponseBody) *GetNetworkDomainResponse
	GetBody() *GetNetworkDomainResponseBody
}

type GetNetworkDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkDomainResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkDomainResponse) GetBody() *GetNetworkDomainResponseBody {
	return s.Body
}

func (s *GetNetworkDomainResponse) SetHeaders(v map[string]*string) *GetNetworkDomainResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkDomainResponse) SetStatusCode(v int32) *GetNetworkDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkDomainResponse) SetBody(v *GetNetworkDomainResponseBody) *GetNetworkDomainResponse {
	s.Body = v
	return s
}

func (s *GetNetworkDomainResponse) Validate() error {
	return dara.Validate(s)
}
