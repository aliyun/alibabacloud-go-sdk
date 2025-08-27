// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDBInstanceSrvNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateDBInstanceSrvNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateDBInstanceSrvNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateDBInstanceSrvNetworkAddressResponseBody) *AllocateDBInstanceSrvNetworkAddressResponse
	GetBody() *AllocateDBInstanceSrvNetworkAddressResponseBody
}

type AllocateDBInstanceSrvNetworkAddressResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateDBInstanceSrvNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateDBInstanceSrvNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateDBInstanceSrvNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) GetBody() *AllocateDBInstanceSrvNetworkAddressResponseBody {
	return s.Body
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocateDBInstanceSrvNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) SetStatusCode(v int32) *AllocateDBInstanceSrvNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) SetBody(v *AllocateDBInstanceSrvNetworkAddressResponseBody) *AllocateDBInstanceSrvNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressResponse) Validate() error {
	return dara.Validate(s)
}
