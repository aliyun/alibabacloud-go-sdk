// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStrategyExecutionRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionRecordList(v []*ListStrategyExecutionRecordResponseBodyExecutionRecordList) *ListStrategyExecutionRecordResponseBody
	GetExecutionRecordList() []*ListStrategyExecutionRecordResponseBodyExecutionRecordList
	SetMaxResults(v int32) *ListStrategyExecutionRecordResponseBody
	GetMaxResults() *int32
	SetRequestId(v string) *ListStrategyExecutionRecordResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListStrategyExecutionRecordResponseBody
	GetTotal() *int64
}

type ListStrategyExecutionRecordResponseBody struct {
	ExecutionRecordList []*ListStrategyExecutionRecordResponseBodyExecutionRecordList `json:"ExecutionRecordList,omitempty" xml:"ExecutionRecordList,omitempty" type:"Repeated"`
	MaxResults          *int32                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total               *int64                                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListStrategyExecutionRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStrategyExecutionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListStrategyExecutionRecordResponseBody) GetExecutionRecordList() []*ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	return s.ExecutionRecordList
}

func (s *ListStrategyExecutionRecordResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStrategyExecutionRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStrategyExecutionRecordResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListStrategyExecutionRecordResponseBody) SetExecutionRecordList(v []*ListStrategyExecutionRecordResponseBodyExecutionRecordList) *ListStrategyExecutionRecordResponseBody {
	s.ExecutionRecordList = v
	return s
}

func (s *ListStrategyExecutionRecordResponseBody) SetMaxResults(v int32) *ListStrategyExecutionRecordResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBody) SetRequestId(v string) *ListStrategyExecutionRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBody) SetTotal(v int64) *ListStrategyExecutionRecordResponseBody {
	s.Total = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStrategyExecutionRecordResponseBodyExecutionRecordList struct {
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExecuteTime  *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	JobParams    *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaName    *string `json:"MediaName,omitempty" xml:"MediaName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s ListStrategyExecutionRecordResponseBodyExecutionRecordList) String() string {
	return dara.Prettify(s)
}

func (s ListStrategyExecutionRecordResponseBodyExecutionRecordList) GoString() string {
	return s.String()
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetJobParams() *string {
	return s.JobParams
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetMediaId() *string {
	return s.MediaId
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetMediaName() *string {
	return s.MediaName
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetStatus() *string {
	return s.Status
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetCompleteTime(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.CompleteTime = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetCreateTime(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.CreateTime = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetErrorCode(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.ErrorCode = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetErrorMessage(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.ErrorMessage = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetExecuteTime(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.ExecuteTime = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetJobParams(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.JobParams = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetMediaId(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.MediaId = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetMediaName(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.MediaName = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetStatus(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.Status = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetStrategyId(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.StrategyId = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetStrategyName(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.StrategyName = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) SetStrategyType(v string) *ListStrategyExecutionRecordResponseBodyExecutionRecordList {
	s.StrategyType = &v
	return s
}

func (s *ListStrategyExecutionRecordResponseBodyExecutionRecordList) Validate() error {
	return dara.Validate(s)
}
