// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentForView interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *IncidentForView
	GetContent() *string
	SetEscalations(v []*IncidentEscalationPolicyForView) *IncidentForView
	GetEscalations() []*IncidentEscalationPolicyForView
	SetGroupUuid(v string) *IncidentForView
	GetGroupUuid() *string
	SetGroupingKeys(v map[string]*string) *IncidentForView
	GetGroupingKeys() map[string]*string
	SetIncidentId(v string) *IncidentForView
	GetIncidentId() *string
	SetNotifyStrategyName(v string) *IncidentForView
	GetNotifyStrategyName() *string
	SetNotifyStrategyUuid(v string) *IncidentForView
	GetNotifyStrategyUuid() *string
	SetOperator(v *ContactForIncidentView) *IncidentForView
	GetOperator() *ContactForIncidentView
	SetOwners(v []*ContactForIncidentView) *IncidentForView
	GetOwners() []*ContactForIncidentView
	SetParticipants(v []*ContactForIncidentView) *IncidentForView
	GetParticipants() []*ContactForIncidentView
	SetPlan(v *IncidentResponsePlanForView) *IncidentForView
	GetPlan() *IncidentResponsePlanForView
	SetRelatedResources(v []*EventResourceForIncidentView) *IncidentForView
	GetRelatedResources() []*EventResourceForIncidentView
	SetRootCauseCategory(v string) *IncidentForView
	GetRootCauseCategory() *string
	SetSeverity(v string) *IncidentForView
	GetSeverity() *string
	SetSolution(v string) *IncidentForView
	GetSolution() *string
	SetState(v string) *IncidentForView
	GetState() *string
	SetSubscriptionName(v string) *IncidentForView
	GetSubscriptionName() *string
	SetSubscriptionUuid(v string) *IncidentForView
	GetSubscriptionUuid() *string
	SetTime(v int64) *IncidentForView
	GetTime() *int64
	SetTitle(v string) *IncidentForView
	GetTitle() *string
	SetUserId(v string) *IncidentForView
	GetUserId() *string
	SetWorkspace(v string) *IncidentForView
	GetWorkspace() *string
}

type IncidentForView struct {
	// The details of the incident.
	//
	// example:
	//
	// "{"description":"ECS实例CPU过高","impact":"支付服务延迟"}"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The list of escalation policies.
	Escalations []*IncidentEscalationPolicyForView `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	// The group identifier.
	//
	// example:
	//
	// "group-123456"
	GroupUuid *string `json:"groupUuid,omitempty" xml:"groupUuid,omitempty"`
	// The key-value pairs for grouping.
	GroupingKeys map[string]*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty"`
	// The unique identifier of the incident.
	//
	// example:
	//
	// "incident-abc123"
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// { "contactId": "contact-123", "name": "张三" }
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// The UUID of the associated notification policy, which is used to trigger notifications.
	//
	// example:
	//
	// "notify-strategy-789"
	NotifyStrategyUuid *string `json:"notifyStrategyUuid,omitempty" xml:"notifyStrategyUuid,omitempty"`
	// The information about the operator.
	Operator *ContactForIncidentView `json:"operator,omitempty" xml:"operator,omitempty"`
	// The list of owners.
	Owners []*ContactForIncidentView `json:"owners,omitempty" xml:"owners,omitempty" type:"Repeated"`
	// The list of participants.
	Participants []*ContactForIncidentView `json:"participants,omitempty" xml:"participants,omitempty" type:"Repeated"`
	// The response plan.
	Plan *IncidentResponsePlanForView `json:"plan,omitempty" xml:"plan,omitempty"`
	// The list of associated resources.
	RelatedResources []*EventResourceForIncidentView `json:"relatedResources,omitempty" xml:"relatedResources,omitempty" type:"Repeated"`
	// The root cause category.
	//
	// example:
	//
	// "Network"
	RootCauseCategory *string `json:"rootCauseCategory,omitempty" xml:"rootCauseCategory,omitempty"`
	// The severity level of the incident.
	//
	// example:
	//
	// "Critical"
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The description of the solution.
	//
	// example:
	//
	// "重启ECS实例后恢复正常"
	Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
	// The current state of the incident.
	//
	// example:
	//
	// "Open"
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The name of the subscription policy.
	//
	// example:
	//
	// "P1-Alert-Notification"
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	// The UUID of the subscription policy.
	//
	// example:
	//
	// "subscription-abc"
	SubscriptionUuid *string `json:"subscriptionUuid,omitempty" xml:"subscriptionUuid,omitempty"`
	// The timestamp when the incident was created.
	//
	// example:
	//
	// 1743876000000
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// The title of the incident.
	//
	// example:
	//
	// "支付服务不可用"
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the user who created the incident.
	//
	// example:
	//
	// "user-abc123"
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// "ws-xyz789"
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentForView) String() string {
	return dara.Prettify(s)
}

func (s IncidentForView) GoString() string {
	return s.String()
}

func (s *IncidentForView) GetContent() *string {
	return s.Content
}

func (s *IncidentForView) GetEscalations() []*IncidentEscalationPolicyForView {
	return s.Escalations
}

func (s *IncidentForView) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *IncidentForView) GetGroupingKeys() map[string]*string {
	return s.GroupingKeys
}

func (s *IncidentForView) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentForView) GetNotifyStrategyName() *string {
	return s.NotifyStrategyName
}

func (s *IncidentForView) GetNotifyStrategyUuid() *string {
	return s.NotifyStrategyUuid
}

func (s *IncidentForView) GetOperator() *ContactForIncidentView {
	return s.Operator
}

func (s *IncidentForView) GetOwners() []*ContactForIncidentView {
	return s.Owners
}

func (s *IncidentForView) GetParticipants() []*ContactForIncidentView {
	return s.Participants
}

func (s *IncidentForView) GetPlan() *IncidentResponsePlanForView {
	return s.Plan
}

func (s *IncidentForView) GetRelatedResources() []*EventResourceForIncidentView {
	return s.RelatedResources
}

func (s *IncidentForView) GetRootCauseCategory() *string {
	return s.RootCauseCategory
}

func (s *IncidentForView) GetSeverity() *string {
	return s.Severity
}

func (s *IncidentForView) GetSolution() *string {
	return s.Solution
}

func (s *IncidentForView) GetState() *string {
	return s.State
}

func (s *IncidentForView) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *IncidentForView) GetSubscriptionUuid() *string {
	return s.SubscriptionUuid
}

func (s *IncidentForView) GetTime() *int64 {
	return s.Time
}

func (s *IncidentForView) GetTitle() *string {
	return s.Title
}

func (s *IncidentForView) GetUserId() *string {
	return s.UserId
}

func (s *IncidentForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentForView) SetContent(v string) *IncidentForView {
	s.Content = &v
	return s
}

func (s *IncidentForView) SetEscalations(v []*IncidentEscalationPolicyForView) *IncidentForView {
	s.Escalations = v
	return s
}

func (s *IncidentForView) SetGroupUuid(v string) *IncidentForView {
	s.GroupUuid = &v
	return s
}

func (s *IncidentForView) SetGroupingKeys(v map[string]*string) *IncidentForView {
	s.GroupingKeys = v
	return s
}

func (s *IncidentForView) SetIncidentId(v string) *IncidentForView {
	s.IncidentId = &v
	return s
}

func (s *IncidentForView) SetNotifyStrategyName(v string) *IncidentForView {
	s.NotifyStrategyName = &v
	return s
}

func (s *IncidentForView) SetNotifyStrategyUuid(v string) *IncidentForView {
	s.NotifyStrategyUuid = &v
	return s
}

func (s *IncidentForView) SetOperator(v *ContactForIncidentView) *IncidentForView {
	s.Operator = v
	return s
}

func (s *IncidentForView) SetOwners(v []*ContactForIncidentView) *IncidentForView {
	s.Owners = v
	return s
}

func (s *IncidentForView) SetParticipants(v []*ContactForIncidentView) *IncidentForView {
	s.Participants = v
	return s
}

func (s *IncidentForView) SetPlan(v *IncidentResponsePlanForView) *IncidentForView {
	s.Plan = v
	return s
}

func (s *IncidentForView) SetRelatedResources(v []*EventResourceForIncidentView) *IncidentForView {
	s.RelatedResources = v
	return s
}

func (s *IncidentForView) SetRootCauseCategory(v string) *IncidentForView {
	s.RootCauseCategory = &v
	return s
}

func (s *IncidentForView) SetSeverity(v string) *IncidentForView {
	s.Severity = &v
	return s
}

func (s *IncidentForView) SetSolution(v string) *IncidentForView {
	s.Solution = &v
	return s
}

func (s *IncidentForView) SetState(v string) *IncidentForView {
	s.State = &v
	return s
}

func (s *IncidentForView) SetSubscriptionName(v string) *IncidentForView {
	s.SubscriptionName = &v
	return s
}

func (s *IncidentForView) SetSubscriptionUuid(v string) *IncidentForView {
	s.SubscriptionUuid = &v
	return s
}

func (s *IncidentForView) SetTime(v int64) *IncidentForView {
	s.Time = &v
	return s
}

func (s *IncidentForView) SetTitle(v string) *IncidentForView {
	s.Title = &v
	return s
}

func (s *IncidentForView) SetUserId(v string) *IncidentForView {
	s.UserId = &v
	return s
}

func (s *IncidentForView) SetWorkspace(v string) *IncidentForView {
	s.Workspace = &v
	return s
}

func (s *IncidentForView) Validate() error {
	if s.Escalations != nil {
		for _, item := range s.Escalations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Operator != nil {
		if err := s.Operator.Validate(); err != nil {
			return err
		}
	}
	if s.Owners != nil {
		for _, item := range s.Owners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Participants != nil {
		for _, item := range s.Participants {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Plan != nil {
		if err := s.Plan.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedResources != nil {
		for _, item := range s.RelatedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
