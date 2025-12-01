// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecentDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetRecentDocumentRequest
	GetDateType() *string
	SetEndDate(v int64) *GetRecentDocumentRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetRecentDocumentRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetRecentDocumentRequest
	GetSuspEventSource() *string
}

type GetRecentDocumentRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time ranges.
	//
	// This parameter is required.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetRecentDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecentDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetRecentDocumentRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetRecentDocumentRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetRecentDocumentRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetRecentDocumentRequest) SetDateType(v string) *GetRecentDocumentRequest {
	s.DateType = &v
	return s
}

func (s *GetRecentDocumentRequest) SetEndDate(v int64) *GetRecentDocumentRequest {
	s.EndDate = &v
	return s
}

func (s *GetRecentDocumentRequest) SetStartDate(v int64) *GetRecentDocumentRequest {
	s.StartDate = &v
	return s
}

func (s *GetRecentDocumentRequest) SetSuspEventSource(v string) *GetRecentDocumentRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetRecentDocumentRequest) Validate() error {
	return dara.Validate(s)
}
