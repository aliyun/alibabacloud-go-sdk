// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunRegion(v string) *UpdateHostGroupRequest
	GetAliyunRegion() *string
	SetEcsLabelKey(v string) *UpdateHostGroupRequest
	GetEcsLabelKey() *string
	SetEcsLabelValue(v string) *UpdateHostGroupRequest
	GetEcsLabelValue() *string
	SetEcsType(v string) *UpdateHostGroupRequest
	GetEcsType() *string
	SetEnvId(v string) *UpdateHostGroupRequest
	GetEnvId() *string
	SetMachineInfos(v string) *UpdateHostGroupRequest
	GetMachineInfos() *string
	SetName(v string) *UpdateHostGroupRequest
	GetName() *string
	SetServiceConnectionId(v int64) *UpdateHostGroupRequest
	GetServiceConnectionId() *int64
	SetTagIds(v string) *UpdateHostGroupRequest
	GetTagIds() *string
	SetType(v string) *UpdateHostGroupRequest
	GetType() *string
}

type UpdateHostGroupRequest struct {
	// example:
	//
	// cn-hangzhou
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
	// 12,23
	TagIds *string `json:"tagIds,omitempty" xml:"tagIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateHostGroupRequest) GetAliyunRegion() *string {
	return s.AliyunRegion
}

func (s *UpdateHostGroupRequest) GetEcsLabelKey() *string {
	return s.EcsLabelKey
}

func (s *UpdateHostGroupRequest) GetEcsLabelValue() *string {
	return s.EcsLabelValue
}

func (s *UpdateHostGroupRequest) GetEcsType() *string {
	return s.EcsType
}

func (s *UpdateHostGroupRequest) GetEnvId() *string {
	return s.EnvId
}

func (s *UpdateHostGroupRequest) GetMachineInfos() *string {
	return s.MachineInfos
}

func (s *UpdateHostGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateHostGroupRequest) GetServiceConnectionId() *int64 {
	return s.ServiceConnectionId
}

func (s *UpdateHostGroupRequest) GetTagIds() *string {
	return s.TagIds
}

func (s *UpdateHostGroupRequest) GetType() *string {
	return s.Type
}

func (s *UpdateHostGroupRequest) SetAliyunRegion(v string) *UpdateHostGroupRequest {
	s.AliyunRegion = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEcsLabelKey(v string) *UpdateHostGroupRequest {
	s.EcsLabelKey = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEcsLabelValue(v string) *UpdateHostGroupRequest {
	s.EcsLabelValue = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEcsType(v string) *UpdateHostGroupRequest {
	s.EcsType = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEnvId(v string) *UpdateHostGroupRequest {
	s.EnvId = &v
	return s
}

func (s *UpdateHostGroupRequest) SetMachineInfos(v string) *UpdateHostGroupRequest {
	s.MachineInfos = &v
	return s
}

func (s *UpdateHostGroupRequest) SetName(v string) *UpdateHostGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateHostGroupRequest) SetServiceConnectionId(v int64) *UpdateHostGroupRequest {
	s.ServiceConnectionId = &v
	return s
}

func (s *UpdateHostGroupRequest) SetTagIds(v string) *UpdateHostGroupRequest {
	s.TagIds = &v
	return s
}

func (s *UpdateHostGroupRequest) SetType(v string) *UpdateHostGroupRequest {
	s.Type = &v
	return s
}

func (s *UpdateHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
