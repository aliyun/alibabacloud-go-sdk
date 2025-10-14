// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfVersions(v []*DescribeEpnInstanceAttributeResponseBodyConfVersions) *DescribeEpnInstanceAttributeResponseBody
	GetConfVersions() []*DescribeEpnInstanceAttributeResponseBodyConfVersions
	SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeResponseBody
	GetEPNInstanceId() *string
	SetEPNInstanceName(v string) *DescribeEpnInstanceAttributeResponseBody
	GetEPNInstanceName() *string
	SetInstances(v []*DescribeEpnInstanceAttributeResponseBodyInstances) *DescribeEpnInstanceAttributeResponseBody
	GetInstances() []*DescribeEpnInstanceAttributeResponseBodyInstances
	SetNetworkingModel(v string) *DescribeEpnInstanceAttributeResponseBody
	GetNetworkingModel() *string
	SetRequestId(v string) *DescribeEpnInstanceAttributeResponseBody
	GetRequestId() *string
	SetVSwitches(v []*DescribeEpnInstanceAttributeResponseBodyVSwitches) *DescribeEpnInstanceAttributeResponseBody
	GetVSwitches() []*DescribeEpnInstanceAttributeResponseBodyVSwitches
}

type DescribeEpnInstanceAttributeResponseBody struct {
	// The information about the EPN configurations.
	ConfVersions []*DescribeEpnInstanceAttributeResponseBodyConfVersions `json:"ConfVersions,omitempty" xml:"ConfVersions,omitempty" type:"Repeated"`
	// The ID of the EPN instance.
	//
	// example:
	//
	// epn-xxxx
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The name of the EPN instance.
	//
	// example:
	//
	// epn-test
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	// The information about the instance.
	Instances []*DescribeEpnInstanceAttributeResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The networking mode. Valid values:
	//
	// 	- SpeedUp: intelligent acceleration network (Internet)
	//
	// 	- Connection: internal network
	//
	// 	- SpeedUpAndConnection: intelligent acceleration network and internal network
	//
	// example:
	//
	// SpeedUp
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the vSwitch.
	VSwitches []*DescribeEpnInstanceAttributeResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s DescribeEpnInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetConfVersions() []*DescribeEpnInstanceAttributeResponseBodyConfVersions {
	return s.ConfVersions
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetEPNInstanceName() *string {
	return s.EPNInstanceName
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetInstances() []*DescribeEpnInstanceAttributeResponseBodyInstances {
	return s.Instances
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetNetworkingModel() *string {
	return s.NetworkingModel
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEpnInstanceAttributeResponseBody) GetVSwitches() []*DescribeEpnInstanceAttributeResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetConfVersions(v []*DescribeEpnInstanceAttributeResponseBodyConfVersions) *DescribeEpnInstanceAttributeResponseBody {
	s.ConfVersions = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetEPNInstanceName(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetInstances(v []*DescribeEpnInstanceAttributeResponseBodyInstances) *DescribeEpnInstanceAttributeResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetNetworkingModel(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetRequestId(v string) *DescribeEpnInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) SetVSwitches(v []*DescribeEpnInstanceAttributeResponseBodyVSwitches) *DescribeEpnInstanceAttributeResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBody) Validate() error {
	if s.ConfVersions != nil {
		for _, item := range s.ConfVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEpnInstanceAttributeResponseBodyConfVersions struct {
	// The version number.
	//
	// example:
	//
	// 2017-10-11
	ConfVersion *string `json:"ConfVersion,omitempty" xml:"ConfVersion,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponseBodyConfVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBodyConfVersions) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) GetConfVersion() *string {
	return s.ConfVersion
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) SetConfVersion(v string) *DescribeEpnInstanceAttributeResponseBodyConfVersions {
	s.ConfVersion = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseBodyConfVersions {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyConfVersions) Validate() error {
	return dara.Validate(s)
}

type DescribeEpnInstanceAttributeResponseBodyInstances struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-telecom-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// epn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// epn-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ISP. Valid values:
	//
	// 	- cmcc: China Mobile
	//
	// 	- unicom: China Unicom
	//
	// 	- telecom: China Telecom
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 192.168.1.12
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 20.3.XX.XX
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- Running
	//
	// 	- Stopped
	//
	// 	- Expired
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetInstanceId(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetInstanceName(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetIsp(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.Isp = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetPrivateIpAddress(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetPublicIpAddress(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) SetStatus(v string) *DescribeEpnInstanceAttributeResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeEpnInstanceAttributeResponseBodyVSwitches struct {
	// The CIDR block.
	//
	// example:
	//
	// 10.0.0.1/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-telecom-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vs-xxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vs-test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeEpnInstanceAttributeResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetCidrBlock(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.CidrBlock = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetVSwitchId(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) SetVSwitchName(v string) *DescribeEpnInstanceAttributeResponseBodyVSwitches {
	s.VSwitchName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseBodyVSwitches) Validate() error {
	return dara.Validate(s)
}
