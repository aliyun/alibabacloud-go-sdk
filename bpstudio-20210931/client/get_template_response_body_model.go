// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTemplateResponseBody
	GetCode() *int32
	SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody
	GetData() *GetTemplateResponseBodyData
	SetMessage(v string) *GetTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
}

type GetTemplateResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the template.
	Data *GetTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The interface returns information
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTemplateResponseBody) GetData() *GetTemplateResponseBodyData {
	return s.Data
}

func (s *GetTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) SetCode(v int32) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetMessage(v string) *GetTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTemplateResponseBodyData struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2020-09-22 17:08:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Template Description
	//
	// example:
	//
	// remark
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The path to the template schema image file
	//
	// example:
	//
	// bp-studio-template/sr-U37UD2YQCRJ75X5V.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template
	//
	// example:
	//
	// cadt-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID
	//
	// example:
	//
	// XFKR6WYRVS24S07R
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The details of the template variables.
	Variables []*GetTemplateResponseBodyDataVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTemplateResponseBodyData) GetDocumentUrl() *string {
	return s.DocumentUrl
}

func (s *GetTemplateResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *GetTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTemplateResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateResponseBodyData) GetVariables() []*GetTemplateResponseBodyDataVariables {
	return s.Variables
}

func (s *GetTemplateResponseBodyData) SetCreateTime(v string) *GetTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetDescription(v string) *GetTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetDocumentUrl(v string) *GetTemplateResponseBodyData {
	s.DocumentUrl = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetImageURL(v string) *GetTemplateResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetName(v string) *GetTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetResourceGroupId(v string) *GetTemplateResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetTemplateId(v string) *GetTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetVariables(v []*GetTemplateResponseBodyDataVariables) *GetTemplateResponseBodyData {
	s.Variables = v
	return s
}

func (s *GetTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTemplateResponseBodyDataVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// instance_name
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The type of the variable.
	//
	// example:
	//
	// String
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value of the variable.
	//
	// example:
	//
	// cadt-app-01
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Options      *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// ${name}
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
}

func (s GetTemplateResponseBodyDataVariables) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyDataVariables) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyDataVariables) GetAttribute() *string {
	return s.Attribute
}

func (s *GetTemplateResponseBodyDataVariables) GetDataType() *string {
	return s.DataType
}

func (s *GetTemplateResponseBodyDataVariables) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetTemplateResponseBodyDataVariables) GetOptions() *string {
	return s.Options
}

func (s *GetTemplateResponseBodyDataVariables) GetVariable() *string {
	return s.Variable
}

func (s *GetTemplateResponseBodyDataVariables) SetAttribute(v string) *GetTemplateResponseBodyDataVariables {
	s.Attribute = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetDataType(v string) *GetTemplateResponseBodyDataVariables {
	s.DataType = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetDefaultValue(v string) *GetTemplateResponseBodyDataVariables {
	s.DefaultValue = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetOptions(v string) *GetTemplateResponseBodyDataVariables {
	s.Options = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetVariable(v string) *GetTemplateResponseBodyDataVariables {
	s.Variable = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) Validate() error {
	return dara.Validate(s)
}
