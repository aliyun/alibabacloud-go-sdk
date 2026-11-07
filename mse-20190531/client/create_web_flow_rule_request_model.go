// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateWebFlowRuleRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *CreateWebFlowRuleRequest
	GetAppId() *string
	SetAppName(v string) *CreateWebFlowRuleRequest
	GetAppName() *string
	SetBurst(v int32) *CreateWebFlowRuleRequest
	GetBurst() *int32
	SetControlBehavior(v int32) *CreateWebFlowRuleRequest
	GetControlBehavior() *int32
	SetEnable(v bool) *CreateWebFlowRuleRequest
	GetEnable() *bool
	SetMaxQueueingTimeMs(v int32) *CreateWebFlowRuleRequest
	GetMaxQueueingTimeMs() *int32
	SetMetricType(v int32) *CreateWebFlowRuleRequest
	GetMetricType() *int32
	SetNamespace(v string) *CreateWebFlowRuleRequest
	GetNamespace() *string
	SetParamItem(v string) *CreateWebFlowRuleRequest
	GetParamItem() *string
	SetRegionId(v string) *CreateWebFlowRuleRequest
	GetRegionId() *string
	SetResource(v string) *CreateWebFlowRuleRequest
	GetResource() *string
	SetResourceMode(v int32) *CreateWebFlowRuleRequest
	GetResourceMode() *int32
	SetResourceType(v int32) *CreateWebFlowRuleRequest
	GetResourceType() *int32
	SetStatIntervalMs(v int32) *CreateWebFlowRuleRequest
	GetStatIntervalMs() *int32
	SetThreshold(v float32) *CreateWebFlowRuleRequest
	GetThreshold() *float32
}

type CreateWebFlowRuleRequest struct {
	// The language in which you want the results to be displayed. Valid values: zh (default): Chinese. en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID
	//
	// example:
	//
	// hkhon1****@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of additional requests allowed for burst requests. The effect is equivalent to the maximum number of tokens of the token bucket algorithm.
	//
	// example:
	//
	// 0
	Burst *int32 `json:"Burst,omitempty" xml:"Burst,omitempty"`
	// The flow control method. Valid values:
	//
	// - 0: fast failure
	//
	// - 2: wait in queue
	//
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// The rule status. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The timeout period. Unit: milliseconds. This parameter is required if the ControlBehavior parameter is set to 2.
	//
	// example:
	//
	// 200
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The threshold type. Valid values:
	//
	// - 1 (number of requests)
	//
	// example:
	//
	// 1
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The namespace
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The parameter description, which is a JSON string containing the following key-value pairs:
	//
	// - parseStrategy: the parameter property (0: client IP address, 1: remote host, 2: header, 3: URL parameter)
	//
	// - fieldName: the header name or URL parameter name. You must specify this key only if parseStrategy is set to 2 or 3
	//
	// - matchStrategy: the method to match the property value. You must specify this key only if parseStrategy is set to 2 or 3 and the property value is matched. Valid values: 0 (exact match), 2 (regular expression match), and 3 (substring match)
	//
	// - pattern: the match string. You must specify this key only if parseStrategy is set to 2 or 3 and the property value is matched
	//
	// example:
	//
	// {"fieldName":"testKey","matchStrategy":2,"parseStrategy":2,"pattern":"testValue"}
	ParamItem *string `json:"ParamItem,omitempty" xml:"ParamItem,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The interface name: The application resources to which the rule applies. The interface name must be the same as the name on the interface details page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// /flow
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Resource mode
	//
	// - 0: resource name mode
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ResourceMode *int32 `json:"ResourceMode,omitempty" xml:"ResourceMode,omitempty"`
	// The resource type. Valid values:
	//
	// - 0: custom interface
	//
	// - 1: WEB
	//
	// - 2: RPC
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The length of the time window. Unit: milliseconds. The default value is 1000, which indicates 1 second.
	//
	// example:
	//
	// 1000
	StatIntervalMs *int32 `json:"StatIntervalMs,omitempty" xml:"StatIntervalMs,omitempty"`
	// Threshold
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateWebFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWebFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateWebFlowRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateWebFlowRuleRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateWebFlowRuleRequest) GetBurst() *int32 {
	return s.Burst
}

func (s *CreateWebFlowRuleRequest) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *CreateWebFlowRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateWebFlowRuleRequest) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *CreateWebFlowRuleRequest) GetMetricType() *int32 {
	return s.MetricType
}

func (s *CreateWebFlowRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateWebFlowRuleRequest) GetParamItem() *string {
	return s.ParamItem
}

func (s *CreateWebFlowRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWebFlowRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateWebFlowRuleRequest) GetResourceMode() *int32 {
	return s.ResourceMode
}

func (s *CreateWebFlowRuleRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *CreateWebFlowRuleRequest) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *CreateWebFlowRuleRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateWebFlowRuleRequest) SetAcceptLanguage(v string) *CreateWebFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetAppId(v string) *CreateWebFlowRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetAppName(v string) *CreateWebFlowRuleRequest {
	s.AppName = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetBurst(v int32) *CreateWebFlowRuleRequest {
	s.Burst = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetControlBehavior(v int32) *CreateWebFlowRuleRequest {
	s.ControlBehavior = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetEnable(v bool) *CreateWebFlowRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetMaxQueueingTimeMs(v int32) *CreateWebFlowRuleRequest {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetMetricType(v int32) *CreateWebFlowRuleRequest {
	s.MetricType = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetNamespace(v string) *CreateWebFlowRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetParamItem(v string) *CreateWebFlowRuleRequest {
	s.ParamItem = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetRegionId(v string) *CreateWebFlowRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetResource(v string) *CreateWebFlowRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetResourceMode(v int32) *CreateWebFlowRuleRequest {
	s.ResourceMode = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetResourceType(v int32) *CreateWebFlowRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetStatIntervalMs(v int32) *CreateWebFlowRuleRequest {
	s.StatIntervalMs = &v
	return s
}

func (s *CreateWebFlowRuleRequest) SetThreshold(v float32) *CreateWebFlowRuleRequest {
	s.Threshold = &v
	return s
}

func (s *CreateWebFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}
