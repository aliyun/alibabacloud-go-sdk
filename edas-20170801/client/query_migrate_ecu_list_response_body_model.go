// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMigrateEcuListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryMigrateEcuListResponseBody
	GetCode() *int32
	SetEcuEntityList(v *QueryMigrateEcuListResponseBodyEcuEntityList) *QueryMigrateEcuListResponseBody
	GetEcuEntityList() *QueryMigrateEcuListResponseBodyEcuEntityList
	SetMessage(v string) *QueryMigrateEcuListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMigrateEcuListResponseBody
	GetRequestId() *string
}

type QueryMigrateEcuListResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about ECUs that can be migrated.
	EcuEntityList *QueryMigrateEcuListResponseBodyEcuEntityList `json:"EcuEntityList,omitempty" xml:"EcuEntityList,omitempty" type:"Struct"`
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
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMigrateEcuListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateEcuListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryMigrateEcuListResponseBody) GetEcuEntityList() *QueryMigrateEcuListResponseBodyEcuEntityList {
	return s.EcuEntityList
}

func (s *QueryMigrateEcuListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMigrateEcuListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMigrateEcuListResponseBody) SetCode(v int32) *QueryMigrateEcuListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMigrateEcuListResponseBody) SetEcuEntityList(v *QueryMigrateEcuListResponseBodyEcuEntityList) *QueryMigrateEcuListResponseBody {
	s.EcuEntityList = v
	return s
}

func (s *QueryMigrateEcuListResponseBody) SetMessage(v string) *QueryMigrateEcuListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMigrateEcuListResponseBody) SetRequestId(v string) *QueryMigrateEcuListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBody) Validate() error {
	if s.EcuEntityList != nil {
		if err := s.EcuEntityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMigrateEcuListResponseBodyEcuEntityList struct {
	EcuEntity []*QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" type:"Repeated"`
}

func (s QueryMigrateEcuListResponseBodyEcuEntityList) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateEcuListResponseBodyEcuEntityList) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityList) GetEcuEntity() []*QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	return s.EcuEntity
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityList) SetEcuEntity(v []*QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) *QueryMigrateEcuListResponseBodyEcuEntityList {
	s.EcuEntity = v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityList) Validate() error {
	if s.EcuEntity != nil {
		for _, item := range s.EcuEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity struct {
	// The number of available CPUs. Unit: cores.
	//
	// example:
	//
	// 2
	AvailableCpu *int32 `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty"`
	// The size of available memory. Unit: MB.
	//
	// example:
	//
	// 2048
	AvailableMem *int32 `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty"`
	// The CPU quota set by the system. Unit: cores. The value 0 indicates that no quota is set by the system.
	//
	// example:
	//
	// 0
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the ECU was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573281041101
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
	// 70ed3f59-b476-49aa-****-************
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The time when the last heartbeat detection was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573281041101
	HeartbeatTime *int64 `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty"`
	// The ID of the ECU.
	//
	// example:
	//
	// i-2zej4i2jd***********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The private IP address of the ECU.
	//
	// example:
	//
	// 192.168.0.150
	IpAddr *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty"`
	// The total size of memory.
	//
	// example:
	//
	// 0
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// product_test003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the ECU is online. Valid values:
	//
	// 	- true: The ECU is online.
	//
	// 	- false: The ECU is offline.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The ID of the region where the ECU resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the ECU was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573281041109
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the Alibaba Cloud account to which the ECU belongs.
	//
	// example:
	//
	// edas_com***_****@******-*****.***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-2zef6ob8m************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone where the ECU resides.
	//
	// example:
	//
	// cn-bei****-*
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetAvailableCpu() *int32 {
	return s.AvailableCpu
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetAvailableMem() *int32 {
	return s.AvailableMem
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetDockerEnv() *bool {
	return s.DockerEnv
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetEcuId() *string {
	return s.EcuId
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetHeartbeatTime() *int64 {
	return s.HeartbeatTime
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetIpAddr() *string {
	return s.IpAddr
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetMem() *int32 {
	return s.Mem
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetName() *string {
	return s.Name
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetOnline() *bool {
	return s.Online
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetUserId() *string {
	return s.UserId
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GetZoneId() *string {
	return s.ZoneId
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetAvailableCpu(v int32) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetAvailableMem(v int32) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetCpu(v int32) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Cpu = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetCreateTime(v int64) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetDockerEnv(v bool) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetEcuId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.EcuId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetHeartbeatTime(v int64) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetInstanceId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetIpAddr(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetMem(v int32) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Mem = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetName(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Name = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetOnline(v bool) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Online = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetRegionId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.RegionId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetUpdateTime(v int64) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetUserId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.UserId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetVpcId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.VpcId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetZoneId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) Validate() error {
	return dara.Validate(s)
}
