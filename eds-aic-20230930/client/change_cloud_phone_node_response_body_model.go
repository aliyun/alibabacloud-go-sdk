// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCloudPhoneNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfos) *ChangeCloudPhoneNodeResponseBody
	GetNodeInfos() []*ChangeCloudPhoneNodeResponseBodyNodeInfos
	SetOrderId(v string) *ChangeCloudPhoneNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ChangeCloudPhoneNodeResponseBody
	GetRequestId() *string
}

type ChangeCloudPhoneNodeResponseBody struct {
	NodeInfos []*ChangeCloudPhoneNodeResponseBodyNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	OrderId   *string                                      `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 4610632D-D661-5982-B3D7-5D3FD183F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeCloudPhoneNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponseBody) GetNodeInfos() []*ChangeCloudPhoneNodeResponseBodyNodeInfos {
	return s.NodeInfos
}

func (s *ChangeCloudPhoneNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ChangeCloudPhoneNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCloudPhoneNodeResponseBody) SetNodeInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfos) *ChangeCloudPhoneNodeResponseBody {
	s.NodeInfos = v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBody) SetOrderId(v string) *ChangeCloudPhoneNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBody) SetRequestId(v string) *ChangeCloudPhoneNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeCloudPhoneNodeResponseBodyNodeInfos struct {
	InstanceInfos []*ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// cpn-e5kxgjyt8s1mb****
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ShareDataVolume *int32  `json:"ShareDataVolume,omitempty" xml:"ShareDataVolume,omitempty"`
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfos) String() string {
	return dara.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfos) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) GetInstanceInfos() []*ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos {
	return s.InstanceInfos
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) GetNodeId() *string {
	return s.NodeId
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) GetShareDataVolume() *int32 {
	return s.ShareDataVolume
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetInstanceInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.InstanceInfos = v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetNodeId(v string) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.NodeId = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetShareDataVolume(v int32) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.ShareDataVolume = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) Validate() error {
	return dara.Validate(s)
}

type ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos struct {
	// example:
	//
	// cpn-jewjt8xryuitu****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneDataVolume *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) String() string {
	return dara.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) SetInstanceId(v string) *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos {
	s.InstanceId = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) SetPhoneDataVolume(v int32) *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos {
	s.PhoneDataVolume = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) Validate() error {
	return dara.Validate(s)
}
