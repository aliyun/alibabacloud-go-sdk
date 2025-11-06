// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFlowRuleResponseBody
	GetCode() *string
	SetData(v *UpdateFlowRuleResponseBodyData) *UpdateFlowRuleResponseBody
	GetData() *UpdateFlowRuleResponseBodyData
	SetMessage(v string) *UpdateFlowRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFlowRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFlowRuleResponseBody
	GetSuccess() *bool
}

type UpdateFlowRuleResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	Data *UpdateFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 52BA6DA6-A702-4362-A32F-DFF79655****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The request was successful.
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The request failed.
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFlowRuleResponseBody) GetData() *UpdateFlowRuleResponseBodyData {
	return s.Data
}

func (s *UpdateFlowRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFlowRuleResponseBody) SetCode(v string) *UpdateFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFlowRuleResponseBody) SetData(v *UpdateFlowRuleResponseBodyData) *UpdateFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFlowRuleResponseBody) SetMessage(v string) *UpdateFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFlowRuleResponseBody) SetRequestId(v string) *UpdateFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowRuleResponseBody) SetSuccess(v bool) *UpdateFlowRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFlowRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFlowRuleResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// hpn9ac29kz@e31a4b871******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The throttling effect.
	//
	// Valid values:
	//
	// 	- 0
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     fast failure
	//
	//     <!-- -->
	//
	// 	- 2
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     in queue
	//
	//     <!-- -->
	//
	// example:
	//
	// 0
	ControlBehavior *int32 `json:"ControlBehavior,omitempty" xml:"ControlBehavior,omitempty"`
	// Indicates whether the rule was enabled.
	//
	// Valid value:
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
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 12
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// The timeout period for queuing when the value of ControlBehavior is set to 2. Unit: milliseconds.
	//
	// example:
	//
	// 500
	MaxQueueingTimeMs *int32 `json:"MaxQueueingTimeMs,omitempty" xml:"MaxQueueingTimeMs,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the API resource.
	//
	// example:
	//
	// /c
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The throttling threshold.
	//
	// example:
	//
	// 5
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateFlowRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFlowRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateFlowRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *UpdateFlowRuleResponseBodyData) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *UpdateFlowRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateFlowRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateFlowRuleResponseBodyData) GetLimitApp() *string {
	return s.LimitApp
}

func (s *UpdateFlowRuleResponseBodyData) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *UpdateFlowRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateFlowRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *UpdateFlowRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateFlowRuleResponseBodyData) SetAppId(v string) *UpdateFlowRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetAppName(v string) *UpdateFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetControlBehavior(v int32) *UpdateFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetEnable(v bool) *UpdateFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetId(v int64) *UpdateFlowRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetLimitApp(v string) *UpdateFlowRuleResponseBodyData {
	s.LimitApp = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *UpdateFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetNamespace(v string) *UpdateFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetResource(v string) *UpdateFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) SetThreshold(v float32) *UpdateFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *UpdateFlowRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
