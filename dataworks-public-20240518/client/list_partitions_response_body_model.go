// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListPartitionsResponseBodyPagingInfo) *ListPartitionsResponseBody
	GetPagingInfo() *ListPartitionsResponseBodyPagingInfo
	SetRequestId(v string) *ListPartitionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPartitionsResponseBody
	GetSuccess() *bool
}

type ListPartitionsResponseBody struct {
	PagingInfo *ListPartitionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// D1E2E5BC-xxxx-xxxx-xxxx-xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPartitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartitionsResponseBody) GetPagingInfo() *ListPartitionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListPartitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPartitionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPartitionsResponseBody) SetPagingInfo(v *ListPartitionsResponseBodyPagingInfo) *ListPartitionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListPartitionsResponseBody) SetRequestId(v string) *ListPartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPartitionsResponseBody) SetSuccess(v bool) *ListPartitionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPartitionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPartitionsResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize      *int32       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PartitionList []*Partition `json:"PartitionList,omitempty" xml:"PartitionList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPartitionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListPartitionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPartitionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPartitionsResponseBodyPagingInfo) GetPartitionList() []*Partition {
	return s.PartitionList
}

func (s *ListPartitionsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPartitionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListPartitionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListPartitionsResponseBodyPagingInfo) SetPageSize(v int32) *ListPartitionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListPartitionsResponseBodyPagingInfo) SetPartitionList(v []*Partition) *ListPartitionsResponseBodyPagingInfo {
	s.PartitionList = v
	return s
}

func (s *ListPartitionsResponseBodyPagingInfo) SetTotalCount(v int64) *ListPartitionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListPartitionsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}
