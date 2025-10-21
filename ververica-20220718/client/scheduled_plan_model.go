// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledPlan interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *ScheduledPlan
	GetCreatedAt() *string
	SetCreator(v string) *ScheduledPlan
	GetCreator() *string
	SetCreatorName(v string) *ScheduledPlan
	GetCreatorName() *string
	SetDeploymentId(v string) *ScheduledPlan
	GetDeploymentId() *string
	SetModifiedAt(v string) *ScheduledPlan
	GetModifiedAt() *string
	SetModifier(v string) *ScheduledPlan
	GetModifier() *string
	SetModifierName(v string) *ScheduledPlan
	GetModifierName() *string
	SetName(v string) *ScheduledPlan
	GetName() *string
	SetNamespace(v string) *ScheduledPlan
	GetNamespace() *string
	SetOrigin(v string) *ScheduledPlan
	GetOrigin() *string
	SetPeriodicSchedulingPolicies(v []*PeriodicSchedulingPolicy) *ScheduledPlan
	GetPeriodicSchedulingPolicies() []*PeriodicSchedulingPolicy
	SetScheduledPlanId(v string) *ScheduledPlan
	GetScheduledPlanId() *string
	SetStatus(v string) *ScheduledPlan
	GetStatus() *string
	SetUpdatedByUser(v bool) *ScheduledPlan
	GetUpdatedByUser() *bool
	SetWorkspace(v string) *ScheduledPlan
	GetWorkspace() *string
}

type ScheduledPlan struct {
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
	// USER_DEFINED
	Origin                     *string                     `json:"origin,omitempty" xml:"origin,omitempty"`
	PeriodicSchedulingPolicies []*PeriodicSchedulingPolicy `json:"periodicSchedulingPolicies,omitempty" xml:"periodicSchedulingPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// f3b4ec1e-85dc-4b1d-9726-1d7f4c37****
	ScheduledPlanId *string `json:"scheduledPlanId,omitempty" xml:"scheduledPlanId,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	UpdatedByUser *bool `json:"updatedByUser,omitempty" xml:"updatedByUser,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ScheduledPlan) String() string {
	return dara.Prettify(s)
}

func (s ScheduledPlan) GoString() string {
	return s.String()
}

func (s *ScheduledPlan) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ScheduledPlan) GetCreator() *string {
	return s.Creator
}

func (s *ScheduledPlan) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ScheduledPlan) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ScheduledPlan) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *ScheduledPlan) GetModifier() *string {
	return s.Modifier
}

func (s *ScheduledPlan) GetModifierName() *string {
	return s.ModifierName
}

func (s *ScheduledPlan) GetName() *string {
	return s.Name
}

func (s *ScheduledPlan) GetNamespace() *string {
	return s.Namespace
}

func (s *ScheduledPlan) GetOrigin() *string {
	return s.Origin
}

func (s *ScheduledPlan) GetPeriodicSchedulingPolicies() []*PeriodicSchedulingPolicy {
	return s.PeriodicSchedulingPolicies
}

func (s *ScheduledPlan) GetScheduledPlanId() *string {
	return s.ScheduledPlanId
}

func (s *ScheduledPlan) GetStatus() *string {
	return s.Status
}

func (s *ScheduledPlan) GetUpdatedByUser() *bool {
	return s.UpdatedByUser
}

func (s *ScheduledPlan) GetWorkspace() *string {
	return s.Workspace
}

func (s *ScheduledPlan) SetCreatedAt(v string) *ScheduledPlan {
	s.CreatedAt = &v
	return s
}

func (s *ScheduledPlan) SetCreator(v string) *ScheduledPlan {
	s.Creator = &v
	return s
}

func (s *ScheduledPlan) SetCreatorName(v string) *ScheduledPlan {
	s.CreatorName = &v
	return s
}

func (s *ScheduledPlan) SetDeploymentId(v string) *ScheduledPlan {
	s.DeploymentId = &v
	return s
}

func (s *ScheduledPlan) SetModifiedAt(v string) *ScheduledPlan {
	s.ModifiedAt = &v
	return s
}

func (s *ScheduledPlan) SetModifier(v string) *ScheduledPlan {
	s.Modifier = &v
	return s
}

func (s *ScheduledPlan) SetModifierName(v string) *ScheduledPlan {
	s.ModifierName = &v
	return s
}

func (s *ScheduledPlan) SetName(v string) *ScheduledPlan {
	s.Name = &v
	return s
}

func (s *ScheduledPlan) SetNamespace(v string) *ScheduledPlan {
	s.Namespace = &v
	return s
}

func (s *ScheduledPlan) SetOrigin(v string) *ScheduledPlan {
	s.Origin = &v
	return s
}

func (s *ScheduledPlan) SetPeriodicSchedulingPolicies(v []*PeriodicSchedulingPolicy) *ScheduledPlan {
	s.PeriodicSchedulingPolicies = v
	return s
}

func (s *ScheduledPlan) SetScheduledPlanId(v string) *ScheduledPlan {
	s.ScheduledPlanId = &v
	return s
}

func (s *ScheduledPlan) SetStatus(v string) *ScheduledPlan {
	s.Status = &v
	return s
}

func (s *ScheduledPlan) SetUpdatedByUser(v bool) *ScheduledPlan {
	s.UpdatedByUser = &v
	return s
}

func (s *ScheduledPlan) SetWorkspace(v string) *ScheduledPlan {
	s.Workspace = &v
	return s
}

func (s *ScheduledPlan) Validate() error {
	if s.PeriodicSchedulingPolicies != nil {
		for _, item := range s.PeriodicSchedulingPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
