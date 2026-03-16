// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryAllSwimmingLaneResponseBodyData) *QueryAllSwimmingLaneResponseBody
	GetData() []*QueryAllSwimmingLaneResponseBodyData
	SetErrorCode(v string) *QueryAllSwimmingLaneResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryAllSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAllSwimmingLaneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAllSwimmingLaneResponseBody
	GetSuccess() *bool
}

type QueryAllSwimmingLaneResponseBody struct {
	// The details of the data.
	//
	// example:
	//
	// [{id:100,name:"test"}]
	Data []*QueryAllSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAllSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBody) GetData() []*QueryAllSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *QueryAllSwimmingLaneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryAllSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAllSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAllSwimmingLaneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAllSwimmingLaneResponseBody) SetData(v []*QueryAllSwimmingLaneResponseBodyData) *QueryAllSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBody) SetErrorCode(v string) *QueryAllSwimmingLaneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBody) SetMessage(v string) *QueryAllSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBody) SetRequestId(v string) *QueryAllSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBody) SetSuccess(v bool) *QueryAllSwimmingLaneResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAllSwimmingLaneResponseBodyData struct {
	// 是否开启。
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ingress application rules.
	EntryRules []*QueryAllSwimmingLaneResponseBodyDataEntryRules `json:"EntryRules,omitempty" xml:"EntryRules,omitempty" type:"Repeated"`
	// 网关路由信息（使用云原生网关为入口时使用）
	GatewaySwimmingLaneRoute *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute `json:"GatewaySwimmingLaneRoute,omitempty" xml:"GatewaySwimmingLaneRoute,omitempty" type:"Struct"`
	// example:
	//
	// {\\"GatewayUniqueId\\":\\"gw-cf815503e0a8441b8956d8e349fa8fb5\\",\\"Percentage\\":0,\\"RouteIdList\\":[32295,32298,32297,32309,32307,32294,32303,32292,32299,35278],\\"CanaryModel\\":1,\\"Conditions\\":[]}
	GatewaySwimmingLaneRouteJson *string `json:"GatewaySwimmingLaneRouteJson,omitempty" xml:"GatewaySwimmingLaneRouteJson,omitempty"`
	// The ID of the lane group.
	//
	// example:
	//
	// 6859
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the lane group.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The side on which canary release for messaging is implemented.
	//
	// example:
	//
	// Client
	MessageQueueFilterSide *string `json:"MessageQueueFilterSide,omitempty" xml:"MessageQueueFilterSide,omitempty"`
	// Indicates whether canary release for messaging is enabled.
	//
	// example:
	//
	// false
	MessageQueueGrayEnable *bool `json:"MessageQueueGrayEnable,omitempty" xml:"MessageQueueGrayEnable,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// swimmingGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Microservices Engine (MSE) namespace to which the lane belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// true
	PathIndependentPercentageEnable *bool `json:"PathIndependentPercentageEnable,omitempty" xml:"PathIndependentPercentageEnable,omitempty"`
	// Indicates whether request details were recorded (metrics of end-to-end canary release).
	//
	// example:
	//
	// false
	RecordCanaryDetail *bool `json:"RecordCanaryDetail,omitempty" xml:"RecordCanaryDetail,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag of the lane.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The user ID of the lane group.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"enableRules,omitempty" xml:"enableRules,omitempty"`
	// The time when the lane was created.
	//
	// example:
	//
	// 1683598484000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the lane was updated.
	//
	// example:
	//
	// 1683598484000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
}

func (s QueryAllSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetEnable() *string {
	return s.Enable
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetEntryRules() []*QueryAllSwimmingLaneResponseBodyDataEntryRules {
	return s.EntryRules
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetGatewaySwimmingLaneRoute() *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	return s.GatewaySwimmingLaneRoute
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetGatewaySwimmingLaneRouteJson() *string {
	return s.GatewaySwimmingLaneRouteJson
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetMessageQueueFilterSide() *string {
	return s.MessageQueueFilterSide
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetMessageQueueGrayEnable() *bool {
	return s.MessageQueueGrayEnable
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetPathIndependentPercentageEnable() *bool {
	return s.PathIndependentPercentageEnable
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetRecordCanaryDetail() *bool {
	return s.RecordCanaryDetail
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryAllSwimmingLaneResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetEnable(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.Enable = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetEntryRules(v []*QueryAllSwimmingLaneResponseBodyDataEntryRules) *QueryAllSwimmingLaneResponseBodyData {
	s.EntryRules = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetGatewaySwimmingLaneRoute(v *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) *QueryAllSwimmingLaneResponseBodyData {
	s.GatewaySwimmingLaneRoute = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetGatewaySwimmingLaneRouteJson(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.GatewaySwimmingLaneRouteJson = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetGroupId(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetId(v int64) *QueryAllSwimmingLaneResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetMessageQueueFilterSide(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.MessageQueueFilterSide = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetMessageQueueGrayEnable(v bool) *QueryAllSwimmingLaneResponseBodyData {
	s.MessageQueueGrayEnable = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetName(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetNamespace(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetPathIndependentPercentageEnable(v bool) *QueryAllSwimmingLaneResponseBodyData {
	s.PathIndependentPercentageEnable = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetRecordCanaryDetail(v bool) *QueryAllSwimmingLaneResponseBodyData {
	s.RecordCanaryDetail = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetRegionId(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetTag(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.Tag = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetUserId(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetEnableRules(v bool) *QueryAllSwimmingLaneResponseBodyData {
	s.EnableRules = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetGmtCreate(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) SetGmtModified(v string) *QueryAllSwimmingLaneResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyData) Validate() error {
	if s.EntryRules != nil {
		for _, item := range s.EntryRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GatewaySwimmingLaneRoute != nil {
		if err := s.GatewaySwimmingLaneRoute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAllSwimmingLaneResponseBodyDataEntryRules struct {
	// The condition.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The path.
	//
	// example:
	//
	// /adump/wxb/prod
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The paths.
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The RESTful matching rules.
	RestItems []*QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems `json:"restItems,omitempty" xml:"restItems,omitempty" type:"Repeated"`
}

func (s QueryAllSwimmingLaneResponseBodyDataEntryRules) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBodyDataEntryRules) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) GetCondition() *string {
	return s.Condition
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) GetPath() *string {
	return s.Path
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) GetPaths() []*string {
	return s.Paths
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) GetRestItems() []*QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	return s.RestItems
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) SetCondition(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRules {
	s.Condition = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) SetPath(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRules {
	s.Path = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) SetPaths(v []*string) *QueryAllSwimmingLaneResponseBodyDataEntryRules {
	s.Paths = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) SetRestItems(v []*QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) *QueryAllSwimmingLaneResponseBodyDataEntryRules {
	s.RestItems = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRules) Validate() error {
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

type QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems struct {
	// The matching character.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The value.
	//
	// example:
	//
	// x
	Datum *string `json:"datum,omitempty" xml:"datum,omitempty"`
	// The divisor.
	//
	// example:
	//
	// 10
	Divisor *int32 `json:"divisor,omitempty" xml:"divisor,omitempty"`
	// The name of the matching rule.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The matching list.
	NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// The operator.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The percentage.
	//
	// example:
	//
	// 30
	Rate *int32 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The remainder.
	//
	// example:
	//
	// 10
	Remainder *int32 `json:"remainder,omitempty" xml:"remainder,omitempty"`
	// The matching type.
	//
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetCond() *string {
	return s.Cond
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetDatum() *string {
	return s.Datum
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetName() *string {
	return s.Name
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetOperator() *string {
	return s.Operator
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetType() *string {
	return s.Type
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) GetValue() *string {
	return s.Value
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetCond(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Cond = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetDatum(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Datum = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetDivisor(v int32) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Divisor = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetName(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Name = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetNameList(v []*string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.NameList = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetOperator(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Operator = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetRate(v int32) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Rate = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetRemainder(v int32) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Remainder = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetType(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Type = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) SetValue(v string) *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Value = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataEntryRulesRestItems) Validate() error {
	return dara.Validate(s)
}

type QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute struct {
	// 全链路灰度模式：
	//
	// - 0: 按请求内容路由
	//
	// - 1: 按比例路由
	//
	// example:
	//
	// 0
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// example:
	//
	// AND、OR
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// 匹配条件列表（按内容路由使用）
	Conditions []*QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// 网关ID
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 网关唯一ID
	//
	// example:
	//
	// gw-84efde2ee1464260bdb17a5b****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// 流量比例（按比例路由使用，取值0-100）
	//
	// example:
	//
	// 20
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// 路由ID列表
	RouteIdList []*int64 `json:"RouteIdList,omitempty" xml:"RouteIdList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	RouteIndependentPercentageEnable *string                                                                                       `json:"RouteIndependentPercentageEnable,omitempty" xml:"RouteIndependentPercentageEnable,omitempty"`
	RouteIndependentPercentageList   []*QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList `json:"RouteIndependentPercentageList,omitempty" xml:"RouteIndependentPercentageList,omitempty" type:"Repeated"`
}

func (s QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetCondition() *string {
	return s.Condition
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetConditions() []*QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions {
	return s.Conditions
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetPercentage() *int32 {
	return s.Percentage
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetRouteIdList() []*int64 {
	return s.RouteIdList
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetRouteIndependentPercentageEnable() *string {
	return s.RouteIndependentPercentageEnable
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) GetRouteIndependentPercentageList() []*QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList {
	return s.RouteIndependentPercentageList
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetCanaryModel(v int32) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.CanaryModel = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetCondition(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.Condition = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetConditions(v []*QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.Conditions = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetGatewayId(v int64) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.GatewayId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetGatewayUniqueId(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.GatewayUniqueId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetPercentage(v int32) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.Percentage = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetRouteIdList(v []*int64) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.RouteIdList = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetRouteIndependentPercentageEnable(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.RouteIndependentPercentageEnable = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) SetRouteIndependentPercentageList(v []*QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute {
	s.RouteIndependentPercentageList = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRoute) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RouteIndependentPercentageList != nil {
		for _, item := range s.RouteIndependentPercentageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions struct {
	// 条件：
	//
	// - PRE：前缀匹配
	//
	// - EQUAL：精确匹配
	//
	// - ERGULAR：正则匹配
	//
	// example:
	//
	// PRE
	Cond *string `json:"Cond,omitempty" xml:"Cond,omitempty"`
	// 参数名称
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// NameList
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
	// 参数类型：
	//
	// - header：Header
	//
	// - param：Parameter
	//
	// example:
	//
	// header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 参数值
	//
	// example:
	//
	// xiaoming
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) GetCond() *string {
	return s.Cond
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) GetName() *string {
	return s.Name
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) GetNameList() []*string {
	return s.NameList
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) GetType() *string {
	return s.Type
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) GetValue() *string {
	return s.Value
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) SetCond(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions {
	s.Cond = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) SetName(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions {
	s.Name = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) SetNameList(v []*string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions {
	s.NameList = v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) SetType(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions {
	s.Type = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) SetValue(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions {
	s.Value = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteConditions) Validate() error {
	return dara.Validate(s)
}

type QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList struct {
	// example:
	//
	// 60
	Percentage *string `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// 12345
	RouteId *string `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) GetPercentage() *string {
	return s.Percentage
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) GetRouteId() *string {
	return s.RouteId
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) SetPercentage(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList {
	s.Percentage = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) SetRouteId(v string) *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList {
	s.RouteId = &v
	return s
}

func (s *QueryAllSwimmingLaneResponseBodyDataGatewaySwimmingLaneRouteRouteIndependentPercentageList) Validate() error {
	return dara.Validate(s)
}
