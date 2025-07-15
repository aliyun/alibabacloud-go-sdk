// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePhysicalConnectionToVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody
	GetRequestId() *string
}

type AssociatePhysicalConnectionToVirtualBorderRouterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 980960B0-2969-40BF-8542-EBB34FD358AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) SetRequestId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
