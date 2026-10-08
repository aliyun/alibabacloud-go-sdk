// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListResourceGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *ListResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroupList(v *ListResourceGroupResponseBodyResourceGroupList) *ListResourceGroupResponseBody
	GetResourceGroupList() *ListResourceGroupResponseBodyResourceGroupList
}

type ListResourceGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// b197-40ab-9155-****
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupList *ListResourceGroupResponseBodyResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Struct"`
}

func (s ListResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupResponseBody) GetResourceGroupList() *ListResourceGroupResponseBodyResourceGroupList {
	return s.ResourceGroupList
}

func (s *ListResourceGroupResponseBody) SetCode(v int32) *ListResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourceGroupResponseBody) SetMessage(v string) *ListResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourceGroupResponseBody) SetRequestId(v string) *ListResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupResponseBody) SetResourceGroupList(v *ListResourceGroupResponseBodyResourceGroupList) *ListResourceGroupResponseBody {
	s.ResourceGroupList = v
	return s
}

func (s *ListResourceGroupResponseBody) Validate() error {
	if s.ResourceGroupList != nil {
		if err := s.ResourceGroupList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupResponseBodyResourceGroupList struct {
	ResGroupEntity []*ListResourceGroupResponseBodyResourceGroupListResGroupEntity `json:"ResGroupEntity,omitempty" xml:"ResGroupEntity,omitempty" type:"Repeated"`
}

func (s ListResourceGroupResponseBodyResourceGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupList) GetResGroupEntity() []*ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	return s.ResGroupEntity
}

func (s *ListResourceGroupResponseBodyResourceGroupList) SetResGroupEntity(v []*ListResourceGroupResponseBodyResourceGroupListResGroupEntity) *ListResourceGroupResponseBodyResourceGroupList {
	s.ResGroupEntity = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupList) Validate() error {
	if s.ResGroupEntity != nil {
		for _, item := range s.ResGroupEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntity struct {
	AdminUserId *string                                                              `json:"AdminUserId,omitempty" xml:"AdminUserId,omitempty"`
	CreateTime  *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlbList     *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList `json:"SlbList,omitempty" xml:"SlbList,omitempty" type:"Struct"`
	UpdateTime  *int64                                                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	EcsList     *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList `json:"ecsList,omitempty" xml:"ecsList,omitempty" type:"Struct"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntity) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetAdminUserId() *string {
	return s.AdminUserId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetDescription() *string {
	return s.Description
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetId() *int64 {
	return s.Id
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetSlbList() *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList {
	return s.SlbList
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GetEcsList() *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList {
	return s.EcsList
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetAdminUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.AdminUserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetCreateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetDescription(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.Description = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetId(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.Id = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.Name = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetSlbList(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.SlbList = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetUpdateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetEcsList(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.EcsList = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) Validate() error {
	if s.SlbList != nil {
		if err := s.SlbList.Validate(); err != nil {
			return err
		}
	}
	if s.EcsList != nil {
		if err := s.EcsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList struct {
	SlbEntity []*ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity `json:"SlbEntity,omitempty" xml:"SlbEntity,omitempty" type:"Repeated"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) GetSlbEntity() []*ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	return s.SlbEntity
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) SetSlbEntity(v []*ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList {
	s.SlbEntity = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) Validate() error {
	if s.SlbEntity != nil {
		for _, item := range s.SlbEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Expired     *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	GroupId     *int32  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlbId       *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbName     *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	SlbStatus   *string `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId   *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetAddress() *string {
	return s.Address
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetAddressType() *string {
	return s.AddressType
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetExpired() *bool {
	return s.Expired
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetSlbId() *string {
	return s.SlbId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetSlbName() *string {
	return s.SlbName
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetSlbStatus() *string {
	return s.SlbStatus
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetAddress(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.Address = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetAddressType(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.AddressType = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetExpired(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.Expired = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetGroupId(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.GroupId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetNetworkType(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.NetworkType = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetSlbId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.SlbId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetSlbName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.SlbName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetSlbStatus(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.SlbStatus = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetVswitchId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.VswitchId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList struct {
	EcsEntity []*ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity `json:"EcsEntity,omitempty" xml:"EcsEntity,omitempty" type:"Repeated"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) GetEcsEntity() []*ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	return s.EcsEntity
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) SetEcsEntity(v []*ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList {
	s.EcsEntity = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) Validate() error {
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

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity struct {
	Cpu          *int32                                                                                 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Description  *string                                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EcuEntity    *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" type:"Struct"`
	Eip          *string                                                                                `json:"Eip,omitempty" xml:"Eip,omitempty"`
	Expired      *bool                                                                                  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	GroupId      *string                                                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HostName     *string                                                                                `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InnerIp      *string                                                                                `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InstanceId   *string                                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string                                                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Mem          *int32                                                                                 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	PrivateIp    *string                                                                                `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp     *string                                                                                `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	RegionId     *string                                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNum    *string                                                                                `json:"SerialNum,omitempty" xml:"SerialNum,omitempty"`
	SgId         *string                                                                                `json:"SgId,omitempty" xml:"SgId,omitempty"`
	Status       *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId       *string                                                                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcEntity    *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity `json:"VpcEntity,omitempty" xml:"VpcEntity,omitempty" type:"Struct"`
	VpcId        *string                                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId       *string                                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetDescription() *string {
	return s.Description
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetEcuEntity() *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	return s.EcuEntity
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetEip() *string {
	return s.Eip
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetExpired() *bool {
	return s.Expired
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetGroupId() *string {
	return s.GroupId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetHostName() *string {
	return s.HostName
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetInnerIp() *string {
	return s.InnerIp
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetMem() *int32 {
	return s.Mem
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetSerialNum() *string {
	return s.SerialNum
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetSgId() *string {
	return s.SgId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetVpcEntity() *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	return s.VpcEntity
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetCpu(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Cpu = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetDescription(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Description = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetEcuEntity(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.EcuEntity = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetEip(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Eip = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetExpired(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Expired = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetGroupId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.GroupId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetHostName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.HostName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetInnerIp(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.InnerIp = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetInstanceId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.InstanceId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetInstanceName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.InstanceName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetMem(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Mem = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetPrivateIp(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.PrivateIp = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetPublicIp(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.PublicIp = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetSerialNum(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.SerialNum = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetSgId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.SgId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetStatus(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Status = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetVpcEntity(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.VpcEntity = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetZoneId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.ZoneId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) Validate() error {
	if s.EcuEntity != nil {
		if err := s.EcuEntity.Validate(); err != nil {
			return err
		}
	}
	if s.VpcEntity != nil {
		if err := s.VpcEntity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity struct {
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

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetAvailableCpu() *int32 {
	return s.AvailableCpu
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetAvailableMem() *int32 {
	return s.AvailableMem
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetDockerEnv() *bool {
	return s.DockerEnv
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetEcuId() *string {
	return s.EcuId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetHeartbeatTime() *int64 {
	return s.HeartbeatTime
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetIpAddr() *string {
	return s.IpAddr
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetMem() *int32 {
	return s.Mem
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetOnline() *bool {
	return s.Online
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetAvailableCpu(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetAvailableMem(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetCpu(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Cpu = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetCreateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetDockerEnv(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetEcuId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.EcuId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetHeartbeatTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetInstanceId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetIpAddr(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetMem(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Mem = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Name = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetOnline(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Online = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetUpdateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetZoneId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity struct {
	Cidrblock   *string `json:"Cidrblock,omitempty" xml:"Cidrblock,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsNum      *int32  `json:"EcsNum,omitempty" xml:"EcsNum,omitempty"`
	Expired     *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName     *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetCidrblock() *string {
	return s.Cidrblock
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetDescription() *string {
	return s.Description
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetEcsNum() *int32 {
	return s.EcsNum
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetExpired() *bool {
	return s.Expired
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GetVpcName() *string {
	return s.VpcName
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetCidrblock(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Cidrblock = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetDescription(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Description = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetEcsNum(v int32) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.EcsNum = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetExpired(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Expired = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetStatus(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Status = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetVpcName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.VpcName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) Validate() error {
	return dara.Validate(s)
}
