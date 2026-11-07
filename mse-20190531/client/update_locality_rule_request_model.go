// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocalityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateLocalityRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateLocalityRuleRequest
	GetAppId() *string
	SetAppName(v string) *UpdateLocalityRuleRequest
	GetAppName() *string
	SetEnable(v bool) *UpdateLocalityRuleRequest
	GetEnable() *bool
	SetNamespace(v string) *UpdateLocalityRuleRequest
	GetNamespace() *string
	SetRegion(v string) *UpdateLocalityRuleRequest
	GetRegion() *string
	SetRules(v string) *UpdateLocalityRuleRequest
	GetRules() *string
	SetSource(v string) *UpdateLocalityRuleRequest
	GetSource() *string
	SetThreshold(v float64) *UpdateLocalityRuleRequest
	GetThreshold() *float64
}

type UpdateLocalityRuleRequest struct {
	// The language of the response. Valid values: zh (default): Chinese. en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to enable zone affinity.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// myNamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The reserved parameter. This parameter is not used.
	//
	// example:
	//
	// No input required
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The source of the rule.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The threshold for zone affinity to take effect. The zone affinity routing feature is enabled for a zone only when the proportion of instances of the current application deployed in the zone exceeds this threshold. That is, when a consumer sends a request to the application, the request is preferentially routed to an instance in the same zone as the consumer.
	//
	// example:
	//
	// 0.2
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLocalityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalityRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocalityRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateLocalityRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateLocalityRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLocalityRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateLocalityRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateLocalityRuleRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateLocalityRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *UpdateLocalityRuleRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateLocalityRuleRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *UpdateLocalityRuleRequest) SetAcceptLanguage(v string) *UpdateLocalityRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetAppId(v string) *UpdateLocalityRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetAppName(v string) *UpdateLocalityRuleRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetEnable(v bool) *UpdateLocalityRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetNamespace(v string) *UpdateLocalityRuleRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetRegion(v string) *UpdateLocalityRuleRequest {
	s.Region = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetRules(v string) *UpdateLocalityRuleRequest {
	s.Rules = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetSource(v string) *UpdateLocalityRuleRequest {
	s.Source = &v
	return s
}

func (s *UpdateLocalityRuleRequest) SetThreshold(v float64) *UpdateLocalityRuleRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateLocalityRuleRequest) Validate() error {
	return dara.Validate(s)
}
