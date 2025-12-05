// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkDomainResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkDomainResponseBody) *ModifyNetworkDomainResponse
	GetBody() *ModifyNetworkDomainResponseBody
}

type ModifyNetworkDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkDomainResponse) GetBody() *ModifyNetworkDomainResponseBody {
	return s.Body
}

func (s *ModifyNetworkDomainResponse) SetHeaders(v map[string]*string) *ModifyNetworkDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkDomainResponse) SetStatusCode(v int32) *ModifyNetworkDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkDomainResponse) SetBody(v *ModifyNetworkDomainResponseBody) *ModifyNetworkDomainResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
