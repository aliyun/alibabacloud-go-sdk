// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKgRelationResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListKgRelationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListKgRelationResponseBody
	GetMessage() *string
	SetPageResult(v *ListKgRelationResponseBodyPageResult) *ListKgRelationResponseBody
	GetPageResult() *ListKgRelationResponseBodyPageResult
	SetRequestId(v string) *ListKgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKgRelationResponseBody
	GetSuccess() *bool
}

type ListKgRelationResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListKgRelationResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ListKgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKgRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListKgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKgRelationResponseBody) GetPageResult() *ListKgRelationResponseBodyPageResult {
	return s.PageResult
}

func (s *ListKgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKgRelationResponseBody) SetCode(v string) *ListKgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *ListKgRelationResponseBody) SetHttpStatusCode(v int32) *ListKgRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListKgRelationResponseBody) SetMessage(v string) *ListKgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *ListKgRelationResponseBody) SetPageResult(v *ListKgRelationResponseBodyPageResult) *ListKgRelationResponseBody {
	s.PageResult = v
	return s
}

func (s *ListKgRelationResponseBody) SetRequestId(v string) *ListKgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKgRelationResponseBody) SetSuccess(v bool) *ListKgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *ListKgRelationResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKgRelationResponseBodyPageResult struct {
	// The paged relationship record list.
	RelationList []*ListKgRelationResponseBodyPageResultRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKgRelationResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListKgRelationResponseBodyPageResult) GetRelationList() []*ListKgRelationResponseBodyPageResultRelationList {
	return s.RelationList
}

func (s *ListKgRelationResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKgRelationResponseBodyPageResult) SetRelationList(v []*ListKgRelationResponseBodyPageResultRelationList) *ListKgRelationResponseBodyPageResult {
	s.RelationList = v
	return s
}

func (s *ListKgRelationResponseBodyPageResult) SetTotalCount(v int32) *ListKgRelationResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResult) Validate() error {
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKgRelationResponseBodyPageResultRelationList struct {
	// The relationship record property list.
	PropertyList []*ListKgRelationResponseBodyPageResultRelationListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The relationship record ID.
	//
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The relationship type code.
	//
	// example:
	//
	// BELONG_TO
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The source entity ID.
	//
	// example:
	//
	// abc-xxx
	SourceEntityId *string `json:"SourceEntityId,omitempty" xml:"SourceEntityId,omitempty"`
	// The target entity ID.
	//
	// example:
	//
	// abd-xxx
	TargetEntityId *string `json:"TargetEntityId,omitempty" xml:"TargetEntityId,omitempty"`
}

func (s ListKgRelationResponseBodyPageResultRelationList) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationResponseBodyPageResultRelationList) GoString() string {
	return s.String()
}

func (s *ListKgRelationResponseBodyPageResultRelationList) GetPropertyList() []*ListKgRelationResponseBodyPageResultRelationListPropertyList {
	return s.PropertyList
}

func (s *ListKgRelationResponseBodyPageResultRelationList) GetRelationId() *string {
	return s.RelationId
}

func (s *ListKgRelationResponseBodyPageResultRelationList) GetRelationType() *string {
	return s.RelationType
}

func (s *ListKgRelationResponseBodyPageResultRelationList) GetSourceEntityId() *string {
	return s.SourceEntityId
}

func (s *ListKgRelationResponseBodyPageResultRelationList) GetTargetEntityId() *string {
	return s.TargetEntityId
}

func (s *ListKgRelationResponseBodyPageResultRelationList) SetPropertyList(v []*ListKgRelationResponseBodyPageResultRelationListPropertyList) *ListKgRelationResponseBodyPageResultRelationList {
	s.PropertyList = v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationList) SetRelationId(v string) *ListKgRelationResponseBodyPageResultRelationList {
	s.RelationId = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationList) SetRelationType(v string) *ListKgRelationResponseBodyPageResultRelationList {
	s.RelationType = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationList) SetSourceEntityId(v string) *ListKgRelationResponseBodyPageResultRelationList {
	s.SourceEntityId = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationList) SetTargetEntityId(v string) *ListKgRelationResponseBodyPageResultRelationList {
	s.TargetEntityId = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationList) Validate() error {
	if s.PropertyList != nil {
		for _, item := range s.PropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKgRelationResponseBodyPageResultRelationListPropertyList struct {
	// The property code.
	//
	// example:
	//
	// company_name
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property data type. Valid values: STRING (string), INTEGER (integer), FLOAT (float), BOOLEAN (boolean), DATE (date), LIST (list), and others.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The property value.
	//
	// example:
	//
	// Alibaba
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListKgRelationResponseBodyPageResultRelationListPropertyList) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationResponseBodyPageResultRelationListPropertyList) GoString() string {
	return s.String()
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) GetCode() *string {
	return s.Code
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) GetValue() *string {
	return s.Value
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) SetCode(v string) *ListKgRelationResponseBodyPageResultRelationListPropertyList {
	s.Code = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) SetDataType(v string) *ListKgRelationResponseBodyPageResultRelationListPropertyList {
	s.DataType = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) SetValue(v string) *ListKgRelationResponseBodyPageResultRelationListPropertyList {
	s.Value = &v
	return s
}

func (s *ListKgRelationResponseBodyPageResultRelationListPropertyList) Validate() error {
	return dara.Validate(s)
}
