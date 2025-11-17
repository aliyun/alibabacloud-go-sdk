// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccelerationLogByCubeIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryAccelerationLogByCubeIdResponseBody
	GetRequestId() *string
	SetResult(v *QueryAccelerationLogByCubeIdResponseBodyResult) *QueryAccelerationLogByCubeIdResponseBody
	GetResult() *QueryAccelerationLogByCubeIdResponseBodyResult
	SetSuccess(v bool) *QueryAccelerationLogByCubeIdResponseBody
	GetSuccess() *bool
}

type QueryAccelerationLogByCubeIdResponseBody struct {
	// example:
	//
	// D8749D********80FF3B4
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryAccelerationLogByCubeIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccelerationLogByCubeIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccelerationLogByCubeIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccelerationLogByCubeIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccelerationLogByCubeIdResponseBody) GetResult() *QueryAccelerationLogByCubeIdResponseBodyResult {
	return s.Result
}

func (s *QueryAccelerationLogByCubeIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccelerationLogByCubeIdResponseBody) SetRequestId(v string) *QueryAccelerationLogByCubeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBody) SetResult(v *QueryAccelerationLogByCubeIdResponseBodyResult) *QueryAccelerationLogByCubeIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBody) SetSuccess(v bool) *QueryAccelerationLogByCubeIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccelerationLogByCubeIdResponseBodyResult struct {
	Data []*QueryAccelerationLogByCubeIdResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	Next *int32 `json:"Next,omitempty" xml:"Next,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// null
	Pre *int32 `json:"Pre,omitempty" xml:"Pre,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryAccelerationLogByCubeIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryAccelerationLogByCubeIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetData() []*QueryAccelerationLogByCubeIdResponseBodyResultData {
	return s.Data
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetNext() *int32 {
	return s.Next
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetPre() *int32 {
	return s.Pre
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetData(v []*QueryAccelerationLogByCubeIdResponseBodyResultData) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetNext(v int32) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.Next = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetPageNum(v int32) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetPageSize(v int32) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetPre(v int32) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.Pre = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetTotalNum(v int32) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) SetTotalPages(v int32) *QueryAccelerationLogByCubeIdResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResult) Validate() error {
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

type QueryAccelerationLogByCubeIdResponseBodyResultData struct {
	// example:
	//
	// 123
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// REST_A***************6a8
	JobHistoryId *string `json:"JobHistoryId,omitempty" xml:"JobHistoryId,omitempty"`
	// example:
	//
	// REST_A***************6a8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// REST_A***************6a8
	JonStartDate *string `json:"JonStartDate,omitempty" xml:"JonStartDate,omitempty"`
	// example:
	//
	// asdav************
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
}

func (s QueryAccelerationLogByCubeIdResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccelerationLogByCubeIdResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) GetDuration() *string {
	return s.Duration
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) GetJobHistoryId() *string {
	return s.JobHistoryId
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) GetJobId() *string {
	return s.JobId
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) GetJonStartDate() *string {
	return s.JonStartDate
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) GetLog() *string {
	return s.Log
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) SetDuration(v string) *QueryAccelerationLogByCubeIdResponseBodyResultData {
	s.Duration = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) SetJobHistoryId(v string) *QueryAccelerationLogByCubeIdResponseBodyResultData {
	s.JobHistoryId = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) SetJobId(v string) *QueryAccelerationLogByCubeIdResponseBodyResultData {
	s.JobId = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) SetJobStatus(v string) *QueryAccelerationLogByCubeIdResponseBodyResultData {
	s.JobStatus = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) SetJonStartDate(v string) *QueryAccelerationLogByCubeIdResponseBodyResultData {
	s.JonStartDate = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) SetLog(v string) *QueryAccelerationLogByCubeIdResponseBodyResultData {
	s.Log = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
