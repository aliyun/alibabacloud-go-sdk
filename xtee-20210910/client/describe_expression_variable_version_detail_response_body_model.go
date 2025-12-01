// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableVersionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExpressionVariableVersionDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeExpressionVariableVersionDetailResponseBodyResultObject) *DescribeExpressionVariableVersionDetailResponseBody
	GetResultObject() *DescribeExpressionVariableVersionDetailResponseBodyResultObject
}

type DescribeExpressionVariableVersionDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object.
	ResultObject *DescribeExpressionVariableVersionDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeExpressionVariableVersionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableVersionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableVersionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressionVariableVersionDetailResponseBody) GetResultObject() *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeExpressionVariableVersionDetailResponseBody) SetRequestId(v string) *DescribeExpressionVariableVersionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBody) SetResultObject(v *DescribeExpressionVariableVersionDetailResponseBodyResultObject) *DescribeExpressionVariableVersionDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExpressionVariableVersionDetailResponseBodyResultObject struct {
	// Creation type.
	//
	// example:
	//
	// MORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Expression.
	//
	// example:
	//
	// @ex_GX9rrlTq4b67 + 1001
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// Expression title.
	//
	// example:
	//
	// @selfvariable_02 + 1001
	ExpressionTitle *string `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	// Expression variable.
	//
	// example:
	//
	// ex_GX9rrlTq4b67
	ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
	// Field ranking.
	//
	// example:
	//
	// 0
	FieldRank *int32 `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1762409015000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1762409026000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Custom variable primary key.
	//
	// example:
	//
	// 397625
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Invoke key.
	//
	// example:
	//
	// deInvokeSelfVariable_v1
	InvokeKey *string `json:"invokeKey,omitempty" xml:"invokeKey,omitempty"`
	// Variable name, a uniquely generated identifier.
	//
	// example:
	//
	// ex_0kWIfZ27c525
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Outlier.
	//
	// example:
	//
	// SYS_ERROR
	Outlier *string `json:"outlier,omitempty" xml:"outlier,omitempty"`
	// Variable return type.
	//
	// example:
	//
	// EXPRESSION
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Variable associated event.
	//
	// example:
	//
	// de_awkhwh0314
	RefObjId *string `json:"refObjId,omitempty" xml:"refObjId,omitempty"`
	// Variable association type.
	//
	// example:
	//
	// EVENT_BY_EXPRESSION
	RefObjType *string `json:"refObjType,omitempty" xml:"refObjType,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Source type.
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Variable title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type.
	//
	// example:
	//
	// EXPRESSION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User UID.
	//
	// example:
	//
	// 151222xxxxxxxxxx
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// Variable version.
	//
	// example:
	//
	// 2
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeExpressionVariableVersionDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableVersionDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetExpression() *string {
	return s.Expression
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetExpressionVariable() *string {
	return s.ExpressionVariable
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetInvokeKey() *string {
	return s.InvokeKey
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetRefObjId() *string {
	return s.RefObjId
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetRefObjType() *string {
	return s.RefObjType
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetRegion() *string {
	return s.Region
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetCreateType(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.CreateType = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetDescription(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetExpression(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Expression = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetExpressionTitle(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetExpressionVariable(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.ExpressionVariable = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetFieldRank(v int32) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.FieldRank = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetGmtCreate(v int64) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetGmtModified(v int64) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetId(v int64) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetInvokeKey(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.InvokeKey = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetName(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetOutlier(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Outlier = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetOutputs(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Outputs = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetRefObjId(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.RefObjId = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetRefObjType(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.RefObjType = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetRegion(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Region = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetSourceType(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.SourceType = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetStatus(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetTitle(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetType(v string) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetUserId(v int64) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) SetVersion(v int64) *DescribeExpressionVariableVersionDetailResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeExpressionVariableVersionDetailResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
