// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPhoneNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackageOrderModel(v *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) *CreateCloudPhoneNodeResponseBody
	GetNetworkPackageOrderModel() *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel
	SetNodeInfos(v []*CreateCloudPhoneNodeResponseBodyNodeInfos) *CreateCloudPhoneNodeResponseBody
	GetNodeInfos() []*CreateCloudPhoneNodeResponseBodyNodeInfos
	SetOrderId(v string) *CreateCloudPhoneNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateCloudPhoneNodeResponseBody
	GetRequestId() *string
}

type CreateCloudPhoneNodeResponseBody struct {
	NetworkPackageOrderModel *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel `json:"NetworkPackageOrderModel,omitempty" xml:"NetworkPackageOrderModel,omitempty" type:"Struct"`
	// The cloud phone matrixes.
	NodeInfos []*CreateCloudPhoneNodeResponseBodyNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	// The order ID.
	//
	// example:
	//
	// 223684716098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudPhoneNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponseBody) GetNetworkPackageOrderModel() *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel {
	return s.NetworkPackageOrderModel
}

func (s *CreateCloudPhoneNodeResponseBody) GetNodeInfos() []*CreateCloudPhoneNodeResponseBodyNodeInfos {
	return s.NodeInfos
}

func (s *CreateCloudPhoneNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateCloudPhoneNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudPhoneNodeResponseBody) SetNetworkPackageOrderModel(v *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) *CreateCloudPhoneNodeResponseBody {
	s.NetworkPackageOrderModel = v
	return s
}

func (s *CreateCloudPhoneNodeResponseBody) SetNodeInfos(v []*CreateCloudPhoneNodeResponseBodyNodeInfos) *CreateCloudPhoneNodeResponseBody {
	s.NodeInfos = v
	return s
}

func (s *CreateCloudPhoneNodeResponseBody) SetOrderId(v string) *CreateCloudPhoneNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCloudPhoneNodeResponseBody) SetRequestId(v string) *CreateCloudPhoneNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudPhoneNodeResponseBody) Validate() error {
	if s.NetworkPackageOrderModel != nil {
		if err := s.NetworkPackageOrderModel.Validate(); err != nil {
			return err
		}
	}
	if s.NodeInfos != nil {
		for _, item := range s.NodeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel struct {
	BandwidthPackageId      *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageOrderId *string `json:"BandwidthPackageOrderId,omitempty" xml:"BandwidthPackageOrderId,omitempty"`
}

func (s CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) GetBandwidthPackageOrderId() *string {
	return s.BandwidthPackageOrderId
}

func (s *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) SetBandwidthPackageId(v string) *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) SetBandwidthPackageOrderId(v string) *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel {
	s.BandwidthPackageOrderId = &v
	return s
}

func (s *CreateCloudPhoneNodeResponseBodyNetworkPackageOrderModel) Validate() error {
	return dara.Validate(s)
}

type CreateCloudPhoneNodeResponseBodyNodeInfos struct {
	// The IDs of the cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the cloud phone matrix.
	//
	// example:
	//
	// cpn-e5kxgjyt8s1mb****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateCloudPhoneNodeResponseBodyNodeInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeResponseBodyNodeInfos) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) SetInstanceIds(v []*string) *CreateCloudPhoneNodeResponseBodyNodeInfos {
	s.InstanceIds = v
	return s
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) SetNodeId(v string) *CreateCloudPhoneNodeResponseBodyNodeInfos {
	s.NodeId = &v
	return s
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) Validate() error {
	return dara.Validate(s)
}
