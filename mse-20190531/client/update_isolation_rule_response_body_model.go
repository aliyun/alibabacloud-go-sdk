// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIsolationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateIsolationRuleResponseBody
	GetCode() *string
	SetData(v *UpdateIsolationRuleResponseBodyData) *UpdateIsolationRuleResponseBody
	GetData() *UpdateIsolationRuleResponseBodyData
	SetMessage(v string) *UpdateIsolationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateIsolationRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateIsolationRuleResponseBody
	GetSuccess() *bool
}

type UpdateIsolationRuleResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The rule information.
	Data *UpdateIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateIsolationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIsolationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateIsolationRuleResponseBody) GetData() *UpdateIsolationRuleResponseBodyData {
	return s.Data
}

func (s *UpdateIsolationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateIsolationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIsolationRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIsolationRuleResponseBody) SetCode(v string) *UpdateIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIsolationRuleResponseBody) SetData(v *UpdateIsolationRuleResponseBodyData) *UpdateIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIsolationRuleResponseBody) SetMessage(v string) *UpdateIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIsolationRuleResponseBody) SetRequestId(v string) *UpdateIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIsolationRuleResponseBody) SetSuccess(v bool) *UpdateIsolationRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIsolationRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIsolationRuleResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
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
	// The ID of the rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The source application of the request.
	//
	// example:
	//
	// spring-cloud-a
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// The microservice namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The application resource to which the rule applies. This value must be the same as the API operation name displayed in the API details on the console.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The concurrency threshold of the isolation rule.
	//
	// example:
	//
	// 3
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateIsolationRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIsolationRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateIsolationRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *UpdateIsolationRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateIsolationRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateIsolationRuleResponseBodyData) GetLimitApp() *string {
	return s.LimitApp
}

func (s *UpdateIsolationRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateIsolationRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *UpdateIsolationRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateIsolationRuleResponseBodyData) SetAppId(v string) *UpdateIsolationRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetAppName(v string) *UpdateIsolationRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetEnable(v bool) *UpdateIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetId(v int64) *UpdateIsolationRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetLimitApp(v string) *UpdateIsolationRuleResponseBodyData {
	s.LimitApp = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetNamespace(v string) *UpdateIsolationRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetResource(v string) *UpdateIsolationRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) SetThreshold(v float32) *UpdateIsolationRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *UpdateIsolationRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
