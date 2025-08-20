// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventCenterRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventChannel(v string) *UpdateEventCenterRuleShrinkRequest
	GetEventChannel() *string
	SetEventConfig(v string) *UpdateEventCenterRuleShrinkRequest
	GetEventConfig() *string
	SetEventScope(v string) *UpdateEventCenterRuleShrinkRequest
	GetEventScope() *string
	SetEventType(v string) *UpdateEventCenterRuleShrinkRequest
	GetEventType() *string
	SetInstanceId(v string) *UpdateEventCenterRuleShrinkRequest
	GetInstanceId() *string
	SetNamespacesShrink(v string) *UpdateEventCenterRuleShrinkRequest
	GetNamespacesShrink() *string
	SetRepoNamesShrink(v string) *UpdateEventCenterRuleShrinkRequest
	GetRepoNamesShrink() *string
	SetRepoTagFilterPattern(v string) *UpdateEventCenterRuleShrinkRequest
	GetRepoTagFilterPattern() *string
	SetRuleId(v string) *UpdateEventCenterRuleShrinkRequest
	GetRuleId() *string
	SetRuleName(v string) *UpdateEventCenterRuleShrinkRequest
	GetRuleName() *string
}

type UpdateEventCenterRuleShrinkRequest struct {
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
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The names of the repositories to which the event notification rule applies.
	//
	// example:
	//
	// reponame
	RepoNamesShrink *string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty"`
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

func (s UpdateEventCenterRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventCenterRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleShrinkRequest) GetEventChannel() *string {
	return s.EventChannel
}

func (s *UpdateEventCenterRuleShrinkRequest) GetEventConfig() *string {
	return s.EventConfig
}

func (s *UpdateEventCenterRuleShrinkRequest) GetEventScope() *string {
	return s.EventScope
}

func (s *UpdateEventCenterRuleShrinkRequest) GetEventType() *string {
	return s.EventType
}

func (s *UpdateEventCenterRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventCenterRuleShrinkRequest) GetNamespacesShrink() *string {
	return s.NamespacesShrink
}

func (s *UpdateEventCenterRuleShrinkRequest) GetRepoNamesShrink() *string {
	return s.RepoNamesShrink
}

func (s *UpdateEventCenterRuleShrinkRequest) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *UpdateEventCenterRuleShrinkRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateEventCenterRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventChannel(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventChannel = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventConfig(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventConfig = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventScope(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventScope = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventType(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventType = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetInstanceId(v string) *UpdateEventCenterRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetNamespacesShrink(v string) *UpdateEventCenterRuleShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRepoNamesShrink(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RepoNamesShrink = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRepoTagFilterPattern(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRuleId(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRuleName(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
