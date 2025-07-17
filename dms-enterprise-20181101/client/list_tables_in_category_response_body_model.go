// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesInCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityList(v *ListTablesInCategoryResponseBodyEntityList) *ListTablesInCategoryResponseBody
	GetEntityList() *ListTablesInCategoryResponseBodyEntityList
	SetErrorCode(v string) *ListTablesInCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTablesInCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTablesInCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTablesInCategoryResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTablesInCategoryResponseBody
	GetTotalCount() *int64
}

type ListTablesInCategoryResponseBody struct {
	EntityList *ListTablesInCategoryResponseBodyEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTablesInCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesInCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesInCategoryResponseBody) GetEntityList() *ListTablesInCategoryResponseBodyEntityList {
	return s.EntityList
}

func (s *ListTablesInCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTablesInCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTablesInCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesInCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTablesInCategoryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTablesInCategoryResponseBody) SetEntityList(v *ListTablesInCategoryResponseBodyEntityList) *ListTablesInCategoryResponseBody {
	s.EntityList = v
	return s
}

func (s *ListTablesInCategoryResponseBody) SetErrorCode(v string) *ListTablesInCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTablesInCategoryResponseBody) SetErrorMessage(v string) *ListTablesInCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTablesInCategoryResponseBody) SetRequestId(v string) *ListTablesInCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesInCategoryResponseBody) SetSuccess(v bool) *ListTablesInCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListTablesInCategoryResponseBody) SetTotalCount(v int64) *ListTablesInCategoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTablesInCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTablesInCategoryResponseBodyEntityList struct {
	MetaCategoryTableEntity []*MetaCategoryTableEntity `json:"MetaCategoryTableEntity,omitempty" xml:"MetaCategoryTableEntity,omitempty" type:"Repeated"`
}

func (s ListTablesInCategoryResponseBodyEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesInCategoryResponseBodyEntityList) GoString() string {
	return s.String()
}

func (s *ListTablesInCategoryResponseBodyEntityList) GetMetaCategoryTableEntity() []*MetaCategoryTableEntity {
	return s.MetaCategoryTableEntity
}

func (s *ListTablesInCategoryResponseBodyEntityList) SetMetaCategoryTableEntity(v []*MetaCategoryTableEntity) *ListTablesInCategoryResponseBodyEntityList {
	s.MetaCategoryTableEntity = v
	return s
}

func (s *ListTablesInCategoryResponseBodyEntityList) Validate() error {
	return dara.Validate(s)
}
