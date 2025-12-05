// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveDatabasesToNetworkDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveDatabasesToNetworkDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveDatabasesToNetworkDomainResponse
	GetStatusCode() *int32
	SetBody(v *MoveDatabasesToNetworkDomainResponseBody) *MoveDatabasesToNetworkDomainResponse
	GetBody() *MoveDatabasesToNetworkDomainResponseBody
}

type MoveDatabasesToNetworkDomainResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveDatabasesToNetworkDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveDatabasesToNetworkDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveDatabasesToNetworkDomainResponse) GoString() string {
	return s.String()
}

func (s *MoveDatabasesToNetworkDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveDatabasesToNetworkDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveDatabasesToNetworkDomainResponse) GetBody() *MoveDatabasesToNetworkDomainResponseBody {
	return s.Body
}

func (s *MoveDatabasesToNetworkDomainResponse) SetHeaders(v map[string]*string) *MoveDatabasesToNetworkDomainResponse {
	s.Headers = v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponse) SetStatusCode(v int32) *MoveDatabasesToNetworkDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponse) SetBody(v *MoveDatabasesToNetworkDomainResponseBody) *MoveDatabasesToNetworkDomainResponse {
	s.Body = v
	return s
}

func (s *MoveDatabasesToNetworkDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
