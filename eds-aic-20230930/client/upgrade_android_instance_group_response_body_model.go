// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAndroidInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *UpgradeAndroidInstanceGroupResponseBody
	GetInstanceIds() *string
	SetOrderId(v string) *UpgradeAndroidInstanceGroupResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradeAndroidInstanceGroupResponseBody
	GetRequestId() *string
}

type UpgradeAndroidInstanceGroupResponseBody struct {
	// The IDs of instances created during the scale-out.
	//
	// example:
	//
	// [\\"acp-2iowaqjyw9nrd****\\", \\"acp-2iowaqjyw9nrd****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 223684716098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 310A783E-CC46-5452-A8A3-71AE5DB59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeAndroidInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAndroidInstanceGroupResponseBody) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *UpgradeAndroidInstanceGroupResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradeAndroidInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeAndroidInstanceGroupResponseBody) SetInstanceIds(v string) *UpgradeAndroidInstanceGroupResponseBody {
	s.InstanceIds = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponseBody) SetOrderId(v string) *UpgradeAndroidInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponseBody) SetRequestId(v string) *UpgradeAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
