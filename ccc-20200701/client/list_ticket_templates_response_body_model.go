// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTicketTemplatesResponseBody
	GetCode() *string
	SetData(v *ListTicketTemplatesResponseBodyData) *ListTicketTemplatesResponseBody
	GetData() *ListTicketTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListTicketTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTicketTemplatesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListTicketTemplatesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListTicketTemplatesResponseBody
	GetRequestId() *string
}

type ListTicketTemplatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListTicketTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 8707EB29-BAED-4302-B999-40BA61877437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTicketTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTicketTemplatesResponseBody) GetData() *ListTicketTemplatesResponseBodyData {
	return s.Data
}

func (s *ListTicketTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTicketTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketTemplatesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListTicketTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketTemplatesResponseBody) SetCode(v string) *ListTicketTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketTemplatesResponseBody) SetData(v *ListTicketTemplatesResponseBodyData) *ListTicketTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketTemplatesResponseBody) SetHttpStatusCode(v int32) *ListTicketTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTicketTemplatesResponseBody) SetMessage(v string) *ListTicketTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketTemplatesResponseBody) SetParams(v []*string) *ListTicketTemplatesResponseBody {
	s.Params = v
	return s
}

func (s *ListTicketTemplatesResponseBody) SetRequestId(v string) *ListTicketTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTicketTemplatesResponseBodyData struct {
	List []*ListTicketTemplatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTicketTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketTemplatesResponseBodyData) GetList() []*ListTicketTemplatesResponseBodyDataList {
	return s.List
}

func (s *ListTicketTemplatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTicketTemplatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketTemplatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTicketTemplatesResponseBodyData) SetList(v []*ListTicketTemplatesResponseBodyDataList) *ListTicketTemplatesResponseBodyData {
	s.List = v
	return s
}

func (s *ListTicketTemplatesResponseBodyData) SetPageNumber(v int32) *ListTicketTemplatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyData) SetPageSize(v int32) *ListTicketTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyData) SetTotalCount(v int32) *ListTicketTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTicketTemplatesResponseBodyDataList struct {
	// example:
	//
	// 0
	AppliedVersion *string `json:"AppliedVersion,omitempty" xml:"AppliedVersion,omitempty"`
	// example:
	//
	// 43c2671b-*****-4223-86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// creator@ccc-test
	Editor *string `json:"Editor,omitempty" xml:"Editor,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1715780670000
	LatestVersion     *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessDefinition *string `json:"ProcessDefinition,omitempty" xml:"ProcessDefinition,omitempty"`
	// example:
	//
	// Enabled
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// b5c21219-3a1e-4bc0-92e7-da66e057d2f6
	TemplateId   *string                                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketFields []*ListTicketTemplatesResponseBodyDataListTicketFields `json:"TicketFields,omitempty" xml:"TicketFields,omitempty" type:"Repeated"`
	// example:
	//
	// 1715780670000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListTicketTemplatesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTemplatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTicketTemplatesResponseBodyDataList) GetAppliedVersion() *string {
	return s.AppliedVersion
}

func (s *ListTicketTemplatesResponseBodyDataList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListTicketTemplatesResponseBodyDataList) GetEditor() *string {
	return s.Editor
}

func (s *ListTicketTemplatesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTicketTemplatesResponseBodyDataList) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListTicketTemplatesResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListTicketTemplatesResponseBodyDataList) GetProcessDefinition() *string {
	return s.ProcessDefinition
}

func (s *ListTicketTemplatesResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListTicketTemplatesResponseBodyDataList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTicketTemplatesResponseBodyDataList) GetTicketFields() []*ListTicketTemplatesResponseBodyDataListTicketFields {
	return s.TicketFields
}

func (s *ListTicketTemplatesResponseBodyDataList) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListTicketTemplatesResponseBodyDataList) SetAppliedVersion(v string) *ListTicketTemplatesResponseBodyDataList {
	s.AppliedVersion = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetCategoryId(v string) *ListTicketTemplatesResponseBodyDataList {
	s.CategoryId = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetEditor(v string) *ListTicketTemplatesResponseBodyDataList {
	s.Editor = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetInstanceId(v string) *ListTicketTemplatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetLatestVersion(v string) *ListTicketTemplatesResponseBodyDataList {
	s.LatestVersion = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetName(v string) *ListTicketTemplatesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetProcessDefinition(v string) *ListTicketTemplatesResponseBodyDataList {
	s.ProcessDefinition = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetState(v string) *ListTicketTemplatesResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetTemplateId(v string) *ListTicketTemplatesResponseBodyDataList {
	s.TemplateId = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetTicketFields(v []*ListTicketTemplatesResponseBodyDataListTicketFields) *ListTicketTemplatesResponseBodyDataList {
	s.TicketFields = v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) SetUpdatedTime(v int64) *ListTicketTemplatesResponseBodyDataList {
	s.UpdatedTime = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListTicketTemplatesResponseBodyDataListTicketFields struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// {}
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// example:
	//
	// 1715780670000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// creator
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// String
	DataType    *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1
	DisplayOrder *int32 `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	// example:
	//
	// textbox
	EditorType *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	// example:
	//
	// 30
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 10
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// example:
	//
	// 1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ^
	Pattern             *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// example:
	//
	// 1715780670000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListTicketTemplatesResponseBodyDataListTicketFields) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTemplatesResponseBodyDataListTicketFields) GoString() string {
	return s.String()
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetArray() *bool {
	return s.Array
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetAttribute() *string {
	return s.Attribute
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetCreator() *string {
	return s.Creator
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetDataType() *string {
	return s.DataType
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetDescription() *string {
	return s.Description
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetDisabled() *bool {
	return s.Disabled
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetEditorType() *string {
	return s.EditorType
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetMaximum() *float64 {
	return s.Maximum
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetMinLength() *int32 {
	return s.MinLength
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetMinimum() *float64 {
	return s.Minimum
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetName() *string {
	return s.Name
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetPattern() *string {
	return s.Pattern
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetRequired() *bool {
	return s.Required
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetSystem() *bool {
	return s.System
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetArray(v bool) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Array = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetAttribute(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Attribute = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetCreatedTime(v int64) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.CreatedTime = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetCreator(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Creator = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetDataType(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.DataType = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetDescription(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Description = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetDisabled(v bool) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Disabled = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetDisplayName(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.DisplayName = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetDisplayOrder(v int32) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.DisplayOrder = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetEditorType(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.EditorType = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetMaxLength(v int32) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.MaxLength = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetMaximum(v float64) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Maximum = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetMinLength(v int32) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.MinLength = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetMinimum(v float64) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Minimum = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetName(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Name = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetPattern(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Pattern = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetPatternErrorMessage(v string) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.PatternErrorMessage = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetReadOnly(v bool) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.ReadOnly = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetRequired(v bool) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.Required = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetSystem(v bool) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.System = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) SetUpdatedTime(v int64) *ListTicketTemplatesResponseBodyDataListTicketFields {
	s.UpdatedTime = &v
	return s
}

func (s *ListTicketTemplatesResponseBodyDataListTicketFields) Validate() error {
	return dara.Validate(s)
}
