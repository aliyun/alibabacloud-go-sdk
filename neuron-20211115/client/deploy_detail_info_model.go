// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDetailInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeployDetailInfo
	GetAccountId() *string
	SetAccountName(v string) *DeployDetailInfo
	GetAccountName() *string
	SetApplicationType(v string) *DeployDetailInfo
	GetApplicationType() *string
	SetChangeOrderId(v string) *DeployDetailInfo
	GetChangeOrderId() *string
	SetContext(v string) *DeployDetailInfo
	GetContext() *string
	SetCpu(v int32) *DeployDetailInfo
	GetCpu() *int32
	SetDeploymentType(v string) *DeployDetailInfo
	GetDeploymentType() *string
	SetDescription(v string) *DeployDetailInfo
	GetDescription() *string
	SetEdasId(v string) *DeployDetailInfo
	GetEdasId() *string
	SetEnterpriseId(v int64) *DeployDetailInfo
	GetEnterpriseId() *int64
	SetEnv(v string) *DeployDetailInfo
	GetEnv() *string
	SetGmtCreate(v string) *DeployDetailInfo
	GetGmtCreate() *string
	SetGmtModified(v string) *DeployDetailInfo
	GetGmtModified() *string
	SetId(v int64) *DeployDetailInfo
	GetId() *int64
	SetImageId(v string) *DeployDetailInfo
	GetImageId() *string
	SetImageTag(v string) *DeployDetailInfo
	GetImageTag() *string
	SetInstanceCount(v int32) *DeployDetailInfo
	GetInstanceCount() *int32
	SetIsMonitorEnable(v int32) *DeployDetailInfo
	GetIsMonitorEnable() *int32
	SetMemory(v int32) *DeployDetailInfo
	GetMemory() *int32
	SetMessage(v string) *DeployDetailInfo
	GetMessage() *string
	SetOrgType(v string) *DeployDetailInfo
	GetOrgType() *string
	SetPbcId(v int64) *DeployDetailInfo
	GetPbcId() *int64
	SetPbcName(v string) *DeployDetailInfo
	GetPbcName() *string
	SetPipelineId(v string) *DeployDetailInfo
	GetPipelineId() *string
	SetPipelineInfos(v []*DeployPipelineInfo) *DeployDetailInfo
	GetPipelineInfos() []*DeployPipelineInfo
	SetPipelineRunId(v string) *DeployDetailInfo
	GetPipelineRunId() *string
	SetRepositoryId(v string) *DeployDetailInfo
	GetRepositoryId() *string
	SetRollbackStatus(v string) *DeployDetailInfo
	GetRollbackStatus() *string
	SetServiceId(v int64) *DeployDetailInfo
	GetServiceId() *int64
	SetServiceName(v string) *DeployDetailInfo
	GetServiceName() *string
	SetStatus(v string) *DeployDetailInfo
	GetStatus() *string
	SetTimes(v int32) *DeployDetailInfo
	GetTimes() *int32
	SetType(v string) *DeployDetailInfo
	GetType() *string
}

type DeployDetailInfo struct {
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// account1
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
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
	// example:
	//
	// sdad
	EdasId *string `json:"edasId,omitempty" xml:"edasId,omitempty"`
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
	InstanceCount   *int32 `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	IsMonitorEnable *int32 `json:"isMonitorEnable,omitempty" xml:"isMonitorEnable,omitempty"`
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
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// management
	PbcName *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	// example:
	//
	// 123141
	PipelineId    *string               `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineInfos []*DeployPipelineInfo `json:"pipelineInfos,omitempty" xml:"pipelineInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 23131
	PipelineRunId *string `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// example:
	//
	// cri-v1d49e57e8ojcwpq
	RepositoryId *string `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
	// example:
	//
	// DEPLOYMENT_PENDING
	RollbackStatus *string `json:"rollbackStatus,omitempty" xml:"rollbackStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// employee
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// example:
	//
	// DEPLOYMENT_PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
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

func (s DeployDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DeployDetailInfo) GoString() string {
	return s.String()
}

func (s *DeployDetailInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *DeployDetailInfo) GetAccountName() *string {
	return s.AccountName
}

func (s *DeployDetailInfo) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DeployDetailInfo) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeployDetailInfo) GetContext() *string {
	return s.Context
}

func (s *DeployDetailInfo) GetCpu() *int32 {
	return s.Cpu
}

func (s *DeployDetailInfo) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeployDetailInfo) GetDescription() *string {
	return s.Description
}

func (s *DeployDetailInfo) GetEdasId() *string {
	return s.EdasId
}

func (s *DeployDetailInfo) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *DeployDetailInfo) GetEnv() *string {
	return s.Env
}

func (s *DeployDetailInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeployDetailInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeployDetailInfo) GetId() *int64 {
	return s.Id
}

func (s *DeployDetailInfo) GetImageId() *string {
	return s.ImageId
}

func (s *DeployDetailInfo) GetImageTag() *string {
	return s.ImageTag
}

func (s *DeployDetailInfo) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DeployDetailInfo) GetIsMonitorEnable() *int32 {
	return s.IsMonitorEnable
}

func (s *DeployDetailInfo) GetMemory() *int32 {
	return s.Memory
}

func (s *DeployDetailInfo) GetMessage() *string {
	return s.Message
}

func (s *DeployDetailInfo) GetOrgType() *string {
	return s.OrgType
}

func (s *DeployDetailInfo) GetPbcId() *int64 {
	return s.PbcId
}

func (s *DeployDetailInfo) GetPbcName() *string {
	return s.PbcName
}

func (s *DeployDetailInfo) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeployDetailInfo) GetPipelineInfos() []*DeployPipelineInfo {
	return s.PipelineInfos
}

func (s *DeployDetailInfo) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *DeployDetailInfo) GetRepositoryId() *string {
	return s.RepositoryId
}

func (s *DeployDetailInfo) GetRollbackStatus() *string {
	return s.RollbackStatus
}

func (s *DeployDetailInfo) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *DeployDetailInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeployDetailInfo) GetStatus() *string {
	return s.Status
}

func (s *DeployDetailInfo) GetTimes() *int32 {
	return s.Times
}

func (s *DeployDetailInfo) GetType() *string {
	return s.Type
}

func (s *DeployDetailInfo) SetAccountId(v string) *DeployDetailInfo {
	s.AccountId = &v
	return s
}

func (s *DeployDetailInfo) SetAccountName(v string) *DeployDetailInfo {
	s.AccountName = &v
	return s
}

func (s *DeployDetailInfo) SetApplicationType(v string) *DeployDetailInfo {
	s.ApplicationType = &v
	return s
}

func (s *DeployDetailInfo) SetChangeOrderId(v string) *DeployDetailInfo {
	s.ChangeOrderId = &v
	return s
}

func (s *DeployDetailInfo) SetContext(v string) *DeployDetailInfo {
	s.Context = &v
	return s
}

func (s *DeployDetailInfo) SetCpu(v int32) *DeployDetailInfo {
	s.Cpu = &v
	return s
}

func (s *DeployDetailInfo) SetDeploymentType(v string) *DeployDetailInfo {
	s.DeploymentType = &v
	return s
}

func (s *DeployDetailInfo) SetDescription(v string) *DeployDetailInfo {
	s.Description = &v
	return s
}

func (s *DeployDetailInfo) SetEdasId(v string) *DeployDetailInfo {
	s.EdasId = &v
	return s
}

func (s *DeployDetailInfo) SetEnterpriseId(v int64) *DeployDetailInfo {
	s.EnterpriseId = &v
	return s
}

func (s *DeployDetailInfo) SetEnv(v string) *DeployDetailInfo {
	s.Env = &v
	return s
}

func (s *DeployDetailInfo) SetGmtCreate(v string) *DeployDetailInfo {
	s.GmtCreate = &v
	return s
}

func (s *DeployDetailInfo) SetGmtModified(v string) *DeployDetailInfo {
	s.GmtModified = &v
	return s
}

func (s *DeployDetailInfo) SetId(v int64) *DeployDetailInfo {
	s.Id = &v
	return s
}

func (s *DeployDetailInfo) SetImageId(v string) *DeployDetailInfo {
	s.ImageId = &v
	return s
}

func (s *DeployDetailInfo) SetImageTag(v string) *DeployDetailInfo {
	s.ImageTag = &v
	return s
}

func (s *DeployDetailInfo) SetInstanceCount(v int32) *DeployDetailInfo {
	s.InstanceCount = &v
	return s
}

func (s *DeployDetailInfo) SetIsMonitorEnable(v int32) *DeployDetailInfo {
	s.IsMonitorEnable = &v
	return s
}

func (s *DeployDetailInfo) SetMemory(v int32) *DeployDetailInfo {
	s.Memory = &v
	return s
}

func (s *DeployDetailInfo) SetMessage(v string) *DeployDetailInfo {
	s.Message = &v
	return s
}

func (s *DeployDetailInfo) SetOrgType(v string) *DeployDetailInfo {
	s.OrgType = &v
	return s
}

func (s *DeployDetailInfo) SetPbcId(v int64) *DeployDetailInfo {
	s.PbcId = &v
	return s
}

func (s *DeployDetailInfo) SetPbcName(v string) *DeployDetailInfo {
	s.PbcName = &v
	return s
}

func (s *DeployDetailInfo) SetPipelineId(v string) *DeployDetailInfo {
	s.PipelineId = &v
	return s
}

func (s *DeployDetailInfo) SetPipelineInfos(v []*DeployPipelineInfo) *DeployDetailInfo {
	s.PipelineInfos = v
	return s
}

func (s *DeployDetailInfo) SetPipelineRunId(v string) *DeployDetailInfo {
	s.PipelineRunId = &v
	return s
}

func (s *DeployDetailInfo) SetRepositoryId(v string) *DeployDetailInfo {
	s.RepositoryId = &v
	return s
}

func (s *DeployDetailInfo) SetRollbackStatus(v string) *DeployDetailInfo {
	s.RollbackStatus = &v
	return s
}

func (s *DeployDetailInfo) SetServiceId(v int64) *DeployDetailInfo {
	s.ServiceId = &v
	return s
}

func (s *DeployDetailInfo) SetServiceName(v string) *DeployDetailInfo {
	s.ServiceName = &v
	return s
}

func (s *DeployDetailInfo) SetStatus(v string) *DeployDetailInfo {
	s.Status = &v
	return s
}

func (s *DeployDetailInfo) SetTimes(v int32) *DeployDetailInfo {
	s.Times = &v
	return s
}

func (s *DeployDetailInfo) SetType(v string) *DeployDetailInfo {
	s.Type = &v
	return s
}

func (s *DeployDetailInfo) Validate() error {
	if s.PipelineInfos != nil {
		for _, item := range s.PipelineInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
