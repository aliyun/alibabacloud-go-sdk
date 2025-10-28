// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScaleOutEcuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListScaleOutEcuResponseBody
	GetCode() *int32
	SetEcuInfoList(v *ListScaleOutEcuResponseBodyEcuInfoList) *ListScaleOutEcuResponseBody
	GetEcuInfoList() *ListScaleOutEcuResponseBodyEcuInfoList
	SetMessage(v string) *ListScaleOutEcuResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScaleOutEcuResponseBody
	GetRequestId() *string
}

type ListScaleOutEcuResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ECUs.
	EcuInfoList *ListScaleOutEcuResponseBodyEcuInfoList `json:"EcuInfoList,omitempty" xml:"EcuInfoList,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AF860D6C-ACE3-4429-9D54-3BD15A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListScaleOutEcuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScaleOutEcuResponseBody) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListScaleOutEcuResponseBody) GetEcuInfoList() *ListScaleOutEcuResponseBodyEcuInfoList {
	return s.EcuInfoList
}

func (s *ListScaleOutEcuResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScaleOutEcuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScaleOutEcuResponseBody) SetCode(v int32) *ListScaleOutEcuResponseBody {
	s.Code = &v
	return s
}

func (s *ListScaleOutEcuResponseBody) SetEcuInfoList(v *ListScaleOutEcuResponseBodyEcuInfoList) *ListScaleOutEcuResponseBody {
	s.EcuInfoList = v
	return s
}

func (s *ListScaleOutEcuResponseBody) SetMessage(v string) *ListScaleOutEcuResponseBody {
	s.Message = &v
	return s
}

func (s *ListScaleOutEcuResponseBody) SetRequestId(v string) *ListScaleOutEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScaleOutEcuResponseBody) Validate() error {
	if s.EcuInfoList != nil {
		if err := s.EcuInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScaleOutEcuResponseBodyEcuInfoList struct {
	EcuInfo []*ListScaleOutEcuResponseBodyEcuInfoListEcuInfo `json:"EcuInfo,omitempty" xml:"EcuInfo,omitempty" type:"Repeated"`
}

func (s ListScaleOutEcuResponseBodyEcuInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListScaleOutEcuResponseBodyEcuInfoList) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponseBodyEcuInfoList) GetEcuInfo() []*ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	return s.EcuInfo
}

func (s *ListScaleOutEcuResponseBodyEcuInfoList) SetEcuInfo(v []*ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) *ListScaleOutEcuResponseBodyEcuInfoList {
	s.EcuInfo = v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoList) Validate() error {
	if s.EcuInfo != nil {
		for _, item := range s.EcuInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScaleOutEcuResponseBodyEcuInfoListEcuInfo struct {
	// The number of available CPU cores for the ECU.
	//
	// example:
	//
	// 2
	AvailableCpu *int32 `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty"`
	// The size of available memory for the ECU. Unit: MB.
	//
	// example:
	//
	// 111
	AvailableMem *int32 `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty"`
	// The time when the ECU was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573281040819
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether Docker is installed. Valid values:
	//
	// 	- true: Docker is installed.
	//
	// 	- false: Docker is not installed.
	//
	// example:
	//
	// false
	DockerEnv *bool `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty"`
	// The unique ID of the ECU. To query the ID, you can run the `dmidecode` command on the ECS instance that corresponds to the ECU.
	//
	// example:
	//
	// 0de2ebdb-9490-4fc4-be41***************
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The time when the last heartbeat detection was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573281040819
	HeartbeatTime *int64 `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty"`
	// The ID of the ECU.
	//
	// example:
	//
	// i-2zej4i2jdf*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The private IP address of the ECU.
	//
	// example:
	//
	// 192.168.XX.XX
	IpAddr *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the ECU is online. If the ECU is online, its corresponding ECS instance is managed in EDAS. Valid values:
	//
	// 	- true: The ECU is online.
	//
	// 	- false: The ECU is offline.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The ID of the region where the ECU is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the ECU was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573281040827
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the Alibaba Cloud account to which the ECU belongs.
	//
	// example:
	//
	// 1172****6608****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the virtual private cloud (VPC) where the ECU is located.
	//
	// example:
	//
	// vpc-2zef6ob8**********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone where the ECU resides.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) String() string {
	return dara.Prettify(s)
}

func (s ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetAvailableCpu() *int32 {
	return s.AvailableCpu
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetAvailableMem() *int32 {
	return s.AvailableMem
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetDockerEnv() *bool {
	return s.DockerEnv
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetEcuId() *string {
	return s.EcuId
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetHeartbeatTime() *int64 {
	return s.HeartbeatTime
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetIpAddr() *string {
	return s.IpAddr
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetName() *string {
	return s.Name
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetOnline() *bool {
	return s.Online
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetUserId() *string {
	return s.UserId
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetAvailableCpu(v int32) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.AvailableCpu = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetAvailableMem(v int32) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.AvailableMem = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetCreateTime(v int64) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.CreateTime = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetDockerEnv(v bool) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.DockerEnv = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetEcuId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.EcuId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetHeartbeatTime(v int64) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.HeartbeatTime = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetInstanceId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.InstanceId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetIpAddr(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.IpAddr = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetName(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.Name = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetOnline(v bool) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.Online = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetRegionId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.RegionId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetUpdateTime(v int64) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.UpdateTime = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetUserId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.UserId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetVpcId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.VpcId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetZoneId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.ZoneId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) Validate() error {
	return dara.Validate(s)
}
