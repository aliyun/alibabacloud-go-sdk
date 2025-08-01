// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySwimmingLaneByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QuerySwimmingLaneByIdResponseBodyData) *QuerySwimmingLaneByIdResponseBody
	GetData() *QuerySwimmingLaneByIdResponseBodyData
	SetErrorCode(v string) *QuerySwimmingLaneByIdResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QuerySwimmingLaneByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySwimmingLaneByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySwimmingLaneByIdResponseBody
	GetSuccess() *bool
}

type QuerySwimmingLaneByIdResponseBody struct {
	// The details of the data.
	//
	// example:
	//
	// {id:102,name:"test"}
	Data *QuerySwimmingLaneByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
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

func (s QuerySwimmingLaneByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySwimmingLaneByIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySwimmingLaneByIdResponseBody) GetData() *QuerySwimmingLaneByIdResponseBodyData {
	return s.Data
}

func (s *QuerySwimmingLaneByIdResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QuerySwimmingLaneByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySwimmingLaneByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySwimmingLaneByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySwimmingLaneByIdResponseBody) SetData(v *QuerySwimmingLaneByIdResponseBodyData) *QuerySwimmingLaneByIdResponseBody {
	s.Data = v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBody) SetErrorCode(v string) *QuerySwimmingLaneByIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBody) SetMessage(v string) *QuerySwimmingLaneByIdResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBody) SetRequestId(v string) *QuerySwimmingLaneByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBody) SetSuccess(v bool) *QuerySwimmingLaneByIdResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySwimmingLaneByIdResponseBodyData struct {
	PathIndependentPercentageEnable *bool                                              `json:"PathIndependentPercentageEnable,omitempty" xml:"PathIndependentPercentageEnable,omitempty"`
	Enable                          *bool                                              `json:"enable,omitempty" xml:"enable,omitempty"`
	EnableRules                     *bool                                              `json:"enableRules,omitempty" xml:"enableRules,omitempty"`
	EntryRule                       *string                                            `json:"entryRule,omitempty" xml:"entryRule,omitempty"`
	EntryRules                      []*QuerySwimmingLaneByIdResponseBodyDataEntryRules `json:"entryRules,omitempty" xml:"entryRules,omitempty" type:"Repeated"`
	GatewaySwimmingLaneRouteJson    *string                                            `json:"gatewaySwimmingLaneRouteJson,omitempty" xml:"gatewaySwimmingLaneRouteJson,omitempty"`
	GmtCreate                       *string                                            `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                            `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupId                         *int64                                             `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id                              *int64                                             `json:"id,omitempty" xml:"id,omitempty"`
	Name                            *string                                            `json:"name,omitempty" xml:"name,omitempty"`
	RegionId                        *string                                            `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status                          *int32                                             `json:"status,omitempty" xml:"status,omitempty"`
	Tag                             *string                                            `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s QuerySwimmingLaneByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySwimmingLaneByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetPathIndependentPercentageEnable() *bool {
	return s.PathIndependentPercentageEnable
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetEntryRule() *string {
	return s.EntryRule
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetEntryRules() []*QuerySwimmingLaneByIdResponseBodyDataEntryRules {
	return s.EntryRules
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetGatewaySwimmingLaneRouteJson() *string {
	return s.GatewaySwimmingLaneRouteJson
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *QuerySwimmingLaneByIdResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetPathIndependentPercentageEnable(v bool) *QuerySwimmingLaneByIdResponseBodyData {
	s.PathIndependentPercentageEnable = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetEnable(v bool) *QuerySwimmingLaneByIdResponseBodyData {
	s.Enable = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetEnableRules(v bool) *QuerySwimmingLaneByIdResponseBodyData {
	s.EnableRules = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetEntryRule(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.EntryRule = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetEntryRules(v []*QuerySwimmingLaneByIdResponseBodyDataEntryRules) *QuerySwimmingLaneByIdResponseBodyData {
	s.EntryRules = v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetGatewaySwimmingLaneRouteJson(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.GatewaySwimmingLaneRouteJson = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetGmtCreate(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetGmtModified(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetGroupId(v int64) *QuerySwimmingLaneByIdResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetId(v int64) *QuerySwimmingLaneByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetName(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetRegionId(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetStatus(v int32) *QuerySwimmingLaneByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) SetTag(v string) *QuerySwimmingLaneByIdResponseBodyData {
	s.Tag = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QuerySwimmingLaneByIdResponseBodyDataEntryRules struct {
	Condition *string                                                     `json:"condition,omitempty" xml:"condition,omitempty"`
	Path      *string                                                     `json:"path,omitempty" xml:"path,omitempty"`
	Paths     []*string                                                   `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	RestItems []*QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems `json:"restItems,omitempty" xml:"restItems,omitempty" type:"Repeated"`
}

func (s QuerySwimmingLaneByIdResponseBodyDataEntryRules) String() string {
	return dara.Prettify(s)
}

func (s QuerySwimmingLaneByIdResponseBodyDataEntryRules) GoString() string {
	return s.String()
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) GetCondition() *string {
	return s.Condition
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) GetPath() *string {
	return s.Path
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) GetPaths() []*string {
	return s.Paths
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) GetRestItems() []*QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	return s.RestItems
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) SetCondition(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRules {
	s.Condition = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) SetPath(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRules {
	s.Path = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) SetPaths(v []*string) *QuerySwimmingLaneByIdResponseBodyDataEntryRules {
	s.Paths = v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) SetRestItems(v []*QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) *QuerySwimmingLaneByIdResponseBodyDataEntryRules {
	s.RestItems = v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRules) Validate() error {
	return dara.Validate(s)
}

type QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems struct {
	Cond      *string   `json:"cond,omitempty" xml:"cond,omitempty"`
	Datum     *string   `json:"datum,omitempty" xml:"datum,omitempty"`
	Divisor   *int32    `json:"divisor,omitempty" xml:"divisor,omitempty"`
	Name      *string   `json:"name,omitempty" xml:"name,omitempty"`
	NameList  []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	Operator  *string   `json:"operator,omitempty" xml:"operator,omitempty"`
	Rate      *int32    `json:"rate,omitempty" xml:"rate,omitempty"`
	Remainder *int32    `json:"remainder,omitempty" xml:"remainder,omitempty"`
	Type      *string   `json:"type,omitempty" xml:"type,omitempty"`
	Value     *string   `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GoString() string {
	return s.String()
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetCond() *string {
	return s.Cond
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetDatum() *string {
	return s.Datum
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetName() *string {
	return s.Name
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetOperator() *string {
	return s.Operator
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetType() *string {
	return s.Type
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) GetValue() *string {
	return s.Value
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetCond(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Cond = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetDatum(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Datum = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetDivisor(v int32) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Divisor = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetName(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Name = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetNameList(v []*string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.NameList = v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetOperator(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Operator = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetRate(v int32) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Rate = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetRemainder(v int32) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Remainder = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetType(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Type = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) SetValue(v string) *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems {
	s.Value = &v
	return s
}

func (s *QuerySwimmingLaneByIdResponseBodyDataEntryRulesRestItems) Validate() error {
	return dara.Validate(s)
}
