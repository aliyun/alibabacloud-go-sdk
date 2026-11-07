// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsWithTagRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListApplicationsWithTagRulesResponseBodyData) *ListApplicationsWithTagRulesResponseBody
	GetData() *ListApplicationsWithTagRulesResponseBodyData
	SetHttpStatusCode(v int32) *ListApplicationsWithTagRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListApplicationsWithTagRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationsWithTagRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApplicationsWithTagRulesResponseBody
	GetSuccess() *bool
}

type ListApplicationsWithTagRulesResponseBody struct {
	// The data overview.
	Data *ListApplicationsWithTagRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, a success message is returned.
	//
	// - If the request fails, the cause of the failure is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7466566F-F30F-4A29-965D-3E0AF21D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBody) GetData() *ListApplicationsWithTagRulesResponseBodyData {
	return s.Data
}

func (s *ListApplicationsWithTagRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApplicationsWithTagRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationsWithTagRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsWithTagRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApplicationsWithTagRulesResponseBody) SetData(v *ListApplicationsWithTagRulesResponseBodyData) *ListApplicationsWithTagRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBody) SetHttpStatusCode(v int32) *ListApplicationsWithTagRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBody) SetMessage(v string) *ListApplicationsWithTagRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBody) SetRequestId(v string) *ListApplicationsWithTagRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBody) SetSuccess(v bool) *ListApplicationsWithTagRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data structure.
	Result []*ListApplicationsWithTagRulesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsWithTagRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsWithTagRulesResponseBodyData) GetResult() []*ListApplicationsWithTagRulesResponseBodyDataResult {
	return s.Result
}

func (s *ListApplicationsWithTagRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListApplicationsWithTagRulesResponseBodyData) SetPageNumber(v int32) *ListApplicationsWithTagRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyData) SetPageSize(v int32) *ListApplicationsWithTagRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyData) SetResult(v []*ListApplicationsWithTagRulesResponseBodyDataResult) *ListApplicationsWithTagRulesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyData) SetTotalSize(v int32) *ListApplicationsWithTagRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyDataResult struct {
	// The ID of the application.
	//
	// example:
	//
	// daqijp6c31@xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// service-lottery-core
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The MSE namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The list of rules.
	RouteRules []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRules `json:"RouteRules,omitempty" xml:"RouteRules,omitempty" type:"Repeated"`
	// The routing status. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 1
	RouteStatus *int64 `json:"RouteStatus,omitempty" xml:"RouteStatus,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) GetRouteRules() []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	return s.RouteRules
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) GetRouteStatus() *int64 {
	return s.RouteStatus
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) SetAppId(v string) *ListApplicationsWithTagRulesResponseBodyDataResult {
	s.AppId = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) SetAppName(v string) *ListApplicationsWithTagRulesResponseBodyDataResult {
	s.AppName = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) SetNamespace(v string) *ListApplicationsWithTagRulesResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) SetRouteRules(v []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) *ListApplicationsWithTagRulesResponseBodyDataResult {
	s.RouteRules = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) SetRouteStatus(v int64) *ListApplicationsWithTagRulesResponseBodyDataResult {
	s.RouteStatus = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResult) Validate() error {
	if s.RouteRules != nil {
		for _, item := range s.RouteRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyDataResultRouteRules struct {
	// The enabling status of the alert rule. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 653
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The rule name.
	//
	// example:
	//
	// dubbo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The proportion.
	//
	// example:
	//
	// 10
	Rate *int32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The applicable routes of the rule.
	//
	// example:
	//
	// {
	//
	//   	"dubbo": [{
	//
	//     "serviceName": "com.taobao.hsf.common.DemoService",
	//
	//     "group": "",
	//
	//     "version": "",
	//
	//     "methodName": "sayHello",
	//
	//     "condition": "AND",
	//
	//     "argumentItems": [{
	//
	//     	"index": 0,
	//
	//     	"expr": "",
	//
	//     	"operator": "rawvalue",
	//
	//     	"value": "jim",
	//
	//     	"cond": "=="
	//
	//     }]
	//
	//   	}]
	//
	//   }
	Rules *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetEnable() *bool {
	return s.Enable
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetId() *int64 {
	return s.Id
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetName() *string {
	return s.Name
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetRate() *int32 {
	return s.Rate
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetRules() *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules {
	return s.Rules
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetStatus() *int32 {
	return s.Status
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) GetTag() *string {
	return s.Tag
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetEnable(v bool) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Enable = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetId(v int64) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Id = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetInstanceNum(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.InstanceNum = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetName(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Name = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetRate(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Rate = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetRules(v *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Rules = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetStatus(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Status = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) SetTag(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules {
	s.Tag = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRules) Validate() error {
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules struct {
	// The Dubbo rule.
	Dubbo []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo `json:"dubbo,omitempty" xml:"dubbo,omitempty" type:"Repeated"`
	// The Spring Cloud rule.
	Springcloud []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud `json:"springcloud,omitempty" xml:"springcloud,omitempty" type:"Repeated"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) GetDubbo() []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	return s.Dubbo
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) GetSpringcloud() []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	return s.Springcloud
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) SetDubbo(v []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules {
	s.Dubbo = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) SetSpringcloud(v []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules {
	s.Springcloud = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRules) Validate() error {
	if s.Dubbo != nil {
		for _, item := range s.Dubbo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Springcloud != nil {
		for _, item := range s.Springcloud {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo struct {
	// The ID of the application.
	//
	// example:
	//
	// ***@***
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The list of parameter arguments.
	ArgumentItems []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems `json:"argumentItems,omitempty" xml:"argumentItems,omitempty" type:"Repeated"`
	// The logical relationship among the following items. Valid values: AND and OR.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The Dubbo service group.
	//
	// example:
	//
	// default
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The Dubbo service method name.
	//
	// example:
	//
	// hello
	MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
	// The list of parameter types.
	ParamTypes []*string `json:"paramTypes,omitempty" xml:"paramTypes,omitempty" type:"Repeated"`
	// The service name (interface).
	//
	// example:
	//
	// HelloService
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The list of tags.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The policy type.
	//
	// example:
	//
	// CONTENT
	TriggerPolicy *string `json:"triggerPolicy,omitempty" xml:"triggerPolicy,omitempty"`
	// The Dubbo service version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetArgumentItems() []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	return s.ArgumentItems
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetCondition() *string {
	return s.Condition
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetGroup() *string {
	return s.Group
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetMethodName() *string {
	return s.MethodName
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetParamTypes() []*string {
	return s.ParamTypes
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetTags() []*string {
	return s.Tags
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetTriggerPolicy() *string {
	return s.TriggerPolicy
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) GetVersion() *string {
	return s.Version
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetAppId(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.AppId = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetArgumentItems(v []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.ArgumentItems = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetCondition(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.Condition = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetGroup(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.Group = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetMethodName(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.MethodName = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetParamTypes(v []*string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.ParamTypes = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetServiceName(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.ServiceName = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetTags(v []*string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.Tags = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetTriggerPolicy(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.TriggerPolicy = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) SetVersion(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo {
	s.Version = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubbo) Validate() error {
	if s.ArgumentItems != nil {
		for _, item := range s.ArgumentItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems struct {
	// The comparison operator. Valid values: >=, <=, >, <, and ==.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The value to be compared for operators such as rawvalue.
	//
	// example:
	//
	// 30
	Datum *string `json:"datum,omitempty" xml:"datum,omitempty"`
	// The field required for the mod operator, which is the divisor.
	//
	// example:
	//
	// 30
	Divisor *int32 `json:"divisor,omitempty" xml:"divisor,omitempty"`
	// The expression. For more information, see the descriptions of other fields.
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// The parameter position, starting from 0.
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The list of names.
	NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// The operator. A value of rawvalue indicates direct comparison, mode indicates modulo operation, and list indicates a whitelist.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The proportion rule, where 20% of traffic is routed to the node corresponding to this tag.
	//
	// example:
	//
	// 10
	Rate *int32 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The field required for the percentage operation, which is the percentage value.
	//
	// example:
	//
	// 30
	Remainder *int32 `json:"remainder,omitempty" xml:"remainder,omitempty"`
	// The value used for comparison. The value obtained from the expression is compared with this value. If the operator is a list whitelist, separate multiple values with commas (,), such as 1,2,3.
	//
	// example:
	//
	// 1,2,3
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetCond() *string {
	return s.Cond
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetDatum() *string {
	return s.Datum
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetExpr() *string {
	return s.Expr
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetIndex() *int32 {
	return s.Index
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetNameList() []*string {
	return s.NameList
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetOperator() *string {
	return s.Operator
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetRate() *int32 {
	return s.Rate
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) GetValue() *string {
	return s.Value
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetCond(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Cond = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetDatum(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Datum = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetDivisor(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Divisor = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetExpr(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Expr = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetIndex(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Index = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetNameList(v []*string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.NameList = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetOperator(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Operator = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetRate(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Rate = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetRemainder(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Remainder = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) SetValue(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems {
	s.Value = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesDubboArgumentItems) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud struct {
	// The ID of the application.
	//
	// example:
	//
	// ***@***
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The logical relationship among the following items. Valid values: AND and OR.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Specifies whether the feature is enabled.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The path.
	//
	// example:
	//
	// /hello
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The list of paths.
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The list of matching conditions for REST requests.
	RestItems []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems `json:"restItems,omitempty" xml:"restItems,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The policy type.
	//
	// example:
	//
	// CONTENT
	TriggerPolicy *string `json:"triggerPolicy,omitempty" xml:"triggerPolicy,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetCondition() *string {
	return s.Condition
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetEnable() *bool {
	return s.Enable
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetPath() *string {
	return s.Path
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetPaths() []*string {
	return s.Paths
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetPriority() *int32 {
	return s.Priority
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetRestItems() []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	return s.RestItems
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetTags() []*string {
	return s.Tags
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) GetTriggerPolicy() *string {
	return s.TriggerPolicy
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetAppId(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.AppId = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetCondition(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.Condition = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetEnable(v bool) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.Enable = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetPath(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.Path = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetPaths(v []*string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.Paths = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetPriority(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.Priority = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetRestItems(v []*ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.RestItems = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetTags(v []*string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.Tags = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) SetTriggerPolicy(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud {
	s.TriggerPolicy = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloud) Validate() error {
	if s.RestItems != nil {
		for _, item := range s.RestItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems struct {
	// The comparison operator. Valid values: >=, <=, >, <, and ==.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The value to be compared for operators such as rawvalue.
	//
	// example:
	//
	// 10
	Datum *string `json:"datum,omitempty" xml:"datum,omitempty"`
	// The field required for the mod operator, which is the divisor.
	//
	// example:
	//
	// 30
	Divisor *int32 `json:"divisor,omitempty" xml:"divisor,omitempty"`
	// The name.
	//
	// example:
	//
	// key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The field required for the list operator.
	NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// The operator. A value of rawvalue indicates direct comparison, mode indicates modulo operation, and list indicates a whitelist.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The proportion rule, where 20% of traffic is routed to the node corresponding to this tag.
	//
	// example:
	//
	// 20
	Rate *int32 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The field required for the percentage operation, which is the percentage value.
	//
	// example:
	//
	// 30
	Remainder *int32 `json:"remainder,omitempty" xml:"remainder,omitempty"`
	// The type.
	//
	// example:
	//
	// test
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetCond() *string {
	return s.Cond
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetDatum() *string {
	return s.Datum
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetName() *string {
	return s.Name
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetOperator() *string {
	return s.Operator
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetType() *string {
	return s.Type
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) GetValue() *string {
	return s.Value
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetCond(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Cond = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetDatum(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Datum = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetDivisor(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Divisor = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetName(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Name = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetNameList(v []*string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.NameList = v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetOperator(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Operator = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetRate(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Rate = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetRemainder(v int32) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Remainder = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetType(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Type = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) SetValue(v string) *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems {
	s.Value = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponseBodyDataResultRouteRulesRulesSpringcloudRestItems) Validate() error {
	return dara.Validate(s)
}
