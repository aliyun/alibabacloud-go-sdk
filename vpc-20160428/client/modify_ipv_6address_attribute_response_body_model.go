// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6AddressAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpv6AddressAttributeResponseBody
	GetRequestId() *string
}

type ModifyIpv6AddressAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D560AF68-4CE8-4A5C-B3FE-469F558094D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpv6AddressAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6AddressAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpv6AddressAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpv6AddressAttributeResponseBody) SetRequestId(v string) *ModifyIpv6AddressAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpv6AddressAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
