// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetHostGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetHostGroupResponseBody
	GetErrorMessage() *string
	SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody
	GetHostGroup() *GetHostGroupResponseBodyHostGroup
	SetRequestId(v string) *GetHostGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHostGroupResponseBody
	GetSuccess() *bool
}

type GetHostGroupResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HostGroup    *GetHostGroupResponseBodyHostGroup `json:"hostGroup,omitempty" xml:"hostGroup,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetHostGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetHostGroupResponseBody) GetHostGroup() *GetHostGroupResponseBodyHostGroup {
	return s.HostGroup
}

func (s *GetHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHostGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHostGroupResponseBody) SetErrorCode(v string) *GetHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetHostGroupResponseBody) SetErrorMessage(v string) *GetHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetHostGroupResponseBody) SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody {
	s.HostGroup = v
	return s
}

func (s *GetHostGroupResponseBody) SetRequestId(v string) *GetHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostGroupResponseBody) SetSuccess(v bool) *GetHostGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetHostGroupResponseBody) Validate() error {
	if s.HostGroup != nil {
		if err := s.HostGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHostGroupResponseBodyHostGroup struct {
	// example:
	//
	// cn-bejing
	AliyunRegion *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 111111
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ecs
	EcsLabelKey *string `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	// example:
	//
	// ecs
	EcsLabelValue *string `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	// example:
	//
	// ECS_ALIYUN
	EcsType   *string                                       `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	HostInfos []*GetHostGroupResponseBodyHostGroupHostInfos `json:"hostInfos,omitempty" xml:"hostInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	HostNum *int64 `json:"hostNum,omitempty" xml:"hostNum,omitempty"`
	// example:
	//
	// 1234
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 11111
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1234
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	// example:
	//
	// ECS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1586863220000
	UpateTIme *int64 `json:"upateTIme,omitempty" xml:"upateTIme,omitempty"`
}

func (s GetHostGroupResponseBodyHostGroup) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupResponseBodyHostGroup) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBodyHostGroup) GetAliyunRegion() *string {
	return s.AliyunRegion
}

func (s *GetHostGroupResponseBodyHostGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetHostGroupResponseBodyHostGroup) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetHostGroupResponseBodyHostGroup) GetDescription() *string {
	return s.Description
}

func (s *GetHostGroupResponseBodyHostGroup) GetEcsLabelKey() *string {
	return s.EcsLabelKey
}

func (s *GetHostGroupResponseBodyHostGroup) GetEcsLabelValue() *string {
	return s.EcsLabelValue
}

func (s *GetHostGroupResponseBodyHostGroup) GetEcsType() *string {
	return s.EcsType
}

func (s *GetHostGroupResponseBodyHostGroup) GetHostInfos() []*GetHostGroupResponseBodyHostGroupHostInfos {
	return s.HostInfos
}

func (s *GetHostGroupResponseBodyHostGroup) GetHostNum() *int64 {
	return s.HostNum
}

func (s *GetHostGroupResponseBodyHostGroup) GetId() *int64 {
	return s.Id
}

func (s *GetHostGroupResponseBodyHostGroup) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *GetHostGroupResponseBodyHostGroup) GetName() *string {
	return s.Name
}

func (s *GetHostGroupResponseBodyHostGroup) GetServiceConnectionId() *int64 {
	return s.ServiceConnectionId
}

func (s *GetHostGroupResponseBodyHostGroup) GetType() *string {
	return s.Type
}

func (s *GetHostGroupResponseBodyHostGroup) GetUpateTIme() *int64 {
	return s.UpateTIme
}

func (s *GetHostGroupResponseBodyHostGroup) SetAliyunRegion(v string) *GetHostGroupResponseBodyHostGroup {
	s.AliyunRegion = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetCreateTime(v int64) *GetHostGroupResponseBodyHostGroup {
	s.CreateTime = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetCreatorAccountId(v string) *GetHostGroupResponseBodyHostGroup {
	s.CreatorAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetDescription(v string) *GetHostGroupResponseBodyHostGroup {
	s.Description = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetEcsLabelKey(v string) *GetHostGroupResponseBodyHostGroup {
	s.EcsLabelKey = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetEcsLabelValue(v string) *GetHostGroupResponseBodyHostGroup {
	s.EcsLabelValue = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetEcsType(v string) *GetHostGroupResponseBodyHostGroup {
	s.EcsType = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostInfos(v []*GetHostGroupResponseBodyHostGroupHostInfos) *GetHostGroupResponseBodyHostGroup {
	s.HostInfos = v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostNum(v int64) *GetHostGroupResponseBodyHostGroup {
	s.HostNum = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetId(v int64) *GetHostGroupResponseBodyHostGroup {
	s.Id = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetModifierAccountId(v string) *GetHostGroupResponseBodyHostGroup {
	s.ModifierAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetName(v string) *GetHostGroupResponseBodyHostGroup {
	s.Name = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetServiceConnectionId(v int64) *GetHostGroupResponseBodyHostGroup {
	s.ServiceConnectionId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetType(v string) *GetHostGroupResponseBodyHostGroup {
	s.Type = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetUpateTIme(v int64) *GetHostGroupResponseBodyHostGroup {
	s.UpateTIme = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) Validate() error {
	if s.HostInfos != nil {
		for _, item := range s.HostInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHostGroupResponseBodyHostGroupHostInfos struct {
	// example:
	//
	// cn-hangzhou
	AliyunRegionId *string `json:"aliyunRegionId,omitempty" xml:"aliyunRegionId,omitempty"`
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 11111
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// ceshi
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// example:
	//
	// 1ssasa
	MachineSn *string `json:"machineSn,omitempty" xml:"machineSn,omitempty"`
	// example:
	//
	// 11111111111
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// example:
	//
	// MachineInfo
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// example:
	//
	// 1586863220000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetHostGroupResponseBodyHostGroupHostInfos) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupResponseBodyHostGroupHostInfos) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetAliyunRegionId() *string {
	return s.AliyunRegionId
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetIp() *string {
	return s.Ip
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetMachineSn() *string {
	return s.MachineSn
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetAliyunRegionId(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.AliyunRegionId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetCreateTime(v int64) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.CreateTime = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetCreatorAccountId(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.CreatorAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetInstanceName(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.InstanceName = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetIp(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.Ip = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetMachineSn(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.MachineSn = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetModifierAccountId(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.ModifierAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetObjectType(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.ObjectType = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetUpdateTime(v int64) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.UpdateTime = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) Validate() error {
	return dara.Validate(s)
}
