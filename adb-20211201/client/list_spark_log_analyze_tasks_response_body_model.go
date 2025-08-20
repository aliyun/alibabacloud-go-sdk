// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkLogAnalyzeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSparkLogAnalyzeTasksResponseBodyData) *ListSparkLogAnalyzeTasksResponseBody
	GetData() *ListSparkLogAnalyzeTasksResponseBodyData
	SetRequestId(v string) *ListSparkLogAnalyzeTasksResponseBody
	GetRequestId() *string
}

type ListSparkLogAnalyzeTasksResponseBody struct {
	// The returned data.
	Data *ListSparkLogAnalyzeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1DF5AF5B-C803-1861-A0FF-63666A557709
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkLogAnalyzeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksResponseBody) GetData() *ListSparkLogAnalyzeTasksResponseBodyData {
	return s.Data
}

func (s *ListSparkLogAnalyzeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSparkLogAnalyzeTasksResponseBody) SetData(v *ListSparkLogAnalyzeTasksResponseBodyData) *ListSparkLogAnalyzeTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBody) SetRequestId(v string) *ListSparkLogAnalyzeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSparkLogAnalyzeTasksResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried Spark log analysis tasks.
	TaskList []*SparkAnalyzeLogTask `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkLogAnalyzeTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) GetTaskList() []*SparkAnalyzeLogTask {
	return s.TaskList
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetPageNumber(v int64) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetPageSize(v int64) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetTaskList(v []*SparkAnalyzeLogTask) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.TaskList = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetTotalCount(v int64) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
