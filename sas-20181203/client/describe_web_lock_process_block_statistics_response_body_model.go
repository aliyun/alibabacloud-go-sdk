// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockProcessBlockStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockProcessBlockStatisticsResponseBody
	GetCurrentPage() *int32
	SetList(v []*DescribeWebLockProcessBlockStatisticsResponseBodyList) *DescribeWebLockProcessBlockStatisticsResponseBody
	GetList() []*DescribeWebLockProcessBlockStatisticsResponseBodyList
	SetPageSize(v int32) *DescribeWebLockProcessBlockStatisticsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWebLockProcessBlockStatisticsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockProcessBlockStatisticsResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockProcessBlockStatisticsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array consisting of the statistics on processes.
	List []*DescribeWebLockProcessBlockStatisticsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of processes.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockProcessBlockStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessBlockStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) GetList() []*DescribeWebLockProcessBlockStatisticsResponseBodyList {
	return s.List
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) SetCurrentPage(v int32) *DescribeWebLockProcessBlockStatisticsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) SetList(v []*DescribeWebLockProcessBlockStatisticsResponseBodyList) *DescribeWebLockProcessBlockStatisticsResponseBody {
	s.List = v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) SetPageSize(v int32) *DescribeWebLockProcessBlockStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) SetRequestId(v string) *DescribeWebLockProcessBlockStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) SetTotalCount(v int32) *DescribeWebLockProcessBlockStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebLockProcessBlockStatisticsResponseBodyList struct {
	// The number of processes that are returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The process.
	//
	// example:
	//
	// cron
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s DescribeWebLockProcessBlockStatisticsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessBlockStatisticsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBodyList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBodyList) GetProcess() *string {
	return s.Process
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBodyList) SetCount(v int32) *DescribeWebLockProcessBlockStatisticsResponseBodyList {
	s.Count = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBodyList) SetProcess(v string) *DescribeWebLockProcessBlockStatisticsResponseBodyList {
	s.Process = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponseBodyList) Validate() error {
	return dara.Validate(s)
}
