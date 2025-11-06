// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFlowRuleResponseBody
	GetCode() *string
	SetData(v *CreateFlowRuleResponseBodyData) *CreateFlowRuleResponseBody
	GetData() *CreateFlowRuleResponseBodyData
	SetMessage(v string) *CreateFlowRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFlowRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFlowRuleResponseBody
	GetSuccess() *bool
}

type CreateFlowRuleResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateFlowRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFlowRuleResponseBody) GetData() *CreateFlowRuleResponseBodyData {
	return s.Data
}

func (s *CreateFlowRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlowRuleResponseBody) SetCode(v string) *CreateFlowRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowRuleResponseBody) SetData(v *CreateFlowRuleResponseBodyData) *CreateFlowRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateFlowRuleResponseBody) SetMessage(v string) *CreateFlowRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowRuleResponseBody) SetRequestId(v string) *CreateFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowRuleResponseBody) SetSuccess(v bool) *CreateFlowRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFlowRuleResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// hkhon1po62@54e1f42f3******
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
	// Indicates whether the rule is enabled.
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
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID.
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the API resource.
	//
	// example:
	//
	// app
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The throttling threshold.
	//
	// example:
	//
	// 200
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateFlowRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateFlowRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateFlowRuleResponseBodyData) GetControlBehavior() *int32 {
	return s.ControlBehavior
}

func (s *CreateFlowRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *CreateFlowRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateFlowRuleResponseBodyData) GetLimitApp() *string {
	return s.LimitApp
}

func (s *CreateFlowRuleResponseBodyData) GetMaxQueueingTimeMs() *int32 {
	return s.MaxQueueingTimeMs
}

func (s *CreateFlowRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateFlowRuleResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFlowRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *CreateFlowRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateFlowRuleResponseBodyData) SetAppId(v string) *CreateFlowRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetAppName(v string) *CreateFlowRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetControlBehavior(v int32) *CreateFlowRuleResponseBodyData {
	s.ControlBehavior = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetEnable(v bool) *CreateFlowRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetId(v int64) *CreateFlowRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetLimitApp(v string) *CreateFlowRuleResponseBodyData {
	s.LimitApp = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetMaxQueueingTimeMs(v int32) *CreateFlowRuleResponseBodyData {
	s.MaxQueueingTimeMs = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetNamespace(v string) *CreateFlowRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetRegionId(v string) *CreateFlowRuleResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetResource(v string) *CreateFlowRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) SetThreshold(v float32) *CreateFlowRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *CreateFlowRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
