// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneRequest
	GetAcceptLanguage() *string
	SetEnable(v bool) *CreateOrUpdateSwimmingLaneRequest
	GetEnable() *bool
	SetEnableRules(v bool) *CreateOrUpdateSwimmingLaneRequest
	GetEnableRules() *bool
	SetEntryRule(v string) *CreateOrUpdateSwimmingLaneRequest
	GetEntryRule() *string
	SetEntryRules(v []*CreateOrUpdateSwimmingLaneRequestEntryRules) *CreateOrUpdateSwimmingLaneRequest
	GetEntryRules() []*CreateOrUpdateSwimmingLaneRequestEntryRules
	SetGatewaySwimmingLaneRouteJson(v *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) *CreateOrUpdateSwimmingLaneRequest
	GetGatewaySwimmingLaneRouteJson() *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson
	SetGroupId(v int64) *CreateOrUpdateSwimmingLaneRequest
	GetGroupId() *int64
	SetId(v int64) *CreateOrUpdateSwimmingLaneRequest
	GetId() *int64
	SetName(v string) *CreateOrUpdateSwimmingLaneRequest
	GetName() *string
	SetNamespace(v string) *CreateOrUpdateSwimmingLaneRequest
	GetNamespace() *string
	SetPathIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneRequest
	GetPathIndependentPercentageEnable() *bool
	SetRegionId(v string) *CreateOrUpdateSwimmingLaneRequest
	GetRegionId() *string
	SetTag(v string) *CreateOrUpdateSwimmingLaneRequest
	GetTag() *string
}

type CreateOrUpdateSwimmingLaneRequest struct {
	// The language of the response. Valid values: zh and en. Default value: zh. The value zh indicates Chinese, and the value en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to enable the lane.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Specifies whether to configure a routing rule for the lane. If an Ingress gateway is used, this parameter is not required.
	//
	// example:
	//
	// false
	EnableRules *bool `json:"EnableRules,omitempty" xml:"EnableRules,omitempty"`
	// Deprecated
	//
	// The JSON string.
	//
	// example:
	//
	// {}
	EntryRule  *string                                        `json:"EntryRule,omitempty" xml:"EntryRule,omitempty"`
	EntryRules []*CreateOrUpdateSwimmingLaneRequestEntryRules `json:"EntryRules,omitempty" xml:"EntryRules,omitempty" type:"Repeated"`
	// The information about the routing rule for the gateway. This parameter is required when a cloud-native gateway is used as the ingress.
	GatewaySwimmingLaneRouteJson *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson `json:"GatewaySwimmingLaneRouteJson,omitempty" xml:"GatewaySwimmingLaneRouteJson,omitempty" type:"Struct"`
	// The language of the response. Valid values:****
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US**: English
	//
	// > Default value: **zh-CN**.
	//
	// example:
	//
	// 115
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the primary key. The value -1 indicates a request that is used to create a lane. A value greater than 0 indicates a request that is used to modify a lane.
	//
	// example:
	//
	// -1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test lane
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	Namespace                       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PathIndependentPercentageEnable *bool   `json:"PathIndependentPercentageEnable,omitempty" xml:"PathIndependentPercentageEnable,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetEntryRule() *string {
	return s.EntryRule
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetEntryRules() []*CreateOrUpdateSwimmingLaneRequestEntryRules {
	return s.EntryRules
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetGatewaySwimmingLaneRouteJson() *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	return s.GatewaySwimmingLaneRouteJson
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetPathIndependentPercentageEnable() *bool {
	return s.PathIndependentPercentageEnable
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateSwimmingLaneRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetEnable(v bool) *CreateOrUpdateSwimmingLaneRequest {
	s.Enable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetEnableRules(v bool) *CreateOrUpdateSwimmingLaneRequest {
	s.EnableRules = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetEntryRule(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.EntryRule = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetEntryRules(v []*CreateOrUpdateSwimmingLaneRequestEntryRules) *CreateOrUpdateSwimmingLaneRequest {
	s.EntryRules = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetGatewaySwimmingLaneRouteJson(v *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) *CreateOrUpdateSwimmingLaneRequest {
	s.GatewaySwimmingLaneRouteJson = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetId(v int64) *CreateOrUpdateSwimmingLaneRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetName(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetNamespace(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.Namespace = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetPathIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneRequest {
	s.PathIndependentPercentageEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetRegionId(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) SetTag(v string) *CreateOrUpdateSwimmingLaneRequest {
	s.Tag = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestEntryRules struct {
	// example:
	//
	// AND
	Condition *string                                                 `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Paths     []*string                                               `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	Priority  *int32                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RestItems []*CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems `json:"RestItems,omitempty" xml:"RestItems,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSwimmingLaneRequestEntryRules) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestEntryRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) GetPaths() []*string {
	return s.Paths
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) GetRestItems() []*CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	return s.RestItems
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) SetCondition(v string) *CreateOrUpdateSwimmingLaneRequestEntryRules {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) SetPaths(v []*string) *CreateOrUpdateSwimmingLaneRequestEntryRules {
	s.Paths = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) SetPriority(v int32) *CreateOrUpdateSwimmingLaneRequestEntryRules {
	s.Priority = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) SetRestItems(v []*CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) *CreateOrUpdateSwimmingLaneRequestEntryRules {
	s.RestItems = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRules) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems struct {
	Cond      *string   `json:"Cond,omitempty" xml:"Cond,omitempty"`
	Datum     *string   `json:"Datum,omitempty" xml:"Datum,omitempty"`
	Divisor   *int32    `json:"Divisor,omitempty" xml:"Divisor,omitempty"`
	Name      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NameList  []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
	Operator  *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Rate      *int32    `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Remainder *int32    `json:"Remainder,omitempty" xml:"Remainder,omitempty"`
	Type      *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string   `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetCond() *string {
	return s.Cond
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetDatum() *string {
	return s.Datum
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetOperator() *string {
	return s.Operator
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetCond(v string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Cond = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetDatum(v string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Datum = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetDivisor(v int32) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Divisor = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetName(v string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetNameList(v []*string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.NameList = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetOperator(v string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetRate(v int32) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Rate = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetRemainder(v int32) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Remainder = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetType(v string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) SetValue(v string) *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestEntryRulesRestItems) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson struct {
	// example:
	//
	// 0
	CanaryModel *int32  `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	Condition   *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The matching conditions.
	Conditions []*CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-84efde2ee1464260bdb17a5b****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 20
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The route IDs.
	RouteIdList                      []*int64                                                                                       `json:"RouteIdList,omitempty" xml:"RouteIdList,omitempty" type:"Repeated"`
	RouteIndependentPercentageEnable *bool                                                                                          `json:"RouteIndependentPercentageEnable,omitempty" xml:"RouteIndependentPercentageEnable,omitempty"`
	RouteIndependentPercentageList   []*CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList `json:"RouteIndependentPercentageList,omitempty" xml:"RouteIndependentPercentageList,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetConditions() []*CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions {
	return s.Conditions
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetRouteIdList() []*int64 {
	return s.RouteIdList
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetRouteIndependentPercentageEnable() *bool {
	return s.RouteIndependentPercentageEnable
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) GetRouteIndependentPercentageList() []*CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList {
	return s.RouteIndependentPercentageList
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.CanaryModel = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetCondition(v string) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetConditions(v []*CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.Conditions = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetGatewayId(v int64) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.GatewayId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetGatewayUniqueId(v string) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.GatewayUniqueId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetPercentage(v int32) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.Percentage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetRouteIdList(v []*int64) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.RouteIdList = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetRouteIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.RouteIndependentPercentageEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) SetRouteIndependentPercentageList(v []*CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson {
	s.RouteIndependentPercentageList = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJson) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions struct {
	// The matching condition. Valid values:
	//
	// 	- PRE: prefix matching
	//
	// 	- EQUAL: exact matching
	//
	// 	- ERGULAR: regular expression matching
	//
	// example:
	//
	// PRE
	Cond *string `json:"Cond,omitempty" xml:"Cond,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the parameter. Valid values:
	//
	// 	- header
	//
	// 	- param
	//
	// example:
	//
	// header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// xiaoming
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) GetCond() *string {
	return s.Cond
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) SetCond(v string) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions {
	s.Cond = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) SetName(v string) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) SetType(v string) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) SetValue(v string) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonConditions) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList struct {
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	RouteId    *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) GetRouteId() *int64 {
	return s.RouteId
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) SetPercentage(v int32) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList {
	s.Percentage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) SetRouteId(v int64) *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList {
	s.RouteId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneRequestGatewaySwimmingLaneRouteJsonRouteIndependentPercentageList) Validate() error {
	return dara.Validate(s)
}
