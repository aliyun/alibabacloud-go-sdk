// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetPageLogRequest
	GetAppName() *string
	SetClusterId(v string) *GetPageLogRequest
	GetClusterId() *string
	SetEndTime(v int64) *GetPageLogRequest
	GetEndTime() *int64
	SetJobExecutionId(v string) *GetPageLogRequest
	GetJobExecutionId() *string
	SetJobName(v string) *GetPageLogRequest
	GetJobName() *string
	SetKeyword(v string) *GetPageLogRequest
	GetKeyword() *string
	SetPageNum(v int32) *GetPageLogRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetPageLogRequest
	GetPageSize() *int32
	SetReverse(v bool) *GetPageLogRequest
	GetReverse() *bool
	SetStartTime(v int64) *GetPageLogRequest
	GetStartTime() *int64
	SetWorkerAddr(v string) *GetPageLogRequest
	GetWorkerAddr() *string
}

type GetPageLogRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1777292662000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// job01
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// hello word
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1777292662000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 192.168.1.100
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s GetPageLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageLogRequest) GoString() string {
	return s.String()
}

func (s *GetPageLogRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetPageLogRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPageLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPageLogRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetPageLogRequest) GetJobName() *string {
	return s.JobName
}

func (s *GetPageLogRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetPageLogRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetPageLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPageLogRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *GetPageLogRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPageLogRequest) GetWorkerAddr() *string {
	return s.WorkerAddr
}

func (s *GetPageLogRequest) SetAppName(v string) *GetPageLogRequest {
	s.AppName = &v
	return s
}

func (s *GetPageLogRequest) SetClusterId(v string) *GetPageLogRequest {
	s.ClusterId = &v
	return s
}

func (s *GetPageLogRequest) SetEndTime(v int64) *GetPageLogRequest {
	s.EndTime = &v
	return s
}

func (s *GetPageLogRequest) SetJobExecutionId(v string) *GetPageLogRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetPageLogRequest) SetJobName(v string) *GetPageLogRequest {
	s.JobName = &v
	return s
}

func (s *GetPageLogRequest) SetKeyword(v string) *GetPageLogRequest {
	s.Keyword = &v
	return s
}

func (s *GetPageLogRequest) SetPageNum(v int32) *GetPageLogRequest {
	s.PageNum = &v
	return s
}

func (s *GetPageLogRequest) SetPageSize(v int32) *GetPageLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetPageLogRequest) SetReverse(v bool) *GetPageLogRequest {
	s.Reverse = &v
	return s
}

func (s *GetPageLogRequest) SetStartTime(v int64) *GetPageLogRequest {
	s.StartTime = &v
	return s
}

func (s *GetPageLogRequest) SetWorkerAddr(v string) *GetPageLogRequest {
	s.WorkerAddr = &v
	return s
}

func (s *GetPageLogRequest) Validate() error {
	return dara.Validate(s)
}
