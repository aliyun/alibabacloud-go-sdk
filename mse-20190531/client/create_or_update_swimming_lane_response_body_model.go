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
	// The status code. The value 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	//
	// example:
	//
	// {}
	Data *CreateOrUpdateSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
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
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneResponseBodyData struct {
	Enable                          *bool                                                   `json:"enable,omitempty" xml:"enable,omitempty"`
	EnableRules                     *bool                                                   `json:"enableRules,omitempty" xml:"enableRules,omitempty"`
	EntryRule                       *string                                                 `json:"entryRule,omitempty" xml:"entryRule,omitempty"`
	EntryRules                      []*CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules `json:"entryRules,omitempty" xml:"entryRules,omitempty" type:"Repeated"`
	GatewaySwimmingLaneRouteJson    *string                                                 `json:"gatewaySwimmingLaneRouteJson,omitempty" xml:"gatewaySwimmingLaneRouteJson,omitempty"`
	GmtCreate                       *string                                                 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupId                         *int64                                                  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id                              *int64                                                  `json:"id,omitempty" xml:"id,omitempty"`
	Name                            *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	PathIndependentPercentageEnable *bool                                                   `json:"pathIndependentPercentageEnable,omitempty" xml:"pathIndependentPercentageEnable,omitempty"`
	RegionId                        *string                                                 `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status                          *int32                                                  `json:"status,omitempty" xml:"status,omitempty"`
	Tag                             *string                                                 `json:"tag,omitempty" xml:"tag,omitempty"`
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
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneResponseBodyDataEntryRules struct {
	Condition *string                                                          `json:"condition,omitempty" xml:"condition,omitempty"`
	Path      *string                                                          `json:"path,omitempty" xml:"path,omitempty"`
	Paths     []*string                                                        `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateOrUpdateSwimmingLaneResponseBodyDataEntryRulesRestItems struct {
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
