// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySrvNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySrvNetworkAddressResponseBody
	GetRequestId() *string
}

type ModifySrvNetworkAddressResponseBody struct {
	// example:
	//
	// 45D2B592-DEBA-4347-BBF3-xxxxxC97DBBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySrvNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySrvNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySrvNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySrvNetworkAddressResponseBody) SetRequestId(v string) *ModifySrvNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySrvNetworkAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
