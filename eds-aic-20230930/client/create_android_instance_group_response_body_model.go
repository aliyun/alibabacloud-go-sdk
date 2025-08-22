// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndroidInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceGroupIds(v []*string) *CreateAndroidInstanceGroupResponseBody
	GetInstanceGroupIds() []*string
	SetInstanceGroupInfos(v []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) *CreateAndroidInstanceGroupResponseBody
	GetInstanceGroupInfos() []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos
	SetNetworkPackageOrderModel(v *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) *CreateAndroidInstanceGroupResponseBody
	GetNetworkPackageOrderModel() *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel
	SetOrderId(v string) *CreateAndroidInstanceGroupResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateAndroidInstanceGroupResponseBody
	GetRequestId() *string
}

type CreateAndroidInstanceGroupResponseBody struct {
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The instance groups.
	InstanceGroupInfos       []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos     `json:"InstanceGroupInfos,omitempty" xml:"InstanceGroupInfos,omitempty" type:"Repeated"`
	NetworkPackageOrderModel *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel `json:"NetworkPackageOrderModel,omitempty" xml:"NetworkPackageOrderModel,omitempty" type:"Struct"`
	// The ID of the order.
	//
	// example:
	//
	// 22365781890****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAndroidInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponseBody) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *CreateAndroidInstanceGroupResponseBody) GetInstanceGroupInfos() []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos {
	return s.InstanceGroupInfos
}

func (s *CreateAndroidInstanceGroupResponseBody) GetNetworkPackageOrderModel() *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel {
	return s.NetworkPackageOrderModel
}

func (s *CreateAndroidInstanceGroupResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAndroidInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndroidInstanceGroupResponseBody) SetInstanceGroupIds(v []*string) *CreateAndroidInstanceGroupResponseBody {
	s.InstanceGroupIds = v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetInstanceGroupInfos(v []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) *CreateAndroidInstanceGroupResponseBody {
	s.InstanceGroupInfos = v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetNetworkPackageOrderModel(v *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) *CreateAndroidInstanceGroupResponseBody {
	s.NetworkPackageOrderModel = v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetOrderId(v string) *CreateAndroidInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetRequestId(v string) *CreateAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos struct {
	// The ID of the instance group.
	//
	// example:
	//
	// ag-cuv4scs4obxch****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The IDs of the instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) SetInstanceGroupId(v string) *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos {
	s.InstanceGroupId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) SetInstanceIds(v []*string) *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos {
	s.InstanceIds = v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) Validate() error {
	return dara.Validate(s)
}

type CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel struct {
	BandwidthPackageId      *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageOrderId *string `json:"BandwidthPackageOrderId,omitempty" xml:"BandwidthPackageOrderId,omitempty"`
}

func (s CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) GetBandwidthPackageOrderId() *string {
	return s.BandwidthPackageOrderId
}

func (s *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) SetBandwidthPackageId(v string) *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) SetBandwidthPackageOrderId(v string) *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel {
	s.BandwidthPackageOrderId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBodyNetworkPackageOrderModel) Validate() error {
	return dara.Validate(s)
}
