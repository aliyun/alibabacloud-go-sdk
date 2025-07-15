// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpsecServerLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListIpsecServerLogsResponseBody
	GetCount() *int32
	SetData(v []*string) *ListIpsecServerLogsResponseBody
	GetData() []*string
	SetIsCompleted(v bool) *ListIpsecServerLogsResponseBody
	GetIsCompleted() *bool
	SetPageNumber(v int32) *ListIpsecServerLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIpsecServerLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListIpsecServerLogsResponseBody
	GetRequestId() *string
}

type ListIpsecServerLogsResponseBody struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Log information list.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether the log is accurate. Valid values:
	//
	// 	- **true**: accurate
	//
	// 	- **false**: inaccurate
	//
	// example:
	//
	// true
	IsCompleted *bool `json:"IsCompleted,omitempty" xml:"IsCompleted,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DEDAC5B1-9292-5BF7-BDDF-61BA58CFB2FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIpsecServerLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpsecServerLogsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListIpsecServerLogsResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListIpsecServerLogsResponseBody) GetIsCompleted() *bool {
	return s.IsCompleted
}

func (s *ListIpsecServerLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIpsecServerLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIpsecServerLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpsecServerLogsResponseBody) SetCount(v int32) *ListIpsecServerLogsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpsecServerLogsResponseBody) SetData(v []*string) *ListIpsecServerLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListIpsecServerLogsResponseBody) SetIsCompleted(v bool) *ListIpsecServerLogsResponseBody {
	s.IsCompleted = &v
	return s
}

func (s *ListIpsecServerLogsResponseBody) SetPageNumber(v int32) *ListIpsecServerLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIpsecServerLogsResponseBody) SetPageSize(v int32) *ListIpsecServerLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIpsecServerLogsResponseBody) SetRequestId(v string) *ListIpsecServerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpsecServerLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
