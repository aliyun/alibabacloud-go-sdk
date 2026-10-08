// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationEcuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListApplicationEcuResponseBody
	GetCode() *int32
	SetEcuInfoList(v *ListApplicationEcuResponseBodyEcuInfoList) *ListApplicationEcuResponseBody
	GetEcuInfoList() *ListApplicationEcuResponseBodyEcuInfoList
	SetMessage(v string) *ListApplicationEcuResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationEcuResponseBody
	GetRequestId() *string
}

type ListApplicationEcuResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code        *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	EcuInfoList *ListApplicationEcuResponseBodyEcuInfoList `json:"EcuInfoList,omitempty" xml:"EcuInfoList,omitempty" type:"Struct"`
	// The message that is returned.
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

func (s ListApplicationEcuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationEcuResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListApplicationEcuResponseBody) GetEcuInfoList() *ListApplicationEcuResponseBodyEcuInfoList {
	return s.EcuInfoList
}

func (s *ListApplicationEcuResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationEcuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationEcuResponseBody) SetCode(v int32) *ListApplicationEcuResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationEcuResponseBody) SetEcuInfoList(v *ListApplicationEcuResponseBodyEcuInfoList) *ListApplicationEcuResponseBody {
	s.EcuInfoList = v
	return s
}

func (s *ListApplicationEcuResponseBody) SetMessage(v string) *ListApplicationEcuResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationEcuResponseBody) SetRequestId(v string) *ListApplicationEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationEcuResponseBody) Validate() error {
	if s.EcuInfoList != nil {
		if err := s.EcuInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationEcuResponseBodyEcuInfoList struct {
	EcuEntity []*ListApplicationEcuResponseBodyEcuInfoListEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" type:"Repeated"`
}

func (s ListApplicationEcuResponseBodyEcuInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationEcuResponseBodyEcuInfoList) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponseBodyEcuInfoList) GetEcuEntity() []*ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	return s.EcuEntity
}

func (s *ListApplicationEcuResponseBodyEcuInfoList) SetEcuEntity(v []*ListApplicationEcuResponseBodyEcuInfoListEcuEntity) *ListApplicationEcuResponseBodyEcuInfoList {
	s.EcuEntity = v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoList) Validate() error {
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

type ListApplicationEcuResponseBodyEcuInfoListEcuEntity struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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

func (s ListApplicationEcuResponseBodyEcuInfoListEcuEntity) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetAvailableCpu() *int32 {
	return s.AvailableCpu
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetAvailableMem() *int32 {
	return s.AvailableMem
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetDockerEnv() *bool {
	return s.DockerEnv
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetEcuId() *string {
	return s.EcuId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetHeartbeatTime() *int64 {
	return s.HeartbeatTime
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetIpAddr() *string {
	return s.IpAddr
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetMem() *int32 {
	return s.Mem
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetName() *string {
	return s.Name
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetOnline() *bool {
	return s.Online
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetAppId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.AppId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetAvailableCpu(v int32) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetAvailableMem(v int32) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetCpu(v int32) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Cpu = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetCreateTime(v int64) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetDockerEnv(v bool) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetEcuId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.EcuId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetHeartbeatTime(v int64) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetInstanceId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetIpAddr(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetMem(v int32) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Mem = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetName(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Name = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetOnline(v bool) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Online = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetRegionId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.RegionId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetUpdateTime(v int64) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetUserId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.UserId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetVpcId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.VpcId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetZoneId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) Validate() error {
	return dara.Validate(s)
}
