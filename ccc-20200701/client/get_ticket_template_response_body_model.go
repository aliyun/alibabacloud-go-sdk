// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTicketTemplateResponseBody
	GetCode() *string
	SetData(v *GetTicketTemplateResponseBodyData) *GetTicketTemplateResponseBody
	GetData() *GetTicketTemplateResponseBodyData
	SetHttpStatusCode(v int32) *GetTicketTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTicketTemplateResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetTicketTemplateResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetTicketTemplateResponseBody
	GetRequestId() *string
}

type GetTicketTemplateResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *GetTicketTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// BF268B34-09C2-43FD-BAC4-5D31EA633111
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTicketTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTicketTemplateResponseBody) GetData() *GetTicketTemplateResponseBodyData {
	return s.Data
}

func (s *GetTicketTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTicketTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTicketTemplateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetTicketTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketTemplateResponseBody) SetCode(v string) *GetTicketTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetData(v *GetTicketTemplateResponseBodyData) *GetTicketTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketTemplateResponseBody) SetHttpStatusCode(v int32) *GetTicketTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetMessage(v string) *GetTicketTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetParams(v []*string) *GetTicketTemplateResponseBody {
	s.Params = v
	return s
}

func (s *GetTicketTemplateResponseBody) SetRequestId(v string) *GetTicketTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTicketTemplateResponseBodyData struct {
	// Ticket category ID.
	//
	// example:
	//
	// 43c2671b-****-4223-86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Template editor.
	//
	// example:
	//
	// editor-xxx@ccc-test
	Editor *string `json:"Editor,omitempty" xml:"Editor,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Template name.
	//
	// example:
	//
	// 测试模板。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Workflow information.
	//
	// example:
	//
	// [{\\"id\\":\\"TICKET_START\\",\\"name\\":\\"开始\\",\\"nodeIndex\\":0,\\"properties\\":{\\"position\\":{\\"x\\":250,\\"y\\":200}},\\"type\\":\\"TICKET_START\\",\\"events\\":[{\\"edgeId\\":\\"8bd07339\\",\\"sourceAnchor\\":0,\\"next\\":\\"APPROVAL__JNBSSREc\\",\\"targetAnchor\\":0}]},{\\"id\\":\\"TICKET_FINISH\\",\\"name\\":\\"结束\\",\\"nodeIndex\\":9999,\\"properties\\":{\\"position\\":{\\"x\\":767,\\"y\\":206}},\\"type\\":\\"TICKET_FINISH\\",\\"events\\":[]},{\\"id\\":\\"APPROVAL__JNBSSREc\\",\\"name\\":\\"流程节点\\",\\"nodeIndex\\":1,\\"properties\\":{\\"skillGroupId\\":\\"chat001@cccV2-kmz\\",\\"position\\":{\\"x\\":537,\\"y\\":164.5}},\\"type\\":\\"APPROVAL\\",\\"events\\":[{\\"edgeId\\":\\"74031613\\",\\"sourceAnchor\\":1,\\"next\\":\\"TICKET_FINISH\\",\\"targetAnchor\\":0}]}]
	ProcessDefinition *string `json:"ProcessDefinition,omitempty" xml:"ProcessDefinition,omitempty"`
	// Status code.
	//
	// example:
	//
	// Enabled
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Template ID.
	//
	// example:
	//
	// 4ca2e2-c8d19b82c-d7ce393ac8197d3ab
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// List of template fields.
	TicketFields []*GetTicketTemplateResponseBodyDataTicketFields `json:"TicketFields,omitempty" xml:"TicketFields,omitempty" type:"Repeated"`
	// Last modified time.
	//
	// example:
	//
	// 1717664210000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetTicketTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetTicketTemplateResponseBodyData) GetEditor() *string {
	return s.Editor
}

func (s *GetTicketTemplateResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTicketTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTicketTemplateResponseBodyData) GetProcessDefinition() *string {
	return s.ProcessDefinition
}

func (s *GetTicketTemplateResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetTicketTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTicketTemplateResponseBodyData) GetTicketFields() []*GetTicketTemplateResponseBodyDataTicketFields {
	return s.TicketFields
}

func (s *GetTicketTemplateResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetTicketTemplateResponseBodyData) SetCategoryId(v string) *GetTicketTemplateResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetEditor(v string) *GetTicketTemplateResponseBodyData {
	s.Editor = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetInstanceId(v string) *GetTicketTemplateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetName(v string) *GetTicketTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetProcessDefinition(v string) *GetTicketTemplateResponseBodyData {
	s.ProcessDefinition = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetState(v string) *GetTicketTemplateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetTemplateId(v string) *GetTicketTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetTicketFields(v []*GetTicketTemplateResponseBodyDataTicketFields) *GetTicketTemplateResponseBodyData {
	s.TicketFields = v
	return s
}

func (s *GetTicketTemplateResponseBodyData) SetUpdatedTime(v int64) *GetTicketTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetTicketTemplateResponseBodyData) Validate() error {
	if s.TicketFields != nil {
		for _, item := range s.TicketFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTicketTemplateResponseBodyDataTicketFields struct {
	// Whether the field is an array.
	//
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// Extension attributes.
	//
	// example:
	//
	// {}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1717664210000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Creator.
	//
	// example:
	//
	// creator@cccV2-kmz
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Data type.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Description.
	//
	// example:
	//
	// 姓名描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the field is disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Display name.
	//
	// example:
	//
	// 姓名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Display order in lists.
	//
	// example:
	//
	// 1
	DisplayOrder *int32 `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	// Editor type.
	//
	// example:
	//
	// textbox
	EditorType *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	// Maximum length.
	//
	// example:
	//
	// 30
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// Numeric maximum value.
	//
	// example:
	//
	// 10
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// Minimum length.
	//
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// Numeric minimum value.
	//
	// example:
	//
	// 1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// Name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Regular expression validation rule.
	//
	// example:
	//
	// ^
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// Error message for regular expression validation.
	//
	// example:
	//
	// 不是有效的email地址
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// Whether the field is read-only.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// Whether the field is required.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// Whether the field is a system field.
	//
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// Update time.
	//
	// example:
	//
	// 1717664210000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetTicketTemplateResponseBodyDataTicketFields) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateResponseBodyDataTicketFields) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetArray() *bool {
	return s.Array
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetAttributes() *string {
	return s.Attributes
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetCreator() *string {
	return s.Creator
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetDataType() *string {
	return s.DataType
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetDescription() *string {
	return s.Description
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetEditorType() *string {
	return s.EditorType
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetMaximum() *float64 {
	return s.Maximum
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetMinLength() *int32 {
	return s.MinLength
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetMinimum() *float64 {
	return s.Minimum
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetName() *string {
	return s.Name
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetPattern() *string {
	return s.Pattern
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetRequired() *bool {
	return s.Required
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetSystem() *bool {
	return s.System
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetArray(v bool) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Array = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetAttributes(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Attributes = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetCreatedTime(v int64) *GetTicketTemplateResponseBodyDataTicketFields {
	s.CreatedTime = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetCreator(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Creator = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetDataType(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.DataType = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetDescription(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Description = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetDisabled(v bool) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Disabled = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetDisplayName(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.DisplayName = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetDisplayOrder(v int32) *GetTicketTemplateResponseBodyDataTicketFields {
	s.DisplayOrder = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetEditorType(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.EditorType = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetMaxLength(v int32) *GetTicketTemplateResponseBodyDataTicketFields {
	s.MaxLength = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetMaximum(v float64) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Maximum = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetMinLength(v int32) *GetTicketTemplateResponseBodyDataTicketFields {
	s.MinLength = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetMinimum(v float64) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Minimum = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetName(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Name = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetPattern(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Pattern = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetPatternErrorMessage(v string) *GetTicketTemplateResponseBodyDataTicketFields {
	s.PatternErrorMessage = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetReadOnly(v bool) *GetTicketTemplateResponseBodyDataTicketFields {
	s.ReadOnly = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetRequired(v bool) *GetTicketTemplateResponseBodyDataTicketFields {
	s.Required = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetSystem(v bool) *GetTicketTemplateResponseBodyDataTicketFields {
	s.System = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) SetUpdatedTime(v int64) *GetTicketTemplateResponseBodyDataTicketFields {
	s.UpdatedTime = &v
	return s
}

func (s *GetTicketTemplateResponseBodyDataTicketFields) Validate() error {
	return dara.Validate(s)
}
