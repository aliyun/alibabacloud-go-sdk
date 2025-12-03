// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunRegion(v string) *CreateHostGroupRequest
	GetAliyunRegion() *string
	SetEcsLabelKey(v string) *CreateHostGroupRequest
	GetEcsLabelKey() *string
	SetEcsLabelValue(v string) *CreateHostGroupRequest
	GetEcsLabelValue() *string
	SetEcsType(v string) *CreateHostGroupRequest
	GetEcsType() *string
	SetEnvId(v string) *CreateHostGroupRequest
	GetEnvId() *string
	SetMachineInfos(v string) *CreateHostGroupRequest
	GetMachineInfos() *string
	SetName(v string) *CreateHostGroupRequest
	GetName() *string
	SetServiceConnectionId(v int64) *CreateHostGroupRequest
	GetServiceConnectionId() *int64
	SetTagIds(v string) *CreateHostGroupRequest
	GetTagIds() *string
	SetType(v string) *CreateHostGroupRequest
	GetType() *string
}

type CreateHostGroupRequest struct {
	// example:
	//
	// cn-beijing
	AliyunRegion *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
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
	EcsType *string `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	// example:
	//
	// 0
	EnvId *string `json:"envId,omitempty" xml:"envId,omitempty"`
	// example:
	//
	// [{"aliyunRegionId":"cn-beijing","machineSn":"i-sssssss","instanceName":"ceshi","ip":"120.0.0.0"}]
	MachineInfos *string `json:"machineInfos,omitempty" xml:"machineInfos,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	// example:
	//
	// 12,234
	TagIds *string `json:"tagIds,omitempty" xml:"tagIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHostGroupRequest) GetAliyunRegion() *string {
	return s.AliyunRegion
}

func (s *CreateHostGroupRequest) GetEcsLabelKey() *string {
	return s.EcsLabelKey
}

func (s *CreateHostGroupRequest) GetEcsLabelValue() *string {
	return s.EcsLabelValue
}

func (s *CreateHostGroupRequest) GetEcsType() *string {
	return s.EcsType
}

func (s *CreateHostGroupRequest) GetEnvId() *string {
	return s.EnvId
}

func (s *CreateHostGroupRequest) GetMachineInfos() *string {
	return s.MachineInfos
}

func (s *CreateHostGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateHostGroupRequest) GetServiceConnectionId() *int64 {
	return s.ServiceConnectionId
}

func (s *CreateHostGroupRequest) GetTagIds() *string {
	return s.TagIds
}

func (s *CreateHostGroupRequest) GetType() *string {
	return s.Type
}

func (s *CreateHostGroupRequest) SetAliyunRegion(v string) *CreateHostGroupRequest {
	s.AliyunRegion = &v
	return s
}

func (s *CreateHostGroupRequest) SetEcsLabelKey(v string) *CreateHostGroupRequest {
	s.EcsLabelKey = &v
	return s
}

func (s *CreateHostGroupRequest) SetEcsLabelValue(v string) *CreateHostGroupRequest {
	s.EcsLabelValue = &v
	return s
}

func (s *CreateHostGroupRequest) SetEcsType(v string) *CreateHostGroupRequest {
	s.EcsType = &v
	return s
}

func (s *CreateHostGroupRequest) SetEnvId(v string) *CreateHostGroupRequest {
	s.EnvId = &v
	return s
}

func (s *CreateHostGroupRequest) SetMachineInfos(v string) *CreateHostGroupRequest {
	s.MachineInfos = &v
	return s
}

func (s *CreateHostGroupRequest) SetName(v string) *CreateHostGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateHostGroupRequest) SetServiceConnectionId(v int64) *CreateHostGroupRequest {
	s.ServiceConnectionId = &v
	return s
}

func (s *CreateHostGroupRequest) SetTagIds(v string) *CreateHostGroupRequest {
	s.TagIds = &v
	return s
}

func (s *CreateHostGroupRequest) SetType(v string) *CreateHostGroupRequest {
	s.Type = &v
	return s
}

func (s *CreateHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
