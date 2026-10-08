// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsNotInClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEcsNotInClusterResponseBody
	GetCode() *int32
	SetEcsEntityList(v *ListEcsNotInClusterResponseBodyEcsEntityList) *ListEcsNotInClusterResponseBody
	GetEcsEntityList() *ListEcsNotInClusterResponseBodyEcsEntityList
	SetMessage(v string) *ListEcsNotInClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEcsNotInClusterResponseBody
	GetRequestId() *string
}

type ListEcsNotInClusterResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code          *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	EcsEntityList *ListEcsNotInClusterResponseBodyEcsEntityList `json:"EcsEntityList,omitempty" xml:"EcsEntityList,omitempty" type:"Struct"`
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
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEcsNotInClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEcsNotInClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEcsNotInClusterResponseBody) GetEcsEntityList() *ListEcsNotInClusterResponseBodyEcsEntityList {
	return s.EcsEntityList
}

func (s *ListEcsNotInClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEcsNotInClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEcsNotInClusterResponseBody) SetCode(v int32) *ListEcsNotInClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ListEcsNotInClusterResponseBody) SetEcsEntityList(v *ListEcsNotInClusterResponseBodyEcsEntityList) *ListEcsNotInClusterResponseBody {
	s.EcsEntityList = v
	return s
}

func (s *ListEcsNotInClusterResponseBody) SetMessage(v string) *ListEcsNotInClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ListEcsNotInClusterResponseBody) SetRequestId(v string) *ListEcsNotInClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBody) Validate() error {
	if s.EcsEntityList != nil {
		if err := s.EcsEntityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEcsNotInClusterResponseBodyEcsEntityList struct {
	EcsEntity []*ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity `json:"EcsEntity,omitempty" xml:"EcsEntity,omitempty" type:"Repeated"`
}

func (s ListEcsNotInClusterResponseBodyEcsEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListEcsNotInClusterResponseBodyEcsEntityList) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityList) GetEcsEntity() []*ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	return s.EcsEntity
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityList) SetEcsEntity(v []*ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) *ListEcsNotInClusterResponseBodyEcsEntityList {
	s.EcsEntity = v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityList) Validate() error {
	if s.EcsEntity != nil {
		for _, item := range s.EcsEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity struct {
	Cpu          *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Eip          *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	Expired      *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	InnerIp      *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Mem          *int32  `json:"Mem,omitempty" xml:"Mem,omitempty"`
	PrivateIp    *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp     *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName      *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) String() string {
	return dara.Prettify(s)
}

func (s ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetEip() *string {
	return s.Eip
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetExpired() *bool {
	return s.Expired
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetInnerIp() *string {
	return s.InnerIp
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetMem() *int32 {
	return s.Mem
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetStatus() *string {
	return s.Status
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GetVpcName() *string {
	return s.VpcName
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetCpu(v int32) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Cpu = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetEip(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Eip = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetExpired(v bool) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Expired = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetInnerIp(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.InnerIp = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetInstanceId(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.InstanceId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetInstanceName(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.InstanceName = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetMem(v int32) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Mem = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetPrivateIp(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.PrivateIp = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetPublicIp(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.PublicIp = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetRegionId(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.RegionId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetStatus(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Status = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetVpcId(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.VpcId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetVpcName(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.VpcName = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) Validate() error {
	return dara.Validate(s)
}
