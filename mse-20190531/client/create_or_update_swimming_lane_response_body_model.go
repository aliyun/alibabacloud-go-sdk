// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateOrUpdateSwimmingLaneResponseBody
	GetCode() *int32
	SetData(v *CreateOrUpdateSwimmingLaneResponseBodyData) *CreateOrUpdateSwimmingLaneResponseBody
	GetData() *CreateOrUpdateSwimmingLaneResponseBodyData
	SetErrorCode(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *CreateOrUpdateSwimmingLaneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrUpdateSwimmingLaneResponseBody
	GetSuccess() *bool
}

type CreateOrUpdateSwimmingLaneResponseBody struct {
	// The code used only to synchronize with the Success parameter.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data overview.
	//
	// example:
	//
	// {}
	Data *CreateOrUpdateSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Succ
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
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

func (s CreateOrUpdateSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetData() *CreateOrUpdateSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetCode(v int32) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetData(v *CreateOrUpdateSwimmingLaneResponseBodyData) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetErrorCode(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetHttpStatusCode(v int32) *CreateOrUpdateSwimmingLaneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetMessage(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetRequestId(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetSuccess(v bool) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneResponseBodyData struct {
	// The swimlane switch.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Specifies whether the rules are enabled.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"enableRules,omitempty" xml:"enableRules,omitempty"`
	// The entry application rule.
	//
	// example:
	//
	// {}
	EntryRule *string `json:"entryRule,omitempty" xml:"entryRule,omitempty"`
	// The name.
	EntryRules []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules `json:"entryRules,omitempty" xml:"entryRules,omitempty" type:"Repeated"`
	// The MSE gateway routing rules.
	//
	// example:
	//
	// {}
	GatewaySwimmingLaneRouteJson *string `json:"gatewaySwimmingLaneRouteJson,omitempty" xml:"gatewaySwimmingLaneRouteJson,omitempty"`
	// The time when the swimlane was created.
	//
	// example:
	//
	// 1545726028000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the swimlane was updated.
	//
	// example:
	//
	// 1565701886000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The swimlane group ID.
	//
	// example:
	//
	// 72586
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The swimlane ID.
	//
	// example:
	//
	// 94765
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The swimlane name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// **[Deprecated]*	- Specifies whether to enable percentage-based canary releases by path when a Java gateway is used as the entry point.
	//
	// example:
	//
	// false
	PathIndependentPercentageEnable *bool `json:"pathIndependentPercentageEnable,omitempty" xml:"pathIndependentPercentageEnable,omitempty"`
	// The region where the application is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The swimlane status.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The tag corresponding to the swimlane.
	//
	// example:
	//
	// release
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetEntryRule() *string {
	return s.EntryRule
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetEntryRules() []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules {
	return s.EntryRules
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetGatewaySwimmingLaneRouteJson() *string {
	return s.GatewaySwimmingLaneRouteJson
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetPathIndependentPercentageEnable() *bool {
	return s.PathIndependentPercentageEnable
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetEnable(v bool) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetEnableRules(v bool) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.EnableRules = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetEntryRule(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.EntryRule = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetEntryRules(v []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.EntryRules = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetGatewaySwimmingLaneRouteJson(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.GatewaySwimmingLaneRouteJson = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetGmtCreate(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetGmtModified(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetId(v int64) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetName(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetPathIndependentPercentageEnable(v bool) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.PathIndependentPercentageEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetRegionId(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetStatus(v int32) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetTag(v string) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.Tag = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) Validate() error {
	if s.EntryRules != nil {
		for _, item := range s.EntryRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules struct {
	// The condition.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The code group path.
	//
	// example:
	//
	// /adump/wxb/prod
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The list of paths.
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The list of REST matching rules.
	RestItems []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems `json:"restItems,omitempty" xml:"restItems,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) GetPath() *string {
	return s.Path
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) GetPaths() []*string {
	return s.Paths
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) GetRestItems() []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	return s.RestItems
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) SetCondition(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) SetPath(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules {
	s.Path = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) SetPaths(v []*string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules {
	s.Paths = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) SetRestItems(v []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules {
	s.RestItems = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules) Validate() error {
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

type CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems struct {
	// The match operator.
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
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of names.
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
	// The match type.
	//
	// example:
	//
	// cookie
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value.
	//
	// example:
	//
	// 9
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetCond() *string {
	return s.Cond
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetDatum() *string {
	return s.Datum
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetOperator() *string {
	return s.Operator
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetType() *string {
	return s.Type
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetCond(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Cond = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetDatum(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Datum = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetDivisor(v int32) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Divisor = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetName(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetNameList(v []*string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.NameList = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetOperator(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetRate(v int32) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Rate = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetRemainder(v int32) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Remainder = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetType(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Type = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) SetValue(v string) *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems) Validate() error {
	return dara.Validate(s)
}
