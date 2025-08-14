// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvanceSearchLeftVariableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAdvanceSearchLeftVariableListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) *DescribeAdvanceSearchLeftVariableListResponseBody
	GetResultObject() []*DescribeAdvanceSearchLeftVariableListResponseBodyResultObject
}

type DescribeAdvanceSearchLeftVariableListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject []*DescribeAdvanceSearchLeftVariableListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeAdvanceSearchLeftVariableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchLeftVariableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBody) GetResultObject() []*DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBody) SetRequestId(v string) *DescribeAdvanceSearchLeftVariableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBody) SetResultObject(v []*DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) *DescribeAdvanceSearchLeftVariableListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAdvanceSearchLeftVariableListResponseBodyResultObject struct {
	// Variable code
	//
	// example:
	//
	// age
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Description.
	//
	// example:
	//
	// 年龄描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Display type and grouping label
	//
	// example:
	//
	// NATIVE
	DisplayType *string `json:"displayType,omitempty" xml:"displayType,omitempty"`
	// Variable return value type
	//
	// example:
	//
	// STRING
	FieldRank *int64 `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	// Field table sorting
	//
	// example:
	//
	// 1
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 2453
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Parent node
	//
	// example:
	//
	// name
	ParentName *string `json:"parentName,omitempty" xml:"parentName,omitempty"`
	// Data source
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetCode() *string {
	return s.Code
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetFieldRank() *int64 {
	return s.FieldRank
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetCode(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.Code = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetDescription(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetDisplayType(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.DisplayType = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetFieldRank(v int64) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.FieldRank = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetFieldType(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.FieldType = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetId(v int64) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetName(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetParentName(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.ParentName = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetSourceType(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.SourceType = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetTitle(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) SetType(v string) *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
