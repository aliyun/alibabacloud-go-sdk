// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDatabasesResponseBodyPagingInfo) *ListDatabasesResponseBody
	GetPagingInfo() *ListDatabasesResponseBodyPagingInfo
	SetRequestId(v string) *ListDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatabasesResponseBody
	GetSuccess() *bool
}

type ListDatabasesResponseBody struct {
	PagingInfo *ListDatabasesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 9DD08926-38B9-XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetPagingInfo() *ListDatabasesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatabasesResponseBody) SetPagingInfo(v *ListDatabasesResponseBodyPagingInfo) *ListDatabasesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetSuccess(v bool) *ListDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyPagingInfo struct {
	Databases []*Database `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabasesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyPagingInfo) GetDatabases() []*Database {
	return s.Databases
}

func (s *ListDatabasesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatabasesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatabasesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabasesResponseBodyPagingInfo) SetDatabases(v []*Database) *ListDatabasesResponseBodyPagingInfo {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDatabasesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesResponseBodyPagingInfo) SetPageSize(v int32) *ListDatabasesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDatabasesResponseBodyPagingInfo) SetTotalCount(v int64) *ListDatabasesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDatabasesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}
