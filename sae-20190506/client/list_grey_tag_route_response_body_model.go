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
	// The status of the interface or the POP error code. Valid values:
	//
	// - **2xx**: The request is successful.
	//
	// - **3xx**: The request is redirected.
	//
	// - **4xx**: A request error occurs.
	//
	// - **5xx**: A server error occurs.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the grayscale rule.
	Data *ListGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9D29CBD0-45D3-410B-9826-52F86F90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the query succeeded.
	//
	// - **true**: The query succeeded.
	//
	// - **false**: The query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to query the details of a call.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGreyTagRouteResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page in a paged query. The value can only be **1**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned result.
	Result []*ListGreyTagRouteResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries. The value can only be **1**.
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

type ListGreyTagRouteResponseBodyDataResult struct {
	// The grayscale rules created for an application for which an Application Load Balancer (ALB) Ingress is configured.
	AlbRules []*ListGreyTagRouteResponseBodyDataResultAlbRules `json:"AlbRules,omitempty" xml:"AlbRules,omitempty" type:"Repeated"`
	// The timestamp when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1619007592013
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The grayscale rules for Dubbo services.
	DubboRules []*ListGreyTagRouteResponseBodyDataResultDubboRules `json:"DubboRules,omitempty" xml:"DubboRules,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 1
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The grayscale rules for Spring Cloud.
	ScRules []*ListGreyTagRouteResponseBodyDataResultScRules `json:"ScRules,omitempty" xml:"ScRules,omitempty" type:"Repeated"`
	// The timestamp when the rule was updated. Unit: milliseconds.
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
	if s.AlbRules != nil {
		for _, item := range s.AlbRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DubboRules != nil {
		for _, item := range s.DubboRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScRules != nil {
		for _, item := range s.ScRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGreyTagRouteResponseBodyDataResultAlbRules struct {
	// The conditional pattern of the grayscale rule. Only AND is supported, which indicates that all conditions in the list must be met.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The Ingress ID.
	//
	// example:
	//
	// 23
	IngressId *string `json:"ingressId,omitempty" xml:"ingressId,omitempty"`
	// The list of conditions.
	Items []*ListGreyTagRouteResponseBodyDataResultAlbRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The service name.
	//
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
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGreyTagRouteResponseBodyDataResultAlbRulesItems struct {
	// Only == is supported.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// This parameter is not applicable to ALB applications.
	//
	// example:
	//
	// N/A
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// This parameter is not applicable to ALB applications.
	//
	// example:
	//
	// N/A
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operator. Valid values: Only rawvalue is supported, which indicates a direct comparison.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The comparison type. Valid values:
	//
	// - **sourceIp**: SourceIp.
	//
	// - **cookie**: Cookie.
	//
	// - **header**: Header.
	//
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The parameter value. The value obtained based on type and name is compared with this value.
	//
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
	// The conditional pattern of the grayscale rule. Valid values:
	//
	// - **AND**: All conditions in the list must be met.
	//
	// - **OR**: Any condition in the list can be met.
	//
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The group of the Dubbo service that corresponds to the grayscale rule.
	//
	// example:
	//
	// DUBBO
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The list of conditions.
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
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGreyTagRouteResponseBodyDataResultDubboRulesItems struct {
	// The comparison operator. Valid values: **>**, **<**, **>=**, **<=**, **==**, and **!=**.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The expression that is used to obtain the parameter value. The syntax follows the Spring Expression Language (SpEL). Valid values:
	//
	// - **Leave it empty**: Obtains the value of the current parameter.
	//
	// - **.name**: Obtains the name property of the parameter. This is equivalent to args0.getName().
	//
	// - **.isEnabled()**: Obtains the enabled property of the parameter. This is equivalent to args0.isEnabled().
	//
	// - **[0]**: Obtains the first value of the array. This is equivalent to args0[0]. Note that the expression does not start with a period (.).
	//
	// - **.get(0)**: Obtains the first value of the list. This is equivalent to args0.get(0).
	//
	// - **.get("key")**: Obtains the value that corresponds to the specified key from the map. This is equivalent to args0.get("key").
	//
	// example:
	//
	// .name
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// The parameter number. 0 indicates the first parameter.
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// This parameter is not applicable to Dubbo services.
	//
	// example:
	//
	// N/A
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operator. Valid values:
	//
	// - **rawvalue**: Direct comparison.
	//
	// - **list**: Whitelist.
	//
	// - **mod**: Modulo 100 operation.
	//
	// - **deterministic_proportional_steaming_division**: Percentage.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// This parameter is not applicable to Dubbo services.
	//
	// example:
	//
	// N/A
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The parameter value. The value obtained based on **expr*	- and **index*	- is compared with this value.
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
	// The conditional pattern of the grayscale rule. Valid values:
	//
	// - **AND**: All conditions in the list must be met.
	//
	// - **OR**: Any condition in the list can be met.
	//
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The list of conditions.
	Items []*ListGreyTagRouteResponseBodyDataResultScRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The path that corresponds to the grayscale rule for the Spring Cloud application.
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
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGreyTagRouteResponseBodyDataResultScRulesItems struct {
	// The comparison operator. Valid values: **>**, **<**, **>=**, **<=**, **==**, and **!=**.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// This parameter is not applicable to Spring Cloud applications.
	//
	// example:
	//
	// N/A
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// This parameter is not applicable to Spring Cloud applications.
	//
	// example:
	//
	// N/A
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operator. Valid values:
	//
	// - **rawvalue**: Direct comparison.
	//
	// - **list**: Whitelist.
	//
	// - **mod**: Modulo 100 operation.
	//
	// - **deterministic_proportional_steaming_division**: Percentage.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The comparison type. Valid values:
	//
	// - **param**: Parameter.
	//
	// - **cookie**: Cookie.
	//
	// - **header**: Header.
	//
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The parameter value. The value obtained based on **type*	- and **name*	- is compared with this value.
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
