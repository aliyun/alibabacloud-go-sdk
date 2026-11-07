// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIsolationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateIsolationRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *CreateIsolationRuleRequest
	GetAppId() *string
	SetAppName(v string) *CreateIsolationRuleRequest
	GetAppName() *string
	SetEnable(v bool) *CreateIsolationRuleRequest
	GetEnable() *bool
	SetLimitApp(v string) *CreateIsolationRuleRequest
	GetLimitApp() *string
	SetNamespace(v string) *CreateIsolationRuleRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateIsolationRuleRequest
	GetRegionId() *string
	SetResource(v string) *CreateIsolationRuleRequest
	GetResource() *string
	SetThreshold(v float32) *CreateIsolationRuleRequest
	GetThreshold() *float32
}

type CreateIsolationRuleRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese. This is the default value.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The enabling status of the rule.
	//
	// This parameter is required.
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
	// The region where the cluster is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the interface to which the rule applies. The name must be the same as the interface name displayed in the interface details on the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The concurrency threshold of the isolation rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateIsolationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateIsolationRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateIsolationRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateIsolationRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateIsolationRuleRequest) GetLimitApp() *string {
	return s.LimitApp
}

func (s *CreateIsolationRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIsolationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIsolationRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateIsolationRuleRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateIsolationRuleRequest) SetAcceptLanguage(v string) *CreateIsolationRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetAppId(v string) *CreateIsolationRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetAppName(v string) *CreateIsolationRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetEnable(v bool) *CreateIsolationRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetLimitApp(v string) *CreateIsolationRuleRequest {
	s.LimitApp = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetNamespace(v string) *CreateIsolationRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetRegionId(v string) *CreateIsolationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetResource(v string) *CreateIsolationRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateIsolationRuleRequest) SetThreshold(v float32) *CreateIsolationRuleRequest {
	s.Threshold = &v
	return s
}

func (s *CreateIsolationRuleRequest) Validate() error {
	return dara.Validate(s)
}
