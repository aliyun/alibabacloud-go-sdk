// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccelerationOfWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAccelerationOfWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *ListAccelerationOfWorkspaceResponseBodyResult) *ListAccelerationOfWorkspaceResponseBody
	GetResult() *ListAccelerationOfWorkspaceResponseBodyResult
	SetSuccess(v bool) *ListAccelerationOfWorkspaceResponseBody
	GetSuccess() *bool
}

type ListAccelerationOfWorkspaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E****************05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The acceleration task model.
	Result *ListAccelerationOfWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccelerationOfWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerationOfWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccelerationOfWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccelerationOfWorkspaceResponseBody) GetResult() *ListAccelerationOfWorkspaceResponseBodyResult {
	return s.Result
}

func (s *ListAccelerationOfWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAccelerationOfWorkspaceResponseBody) SetRequestId(v string) *ListAccelerationOfWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBody) SetResult(v *ListAccelerationOfWorkspaceResponseBodyResult) *ListAccelerationOfWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBody) SetSuccess(v bool) *ListAccelerationOfWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAccelerationOfWorkspaceResponseBodyResult struct {
	// An array of acceleration task information.
	Data []*ListAccelerationOfWorkspaceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The next page. A value of NULL or 0 indicates that there is no next page.
	//
	// example:
	//
	// null
	Next *int32 `json:"Next,omitempty" xml:"Next,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page that you set when you call the operation.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The previous page. A value of NULL or 0 indicates that there is no previous page.
	//
	// example:
	//
	// 0
	Pre *int32 `json:"Pre,omitempty" xml:"Pre,omitempty"`
	// The total number of rows.
	//
	// example:
	//
	// 18
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListAccelerationOfWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerationOfWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetData() []*ListAccelerationOfWorkspaceResponseBodyResultData {
	return s.Data
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetNext() *int32 {
	return s.Next
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetPre() *int32 {
	return s.Pre
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetData(v []*ListAccelerationOfWorkspaceResponseBodyResultData) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetNext(v int32) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.Next = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetPageNum(v int32) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetPageSize(v int32) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetPre(v int32) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.Pre = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetTotalNum(v int32) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) SetTotalPages(v int32) *ListAccelerationOfWorkspaceResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccelerationOfWorkspaceResponseBodyResultData struct {
	// The name of the creator.
	//
	// example:
	//
	// system
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// d14e*********fef8de29fd
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// test
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// The time when acceleration was enabled.
	//
	// example:
	//
	// 20250911 00:00:00
	EnableQuickindexTime *string `json:"EnableQuickindexTime,omitempty" xml:"EnableQuickindexTime,omitempty"`
	// The task history ID.
	//
	// example:
	//
	// QWDAASG*******8SAD
	JobHistoryId *string `json:"JobHistoryId,omitempty" xml:"JobHistoryId,omitempty"`
	// The acceleration task ID.
	//
	// example:
	//
	// b30b74**********b3b
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status of the task. Valid values:
	//
	// - 0: Pending
	//
	// - 1: Running
	//
	// - 2: Succeeded
	//
	// - 3: Failed
	//
	// - 4: Canceled
	//
	// - 5: Warning
	//
	// example:
	//
	// 0
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The time when the task was last modified.
	//
	// example:
	//
	// 20250911 00:00:00
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The size.
	//
	// example:
	//
	// 47045632
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListAccelerationOfWorkspaceResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerationOfWorkspaceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetCubeId() *string {
	return s.CubeId
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetCubeName() *string {
	return s.CubeName
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetEnableQuickindexTime() *string {
	return s.EnableQuickindexTime
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetJobHistoryId() *string {
	return s.JobHistoryId
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetJobId() *string {
	return s.JobId
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) GetSize() *string {
	return s.Size
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetCreatorName(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.CreatorName = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetCubeId(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.CubeId = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetCubeName(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.CubeName = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetEnableQuickindexTime(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.EnableQuickindexTime = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetJobHistoryId(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.JobHistoryId = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetJobId(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.JobId = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetJobStatus(v int32) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.JobStatus = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetLastModifyTime(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.LastModifyTime = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) SetSize(v string) *ListAccelerationOfWorkspaceResponseBodyResultData {
	s.Size = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
