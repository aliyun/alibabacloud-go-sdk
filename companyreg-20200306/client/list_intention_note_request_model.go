// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentionNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListIntentionNoteRequest
	GetBeginTime() *int64
	SetBizType(v string) *ListIntentionNoteRequest
	GetBizType() *string
	SetEndTime(v int64) *ListIntentionNoteRequest
	GetEndTime() *int64
	SetIntentionBizId(v string) *ListIntentionNoteRequest
	GetIntentionBizId() *string
	SetPageNumber(v int32) *ListIntentionNoteRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIntentionNoteRequest
	GetPageSize() *int32
}

type ListIntentionNoteRequest struct {
	// example:
	//
	// 1640456765459
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1631635199999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I20210420142416000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIntentionNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionNoteRequest) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListIntentionNoteRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListIntentionNoteRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntentionNoteRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListIntentionNoteRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIntentionNoteRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentionNoteRequest) SetBeginTime(v int64) *ListIntentionNoteRequest {
	s.BeginTime = &v
	return s
}

func (s *ListIntentionNoteRequest) SetBizType(v string) *ListIntentionNoteRequest {
	s.BizType = &v
	return s
}

func (s *ListIntentionNoteRequest) SetEndTime(v int64) *ListIntentionNoteRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntentionNoteRequest) SetIntentionBizId(v string) *ListIntentionNoteRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListIntentionNoteRequest) SetPageNumber(v int32) *ListIntentionNoteRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIntentionNoteRequest) SetPageSize(v int32) *ListIntentionNoteRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntentionNoteRequest) Validate() error {
	return dara.Validate(s)
}
