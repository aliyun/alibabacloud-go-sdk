// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *CreateDesktopGroupResponseBody
	GetDesktopGroupId() *string
	SetDesktopGroupIds(v []*string) *CreateDesktopGroupResponseBody
	GetDesktopGroupIds() []*string
	SetOrderIds(v []*string) *CreateDesktopGroupResponseBody
	GetOrderIds() []*string
	SetRequestId(v string) *CreateDesktopGroupResponseBody
	GetRequestId() *string
}

type CreateDesktopGroupResponseBody struct {
	// The ID of the shared group.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The IDs of the shared groups.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The IDs of the orders.
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3EB7FCEE-D731-4948-85A3-4B2C341CA983
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponseBody) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *CreateDesktopGroupResponseBody) GetDesktopGroupIds() []*string {
	return s.DesktopGroupIds
}

func (s *CreateDesktopGroupResponseBody) GetOrderIds() []*string {
	return s.OrderIds
}

func (s *CreateDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDesktopGroupResponseBody) SetDesktopGroupId(v string) *CreateDesktopGroupResponseBody {
	s.DesktopGroupId = &v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetDesktopGroupIds(v []*string) *CreateDesktopGroupResponseBody {
	s.DesktopGroupIds = v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetOrderIds(v []*string) *CreateDesktopGroupResponseBody {
	s.OrderIds = v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetRequestId(v string) *CreateDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
