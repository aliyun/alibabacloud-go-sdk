// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHostsToNetworkDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveHostsToNetworkDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveHostsToNetworkDomainResponse
	GetStatusCode() *int32
	SetBody(v *MoveHostsToNetworkDomainResponseBody) *MoveHostsToNetworkDomainResponse
	GetBody() *MoveHostsToNetworkDomainResponseBody
}

type MoveHostsToNetworkDomainResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveHostsToNetworkDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveHostsToNetworkDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveHostsToNetworkDomainResponse) GoString() string {
	return s.String()
}

func (s *MoveHostsToNetworkDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveHostsToNetworkDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveHostsToNetworkDomainResponse) GetBody() *MoveHostsToNetworkDomainResponseBody {
	return s.Body
}

func (s *MoveHostsToNetworkDomainResponse) SetHeaders(v map[string]*string) *MoveHostsToNetworkDomainResponse {
	s.Headers = v
	return s
}

func (s *MoveHostsToNetworkDomainResponse) SetStatusCode(v int32) *MoveHostsToNetworkDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveHostsToNetworkDomainResponse) SetBody(v *MoveHostsToNetworkDomainResponseBody) *MoveHostsToNetworkDomainResponse {
	s.Body = v
	return s
}

func (s *MoveHostsToNetworkDomainResponse) Validate() error {
	return dara.Validate(s)
}
