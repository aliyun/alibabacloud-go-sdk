// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardingRules(v []*ListForwardingRulesResponseBodyForwardingRules) *ListForwardingRulesResponseBody
	GetForwardingRules() []*ListForwardingRulesResponseBodyForwardingRules
	SetMaxResults(v int32) *ListForwardingRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListForwardingRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListForwardingRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListForwardingRulesResponseBody
	GetTotalCount() *int32
}

type ListForwardingRulesResponseBody struct {
	// The forwarding rules.
	ForwardingRules []*ListForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If **NextToken*	- is not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CFC67ED9-4AB1-431F-B6E3-A752B7B8CCD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListForwardingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBody) GetForwardingRules() []*ListForwardingRulesResponseBodyForwardingRules {
	return s.ForwardingRules
}

func (s *ListForwardingRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListForwardingRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListForwardingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListForwardingRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListForwardingRulesResponseBody) SetForwardingRules(v []*ListForwardingRulesResponseBodyForwardingRules) *ListForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *ListForwardingRulesResponseBody) SetMaxResults(v int32) *ListForwardingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetNextToken(v string) *ListForwardingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetRequestId(v string) *ListForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetTotalCount(v int32) *ListForwardingRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListForwardingRulesResponseBody) Validate() error {
	if s.ForwardingRules != nil {
		for _, item := range s.ForwardingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListForwardingRulesResponseBodyForwardingRules struct {
	// The direction in which the forwarding rule takes effect.
	//
	// By default, **request*	- is returned, which indicates that the forwarding rule takes effect on requests.
	//
	// example:
	//
	// request
	ForwardingRuleDirection *string `json:"ForwardingRuleDirection,omitempty" xml:"ForwardingRuleDirection,omitempty"`
	// The forwarding rule ID.
	//
	// example:
	//
	// frule-bp19a3t3yzr21q3****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	// The forwarding rule name.
	//
	// example:
	//
	// auto_named_rule
	ForwardingRuleName *string `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	// The state of the forwarding rule. Valid values:
	//
	// 	- **active:*	- The forwarding rule is normal.
	//
	// 	- **configuring:*	- The forwarding rule is being modified.
	//
	// 	- **deleting:*	- The forwarding rule is being deleted.
	//
	// example:
	//
	// active
	ForwardingRuleStatus *string `json:"ForwardingRuleStatus,omitempty" xml:"ForwardingRuleStatus,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The priority of the forwarding rule.
	//
	// A value between **1*	- and **10000*	- is returned. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1000
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The forwarding actions.
	RuleActions []*ListForwardingRulesResponseBodyForwardingRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The conditions that trigger the forwarding rule.
	RuleConditions []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the GA instance is managed. Valid values:
	//
	// 	- **true**: The GA instance is managed.
	//
	// 	- **false**: The GA instance is not managed.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that you can perform on the managed instance.
	//
	// >  This parameter takes effect only if **ServiceManaged*	- is set to **True**.
	//
	// 	- You can perform only specific actions on the managed instance.
	ServiceManagedInfos []*ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetForwardingRuleDirection() *string {
	return s.ForwardingRuleDirection
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetForwardingRuleId() *string {
	return s.ForwardingRuleId
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetForwardingRuleName() *string {
	return s.ForwardingRuleName
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetForwardingRuleStatus() *string {
	return s.ForwardingRuleStatus
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetRuleActions() []*ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	return s.RuleActions
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetRuleConditions() []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	return s.RuleConditions
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListForwardingRulesResponseBodyForwardingRules) GetServiceManagedInfos() []*ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleDirection(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleDirection = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleName(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleStatus(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleStatus = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetListenerId(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ListenerId = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetPriority(v int32) *ListForwardingRulesResponseBodyForwardingRules {
	s.Priority = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetRuleActions(v []*ListForwardingRulesResponseBodyForwardingRulesRuleActions) *ListForwardingRulesResponseBodyForwardingRules {
	s.RuleActions = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetRuleConditions(v []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions) *ListForwardingRulesResponseBodyForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetServiceId(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ServiceId = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetServiceManaged(v bool) *ListForwardingRulesResponseBodyForwardingRules {
	s.ServiceManaged = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetServiceManagedInfos(v []*ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) *ListForwardingRulesResponseBodyForwardingRules {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) Validate() error {
	if s.RuleActions != nil {
		for _, item := range s.RuleActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleConditions != nil {
		for _, item := range s.RuleConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceManagedInfos != nil {
		for _, item := range s.ServiceManagedInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActions struct {
	// The forwarding action configuration.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	ForwardGroupConfig *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The forwarding priority.
	//
	// >  This parameter does not take effect.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The type of the forwarding action. Valid values:
	//
	// 	- **ForwardGroup**: forwards a request.
	//
	// 	- **Redirect**: redirects a request.
	//
	// 	- **FixResponse**: returns a fixed response.
	//
	// 	- **Rewrite**: rewrites a request.
	//
	// 	- **AddHeader**: adds a header to a request.
	//
	// 	- **RemoveHeaderConfig**: deletes the header from a request.
	//
	// 	- **Drop**: drops a request.
	//
	// example:
	//
	// ForwardGroup
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The value of the forwarding action.
	//
	// Different JSON strings are returned based on the value of **RuleActionType**.
	//
	// 	- If you set **RuleActionType*	- to **ForwardGroup**, the information about a virtual endpoint group is returned. The following section describes the parameters:
	//
	//     	- `type`: `endpointgroup` is returned.
	//
	//     	- `value`: the ID of the virtual endpoint group.
	//
	// 	- If you set **RuleActionType*	- to **Redirect**, the redirecting configuration is returned. The following section describes the parameters:
	//
	//     	- `protocol`: the protocol of requests after the requests are redirected.
	//
	//     	- `domain`: the domain name to which requests are redirected.
	//
	//     	- `port`: the port to which requests are redirected.
	//
	//     	- `path`: the path to which requests are redirected.
	//
	//     	- `query`: the query string of the requests that are redirected.
	//
	//     	- `code`: the redirecting code.
	//
	// 	- If you set **RuleActionType*	- to **FixResponse**, the information about the fixed response that you configured is returned. The following section describes the parameters:
	//
	//     	- `code`: the HTTP status code.
	//
	//     	- `type`: the content type of the response.
	//
	//     	- `content`: the content of the response.
	//
	// 	- If **RuleActionType*	- is set to **AddHeader**, the information about the HTTP header that is added is returned. The following section describes the parameters:
	//
	//     	- `name`: the name of the HTTP header.
	//
	//     	- `type`: the content type of the HTTP header.
	//
	//     	- `value`: the content of the HTTP header.
	//
	// 	- If you set **RuleActionType*	- to **RemoveHeader**, the information about the HTTP header that is deleted is returned.
	//
	// 	- If you set **RuleActionType*	- to **Rewrite**, the rewriting configuration is returned. The following section describes the parameters:
	//
	//     	- `domain`: the domain name to which requests are redirected.
	//
	//     	- `path`: the path to which requests are redirected.
	//
	//     	- `query`: the query string of the requests that are redirected.
	//
	// 	- If you set **RuleActionType*	- to **Drop**, an empty string is returned.
	//
	// example:
	//
	// [{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}]
	RuleActionValue *string `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) GetForwardGroupConfig() *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) GetRuleActionValue() *string {
	return s.RuleActionValue
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetForwardGroupConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetOrder(v int32) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetRuleActionType(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetRuleActionValue(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.RuleActionValue = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) Validate() error {
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig struct {
	// The information about the endpoint groups.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	ServerGroupTuples []*ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) Validate() error {
	if s.ServerGroupTuples != nil {
		for _, item := range s.ServerGroupTuples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The endpoint group ID.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	//
	// example:
	//
	// epg-bp1enpdcrqhl78g6r****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type ListForwardingRulesResponseBodyForwardingRulesRuleConditions struct {
	// The domain name configuration.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	HostConfig *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The path configuration.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	PathConfig *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The type of the forwarding condition. Valid values:
	//
	// 	- **Host:*	- domain name.
	//
	// 	- **Path:*	- path.
	//
	// 	- **RequestHeader:*	- HTTP header.
	//
	// 	- **Query:*	- query string.
	//
	// 	- **Method:*	- HTTP method.
	//
	// 	- **Cookie:*	- cookie.
	//
	// 	- **SourceIP:*	- source IP address.
	//
	// example:
	//
	// Host
	RuleConditionType *string `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	// The value of the forwarding condition type.
	//
	// Different JSON strings are returned based on the value of the **RuleConditionType*	- parameter.
	//
	// 	- If you set **RuleConditionType*	- to **Host**, a domain name condition is returned. If multiple domain names are returned in a forwarding condition, the relationship between the domain names is OR.
	//
	// 	- If you set **RuleConditionType*	- to **Path**, a path condition is returned. If multiple forwarding conditions of the path type are returned in a forwarding rule, the relationship between the forwarding conditions is OR. If multiple paths are returned in a forwarding condition, the relationship between the paths is OR.
	//
	// 	- If you set **RuleConditionType*	- to **RequestHeader**, an HTTP header condition that consists of key-value pairs is returned.
	//
	// 	- If you set **RuleConditionType*	- to **Query**, a query string condition that consists of key-value pairs is returned.
	//
	// 	- If you set **RuleConditionType*	- to **Method**, an HTTP method condition is returned.
	//
	// 	- If you set **RuleConditionType*	- to **Cookie**, a cookie condition that consists of key-value pairs is returned.
	//
	// 	- If you set **RuleConditionType*	- to **SourceIP**, a source IP address condition is returned. If multiple source IP addresses are returned in a forwarding condition, the relationship between the source IP addresses is OR.
	//
	// example:
	//
	// ["www.example.com", "www.aliyun.com"]
	RuleConditionValue *string `json:"RuleConditionValue,omitempty" xml:"RuleConditionValue,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GetHostConfig() *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GetPathConfig() *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GetRuleConditionType() *string {
	return s.RuleConditionType
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GetRuleConditionValue() *string {
	return s.RuleConditionValue
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetHostConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetPathConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetRuleConditionType(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetRuleConditionValue(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) Validate() error {
	if s.HostConfig != nil {
		if err := s.HostConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PathConfig != nil {
		if err := s.PathConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig struct {
	// The domain name configuration.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig struct {
	// The path configuration.
	//
	// >  GA instances created after July 12, 2022 support all forwarding condition types and action types. We recommend that you query forwarding conditions and actions by calling the **RuleActionType*	- and **RuleActionValue*	- operations.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}

type ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos struct {
	// The name of the action that you can perform on the managed instance. Valid values:
	//
	// 	- **Create**: Create an instance.
	//
	// 	- **Update**: Update the current instance.
	//
	// 	- **Delete**: Delete the current instance.
	//
	// 	- **Associate**: Reference the current instance.
	//
	// 	- **UserUnmanaged**: Unmanage the instance.
	//
	// 	- **CreateChild**: Create a child resource on the current instance.
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// 	- **Listener**: listener.
	//
	// 	- **IpSet**: acceleration region.
	//
	// 	- **EndpointGroup**: endpoint group.
	//
	// 	- **ForwardingRule**: forwarding rule.
	//
	// 	- **Endpoint**: endpoint.
	//
	// 	- **EndpointGroupDestination**: the protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: the traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter is returned only if the value of **Action*	- is **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) SetAction(v string) *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) SetChildType(v string) *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) SetIsManaged(v bool) *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
