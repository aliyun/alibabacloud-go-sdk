// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGreyTagRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGreyTagRouteResponseBody
	GetCode() *string
	SetData(v *ListGreyTagRouteResponseBodyData) *ListGreyTagRouteResponseBody
	GetData() *ListGreyTagRouteResponseBodyData
	SetErrorCode(v string) *ListGreyTagRouteResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListGreyTagRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGreyTagRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGreyTagRouteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListGreyTagRouteResponseBody
	GetTraceId() *string
}

type ListGreyTagRouteResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The call was successful.
	//
	// - **3xx**: The call was redirected.
	//
	// - **4xx**: The call failed.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the canary release rule.
	Data *ListGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error code. Valid values:
	//
	// - If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned information. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D29CBD0-45D3-410B-9826-52F86F90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the information of the change order was queried. Valid values:
	//
	// - **true**: The information was queried.
	//
	// - **false**: The information failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListGreyTagRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGreyTagRouteResponseBody) GetData() *ListGreyTagRouteResponseBodyData {
	return s.Data
}

func (s *ListGreyTagRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListGreyTagRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGreyTagRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGreyTagRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGreyTagRouteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListGreyTagRouteResponseBody) SetCode(v string) *ListGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetData(v *ListGreyTagRouteResponseBodyData) *ListGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetErrorCode(v string) *ListGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetMessage(v string) *ListGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetRequestId(v string) *ListGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetSuccess(v bool) *ListGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetTraceId(v string) *ListGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned on each page. Valid value: **1**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned result.
	Result []*ListGreyTagRouteResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of canary release rules. Valid value: **1**.
	//
	// example:
	//
	// 1
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGreyTagRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListGreyTagRouteResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGreyTagRouteResponseBodyData) GetResult() []*ListGreyTagRouteResponseBodyDataResult {
	return s.Result
}

func (s *ListGreyTagRouteResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGreyTagRouteResponseBodyData) SetCurrentPage(v int32) *ListGreyTagRouteResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) SetPageSize(v int32) *ListGreyTagRouteResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) SetResult(v []*ListGreyTagRouteResponseBodyDataResult) *ListGreyTagRouteResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) SetTotalSize(v int64) *ListGreyTagRouteResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResult struct {
	AlbRules []*ListGreyTagRouteResponseBodyDataResultAlbRules `json:"AlbRules,omitempty" xml:"AlbRules,omitempty" type:"Repeated"`
	// The timestamp when the canary release rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1619007592013
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the canary release rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The canary release rule of the Dubbo service.
	DubboRules []*ListGreyTagRouteResponseBodyDataResultDubboRules `json:"DubboRules,omitempty" xml:"DubboRules,omitempty" type:"Repeated"`
	// The ID of the canary release rule.
	//
	// example:
	//
	// 1
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	// The name of the canary release rule.
	//
	// example:
	//
	// rule-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The canary release rule of the Spring Cloud application.
	ScRules []*ListGreyTagRouteResponseBodyDataResultScRules `json:"ScRules,omitempty" xml:"ScRules,omitempty" type:"Repeated"`
	// The timestamp when the canary release rule was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1609434061000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetAlbRules() []*ListGreyTagRouteResponseBodyDataResultAlbRules {
	return s.AlbRules
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetDubboRules() []*ListGreyTagRouteResponseBodyDataResultDubboRules {
	return s.DubboRules
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetScRules() []*ListGreyTagRouteResponseBodyDataResultScRules {
	return s.ScRules
}

func (s *ListGreyTagRouteResponseBodyDataResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetAlbRules(v []*ListGreyTagRouteResponseBodyDataResultAlbRules) *ListGreyTagRouteResponseBodyDataResult {
	s.AlbRules = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetCreateTime(v int64) *ListGreyTagRouteResponseBodyDataResult {
	s.CreateTime = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetDescription(v string) *ListGreyTagRouteResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetDubboRules(v []*ListGreyTagRouteResponseBodyDataResultDubboRules) *ListGreyTagRouteResponseBodyDataResult {
	s.DubboRules = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetGreyTagRouteId(v int64) *ListGreyTagRouteResponseBodyDataResult {
	s.GreyTagRouteId = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetName(v string) *ListGreyTagRouteResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetScRules(v []*ListGreyTagRouteResponseBodyDataResultScRules) *ListGreyTagRouteResponseBodyDataResult {
	s.ScRules = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetUpdateTime(v int64) *ListGreyTagRouteResponseBodyDataResult {
	s.UpdateTime = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResultAlbRules struct {
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// example:
	//
	// 23
	IngressId *string                                                `json:"ingressId,omitempty" xml:"ingressId,omitempty"`
	Items     []*ListGreyTagRouteResponseBodyDataResultAlbRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// s-6366-e3****-99**
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultAlbRules) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultAlbRules) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) GetCondition() *string {
	return s.Condition
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) GetIngressId() *string {
	return s.IngressId
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) GetItems() []*ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	return s.Items
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) SetCondition(v string) *ListGreyTagRouteResponseBodyDataResultAlbRules {
	s.Condition = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) SetIngressId(v string) *ListGreyTagRouteResponseBodyDataResultAlbRules {
	s.IngressId = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) SetItems(v []*ListGreyTagRouteResponseBodyDataResultAlbRulesItems) *ListGreyTagRouteResponseBodyDataResultAlbRules {
	s.Items = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) SetServiceName(v string) *ListGreyTagRouteResponseBodyDataResultAlbRules {
	s.ServiceName = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRules) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResultAlbRulesItems struct {
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// example:
	//
	// N/A
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// example:
	//
	// N/A
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultAlbRulesItems) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetCond() *string {
	return s.Cond
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetExpr() *string {
	return s.Expr
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetIndex() *int32 {
	return s.Index
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetName() *string {
	return s.Name
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetOperator() *string {
	return s.Operator
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetType() *string {
	return s.Type
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) GetValue() *string {
	return s.Value
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetCond(v string) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Cond = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetExpr(v string) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Expr = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetIndex(v int32) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Index = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetName(v string) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetOperator(v string) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Operator = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetType(v string) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Type = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) SetValue(v string) *ListGreyTagRouteResponseBodyDataResultAlbRulesItems {
	s.Value = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultAlbRulesItems) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResultDubboRules struct {
	// The relationship between the conditions in the canary release rule. Valid values:
	//
	// - **AND**: The conditions are in the logical AND relation. All conditions must be met at the same time.
	//
	// - **OR**: The conditions are in the logical OR relation. At least one of the conditions must be met.
	//
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The group of the Dubbo service that corresponds to the canary release rule.
	//
	// example:
	//
	// DUBBO
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The conditions.
	Items []*ListGreyTagRouteResponseBodyDataResultDubboRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The method name of the Dubbo service.
	//
	// example:
	//
	// echo
	MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
	// The name of the Dubbo service.
	//
	// example:
	//
	// com.alibaba.edas.boot.EchoService
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The version of the Dubbo service.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRules) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRules) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) GetCondition() *string {
	return s.Condition
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) GetGroup() *string {
	return s.Group
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) GetItems() []*ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	return s.Items
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) GetMethodName() *string {
	return s.MethodName
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) GetVersion() *string {
	return s.Version
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetCondition(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Condition = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetGroup(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Group = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetItems(v []*ListGreyTagRouteResponseBodyDataResultDubboRulesItems) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Items = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetMethodName(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.MethodName = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetServiceName(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.ServiceName = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetVersion(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Version = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResultDubboRulesItems struct {
	// The comparison operator. Valid values: **>**, **<**, **>=**, **<=**, **==**, and **! =**.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The expression that is used to obtain the value of the parameter. The syntax of the expression must follow the standard of the SpEL. Valid values:
	//
	// - **Empty**: obtains the value of the parameter.
	//
	// - **.name**: obtains the name property of the parameter. This expression works the same way as args0.getName().
	//
	// - **.isEnabled()**: obtains the enabled property of the parameter. This expression works the same way as args0.isEnabled().
	//
	// - **[0]**: indicates that the value of the parameter is an array and obtains the first value of the array. This expression works the same way as args0[0]. This expression does not start with a period (.).
	//
	// - **.get(0)**: indicates that the value of the parameter is a list and obtains the first value of the list. This expression works the same way as args0.get(0).
	//
	// - **.get("key")**: indicates that the value of the parameter is a map and obtains the value of the key in the map. This expression works the same way as args0.get("key").  >  For more information about the expressions that are used to obtain parameter values, see  [Spring Expression Language (SpEL)](https://docs.spring.io/spring-integration/docs/current/reference/html/spel.html).
	//
	// example:
	//
	// .name
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// The index of the parameter. The value 0 indicates the first parameter.
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// This parameter is not returned for Dubbo services.
	//
	// example:
	//
	// N/A
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operator. Valid values:
	//
	// - **rawvalue**: direct comparison.
	//
	// - **list**: whitelist.
	//
	// - **mod**: mods 100.
	//
	// - **deterministic_proportional_steaming_division**: percentage.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// This parameter is not returned for Dubbo services.
	//
	// example:
	//
	// N/A
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the parameter. This value is compared with the value that is obtained based on the **expr*	- and **index*	- parameters.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRulesItems) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetCond() *string {
	return s.Cond
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetExpr() *string {
	return s.Expr
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetIndex() *int32 {
	return s.Index
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetName() *string {
	return s.Name
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetOperator() *string {
	return s.Operator
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetType() *string {
	return s.Type
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GetValue() *string {
	return s.Value
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetCond(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Cond = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetExpr(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Expr = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetIndex(v int32) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Index = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetName(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetOperator(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Operator = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetType(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Type = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetValue(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Value = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResultScRules struct {
	// The relationship between the conditions in the canary release rule. Valid values:
	//
	// 	- **AND**: The conditions are in the logical AND relation. All conditions must be met at the same time.
	//
	// 	- **OR**: The conditions are in the logical OR relation. At least one of the conditions must be met.
	//
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The conditions.
	Items []*ListGreyTagRouteResponseBodyDataResultScRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The path of the canary release rule of the Spring Cloud application.
	//
	// example:
	//
	// /path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultScRules) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultScRules) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) GetCondition() *string {
	return s.Condition
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) GetItems() []*ListGreyTagRouteResponseBodyDataResultScRulesItems {
	return s.Items
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) GetPath() *string {
	return s.Path
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) SetCondition(v string) *ListGreyTagRouteResponseBodyDataResultScRules {
	s.Condition = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) SetItems(v []*ListGreyTagRouteResponseBodyDataResultScRulesItems) *ListGreyTagRouteResponseBodyDataResultScRules {
	s.Items = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) SetPath(v string) *ListGreyTagRouteResponseBodyDataResultScRules {
	s.Path = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) Validate() error {
	return dara.Validate(s)
}

type ListGreyTagRouteResponseBodyDataResultScRulesItems struct {
	// The comparison operator. Valid values: **>**, **<**, **>=**, **<=**, **==**, and **! =**.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// This parameter is not returned for Spring Cloud applications.
	//
	// example:
	//
	// N/A
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// This parameter is not returned for Spring Cloud applications.
	//
	// example:
	//
	// N/A
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operator. Valid values:
	//
	// 	- **rawvalue**: direct comparison.
	//
	// 	- **list**: whitelist.
	//
	// 	- **mod**: mods 100.
	//
	// 	- **deterministic_proportional_steaming_division**: percentage.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The type of the comparison. Valid values:
	//
	// 	- **param**: parameter
	//
	// 	- **cookie**: cookie
	//
	// 	- **header**: header
	//
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the parameter. This value is compared with the value that is obtained based on the **type*	- and **name*	- parameters.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultScRulesItems) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultScRulesItems) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetCond() *string {
	return s.Cond
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetExpr() *string {
	return s.Expr
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetIndex() *int32 {
	return s.Index
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetName() *string {
	return s.Name
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetOperator() *string {
	return s.Operator
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetType() *string {
	return s.Type
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) GetValue() *string {
	return s.Value
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetCond(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Cond = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetExpr(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Expr = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetIndex(v int32) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Index = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetName(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetOperator(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Operator = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetType(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Type = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetValue(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Value = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) Validate() error {
	return dara.Validate(s)
}
