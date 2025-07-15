// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateRouteTableResponseBody
	GetRequestId() *string
}

type UnassociateRouteTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 62172DD5-6BAC-45DF-8D44-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateRouteTableResponseBody) SetRequestId(v string) *UnassociateRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
