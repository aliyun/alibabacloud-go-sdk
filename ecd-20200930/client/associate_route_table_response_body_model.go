// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateRouteTableResponseBody
	GetRequestId() *string
}

type AssociateRouteTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateRouteTableResponseBody) SetRequestId(v string) *AssociateRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
