// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGreyTagRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeGreyTagRouteResponseBody
	GetCode() *string
	SetData(v *DescribeGreyTagRouteResponseBodyData) *DescribeGreyTagRouteResponseBody
	GetData() *DescribeGreyTagRouteResponseBodyData
	SetErrorCode(v string) *DescribeGreyTagRouteResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeGreyTagRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeGreyTagRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeGreyTagRouteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeGreyTagRouteResponseBody
	GetTraceId() *string
}

type DescribeGreyTagRouteResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the canary release rule.
	Data *DescribeGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
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
	// Indicates whether the information of the change order was queried. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The information failed to be queried.
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

func (s DescribeGreyTagRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeGreyTagRouteResponseBody) GetData() *DescribeGreyTagRouteResponseBodyData {
	return s.Data
}

func (s *DescribeGreyTagRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeGreyTagRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeGreyTagRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGreyTagRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeGreyTagRouteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeGreyTagRouteResponseBody) SetCode(v string) *DescribeGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetData(v *DescribeGreyTagRouteResponseBodyData) *DescribeGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetErrorCode(v string) *DescribeGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetMessage(v string) *DescribeGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetRequestId(v string) *DescribeGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetSuccess(v bool) *DescribeGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetTraceId(v string) *DescribeGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGreyTagRouteResponseBodyData struct {
	AlbRules []*DescribeGreyTagRouteResponseBodyDataAlbRules `json:"AlbRules,omitempty" xml:"AlbRules,omitempty" type:"Repeated"`
	// The ID of the application.
	//
	// example:
	//
	// 3faaf993-7aed-4bcd-b189-625e6a5a****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	DubboRules []*DescribeGreyTagRouteResponseBodyDataDubboRules `json:"DubboRules,omitempty" xml:"DubboRules,omitempty" type:"Repeated"`
	// The ID of the canary release rule. The ID is globally unique.
	//
	// example:
	//
	// 16
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	// The name of the canary release rule.
	//
	// example:
	//
	// rule-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The canary release rule of the Spring Cloud application.
	ScRules []*DescribeGreyTagRouteResponseBodyDataScRules `json:"ScRules,omitempty" xml:"ScRules,omitempty" type:"Repeated"`
	// The timestamp when the canary release rule was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1609434061000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyData) GetAlbRules() []*DescribeGreyTagRouteResponseBodyDataAlbRules {
	return s.AlbRules
}

func (s *DescribeGreyTagRouteResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeGreyTagRouteResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGreyTagRouteResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeGreyTagRouteResponseBodyData) GetDubboRules() []*DescribeGreyTagRouteResponseBodyDataDubboRules {
	return s.DubboRules
}

func (s *DescribeGreyTagRouteResponseBodyData) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *DescribeGreyTagRouteResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeGreyTagRouteResponseBodyData) GetScRules() []*DescribeGreyTagRouteResponseBodyDataScRules {
	return s.ScRules
}

func (s *DescribeGreyTagRouteResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeGreyTagRouteResponseBodyData) SetAlbRules(v []*DescribeGreyTagRouteResponseBodyDataAlbRules) *DescribeGreyTagRouteResponseBodyData {
	s.AlbRules = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetAppId(v string) *DescribeGreyTagRouteResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetCreateTime(v int64) *DescribeGreyTagRouteResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetDescription(v string) *DescribeGreyTagRouteResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetDubboRules(v []*DescribeGreyTagRouteResponseBodyDataDubboRules) *DescribeGreyTagRouteResponseBodyData {
	s.DubboRules = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *DescribeGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetName(v string) *DescribeGreyTagRouteResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetScRules(v []*DescribeGreyTagRouteResponseBodyDataScRules) *DescribeGreyTagRouteResponseBodyData {
	s.ScRules = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetUpdateTime(v int64) *DescribeGreyTagRouteResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) Validate() error {
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

type DescribeGreyTagRouteResponseBodyDataAlbRules struct {
	// The condition mode of the canary release rule. Valid value: AND. This value indicates that that all conditions must be met.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The ID of the gateway routing rule.
	//
	// example:
	//
	// 23
	IngressId *string                                              `json:"ingressId,omitempty" xml:"ingressId,omitempty"`
	Items     []*DescribeGreyTagRouteResponseBodyDataAlbRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The service ID.
	//
	// example:
	//
	// 22
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataAlbRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataAlbRules) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) GetCondition() *string {
	return s.Condition
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) GetIngressId() *string {
	return s.IngressId
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) GetItems() []*DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	return s.Items
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) SetCondition(v string) *DescribeGreyTagRouteResponseBodyDataAlbRules {
	s.Condition = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) SetIngressId(v string) *DescribeGreyTagRouteResponseBodyDataAlbRules {
	s.IngressId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) SetItems(v []*DescribeGreyTagRouteResponseBodyDataAlbRulesItems) *DescribeGreyTagRouteResponseBodyDataAlbRules {
	s.Items = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) SetServiceId(v string) *DescribeGreyTagRouteResponseBodyDataAlbRules {
	s.ServiceId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRules) Validate() error {
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

type DescribeGreyTagRouteResponseBodyDataAlbRulesItems struct {
	// Valid value: ==.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// This parameter is not returned for applications that are associated with ALB instances.
	//
	// example:
	//
	// N/A
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// This parameter is not returned for applications that are associated with Application Load Balancer (ALB) instances.
	//
	// example:
	//
	// N/A
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operator. Valid value: **rawvalue**. This value indicates direct comparison.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The type of the comparison. Valid values:
	//
	// 	- **sourceIp**: SourceIp
	//
	// 	- **cookie**: cookie
	//
	// 	- **header**: header
	//
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the parameter. This value is compared with the value that is obtained based on the type and name parameters.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataAlbRulesItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetCond() *string {
	return s.Cond
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetExpr() *string {
	return s.Expr
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetName() *string {
	return s.Name
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetOperator() *string {
	return s.Operator
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetType() *string {
	return s.Type
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) GetValue() *string {
	return s.Value
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetCond(v string) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Cond = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetExpr(v string) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Expr = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetIndex(v int32) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Index = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetName(v string) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetOperator(v string) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Operator = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetType(v string) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Type = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) SetValue(v string) *DescribeGreyTagRouteResponseBodyDataAlbRulesItems {
	s.Value = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataAlbRulesItems) Validate() error {
	return dara.Validate(s)
}

type DescribeGreyTagRouteResponseBodyDataDubboRules struct {
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
	// The group of the Dubbo service that corresponds to the canary release rule.
	//
	// example:
	//
	// DUBBO
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The conditions.
	Items []*DescribeGreyTagRouteResponseBodyDataDubboRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s DescribeGreyTagRouteResponseBodyDataDubboRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataDubboRules) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) GetCondition() *string {
	return s.Condition
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) GetGroup() *string {
	return s.Group
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) GetItems() []*DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	return s.Items
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) GetMethodName() *string {
	return s.MethodName
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) GetVersion() *string {
	return s.Version
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetCondition(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Condition = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetGroup(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Group = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetItems(v []*DescribeGreyTagRouteResponseBodyDataDubboRulesItems) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Items = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetMethodName(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.MethodName = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetServiceName(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.ServiceName = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetVersion(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Version = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) Validate() error {
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

type DescribeGreyTagRouteResponseBodyDataDubboRulesItems struct {
	// The comparison operator. Valid values: **>**, **<**, **>=**, **<=**, **==**, and **! =**.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The expression that is used to obtain the value of the parameter. Valid values:
	//
	// 	- **Empty**: obtains the value of the parameter.
	//
	// 	- **.name**: obtains the name property of the parameter. This expression works the same way as args0.getName().
	//
	// 	- **.isEnabled()**: obtains the enabled property of the parameter. This expression works the same way as args0.isEnabled().
	//
	// 	- **[0]**: indicates that the value of the parameter is an array and obtains the first value of the array. This expression works the same way as args0[0]. This expression does not start with a period (.).
	//
	// 	- **.get(0)**: indicates that the value of the parameter is a list and obtains the first value of the list. This expression works the same way as args0.get(0).
	//
	// 	- **.get("key")**: indicates that the value of the parameter is a map and obtains the value of the key in the map. This expression works the same way as args0.get("key").
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

func (s DescribeGreyTagRouteResponseBodyDataDubboRulesItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetCond() *string {
	return s.Cond
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetExpr() *string {
	return s.Expr
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetName() *string {
	return s.Name
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetOperator() *string {
	return s.Operator
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetType() *string {
	return s.Type
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GetValue() *string {
	return s.Value
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetCond(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Cond = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetExpr(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Expr = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetIndex(v int32) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Index = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetName(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetOperator(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Operator = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetType(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Type = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetValue(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Value = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) Validate() error {
	return dara.Validate(s)
}

type DescribeGreyTagRouteResponseBodyDataScRules struct {
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
	Items []*DescribeGreyTagRouteResponseBodyDataScRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The path of the canary release rule of the Spring Cloud application.
	//
	// example:
	//
	// /path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataScRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataScRules) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) GetCondition() *string {
	return s.Condition
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) GetItems() []*DescribeGreyTagRouteResponseBodyDataScRulesItems {
	return s.Items
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) GetPath() *string {
	return s.Path
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) SetCondition(v string) *DescribeGreyTagRouteResponseBodyDataScRules {
	s.Condition = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) SetItems(v []*DescribeGreyTagRouteResponseBodyDataScRulesItems) *DescribeGreyTagRouteResponseBodyDataScRules {
	s.Items = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) SetPath(v string) *DescribeGreyTagRouteResponseBodyDataScRules {
	s.Path = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) Validate() error {
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

type DescribeGreyTagRouteResponseBodyDataScRulesItems struct {
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

func (s DescribeGreyTagRouteResponseBodyDataScRulesItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataScRulesItems) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetCond() *string {
	return s.Cond
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetExpr() *string {
	return s.Expr
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetName() *string {
	return s.Name
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetOperator() *string {
	return s.Operator
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetType() *string {
	return s.Type
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) GetValue() *string {
	return s.Value
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetCond(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Cond = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetExpr(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Expr = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetIndex(v int32) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Index = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetName(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetOperator(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Operator = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetType(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Type = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetValue(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Value = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) Validate() error {
	return dara.Validate(s)
}
