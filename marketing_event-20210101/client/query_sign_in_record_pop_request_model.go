// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySignInRecordPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *QuerySignInRecordPopRequest
	GetActivityId() *int64
	SetEndTime(v string) *QuerySignInRecordPopRequest
	GetEndTime() *string
	SetPageNum(v int32) *QuerySignInRecordPopRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySignInRecordPopRequest
	GetPageSize() *int32
	SetStartTime(v string) *QuerySignInRecordPopRequest
	GetStartTime() *string
}

type QuerySignInRecordPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QuerySignInRecordPopRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySignInRecordPopRequest) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *QuerySignInRecordPopRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySignInRecordPopRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySignInRecordPopRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySignInRecordPopRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySignInRecordPopRequest) SetActivityId(v int64) *QuerySignInRecordPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetEndTime(v string) *QuerySignInRecordPopRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetPageNum(v int32) *QuerySignInRecordPopRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetPageSize(v int32) *QuerySignInRecordPopRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySignInRecordPopRequest) SetStartTime(v string) *QuerySignInRecordPopRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySignInRecordPopRequest) Validate() error {
	return dara.Validate(s)
}
