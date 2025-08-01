// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateFlowRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *CreateFlowRuleRequest
	GetAppId() *string
	SetAppName(v string) *CreateFlowRuleRequest
	GetAppName() *string
	SetControlBehavior(v int32) *CreateFlowRuleRequest
	GetControlBehavior() *int32
	SetEnable(v bool) *CreateFlowRuleRequest
	GetEnable() *bool
	SetLimitApp(v string) *CreateFlowRuleRequest
	GetLimitApp() *string
	SetMaxQueueingTimeMs(v int32) *CreateFlowRuleRequest
	GetMaxQueueingTimeMs() *int32
	SetNamespace(v string) *CreateFlowRuleRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateFlowRuleRequest
	GetRegionId() *string
	SetResource(v string) *CreateFlowRuleRequest
	GetResource() *string
	SetResourceType(v int32) *CreateFlowRuleRequest
	GetResourceType() *int32
	SetThreshold(v int32) *CreateFlowRuleRequest
	GetThreshold() *int32
}

type CreateFlowRuleRequest struct {
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
	// The application ID.
	//
	// example:
	//
	// ib09eblv6p@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The throttling effect.
	//
	// Valid values:
	//
	// 	- 0: fast failure
	//
	// 	- 2: in queue
	//
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// Specifies whether to enable the rule.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enable   *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// The timeout period. Unit: milliseconds. This value is required if the ControlBehavior parameter is set to 2.
	//
	// example:
	//
	// 10
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the API resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// /b
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceType *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The throttling threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateFlowRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateFlowRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateFlowRuleRequest) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *CreateFlowRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateFlowRuleRequest) GetLimitApp() *string {
	return s.LimitApp
}

func (s *CreateFlowRuleRequest) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *CreateFlowRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateFlowRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFlowRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateFlowRuleRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *CreateFlowRuleRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *CreateFlowRuleRequest) SetAcceptLanguage(v string) *CreateFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateFlowRuleRequest) SetAppId(v string) *CreateFlowRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateFlowRuleRequest) SetAppName(v string) *CreateFlowRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateFlowRuleRequest) SetControlBehavior(v int32) *CreateFlowRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *CreateFlowRuleRequest) SetEnable(v bool) *CreateFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateFlowRuleRequest) SetLimitApp(v string) *CreateFlowRuleRequest {
	s.LimitApp = &v
	return s
}

func (s *CreateFlowRuleRequest) SetMaxQueueingTimeMs(v int32) *CreateFlowRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateFlowRuleRequest) SetNamespace(v string) *CreateFlowRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateFlowRuleRequest) SetRegionId(v string) *CreateFlowRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowRuleRequest) SetResource(v string) *CreateFlowRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateFlowRuleRequest) SetResourceType(v int32) *CreateFlowRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateFlowRuleRequest) SetThreshold(v int32) *CreateFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *CreateFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}
