// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcuByRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEcuByRegionResponseBody
	GetCode() *int32
	SetEcuEntityList(v *ListEcuByRegionResponseBodyEcuEntityList) *ListEcuByRegionResponseBody
	GetEcuEntityList() *ListEcuByRegionResponseBodyEcuEntityList
	SetMessage(v string) *ListEcuByRegionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEcuByRegionResponseBody
	GetRequestId() *string
}

type ListEcuByRegionResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code          *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	EcuEntityList *ListEcuByRegionResponseBodyEcuEntityList `json:"EcuEntityList,omitempty" xml:"EcuEntityList,omitempty" type:"Struct"`
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
	// 00000000-0000-0000-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEcuByRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEcuByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEcuByRegionResponseBody) GetEcuEntityList() *ListEcuByRegionResponseBodyEcuEntityList {
	return s.EcuEntityList
}

func (s *ListEcuByRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEcuByRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEcuByRegionResponseBody) SetCode(v int32) *ListEcuByRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListEcuByRegionResponseBody) SetEcuEntityList(v *ListEcuByRegionResponseBodyEcuEntityList) *ListEcuByRegionResponseBody {
	s.EcuEntityList = v
	return s
}

func (s *ListEcuByRegionResponseBody) SetMessage(v string) *ListEcuByRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListEcuByRegionResponseBody) SetRequestId(v string) *ListEcuByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcuByRegionResponseBody) Validate() error {
	if s.EcuEntityList != nil {
		if err := s.EcuEntityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEcuByRegionResponseBodyEcuEntityList struct {
	EcuEntity []*ListEcuByRegionResponseBodyEcuEntityListEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" type:"Repeated"`
}

func (s ListEcuByRegionResponseBodyEcuEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListEcuByRegionResponseBodyEcuEntityList) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponseBodyEcuEntityList) GetEcuEntity() []*ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	return s.EcuEntity
}

func (s *ListEcuByRegionResponseBodyEcuEntityList) SetEcuEntity(v []*ListEcuByRegionResponseBodyEcuEntityListEcuEntity) *ListEcuByRegionResponseBodyEcuEntityList {
	s.EcuEntity = v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityList) Validate() error {
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

type ListEcuByRegionResponseBodyEcuEntityListEcuEntity struct {
	AvailableCpu  *int32  `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty"`
	AvailableMem  *int32  `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty"`
	Cpu           *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty"`
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty"`
	Mem           *int32  `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListEcuByRegionResponseBodyEcuEntityListEcuEntity) String() string {
	return dara.Prettify(s)
}

func (s ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetAvailableCpu() *int32 {
	return s.AvailableCpu
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetAvailableMem() *int32 {
	return s.AvailableMem
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetDockerEnv() *bool {
	return s.DockerEnv
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetEcuId() *string {
	return s.EcuId
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetHeartbeatTime() *int64 {
	return s.HeartbeatTime
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetIpAddr() *string {
	return s.IpAddr
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetMem() *int32 {
	return s.Mem
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetName() *string {
	return s.Name
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetOnline() *bool {
	return s.Online
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetAvailableCpu(v int32) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetAvailableMem(v int32) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetCpu(v int32) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Cpu = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetCreateTime(v int64) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetDockerEnv(v bool) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetEcuId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.EcuId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetHeartbeatTime(v int64) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetInstanceId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetIpAddr(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetMem(v int32) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Mem = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetName(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Name = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetOnline(v bool) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Online = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetRegionId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.RegionId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetUpdateTime(v int64) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetUserId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.UserId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetVpcId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.VpcId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetZoneId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) Validate() error {
	return dara.Validate(s)
}
