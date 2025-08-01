// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTagPoliciesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ApplyTagPoliciesShrinkRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ApplyTagPoliciesShrinkRequest
	GetAppId() *string
	SetAppName(v string) *ApplyTagPoliciesShrinkRequest
	GetAppName() *string
	SetEnable(v bool) *ApplyTagPoliciesShrinkRequest
	GetEnable() *bool
	SetNamespace(v string) *ApplyTagPoliciesShrinkRequest
	GetNamespace() *string
	SetNamespaceId(v string) *ApplyTagPoliciesShrinkRequest
	GetNamespaceId() *string
	SetRegion(v string) *ApplyTagPoliciesShrinkRequest
	GetRegion() *string
	SetRulesShrink(v string) *ApplyTagPoliciesShrinkRequest
	GetRulesShrink() *string
}

type ApplyTagPoliciesShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Deprecated
	//
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// h5pohqu3gd@xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to enable the routing rule.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The Microservices Engine (MSE) namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Optional. The ID of the namespace.
	//
	// example:
	//
	// 12233****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The details of the routing rule.
	//
	// example:
	//
	// { "blue": { # Tag key. "rate": 20, # Rate. A value of 20 indicates that 20% of the traffic is routed to the node identified by the tag. "name": "luoye-route-test", # Routing rule name. "carryData": false, # Specifies whether to enable pass-through. This parameter is applicable to end-to-end canary release. "rules": { # Rule information. "dubbo": [{ # Dubbo rule. The system also supports Spring Cloud and Istio rule. "serviceName": "com.taobao.hsf.common.DemoService", # Service name (operation name). "group": "", # Group name. "version": "", # Service version. "methodName": "sayHello", # Method name. "condition": "AND", # Logical operator used by the following items. Valid values: AND and OR. "argumentItems": [{ # Array of rule details. "index": 0, # Index field. "expr": "", # Expression. Its details is described below. "operator": "rawvalue", # The comparison mode. A value of rawvalue indicates direct comparison. A value of mode indicates the modulo operation. A value of list indicates using a whitelist. "value": "jim", # Base value. The value obtained by the expression will be compared with this value. If operator is set to list, separate the items specified for this parameter with commas (,). Example: 1,2,3. "cond": "==" # Comparison operator. Valid values: >=, <=, >, <, and ==. }] }] } }, "_base": { # Another tag key. "rate": 80 # Rate. A value of 20 indicates that 20% of the traffic is routed to the node identified by the tag. } }
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ApplyTagPoliciesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyTagPoliciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyTagPoliciesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ApplyTagPoliciesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ApplyTagPoliciesShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *ApplyTagPoliciesShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ApplyTagPoliciesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ApplyTagPoliciesShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ApplyTagPoliciesShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ApplyTagPoliciesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *ApplyTagPoliciesShrinkRequest) SetAcceptLanguage(v string) *ApplyTagPoliciesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetAppId(v string) *ApplyTagPoliciesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetAppName(v string) *ApplyTagPoliciesShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetEnable(v bool) *ApplyTagPoliciesShrinkRequest {
	s.Enable = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetNamespace(v string) *ApplyTagPoliciesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetNamespaceId(v string) *ApplyTagPoliciesShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetRegion(v string) *ApplyTagPoliciesShrinkRequest {
	s.Region = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) SetRulesShrink(v string) *ApplyTagPoliciesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *ApplyTagPoliciesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
