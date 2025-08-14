// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventVariableTemplateListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeEventVariableTemplateListResponseBodyResultObject) *DescribeEventVariableTemplateListResponseBody
	GetResultObject() []*DescribeEventVariableTemplateListResponseBodyResultObject
}

type DescribeEventVariableTemplateListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeEventVariableTemplateListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeEventVariableTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventVariableTemplateListResponseBody) GetResultObject() []*DescribeEventVariableTemplateListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventVariableTemplateListResponseBody) SetRequestId(v string) *DescribeEventVariableTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBody) SetResultObject(v []*DescribeEventVariableTemplateListResponseBodyResultObject) *DescribeEventVariableTemplateListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableTemplateListResponseBodyResultObject struct {
	// Template code.
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Template name.
	//
	// example:
	//
	// 注册模版
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// Variable list.
	Variables []*DescribeEventVariableTemplateListResponseBodyResultObjectVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s DescribeEventVariableTemplateListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) GetVariables() []*DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	return s.Variables
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) SetTemplateCode(v string) *DescribeEventVariableTemplateListResponseBodyResultObject {
	s.TemplateCode = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) SetTemplateName(v string) *DescribeEventVariableTemplateListResponseBodyResultObject {
	s.TemplateName = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) SetVariables(v []*DescribeEventVariableTemplateListResponseBodyResultObjectVariables) *DescribeEventVariableTemplateListResponseBodyResultObject {
	s.Variables = v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableTemplateListResponseBodyResultObjectVariables struct {
	// Variable code
	//
	// example:
	//
	// age
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Variable input type
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 454
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventVariableTemplateListResponseBodyResultObjectVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetCode(v string) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetDescription(v string) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetFieldType(v string) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetId(v int64) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetName(v string) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetTitle(v string) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) SetType(v string) *DescribeEventVariableTemplateListResponseBodyResultObjectVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponseBodyResultObjectVariables) Validate() error {
	return dara.Validate(s)
}
