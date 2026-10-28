// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUuid(v string) *ListSubscriptionsRequest
	GetAgentUuid() *string
	SetDirection(v string) *ListSubscriptionsRequest
	GetDirection() *string
	SetEnable(v bool) *ListSubscriptionsRequest
	GetEnable() *bool
	SetMaxResults(v int32) *ListSubscriptionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSubscriptionsRequest
	GetNextToken() *string
	SetNotifyStrategyUuid(v string) *ListSubscriptionsRequest
	GetNotifyStrategyUuid() *string
	SetOrderBy(v string) *ListSubscriptionsRequest
	GetOrderBy() *string
	SetSubscriptionId(v string) *ListSubscriptionsRequest
	GetSubscriptionId() *string
	SetSubscriptionName(v string) *ListSubscriptionsRequest
	GetSubscriptionName() *string
	SetSubscriptionType(v string) *ListSubscriptionsRequest
	GetSubscriptionType() *string
	SetUnbound(v bool) *ListSubscriptionsRequest
	GetUnbound() *bool
	SetWorkspace(v string) *ListSubscriptionsRequest
	GetWorkspace() *string
}

type ListSubscriptionsRequest struct {
	// The UUID of the AI agent. This parameter is used to filter subscriptions of a specified AI agent.
	//
	// example:
	//
	// agent-rca-001
	AgentUuid *string `json:"agentUuid,omitempty" xml:"agentUuid,omitempty"`
	// The sort order of the returned versions. Valid values:
	//
	// - asc: Ascending order.
	//
	// - desc: Descending order.
	//
	// example:
	//
	// asc
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// Specifies whether the subscription is enabled.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The maximum number of results per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 2-3408-4595-9d80-****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The UUID of the notification policy. This parameter is used to filter subscriptions that are bound to a specified notification policy. The value is a 32-character string.
	//
	// example:
	//
	// 1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d
	NotifyStrategyUuid *string `json:"notifyStrategyUuid,omitempty" xml:"notifyStrategyUuid,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// id
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The ID of the subscription.
	//
	// example:
	//
	// 123123123123
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// Test subscription
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	// The type of the subscription. Valid values:
	//
	// - NORMAL: Standard subscription.
	//
	// - AI_AGENT: AI agent subscription.
	//
	// example:
	//
	// NORMAL
	SubscriptionType *string `json:"subscriptionType,omitempty" xml:"subscriptionType,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This is an internal reserved parameter. Do not use it.
	//
	// example:
	//
	// true
	Unbound *bool `json:"unbound,omitempty" xml:"unbound,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsRequest) GetAgentUuid() *string {
	return s.AgentUuid
}

func (s *ListSubscriptionsRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListSubscriptionsRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ListSubscriptionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSubscriptionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSubscriptionsRequest) GetNotifyStrategyUuid() *string {
	return s.NotifyStrategyUuid
}

func (s *ListSubscriptionsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSubscriptionsRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListSubscriptionsRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *ListSubscriptionsRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *ListSubscriptionsRequest) GetUnbound() *bool {
	return s.Unbound
}

func (s *ListSubscriptionsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListSubscriptionsRequest) SetAgentUuid(v string) *ListSubscriptionsRequest {
	s.AgentUuid = &v
	return s
}

func (s *ListSubscriptionsRequest) SetDirection(v string) *ListSubscriptionsRequest {
	s.Direction = &v
	return s
}

func (s *ListSubscriptionsRequest) SetEnable(v bool) *ListSubscriptionsRequest {
	s.Enable = &v
	return s
}

func (s *ListSubscriptionsRequest) SetMaxResults(v int32) *ListSubscriptionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSubscriptionsRequest) SetNextToken(v string) *ListSubscriptionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSubscriptionsRequest) SetNotifyStrategyUuid(v string) *ListSubscriptionsRequest {
	s.NotifyStrategyUuid = &v
	return s
}

func (s *ListSubscriptionsRequest) SetOrderBy(v string) *ListSubscriptionsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSubscriptionsRequest) SetSubscriptionId(v string) *ListSubscriptionsRequest {
	s.SubscriptionId = &v
	return s
}

func (s *ListSubscriptionsRequest) SetSubscriptionName(v string) *ListSubscriptionsRequest {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionsRequest) SetSubscriptionType(v string) *ListSubscriptionsRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ListSubscriptionsRequest) SetUnbound(v bool) *ListSubscriptionsRequest {
	s.Unbound = &v
	return s
}

func (s *ListSubscriptionsRequest) SetWorkspace(v string) *ListSubscriptionsRequest {
	s.Workspace = &v
	return s
}

func (s *ListSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
