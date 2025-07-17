// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListSchemasResponseBodyPagingInfo) *ListSchemasResponseBody
	GetPagingInfo() *ListSchemasResponseBodyPagingInfo
	SetRequestId(v string) *ListSchemasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSchemasResponseBody
	GetSuccess() *bool
}

type ListSchemasResponseBody struct {
	// The pagination information.
	PagingInfo *ListSchemasResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 235BBA5E-3428-XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBody) GetPagingInfo() *ListSchemasResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchemasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSchemasResponseBody) SetPagingInfo(v *ListSchemasResponseBodyPagingInfo) *ListSchemasResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListSchemasResponseBody) SetRequestId(v string) *ListSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemasResponseBody) SetSuccess(v bool) *ListSchemasResponseBody {
	s.Success = &v
	return s
}

func (s *ListSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSchemasResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The schemas.
	Schemas []*Schema `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSchemasResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSchemasResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchemasResponseBodyPagingInfo) GetSchemas() []*Schema {
	return s.Schemas
}

func (s *ListSchemasResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSchemasResponseBodyPagingInfo) SetPageNumber(v int32) *ListSchemasResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListSchemasResponseBodyPagingInfo) SetPageSize(v int32) *ListSchemasResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListSchemasResponseBodyPagingInfo) SetSchemas(v []*Schema) *ListSchemasResponseBodyPagingInfo {
	s.Schemas = v
	return s
}

func (s *ListSchemasResponseBodyPagingInfo) SetTotalCount(v int64) *ListSchemasResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSchemasResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}
