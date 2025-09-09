// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkDomainResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkDomainResponseBody) *CreateNetworkDomainResponse
	GetBody() *CreateNetworkDomainResponseBody
}

type CreateNetworkDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkDomainResponse) GetBody() *CreateNetworkDomainResponseBody {
	return s.Body
}

func (s *CreateNetworkDomainResponse) SetHeaders(v map[string]*string) *CreateNetworkDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkDomainResponse) SetStatusCode(v int32) *CreateNetworkDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkDomainResponse) SetBody(v *CreateNetworkDomainResponseBody) *CreateNetworkDomainResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkDomainResponse) Validate() error {
	return dara.Validate(s)
}
