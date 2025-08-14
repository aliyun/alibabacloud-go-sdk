// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableTemplateBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventVariableTemplateBindResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeEventVariableTemplateBindResponseBodyResultObject) *DescribeEventVariableTemplateBindResponseBody
	GetResultObject() *DescribeEventVariableTemplateBindResponseBodyResultObject
}

type DescribeEventVariableTemplateBindResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeEventVariableTemplateBindResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeEventVariableTemplateBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateBindResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventVariableTemplateBindResponseBody) GetResultObject() *DescribeEventVariableTemplateBindResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventVariableTemplateBindResponseBody) SetRequestId(v string) *DescribeEventVariableTemplateBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBody) SetResultObject(v *DescribeEventVariableTemplateBindResponseBodyResultObject) *DescribeEventVariableTemplateBindResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableTemplateBindResponseBodyResultObject struct {
	// List of chargeable variables
	ChargeVariables []*DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables `json:"chargeVariables,omitempty" xml:"chargeVariables,omitempty" type:"Repeated"`
	// List of free variables
	FreeVariables []*DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables `json:"freeVariables,omitempty" xml:"freeVariables,omitempty" type:"Repeated"`
	// Template code
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Total count
	//
	// example:
	//
	// 38
	TotalCount *string `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeEventVariableTemplateBindResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateBindResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) GetChargeVariables() []*DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	return s.ChargeVariables
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) GetFreeVariables() []*DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	return s.FreeVariables
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) SetChargeVariables(v []*DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) *DescribeEventVariableTemplateBindResponseBodyResultObject {
	s.ChargeVariables = v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) SetFreeVariables(v []*DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) *DescribeEventVariableTemplateBindResponseBodyResultObject {
	s.FreeVariables = v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) SetTemplateCode(v string) *DescribeEventVariableTemplateBindResponseBodyResultObject {
	s.TemplateCode = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) SetTotalCount(v string) *DescribeEventVariableTemplateBindResponseBodyResultObject {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables struct {
	// Variable code
	//
	// example:
	//
	// ip
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Description of the variable.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Field type.
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 456
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable name
	//
	// example:
	//
	// ip
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Title.
	//
	// example:
	//
	// ip
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetCode(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetDescription(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetFieldType(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetId(v int64) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetName(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetTitle(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) SetType(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectChargeVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables struct {
	// Variable code
	//
	// example:
	//
	// age
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Variable description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Field type.
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 234
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

func (s DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetCode(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetDescription(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetFieldType(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetId(v int64) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetName(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetTitle(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) SetType(v string) *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponseBodyResultObjectFreeVariables) Validate() error {
	return dara.Validate(s)
}
