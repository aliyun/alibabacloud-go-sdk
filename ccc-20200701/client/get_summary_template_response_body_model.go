// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSummaryTemplateResponseBody
	GetCode() *string
	SetData(v *GetSummaryTemplateResponseBodyData) *GetSummaryTemplateResponseBody
	GetData() *GetSummaryTemplateResponseBodyData
	SetHttpStatusCode(v int32) *GetSummaryTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSummaryTemplateResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetSummaryTemplateResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetSummaryTemplateResponseBody
	GetRequestId() *string
}

type GetSummaryTemplateResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSummaryTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 30C7D235-DDCF-4C7F-A462-5E2598252C2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSummaryTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSummaryTemplateResponseBody) GetData() *GetSummaryTemplateResponseBodyData {
	return s.Data
}

func (s *GetSummaryTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSummaryTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSummaryTemplateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetSummaryTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSummaryTemplateResponseBody) SetCode(v string) *GetSummaryTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetSummaryTemplateResponseBody) SetData(v *GetSummaryTemplateResponseBodyData) *GetSummaryTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetSummaryTemplateResponseBody) SetHttpStatusCode(v int32) *GetSummaryTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSummaryTemplateResponseBody) SetMessage(v string) *GetSummaryTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetSummaryTemplateResponseBody) SetParams(v []*string) *GetSummaryTemplateResponseBody {
	s.Params = v
	return s
}

func (s *GetSummaryTemplateResponseBody) SetRequestId(v string) *GetSummaryTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSummaryTemplateResponseBodyData struct {
	// example:
	//
	// 8939-4223-86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// editor-xxx@ccc-test
	Editor *string `json:"Editor,omitempty" xml:"Editor,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId   *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyList []*GetSummaryTemplateResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 43c2671b-8939-4223-86d0-6bd187905cc8_1717664210492
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetSummaryTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSummaryTemplateResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetSummaryTemplateResponseBodyData) GetEditor() *string {
	return s.Editor
}

func (s *GetSummaryTemplateResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSummaryTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSummaryTemplateResponseBodyData) GetPropertyList() []*GetSummaryTemplateResponseBodyDataPropertyList {
	return s.PropertyList
}

func (s *GetSummaryTemplateResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetSummaryTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetSummaryTemplateResponseBodyData) SetCategoryId(v string) *GetSummaryTemplateResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) SetEditor(v string) *GetSummaryTemplateResponseBodyData {
	s.Editor = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) SetInstanceId(v string) *GetSummaryTemplateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) SetName(v string) *GetSummaryTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) SetPropertyList(v []*GetSummaryTemplateResponseBodyDataPropertyList) *GetSummaryTemplateResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) SetState(v string) *GetSummaryTemplateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) SetTemplateId(v string) *GetSummaryTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSummaryTemplateResponseBodyDataPropertyList struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// {}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// 1717664210000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// cretor-xxx@ccc-test
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// Description-xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// DisplayName-A
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
	// Name-A
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
	// 1717664210000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetSummaryTemplateResponseBodyDataPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTemplateResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetArray() *bool {
	return s.Array
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetAttributes() *string {
	return s.Attributes
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetCreator() *string {
	return s.Creator
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetDescription() *string {
	return s.Description
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetEditorType() *string {
	return s.EditorType
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetMaximum() *float64 {
	return s.Maximum
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetMinLength() *int32 {
	return s.MinLength
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetMinimum() *float64 {
	return s.Minimum
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetName() *string {
	return s.Name
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetPattern() *string {
	return s.Pattern
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetRequired() *bool {
	return s.Required
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetSystem() *bool {
	return s.System
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetArray(v bool) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Array = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetAttributes(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Attributes = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetCreatedTime(v int64) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.CreatedTime = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetCreator(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Creator = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetDataType(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.DataType = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetDescription(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Description = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetDisabled(v bool) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Disabled = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetDisplayName(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.DisplayName = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetDisplayOrder(v int32) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.DisplayOrder = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetEditorType(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.EditorType = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetMaxLength(v int32) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.MaxLength = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetMaximum(v float64) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Maximum = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetMinLength(v int32) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.MinLength = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetMinimum(v float64) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Minimum = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetName(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetPattern(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Pattern = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetPatternErrorMessage(v string) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.PatternErrorMessage = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetReadOnly(v bool) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.ReadOnly = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetRequired(v bool) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.Required = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetSystem(v bool) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.System = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) SetUpdatedTime(v int64) *GetSummaryTemplateResponseBodyDataPropertyList {
	s.UpdatedTime = &v
	return s
}

func (s *GetSummaryTemplateResponseBodyDataPropertyList) Validate() error {
	return dara.Validate(s)
}
