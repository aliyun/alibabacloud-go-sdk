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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListKgRelationResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	RelationList []*ListKgRelationResponseBodyPageResultRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
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
	PropertyList []*ListKgRelationResponseBodyPageResultRelationListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// example:
	//
	// BELONG_TO
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// example:
	//
	// abc-xxx
	SourceEntityId *string `json:"SourceEntityId,omitempty" xml:"SourceEntityId,omitempty"`
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
	// example:
	//
	// company_name
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
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
