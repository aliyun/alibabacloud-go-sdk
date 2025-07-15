// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEntitlementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEntitlementResponseBody
	GetRequestId() *string
}

type ModifyEntitlementResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEntitlementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEntitlementResponseBody) SetRequestId(v string) *ModifyEntitlementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEntitlementResponseBody) Validate() error {
	return dara.Validate(s)
}
