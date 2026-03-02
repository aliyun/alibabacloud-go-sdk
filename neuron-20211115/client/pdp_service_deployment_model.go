// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceDeployment interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpServiceDeployment
	GetAccountId() *string
	SetApplicationType(v string) *PdpServiceDeployment
	GetApplicationType() *string
	SetChangeOrderId(v string) *PdpServiceDeployment
	GetChangeOrderId() *string
	SetContext(v string) *PdpServiceDeployment
	GetContext() *string
	SetCpu(v int32) *PdpServiceDeployment
	GetCpu() *int32
	SetDeploymentType(v string) *PdpServiceDeployment
	GetDeploymentType() *string
	SetDescription(v string) *PdpServiceDeployment
	GetDescription() *string
	SetEdasId(v string) *PdpServiceDeployment
	GetEdasId() *string
	SetEnterpriseId(v int64) *PdpServiceDeployment
	GetEnterpriseId() *int64
	SetEnv(v string) *PdpServiceDeployment
	GetEnv() *string
	SetGmtCreate(v string) *PdpServiceDeployment
	GetGmtCreate() *string
	SetGmtModified(v string) *PdpServiceDeployment
	GetGmtModified() *string
	SetId(v int64) *PdpServiceDeployment
	GetId() *int64
	SetImageId(v string) *PdpServiceDeployment
	GetImageId() *string
	SetImageTag(v string) *PdpServiceDeployment
	GetImageTag() *string
	SetInstanceCount(v int32) *PdpServiceDeployment
	GetInstanceCount() *int32
	SetIsServiceGroupEnable(v int32) *PdpServiceDeployment
	GetIsServiceGroupEnable() *int32
	SetMemory(v int32) *PdpServiceDeployment
	GetMemory() *int32
	SetMessage(v string) *PdpServiceDeployment
	GetMessage() *string
	SetOrgType(v string) *PdpServiceDeployment
	GetOrgType() *string
	SetPbcId(v int64) *PdpServiceDeployment
	GetPbcId() *int64
	SetPbcName(v string) *PdpServiceDeployment
	GetPbcName() *string
	SetPipelineTimes(v int32) *PdpServiceDeployment
	GetPipelineTimes() *int32
	SetRepositoryId(v string) *PdpServiceDeployment
	GetRepositoryId() *string
	SetRequestId(v string) *PdpServiceDeployment
	GetRequestId() *string
	SetRollbackStatus(v string) *PdpServiceDeployment
	GetRollbackStatus() *string
	SetServiceGroupId(v int64) *PdpServiceDeployment
	GetServiceGroupId() *int64
	SetServiceId(v int64) *PdpServiceDeployment
	GetServiceId() *int64
	SetServiceName(v string) *PdpServiceDeployment
	GetServiceName() *string
	SetStatus(v string) *PdpServiceDeployment
	GetStatus() *string
	SetTimeout(v int32) *PdpServiceDeployment
	GetTimeout() *int32
	SetTimes(v int32) *PdpServiceDeployment
	GetTimes() *int32
	SetType(v string) *PdpServiceDeployment
	GetType() *string
}

type PdpServiceDeployment struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// MULTI_TENANT_DAPR
	ApplicationType *string `json:"applicationType,omitempty" xml:"applicationType,omitempty"`
	// example:
	//
	// cd65b247-****-475b-ad4b-7039040d625c
	ChangeOrderId *string `json:"changeOrderId,omitempty" xml:"changeOrderId,omitempty"`
	Context       *string `json:"context,omitempty" xml:"context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// AUTO
	DeploymentType *string `json:"deploymentType,omitempty" xml:"deploymentType,omitempty"`
	// example:
	//
	// 部署员工管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdad
	EdasId *string `json:"edasId,omitempty" xml:"edasId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 2022-2-22 11:11:2
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-2-22 11:11:2
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-v12wpq
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220421100717_prepub
	ImageTag *string `json:"imageTag,omitempty" xml:"imageTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	InstanceCount *int32 `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	// example:
	//
	// 0
	IsServiceGroupEnable *int32 `json:"isServiceGroupEnable,omitempty" xml:"isServiceGroupEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	Memory  *int32  `json:"memory,omitempty" xml:"memory,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// INNER
	OrgType *string `json:"orgType,omitempty" xml:"orgType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// management
	PbcName *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	// example:
	//
	// 12
	PipelineTimes *int32 `json:"pipelineTimes,omitempty" xml:"pipelineTimes,omitempty"`
	// example:
	//
	// cri-v1d49e57e8ojcwpq
	RepositoryId *string `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DEPLOYMENT_PENDING
	RollbackStatus *string `json:"rollbackStatus,omitempty" xml:"rollbackStatus,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// employee
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEPLOYMENT_PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 123
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// 2
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SERVICE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpServiceDeployment) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceDeployment) GoString() string {
	return s.String()
}

func (s *PdpServiceDeployment) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpServiceDeployment) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *PdpServiceDeployment) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *PdpServiceDeployment) GetContext() *string {
	return s.Context
}

func (s *PdpServiceDeployment) GetCpu() *int32 {
	return s.Cpu
}

func (s *PdpServiceDeployment) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *PdpServiceDeployment) GetDescription() *string {
	return s.Description
}

func (s *PdpServiceDeployment) GetEdasId() *string {
	return s.EdasId
}

func (s *PdpServiceDeployment) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *PdpServiceDeployment) GetEnv() *string {
	return s.Env
}

func (s *PdpServiceDeployment) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpServiceDeployment) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PdpServiceDeployment) GetId() *int64 {
	return s.Id
}

func (s *PdpServiceDeployment) GetImageId() *string {
	return s.ImageId
}

func (s *PdpServiceDeployment) GetImageTag() *string {
	return s.ImageTag
}

func (s *PdpServiceDeployment) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *PdpServiceDeployment) GetIsServiceGroupEnable() *int32 {
	return s.IsServiceGroupEnable
}

func (s *PdpServiceDeployment) GetMemory() *int32 {
	return s.Memory
}

func (s *PdpServiceDeployment) GetMessage() *string {
	return s.Message
}

func (s *PdpServiceDeployment) GetOrgType() *string {
	return s.OrgType
}

func (s *PdpServiceDeployment) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpServiceDeployment) GetPbcName() *string {
	return s.PbcName
}

func (s *PdpServiceDeployment) GetPipelineTimes() *int32 {
	return s.PipelineTimes
}

func (s *PdpServiceDeployment) GetRepositoryId() *string {
	return s.RepositoryId
}

func (s *PdpServiceDeployment) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpServiceDeployment) GetRollbackStatus() *string {
	return s.RollbackStatus
}

func (s *PdpServiceDeployment) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *PdpServiceDeployment) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpServiceDeployment) GetServiceName() *string {
	return s.ServiceName
}

func (s *PdpServiceDeployment) GetStatus() *string {
	return s.Status
}

func (s *PdpServiceDeployment) GetTimeout() *int32 {
	return s.Timeout
}

func (s *PdpServiceDeployment) GetTimes() *int32 {
	return s.Times
}

func (s *PdpServiceDeployment) GetType() *string {
	return s.Type
}

func (s *PdpServiceDeployment) SetAccountId(v string) *PdpServiceDeployment {
	s.AccountId = &v
	return s
}

func (s *PdpServiceDeployment) SetApplicationType(v string) *PdpServiceDeployment {
	s.ApplicationType = &v
	return s
}

func (s *PdpServiceDeployment) SetChangeOrderId(v string) *PdpServiceDeployment {
	s.ChangeOrderId = &v
	return s
}

func (s *PdpServiceDeployment) SetContext(v string) *PdpServiceDeployment {
	s.Context = &v
	return s
}

func (s *PdpServiceDeployment) SetCpu(v int32) *PdpServiceDeployment {
	s.Cpu = &v
	return s
}

func (s *PdpServiceDeployment) SetDeploymentType(v string) *PdpServiceDeployment {
	s.DeploymentType = &v
	return s
}

func (s *PdpServiceDeployment) SetDescription(v string) *PdpServiceDeployment {
	s.Description = &v
	return s
}

func (s *PdpServiceDeployment) SetEdasId(v string) *PdpServiceDeployment {
	s.EdasId = &v
	return s
}

func (s *PdpServiceDeployment) SetEnterpriseId(v int64) *PdpServiceDeployment {
	s.EnterpriseId = &v
	return s
}

func (s *PdpServiceDeployment) SetEnv(v string) *PdpServiceDeployment {
	s.Env = &v
	return s
}

func (s *PdpServiceDeployment) SetGmtCreate(v string) *PdpServiceDeployment {
	s.GmtCreate = &v
	return s
}

func (s *PdpServiceDeployment) SetGmtModified(v string) *PdpServiceDeployment {
	s.GmtModified = &v
	return s
}

func (s *PdpServiceDeployment) SetId(v int64) *PdpServiceDeployment {
	s.Id = &v
	return s
}

func (s *PdpServiceDeployment) SetImageId(v string) *PdpServiceDeployment {
	s.ImageId = &v
	return s
}

func (s *PdpServiceDeployment) SetImageTag(v string) *PdpServiceDeployment {
	s.ImageTag = &v
	return s
}

func (s *PdpServiceDeployment) SetInstanceCount(v int32) *PdpServiceDeployment {
	s.InstanceCount = &v
	return s
}

func (s *PdpServiceDeployment) SetIsServiceGroupEnable(v int32) *PdpServiceDeployment {
	s.IsServiceGroupEnable = &v
	return s
}

func (s *PdpServiceDeployment) SetMemory(v int32) *PdpServiceDeployment {
	s.Memory = &v
	return s
}

func (s *PdpServiceDeployment) SetMessage(v string) *PdpServiceDeployment {
	s.Message = &v
	return s
}

func (s *PdpServiceDeployment) SetOrgType(v string) *PdpServiceDeployment {
	s.OrgType = &v
	return s
}

func (s *PdpServiceDeployment) SetPbcId(v int64) *PdpServiceDeployment {
	s.PbcId = &v
	return s
}

func (s *PdpServiceDeployment) SetPbcName(v string) *PdpServiceDeployment {
	s.PbcName = &v
	return s
}

func (s *PdpServiceDeployment) SetPipelineTimes(v int32) *PdpServiceDeployment {
	s.PipelineTimes = &v
	return s
}

func (s *PdpServiceDeployment) SetRepositoryId(v string) *PdpServiceDeployment {
	s.RepositoryId = &v
	return s
}

func (s *PdpServiceDeployment) SetRequestId(v string) *PdpServiceDeployment {
	s.RequestId = &v
	return s
}

func (s *PdpServiceDeployment) SetRollbackStatus(v string) *PdpServiceDeployment {
	s.RollbackStatus = &v
	return s
}

func (s *PdpServiceDeployment) SetServiceGroupId(v int64) *PdpServiceDeployment {
	s.ServiceGroupId = &v
	return s
}

func (s *PdpServiceDeployment) SetServiceId(v int64) *PdpServiceDeployment {
	s.ServiceId = &v
	return s
}

func (s *PdpServiceDeployment) SetServiceName(v string) *PdpServiceDeployment {
	s.ServiceName = &v
	return s
}

func (s *PdpServiceDeployment) SetStatus(v string) *PdpServiceDeployment {
	s.Status = &v
	return s
}

func (s *PdpServiceDeployment) SetTimeout(v int32) *PdpServiceDeployment {
	s.Timeout = &v
	return s
}

func (s *PdpServiceDeployment) SetTimes(v int32) *PdpServiceDeployment {
	s.Times = &v
	return s
}

func (s *PdpServiceDeployment) SetType(v string) *PdpServiceDeployment {
	s.Type = &v
	return s
}

func (s *PdpServiceDeployment) Validate() error {
	return dara.Validate(s)
}
