// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventCenterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventChannel(v string) *UpdateEventCenterRuleRequest
	GetEventChannel() *string
	SetEventConfig(v string) *UpdateEventCenterRuleRequest
	GetEventConfig() *string
	SetEventScope(v string) *UpdateEventCenterRuleRequest
	GetEventScope() *string
	SetEventType(v string) *UpdateEventCenterRuleRequest
	GetEventType() *string
	SetInstanceId(v string) *UpdateEventCenterRuleRequest
	GetInstanceId() *string
	SetNamespaces(v []*string) *UpdateEventCenterRuleRequest
	GetNamespaces() []*string
	SetRepoNames(v []*string) *UpdateEventCenterRuleRequest
	GetRepoNames() []*string
	SetRepoTagFilterPattern(v string) *UpdateEventCenterRuleRequest
	GetRepoTagFilterPattern() *string
	SetRuleId(v string) *UpdateEventCenterRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *UpdateEventCenterRuleRequest
	GetRuleName() *string
}

type UpdateEventCenterRuleRequest struct {
	// The event notification channel.
	//
	// example:
	//
	// EVENT_BRIDGE
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The event configuration.
	//
	// example:
	//
	// {
	//
	//         "notifyMethod":"http",
	//
	//         "notifyConfig":{
	//
	//             "Url":"http://www.aliyundoc.com",
	//
	//             "id":"MaAV3HgTkO5Fh8l1V********",
	//
	//         },
	//
	//         "notifyFilter":{}
	//
	//     }
	EventConfig *string `json:"EventConfig,omitempty" xml:"EventConfig,omitempty"`
	// The event scope. Valid values:
	//
	// 	- `INSTANCE`
	//
	// 	- `NAMESPACE`
	//
	// 	- `REPO`
	//
	// Default value: `INSTANCE`
	//
	// example:
	//
	// INSTANCE
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	//
	// 	- `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	//
	// 	- `cr:Artifact:BuildCompleted`: The image is built.
	//
	// 	- `cr:Artifact:ScanCompleted`: The image is scanned.
	//
	// 	- `cr:Artifact:SigningCompleted`: The image is signed.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespaces to which the event notification rule applies.
	//
	// example:
	//
	// ns
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The names of the repositories to which the event notification rule applies.
	//
	// example:
	//
	// reponame
	RepoNames []*string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty" type:"Repeated"`
	// The regular expression for image tags.
	//
	// example:
	//
	// .*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The ID of the event notification rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crecr-n6pbhgjxt*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// chain-demo
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateEventCenterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleRequest) GetEventChannel() *string {
	return s.EventChannel
}

func (s *UpdateEventCenterRuleRequest) GetEventConfig() *string {
	return s.EventConfig
}

func (s *UpdateEventCenterRuleRequest) GetEventScope() *string {
	return s.EventScope
}

func (s *UpdateEventCenterRuleRequest) GetEventType() *string {
	return s.EventType
}

func (s *UpdateEventCenterRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventCenterRuleRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *UpdateEventCenterRuleRequest) GetRepoNames() []*string {
	return s.RepoNames
}

func (s *UpdateEventCenterRuleRequest) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *UpdateEventCenterRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateEventCenterRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateEventCenterRuleRequest) SetEventChannel(v string) *UpdateEventCenterRuleRequest {
	s.EventChannel = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetEventConfig(v string) *UpdateEventCenterRuleRequest {
	s.EventConfig = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetEventScope(v string) *UpdateEventCenterRuleRequest {
	s.EventScope = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetEventType(v string) *UpdateEventCenterRuleRequest {
	s.EventType = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetInstanceId(v string) *UpdateEventCenterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetNamespaces(v []*string) *UpdateEventCenterRuleRequest {
	s.Namespaces = v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRepoNames(v []*string) *UpdateEventCenterRuleRequest {
	s.RepoNames = v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRepoTagFilterPattern(v string) *UpdateEventCenterRuleRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRuleId(v string) *UpdateEventCenterRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRuleName(v string) *UpdateEventCenterRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) Validate() error {
	return dara.Validate(s)
}
