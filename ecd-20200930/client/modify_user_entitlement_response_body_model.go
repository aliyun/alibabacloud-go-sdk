// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserEntitlementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserEntitlementResponseBody
	GetRequestId() *string
}

type ModifyUserEntitlementResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserEntitlementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserEntitlementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserEntitlementResponseBody) SetRequestId(v string) *ModifyUserEntitlementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserEntitlementResponseBody) Validate() error {
	return dara.Validate(s)
}
