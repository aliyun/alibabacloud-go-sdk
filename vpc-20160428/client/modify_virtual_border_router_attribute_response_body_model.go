// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBorderRouterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVirtualBorderRouterAttributeResponseBody
	GetRequestId() *string
}

type ModifyVirtualBorderRouterAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 980960B0-2969-40BF-8542-EBB34FD358AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVirtualBorderRouterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBorderRouterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBorderRouterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVirtualBorderRouterAttributeResponseBody) SetRequestId(v string) *ModifyVirtualBorderRouterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
