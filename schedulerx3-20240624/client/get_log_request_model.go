// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetLogRequest
	GetAppName() *string
	SetClusterId(v string) *GetLogRequest
	GetClusterId() *string
	SetEndTime(v int64) *GetLogRequest
	GetEndTime() *int64
	SetJobExecutionId(v string) *GetLogRequest
	GetJobExecutionId() *string
	SetKeyword(v string) *GetLogRequest
	GetKeyword() *string
	SetLevel(v string) *GetLogRequest
	GetLevel() *string
	SetLineNum(v int32) *GetLogRequest
	GetLineNum() *int32
	SetLogId(v int64) *GetLogRequest
	GetLogId() *int64
	SetOffset(v int32) *GetLogRequest
	GetOffset() *int32
	SetReverse(v bool) *GetLogRequest
	GetReverse() *bool
	SetStartTime(v int64) *GetLogRequest
	GetStartTime() *int64
}

type GetLogRequest struct {
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
	// 1721636220
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// hello word
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// LineNum
	//
	// example:
	//
	// 2
	LineNum *int32 `json:"LineNum,omitempty" xml:"LineNum,omitempty"`
	// example:
	//
	// 344008
	LogId *int64 `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1721636220
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogRequest) GoString() string {
	return s.String()
}

func (s *GetLogRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetLogRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetLogRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetLogRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetLogRequest) GetLevel() *string {
	return s.Level
}

func (s *GetLogRequest) GetLineNum() *int32 {
	return s.LineNum
}

func (s *GetLogRequest) GetLogId() *int64 {
	return s.LogId
}

func (s *GetLogRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *GetLogRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *GetLogRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetLogRequest) SetAppName(v string) *GetLogRequest {
	s.AppName = &v
	return s
}

func (s *GetLogRequest) SetClusterId(v string) *GetLogRequest {
	s.ClusterId = &v
	return s
}

func (s *GetLogRequest) SetEndTime(v int64) *GetLogRequest {
	s.EndTime = &v
	return s
}

func (s *GetLogRequest) SetJobExecutionId(v string) *GetLogRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetLogRequest) SetKeyword(v string) *GetLogRequest {
	s.Keyword = &v
	return s
}

func (s *GetLogRequest) SetLevel(v string) *GetLogRequest {
	s.Level = &v
	return s
}

func (s *GetLogRequest) SetLineNum(v int32) *GetLogRequest {
	s.LineNum = &v
	return s
}

func (s *GetLogRequest) SetLogId(v int64) *GetLogRequest {
	s.LogId = &v
	return s
}

func (s *GetLogRequest) SetOffset(v int32) *GetLogRequest {
	s.Offset = &v
	return s
}

func (s *GetLogRequest) SetReverse(v bool) *GetLogRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogRequest) SetStartTime(v int64) *GetLogRequest {
	s.StartTime = &v
	return s
}

func (s *GetLogRequest) Validate() error {
	return dara.Validate(s)
}
