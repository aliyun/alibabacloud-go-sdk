// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWebFlowRuleResponseBody
	GetCode() *string
	SetData(v *CreateWebFlowRuleResponseBodyData) *CreateWebFlowRuleResponseBody
	GetData() *CreateWebFlowRuleResponseBodyData
	SetMessage(v string) *CreateWebFlowRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWebFlowRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWebFlowRuleResponseBody
	GetSuccess() *bool
}

type CreateWebFlowRuleResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *CreateWebFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A73AC37C-C617-4E3A-8049-372CF49C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWebFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebFlowRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWebFlowRuleResponseBody) GetData() *CreateWebFlowRuleResponseBodyData {
	return s.Data
}

func (s *CreateWebFlowRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWebFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebFlowRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWebFlowRuleResponseBody) SetCode(v string) *CreateWebFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWebFlowRuleResponseBody) SetData(v *CreateWebFlowRuleResponseBodyData) *CreateWebFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateWebFlowRuleResponseBody) SetMessage(v string) *CreateWebFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWebFlowRuleResponseBody) SetRequestId(v string) *CreateWebFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebFlowRuleResponseBody) SetSuccess(v bool) *CreateWebFlowRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWebFlowRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWebFlowRuleResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// hkhon1****@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
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
	// The rule ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timeout period for queuing when the ControlBehavior parameter is set to 2. Unit: milliseconds.
	//
	// example:
	//
	// 200
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The threshold type. Valid values:
	//
	// - 1: number of requests
	//
	// example:
	//
	// 1
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The namespace.
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
	// Resource mode
	//
	// 0: resource name mode
	//
	// example:
	//
	// 0
	ReourceMode *int32 `json:"ReourceMode,omitempty" xml:"ReourceMode,omitempty"`
	// Interface name: the application resources to which the rule applies. The interface name must be the same as that displayed on the API details page in the console.
	//
	// example:
	//
	// /flow
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The length of the time window. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	StatIntervalMs *int32 `json:"StatIntervalMs,omitempty" xml:"StatIntervalMs,omitempty"`
	// The threshold.
	//
	// example:
	//
	// 20
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateWebFlowRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWebFlowRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateWebFlowRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateWebFlowRuleResponseBodyData) GetBurst() *int32 {
	return s.Burst
}

func (s *CreateWebFlowRuleResponseBodyData) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *CreateWebFlowRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *CreateWebFlowRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateWebFlowRuleResponseBodyData) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *CreateWebFlowRuleResponseBodyData) GetMetricType() *int32 {
	return s.MetricType
}

func (s *CreateWebFlowRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateWebFlowRuleResponseBodyData) GetParamItem() *string {
	return s.ParamItem
}

func (s *CreateWebFlowRuleResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWebFlowRuleResponseBodyData) GetReourceMode() *int32 {
	return s.ReourceMode
}

func (s *CreateWebFlowRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *CreateWebFlowRuleResponseBodyData) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *CreateWebFlowRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateWebFlowRuleResponseBodyData) SetAppId(v string) *CreateWebFlowRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetAppName(v string) *CreateWebFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetBurst(v int32) *CreateWebFlowRuleResponseBodyData {
	s.Burst = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetControlBehavior(v int32) *CreateWebFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetEnable(v bool) *CreateWebFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetId(v int64) *CreateWebFlowRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *CreateWebFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetMetricType(v int32) *CreateWebFlowRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetNamespace(v string) *CreateWebFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetParamItem(v string) *CreateWebFlowRuleResponseBodyData {
	s.ParamItem = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetRegionId(v string) *CreateWebFlowRuleResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetReourceMode(v int32) *CreateWebFlowRuleResponseBodyData {
	s.ReourceMode = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetResource(v string) *CreateWebFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetStatIntervalMs(v int32) *CreateWebFlowRuleResponseBodyData {
	s.StatIntervalMs = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) SetThreshold(v float32) *CreateWebFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *CreateWebFlowRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
