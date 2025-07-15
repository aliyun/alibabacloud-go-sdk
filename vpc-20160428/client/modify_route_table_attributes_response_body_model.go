// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteTableAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRouteTableAttributesResponseBody
	GetRequestId() *string
}

type ModifyRouteTableAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 62172DD5-6BAC-45DF-8D44
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRouteTableAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteTableAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRouteTableAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRouteTableAttributesResponseBody) SetRequestId(v string) *ModifyRouteTableAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRouteTableAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
