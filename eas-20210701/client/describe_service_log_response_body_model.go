// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*string) *DescribeServiceLogResponseBody
	GetLogs() []*string
	SetPageNum(v int64) *DescribeServiceLogResponseBody
	GetPageNum() *int64
	SetRequestId(v string) *DescribeServiceLogResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeServiceLogResponseBody
	GetTotalCount() *int64
	SetTotalPageNum(v int64) *DescribeServiceLogResponseBody
	GetTotalPageNum() *int64
}

type DescribeServiceLogResponseBody struct {
	// The returned logs.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 500
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s DescribeServiceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *DescribeServiceLogResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeServiceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeServiceLogResponseBody) GetTotalPageNum() *int64 {
	return s.TotalPageNum
}

func (s *DescribeServiceLogResponseBody) SetLogs(v []*string) *DescribeServiceLogResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeServiceLogResponseBody) SetPageNum(v int64) *DescribeServiceLogResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetRequestId(v string) *DescribeServiceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetTotalCount(v int64) *DescribeServiceLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetTotalPageNum(v int64) *DescribeServiceLogResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *DescribeServiceLogResponseBody) Validate() error {
	return dara.Validate(s)
}
