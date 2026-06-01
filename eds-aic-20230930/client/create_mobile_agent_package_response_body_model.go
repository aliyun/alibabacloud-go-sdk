// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMobileAgentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMobileAgentPackageIds(v []*string) *CreateMobileAgentPackageResponseBody
	GetMobileAgentPackageIds() []*string
	SetOrderId(v string) *CreateMobileAgentPackageResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateMobileAgentPackageResponseBody
	GetRequestId() *string
}

type CreateMobileAgentPackageResponseBody struct {
	MobileAgentPackageIds []*string `json:"MobileAgentPackageIds,omitempty" xml:"MobileAgentPackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMobileAgentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentPackageResponseBody) GetMobileAgentPackageIds() []*string {
	return s.MobileAgentPackageIds
}

func (s *CreateMobileAgentPackageResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateMobileAgentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMobileAgentPackageResponseBody) SetMobileAgentPackageIds(v []*string) *CreateMobileAgentPackageResponseBody {
	s.MobileAgentPackageIds = v
	return s
}

func (s *CreateMobileAgentPackageResponseBody) SetOrderId(v string) *CreateMobileAgentPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateMobileAgentPackageResponseBody) SetRequestId(v string) *CreateMobileAgentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMobileAgentPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
