// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStrategyExecutionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListStrategyExecutionRecordRequest
	GetEndTime() *string
	SetMediaId(v string) *ListStrategyExecutionRecordRequest
	GetMediaId() *string
	SetPageNo(v int32) *ListStrategyExecutionRecordRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListStrategyExecutionRecordRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListStrategyExecutionRecordRequest
	GetStartTime() *string
	SetStatus(v string) *ListStrategyExecutionRecordRequest
	GetStatus() *string
	SetStrategyId(v string) *ListStrategyExecutionRecordRequest
	GetStrategyId() *string
	SetStrategyType(v string) *ListStrategyExecutionRecordRequest
	GetStrategyType() *string
}

type ListStrategyExecutionRecordRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s ListStrategyExecutionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStrategyExecutionRecordRequest) GoString() string {
	return s.String()
}

func (s *ListStrategyExecutionRecordRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListStrategyExecutionRecordRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *ListStrategyExecutionRecordRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListStrategyExecutionRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStrategyExecutionRecordRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListStrategyExecutionRecordRequest) GetStatus() *string {
	return s.Status
}

func (s *ListStrategyExecutionRecordRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListStrategyExecutionRecordRequest) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListStrategyExecutionRecordRequest) SetEndTime(v string) *ListStrategyExecutionRecordRequest {
	s.EndTime = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetMediaId(v string) *ListStrategyExecutionRecordRequest {
	s.MediaId = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetPageNo(v int32) *ListStrategyExecutionRecordRequest {
	s.PageNo = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetPageSize(v int32) *ListStrategyExecutionRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetStartTime(v string) *ListStrategyExecutionRecordRequest {
	s.StartTime = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetStatus(v string) *ListStrategyExecutionRecordRequest {
	s.Status = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetStrategyId(v string) *ListStrategyExecutionRecordRequest {
	s.StrategyId = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) SetStrategyType(v string) *ListStrategyExecutionRecordRequest {
	s.StrategyType = &v
	return s
}

func (s *ListStrategyExecutionRecordRequest) Validate() error {
	return dara.Validate(s)
}
