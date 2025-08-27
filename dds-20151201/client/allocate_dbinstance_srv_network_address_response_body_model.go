// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDBInstanceSrvNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateDBInstanceSrvNetworkAddressResponseBody
	GetRequestId() *string
}

type AllocateDBInstanceSrvNetworkAddressResponseBody struct {
	// example:
	//
	// 6B82A9EF-9961-5A31-A19F-009B7098ACCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateDBInstanceSrvNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateDBInstanceSrvNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateDBInstanceSrvNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateDBInstanceSrvNetworkAddressResponseBody) SetRequestId(v string) *AllocateDBInstanceSrvNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
