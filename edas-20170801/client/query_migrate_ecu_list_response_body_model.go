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
	Code          *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
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
