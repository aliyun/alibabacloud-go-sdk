// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKgEntityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListKgEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListKgEntityResponseBody
	GetMessage() *string
	SetPageResult(v *ListKgEntityResponseBodyPageResult) *ListKgEntityResponseBody
	GetPageResult() *ListKgEntityResponseBodyPageResult
	SetRequestId(v string) *ListKgEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKgEntityResponseBody
	GetSuccess() *bool
}

type ListKgEntityResponseBody struct {
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
	PageResult *ListKgEntityResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKgEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityResponseBody) GoString() string {
	return s.String()
}

func (s *ListKgEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKgEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListKgEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKgEntityResponseBody) GetPageResult() *ListKgEntityResponseBodyPageResult {
	return s.PageResult
}

func (s *ListKgEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKgEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKgEntityResponseBody) SetCode(v string) *ListKgEntityResponseBody {
	s.Code = &v
	return s
}

func (s *ListKgEntityResponseBody) SetHttpStatusCode(v int32) *ListKgEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListKgEntityResponseBody) SetMessage(v string) *ListKgEntityResponseBody {
	s.Message = &v
	return s
}

func (s *ListKgEntityResponseBody) SetPageResult(v *ListKgEntityResponseBodyPageResult) *ListKgEntityResponseBody {
	s.PageResult = v
	return s
}

func (s *ListKgEntityResponseBody) SetRequestId(v string) *ListKgEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKgEntityResponseBody) SetSuccess(v bool) *ListKgEntityResponseBody {
	s.Success = &v
	return s
}

func (s *ListKgEntityResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKgEntityResponseBodyPageResult struct {
	// The paged entity record list.
	EntityList []*ListKgEntityResponseBodyPageResultEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKgEntityResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListKgEntityResponseBodyPageResult) GetEntityList() []*ListKgEntityResponseBodyPageResultEntityList {
	return s.EntityList
}

func (s *ListKgEntityResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKgEntityResponseBodyPageResult) SetEntityList(v []*ListKgEntityResponseBodyPageResultEntityList) *ListKgEntityResponseBodyPageResult {
	s.EntityList = v
	return s
}

func (s *ListKgEntityResponseBodyPageResult) SetTotalCount(v int32) *ListKgEntityResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListKgEntityResponseBodyPageResult) Validate() error {
	if s.EntityList != nil {
		for _, item := range s.EntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKgEntityResponseBodyPageResultEntityList struct {
	// The entity record ID.
	//
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity type code.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity record property list.
	PropertyList []*ListKgEntityResponseBodyPageResultEntityListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
}

func (s ListKgEntityResponseBodyPageResultEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityResponseBodyPageResultEntityList) GoString() string {
	return s.String()
}

func (s *ListKgEntityResponseBodyPageResultEntityList) GetEntityId() *string {
	return s.EntityId
}

func (s *ListKgEntityResponseBodyPageResultEntityList) GetEntityType() *string {
	return s.EntityType
}

func (s *ListKgEntityResponseBodyPageResultEntityList) GetPropertyList() []*ListKgEntityResponseBodyPageResultEntityListPropertyList {
	return s.PropertyList
}

func (s *ListKgEntityResponseBodyPageResultEntityList) SetEntityId(v string) *ListKgEntityResponseBodyPageResultEntityList {
	s.EntityId = &v
	return s
}

func (s *ListKgEntityResponseBodyPageResultEntityList) SetEntityType(v string) *ListKgEntityResponseBodyPageResultEntityList {
	s.EntityType = &v
	return s
}

func (s *ListKgEntityResponseBodyPageResultEntityList) SetPropertyList(v []*ListKgEntityResponseBodyPageResultEntityListPropertyList) *ListKgEntityResponseBodyPageResultEntityList {
	s.PropertyList = v
	return s
}

func (s *ListKgEntityResponseBodyPageResultEntityList) Validate() error {
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

type ListKgEntityResponseBodyPageResultEntityListPropertyList struct {
	// The property code.
	//
	// example:
	//
	// company_name
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property data type. Valid values:
	//
	// - STRING: string.
	//
	// - INTEGER: integer.
	//
	// - FLOAT: floating-point number.
	//
	// - BOOLEAN: Boolean.
	//
	// - DATE: date.
	//
	// - LIST: list.
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

func (s ListKgEntityResponseBodyPageResultEntityListPropertyList) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityResponseBodyPageResultEntityListPropertyList) GoString() string {
	return s.String()
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) GetCode() *string {
	return s.Code
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) GetDataType() *string {
	return s.DataType
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) GetValue() *string {
	return s.Value
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) SetCode(v string) *ListKgEntityResponseBodyPageResultEntityListPropertyList {
	s.Code = &v
	return s
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) SetDataType(v string) *ListKgEntityResponseBodyPageResultEntityListPropertyList {
	s.DataType = &v
	return s
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) SetValue(v string) *ListKgEntityResponseBodyPageResultEntityListPropertyList {
	s.Value = &v
	return s
}

func (s *ListKgEntityResponseBodyPageResultEntityListPropertyList) Validate() error {
	return dara.Validate(s)
}
