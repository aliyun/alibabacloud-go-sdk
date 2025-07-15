// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoolsResponseBody
	GetPageSize() *int32
	SetPoolList(v []*ListPoolsResponseBodyPoolList) *ListPoolsResponseBody
	GetPoolList() []*ListPoolsResponseBodyPoolList
	SetRequestId(v string) *ListPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPoolsResponseBody
	GetTotalCount() *int32
}

type ListPoolsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PoolList []*ListPoolsResponseBodyPoolList `json:"PoolList,omitempty" xml:"PoolList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoolsResponseBody) GetPoolList() []*ListPoolsResponseBodyPoolList {
	return s.PoolList
}

func (s *ListPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPoolsResponseBody) SetPageNumber(v int32) *ListPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsResponseBody) SetPageSize(v int32) *ListPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPoolsResponseBody) SetPoolList(v []*ListPoolsResponseBodyPoolList) *ListPoolsResponseBody {
	s.PoolList = v
	return s
}

func (s *ListPoolsResponseBody) SetRequestId(v string) *ListPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoolsResponseBody) SetTotalCount(v int32) *ListPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPoolsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPoolsResponseBodyPoolList struct {
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPoolsResponseBodyPoolList) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponseBodyPoolList) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBodyPoolList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListPoolsResponseBodyPoolList) GetMaxExectorNum() *int32 {
	return s.MaxExectorNum
}

func (s *ListPoolsResponseBodyPoolList) GetPoolName() *string {
	return s.PoolName
}

func (s *ListPoolsResponseBodyPoolList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPoolsResponseBodyPoolList) GetStatus() *string {
	return s.Status
}

func (s *ListPoolsResponseBodyPoolList) SetIsDefault(v bool) *ListPoolsResponseBodyPoolList {
	s.IsDefault = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetMaxExectorNum(v int32) *ListPoolsResponseBodyPoolList {
	s.MaxExectorNum = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetPoolName(v string) *ListPoolsResponseBodyPoolList {
	s.PoolName = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetPriority(v int32) *ListPoolsResponseBodyPoolList {
	s.Priority = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetStatus(v string) *ListPoolsResponseBodyPoolList {
	s.Status = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) Validate() error {
	return dara.Validate(s)
}
