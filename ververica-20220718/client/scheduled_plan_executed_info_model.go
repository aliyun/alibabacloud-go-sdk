// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledPlanExecutedInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *ScheduledPlanExecutedInfo
	GetCreatedAt() *string
	SetCreator(v string) *ScheduledPlanExecutedInfo
	GetCreator() *string
	SetCreatorName(v string) *ScheduledPlanExecutedInfo
	GetCreatorName() *string
	SetDeploymentId(v string) *ScheduledPlanExecutedInfo
	GetDeploymentId() *string
	SetJobResourceUpgradingId(v string) *ScheduledPlanExecutedInfo
	GetJobResourceUpgradingId() *string
	SetModifiedAt(v string) *ScheduledPlanExecutedInfo
	GetModifiedAt() *string
	SetModifier(v string) *ScheduledPlanExecutedInfo
	GetModifier() *string
	SetModifierName(v string) *ScheduledPlanExecutedInfo
	GetModifierName() *string
	SetName(v string) *ScheduledPlanExecutedInfo
	GetName() *string
	SetNamespace(v string) *ScheduledPlanExecutedInfo
	GetNamespace() *string
	SetOrigin(v string) *ScheduledPlanExecutedInfo
	GetOrigin() *string
	SetOriginJobId(v string) *ScheduledPlanExecutedInfo
	GetOriginJobId() *string
	SetStatus(v *ScheduledPlanExecutedStatus) *ScheduledPlanExecutedInfo
	GetStatus() *ScheduledPlanExecutedStatus
	SetWorkspace(v string) *ScheduledPlanExecutedInfo
	GetWorkspace() *string
}

type ScheduledPlanExecutedInfo struct {
	// example:
	//
	// 1723197248
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000001
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 0e6d3bab-2277-4ed1-b573-9de6413d****
	JobResourceUpgradingId *string `json:"jobResourceUpgradingId,omitempty" xml:"jobResourceUpgradingId,omitempty"`
	// example:
	//
	// 1723197248
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// test-scheduled-plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// SCHEDULED_PLAN
	Origin *string `json:"origin,omitempty" xml:"origin,omitempty"`
	// example:
	//
	// f8a2d5d9-9fc5-4273-bfcc-2a3cd354****
	OriginJobId *string                      `json:"originJobId,omitempty" xml:"originJobId,omitempty"`
	Status      *ScheduledPlanExecutedStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ScheduledPlanExecutedInfo) String() string {
	return dara.Prettify(s)
}

func (s ScheduledPlanExecutedInfo) GoString() string {
	return s.String()
}

func (s *ScheduledPlanExecutedInfo) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ScheduledPlanExecutedInfo) GetCreator() *string {
	return s.Creator
}

func (s *ScheduledPlanExecutedInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ScheduledPlanExecutedInfo) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ScheduledPlanExecutedInfo) GetJobResourceUpgradingId() *string {
	return s.JobResourceUpgradingId
}

func (s *ScheduledPlanExecutedInfo) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *ScheduledPlanExecutedInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ScheduledPlanExecutedInfo) GetModifierName() *string {
	return s.ModifierName
}

func (s *ScheduledPlanExecutedInfo) GetName() *string {
	return s.Name
}

func (s *ScheduledPlanExecutedInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *ScheduledPlanExecutedInfo) GetOrigin() *string {
	return s.Origin
}

func (s *ScheduledPlanExecutedInfo) GetOriginJobId() *string {
	return s.OriginJobId
}

func (s *ScheduledPlanExecutedInfo) GetStatus() *ScheduledPlanExecutedStatus {
	return s.Status
}

func (s *ScheduledPlanExecutedInfo) GetWorkspace() *string {
	return s.Workspace
}

func (s *ScheduledPlanExecutedInfo) SetCreatedAt(v string) *ScheduledPlanExecutedInfo {
	s.CreatedAt = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetCreator(v string) *ScheduledPlanExecutedInfo {
	s.Creator = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetCreatorName(v string) *ScheduledPlanExecutedInfo {
	s.CreatorName = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetDeploymentId(v string) *ScheduledPlanExecutedInfo {
	s.DeploymentId = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetJobResourceUpgradingId(v string) *ScheduledPlanExecutedInfo {
	s.JobResourceUpgradingId = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetModifiedAt(v string) *ScheduledPlanExecutedInfo {
	s.ModifiedAt = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetModifier(v string) *ScheduledPlanExecutedInfo {
	s.Modifier = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetModifierName(v string) *ScheduledPlanExecutedInfo {
	s.ModifierName = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetName(v string) *ScheduledPlanExecutedInfo {
	s.Name = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetNamespace(v string) *ScheduledPlanExecutedInfo {
	s.Namespace = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetOrigin(v string) *ScheduledPlanExecutedInfo {
	s.Origin = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetOriginJobId(v string) *ScheduledPlanExecutedInfo {
	s.OriginJobId = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetStatus(v *ScheduledPlanExecutedStatus) *ScheduledPlanExecutedInfo {
	s.Status = v
	return s
}

func (s *ScheduledPlanExecutedInfo) SetWorkspace(v string) *ScheduledPlanExecutedInfo {
	s.Workspace = &v
	return s
}

func (s *ScheduledPlanExecutedInfo) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
