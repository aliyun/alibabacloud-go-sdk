// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpAddressesFromEgressProxyIpPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveIpAddressesFromEgressProxyIpPoolResponseBody
	GetRequestId() *string
}

type RemoveIpAddressesFromEgressProxyIpPoolResponseBody struct {
	// example:
	//
	// 017E7009-8930-5511-BAE6-BFFDFA21A48A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIpAddressesFromEgressProxyIpPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpAddressesFromEgressProxyIpPoolResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponseBody) SetRequestId(v string) *RemoveIpAddressesFromEgressProxyIpPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
