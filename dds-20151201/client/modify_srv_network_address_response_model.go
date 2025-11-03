// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySrvNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySrvNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySrvNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *ModifySrvNetworkAddressResponseBody) *ModifySrvNetworkAddressResponse
	GetBody() *ModifySrvNetworkAddressResponseBody
}

type ModifySrvNetworkAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySrvNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySrvNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySrvNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifySrvNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySrvNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySrvNetworkAddressResponse) GetBody() *ModifySrvNetworkAddressResponseBody {
	return s.Body
}

func (s *ModifySrvNetworkAddressResponse) SetHeaders(v map[string]*string) *ModifySrvNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifySrvNetworkAddressResponse) SetStatusCode(v int32) *ModifySrvNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySrvNetworkAddressResponse) SetBody(v *ModifySrvNetworkAddressResponseBody) *ModifySrvNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *ModifySrvNetworkAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
