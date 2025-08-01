// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetAcceptLanguage() *string
	SetEnable(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetEnable() *bool
	SetEnableRules(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetEnableRules() *bool
	SetEntryRule(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetEntryRule() *string
	SetEntryRules(v []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetEntryRules() []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRules
	SetGatewaySwimmingLaneRouteJsonShrink(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetGatewaySwimmingLaneRouteJsonShrink() *string
	SetGroupId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetGroupId() *int64
	SetId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetId() *int64
	SetName(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetName() *string
	SetNamespace(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetNamespace() *string
	SetPathIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetPathIndependentPercentageEnable() *bool
	SetRegionId(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetRegionId() *string
	SetTag(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetTag() *string
}

type CreateOrUpdateSwimmingLaneShrinkRequest struct {
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
	EntryRule  *string                                              `json:"EntryRule,omitempty" xml:"EntryRule,omitempty"`
	EntryRules []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRules `json:"EntryRules,omitempty" xml:"EntryRules,omitempty" type:"Repeated"`
	// The information about the routing rule for the gateway. This parameter is required when a cloud-native gateway is used as the ingress.
	GatewaySwimmingLaneRouteJsonShrink *string `json:"GatewaySwimmingLaneRouteJson,omitempty" xml:"GatewaySwimmingLaneRouteJson,omitempty"`
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

func (s CreateOrUpdateSwimmingLaneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetEntryRule() *string {
	return s.EntryRule
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetEntryRules() []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRules {
	return s.EntryRules
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetGatewaySwimmingLaneRouteJsonShrink() *string {
	return s.GatewaySwimmingLaneRouteJsonShrink
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetPathIndependentPercentageEnable() *bool {
	return s.PathIndependentPercentageEnable
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetEnable(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetEnableRules(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.EnableRules = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetEntryRule(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.EntryRule = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetEntryRules(v []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.EntryRules = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetGatewaySwimmingLaneRouteJsonShrink(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.GatewaySwimmingLaneRouteJsonShrink = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetName(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetNamespace(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetPathIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.PathIndependentPercentageEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetRegionId(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetTag(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.Tag = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneShrinkRequestEntryRules struct {
	// example:
	//
	// AND
	Condition *string                                                       `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Paths     []*string                                                     `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	Priority  *int32                                                        `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RestItems []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems `json:"RestItems,omitempty" xml:"RestItems,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) GetPaths() []*string {
	return s.Paths
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) GetRestItems() []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	return s.RestItems
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) SetCondition(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) SetPaths(v []*string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules {
	s.Paths = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) SetPriority(v int32) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules {
	s.Priority = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) SetRestItems(v []*CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules {
	s.RestItems = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRules) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems struct {
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

func (s CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetCond() *string {
	return s.Cond
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetDatum() *string {
	return s.Datum
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetOperator() *string {
	return s.Operator
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetCond(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Cond = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetDatum(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Datum = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetDivisor(v int32) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Divisor = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetName(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetNameList(v []*string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.NameList = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetOperator(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetRate(v int32) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Rate = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetRemainder(v int32) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Remainder = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetType(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) SetValue(v string) *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequestEntryRulesRestItems) Validate() error {
	return dara.Validate(s)
}
