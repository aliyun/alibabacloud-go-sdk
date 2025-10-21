// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledPlanAppliedInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *ScheduledPlanAppliedInfo
	GetDeploymentId() *string
	SetExpectedState(v string) *ScheduledPlanAppliedInfo
	GetExpectedState() *string
	SetModifiedAt(v string) *ScheduledPlanAppliedInfo
	GetModifiedAt() *string
	SetModifier(v string) *ScheduledPlanAppliedInfo
	GetModifier() *string
	SetModifierName(v string) *ScheduledPlanAppliedInfo
	GetModifierName() *string
	SetNamespace(v string) *ScheduledPlanAppliedInfo
	GetNamespace() *string
	SetScheduledPlanId(v string) *ScheduledPlanAppliedInfo
	GetScheduledPlanId() *string
	SetScheduledPlanName(v string) *ScheduledPlanAppliedInfo
	GetScheduledPlanName() *string
	SetStatusState(v string) *ScheduledPlanAppliedInfo
	GetStatusState() *string
	SetWorkspace(v string) *ScheduledPlanAppliedInfo
	GetWorkspace() *string
}

type ScheduledPlanAppliedInfo struct {
	// example:
	//
	// 00000000-0000-0000-0000-000000000001
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// RUNNING
	ExpectedState *string `json:"expectedState,omitempty" xml:"expectedState,omitempty"`
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
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// f3b4ec1e-85dc-4b1d-9726-1d7f4c37****
	ScheduledPlanId *string `json:"scheduledPlanId,omitempty" xml:"scheduledPlanId,omitempty"`
	// example:
	//
	// test-scheduled-plan
	ScheduledPlanName *string `json:"scheduledPlanName,omitempty" xml:"scheduledPlanName,omitempty"`
	// example:
	//
	// WAITING
	StatusState *string `json:"statusState,omitempty" xml:"statusState,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ScheduledPlanAppliedInfo) String() string {
	return dara.Prettify(s)
}

func (s ScheduledPlanAppliedInfo) GoString() string {
	return s.String()
}

func (s *ScheduledPlanAppliedInfo) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ScheduledPlanAppliedInfo) GetExpectedState() *string {
	return s.ExpectedState
}

func (s *ScheduledPlanAppliedInfo) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *ScheduledPlanAppliedInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ScheduledPlanAppliedInfo) GetModifierName() *string {
	return s.ModifierName
}

func (s *ScheduledPlanAppliedInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *ScheduledPlanAppliedInfo) GetScheduledPlanId() *string {
	return s.ScheduledPlanId
}

func (s *ScheduledPlanAppliedInfo) GetScheduledPlanName() *string {
	return s.ScheduledPlanName
}

func (s *ScheduledPlanAppliedInfo) GetStatusState() *string {
	return s.StatusState
}

func (s *ScheduledPlanAppliedInfo) GetWorkspace() *string {
	return s.Workspace
}

func (s *ScheduledPlanAppliedInfo) SetDeploymentId(v string) *ScheduledPlanAppliedInfo {
	s.DeploymentId = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetExpectedState(v string) *ScheduledPlanAppliedInfo {
	s.ExpectedState = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetModifiedAt(v string) *ScheduledPlanAppliedInfo {
	s.ModifiedAt = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetModifier(v string) *ScheduledPlanAppliedInfo {
	s.Modifier = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetModifierName(v string) *ScheduledPlanAppliedInfo {
	s.ModifierName = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetNamespace(v string) *ScheduledPlanAppliedInfo {
	s.Namespace = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetScheduledPlanId(v string) *ScheduledPlanAppliedInfo {
	s.ScheduledPlanId = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetScheduledPlanName(v string) *ScheduledPlanAppliedInfo {
	s.ScheduledPlanName = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetStatusState(v string) *ScheduledPlanAppliedInfo {
	s.StatusState = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) SetWorkspace(v string) *ScheduledPlanAppliedInfo {
	s.Workspace = &v
	return s
}

func (s *ScheduledPlanAppliedInfo) Validate() error {
	return dara.Validate(s)
}
