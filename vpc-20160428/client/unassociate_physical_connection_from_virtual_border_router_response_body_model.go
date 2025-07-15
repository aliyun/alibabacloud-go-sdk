// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody
	GetRequestId() *string
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 980960B0-2969-40BF-8542-EBB34FD358AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) SetRequestId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
