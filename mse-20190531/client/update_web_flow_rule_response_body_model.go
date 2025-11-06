// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWebFlowRuleResponseBody
	GetCode() *string
	SetData(v *UpdateWebFlowRuleResponseBodyData) *UpdateWebFlowRuleResponseBody
	GetData() *UpdateWebFlowRuleResponseBodyData
	SetMessage(v string) *UpdateWebFlowRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWebFlowRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWebFlowRuleResponseBody
	GetSuccess() *bool
}

type UpdateWebFlowRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateWebFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 54973C90-F379-4372-9AA5-053A3F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWebFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebFlowRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWebFlowRuleResponseBody) GetData() *UpdateWebFlowRuleResponseBodyData {
	return s.Data
}

func (s *UpdateWebFlowRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWebFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWebFlowRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWebFlowRuleResponseBody) SetCode(v string) *UpdateWebFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBody) SetData(v *UpdateWebFlowRuleResponseBodyData) *UpdateWebFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWebFlowRuleResponseBody) SetMessage(v string) *UpdateWebFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBody) SetRequestId(v string) *UpdateWebFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBody) SetSuccess(v bool) *UpdateWebFlowRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWebFlowRuleResponseBodyData struct {
	// example:
	//
	// hkhon1XXXX@54e1f42f37cXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 0
	Burst *int32 `json:"Burst,omitempty" xml:"Burst,omitempty"`
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 200
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// example:
	//
	// 1
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// {"fieldName":"testKey","matchStrategy":2,"parseStrategy":2,"pattern":"testValue"}
	ParamItem *string `json:"ParamItem,omitempty" xml:"ParamItem,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	ReourceMode *int32 `json:"ReourceMode,omitempty" xml:"ReourceMode,omitempty"`
	// example:
	//
	// /flow
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// 1000
	StatIntervalMs *int32 `json:"StatIntervalMs,omitempty" xml:"StatIntervalMs,omitempty"`
	// example:
	//
	// 20
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateWebFlowRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWebFlowRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateWebFlowRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWebFlowRuleResponseBodyData) GetBurst() *int32 {
	return s.Burst
}

func (s *UpdateWebFlowRuleResponseBodyData) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *UpdateWebFlowRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateWebFlowRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateWebFlowRuleResponseBodyData) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *UpdateWebFlowRuleResponseBodyData) GetMetricType() *int32 {
	return s.MetricType
}

func (s *UpdateWebFlowRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateWebFlowRuleResponseBodyData) GetParamItem() *string {
	return s.ParamItem
}

func (s *UpdateWebFlowRuleResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWebFlowRuleResponseBodyData) GetReourceMode() *int32 {
	return s.ReourceMode
}

func (s *UpdateWebFlowRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *UpdateWebFlowRuleResponseBodyData) GetStatIntervalMs() *int32 {
	return s.StatIntervalMs
}

func (s *UpdateWebFlowRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateWebFlowRuleResponseBodyData) SetAppId(v string) *UpdateWebFlowRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetAppName(v string) *UpdateWebFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetBurst(v int32) *UpdateWebFlowRuleResponseBodyData {
	s.Burst = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetControlBehavior(v int32) *UpdateWebFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetEnable(v bool) *UpdateWebFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetId(v int64) *UpdateWebFlowRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *UpdateWebFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetMetricType(v int32) *UpdateWebFlowRuleResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetNamespace(v string) *UpdateWebFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetParamItem(v string) *UpdateWebFlowRuleResponseBodyData {
	s.ParamItem = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetRegionId(v string) *UpdateWebFlowRuleResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetReourceMode(v int32) *UpdateWebFlowRuleResponseBodyData {
	s.ReourceMode = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetResource(v string) *UpdateWebFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetStatIntervalMs(v int32) *UpdateWebFlowRuleResponseBodyData {
	s.StatIntervalMs = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) SetThreshold(v float32) *UpdateWebFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *UpdateWebFlowRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
