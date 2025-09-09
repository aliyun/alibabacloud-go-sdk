// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkDomainResponseBody) *DeleteNetworkDomainResponse
	GetBody() *DeleteNetworkDomainResponseBody
}

type DeleteNetworkDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkDomainResponse) GetBody() *DeleteNetworkDomainResponseBody {
	return s.Body
}

func (s *DeleteNetworkDomainResponse) SetHeaders(v map[string]*string) *DeleteNetworkDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkDomainResponse) SetStatusCode(v int32) *DeleteNetworkDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkDomainResponse) SetBody(v *DeleteNetworkDomainResponseBody) *DeleteNetworkDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkDomainResponse) Validate() error {
	return dara.Validate(s)
}
