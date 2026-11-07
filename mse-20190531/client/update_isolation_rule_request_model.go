// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIsolationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateIsolationRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateIsolationRuleRequest
	GetAppId() *string
	SetAppName(v string) *UpdateIsolationRuleRequest
	GetAppName() *string
	SetEnable(v bool) *UpdateIsolationRuleRequest
	GetEnable() *bool
	SetLimitApp(v string) *UpdateIsolationRuleRequest
	GetLimitApp() *string
	SetNamespace(v string) *UpdateIsolationRuleRequest
	GetNamespace() *string
	SetRuleId(v int64) *UpdateIsolationRuleRequest
	GetRuleId() *int64
	SetThreshold(v float32) *UpdateIsolationRuleRequest
	GetThreshold() *float32
}

type UpdateIsolationRuleRequest struct {
	// The language in which the response is displayed. Valid values: zh (default): Chinese. en: English.
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
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The source application of the request.
	//
	// example:
	//
	// spring-cloud-a
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// The microservice namespace to which the application belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The concurrency threshold of the isolation rule.
	//
	// example:
	//
	// 3
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateIsolationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateIsolationRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateIsolationRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateIsolationRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateIsolationRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateIsolationRuleRequest) GetLimitApp() *string {
	return s.LimitApp
}

func (s *UpdateIsolationRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateIsolationRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *UpdateIsolationRuleRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateIsolationRuleRequest) SetAcceptLanguage(v string) *UpdateIsolationRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetAppId(v string) *UpdateIsolationRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetAppName(v string) *UpdateIsolationRuleRequest {
	s.AppName = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetEnable(v bool) *UpdateIsolationRuleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetLimitApp(v string) *UpdateIsolationRuleRequest {
	s.LimitApp = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetNamespace(v string) *UpdateIsolationRuleRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetRuleId(v int64) *UpdateIsolationRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateIsolationRuleRequest) SetThreshold(v float32) *UpdateIsolationRuleRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateIsolationRuleRequest) Validate() error {
	return dara.Validate(s)
}
