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
	SetRequestId(v string) *ChangeCloudPhoneNodeResponseBody
	GetRequestId() *string
}

type ChangeCloudPhoneNodeResponseBody struct {
	NodeInfos []*ChangeCloudPhoneNodeResponseBodyNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
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

func (s *ChangeCloudPhoneNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCloudPhoneNodeResponseBody) SetNodeInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfos) *ChangeCloudPhoneNodeResponseBody {
	s.NodeInfos = v
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
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetInstanceInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.InstanceInfos = v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetNodeId(v string) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.NodeId = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) Validate() error {
	return dara.Validate(s)
}

type ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos struct {
	// example:
	//
	// cpn-jewjt8xryuitu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) SetInstanceId(v string) *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos {
	s.InstanceId = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) Validate() error {
	return dara.Validate(s)
}
