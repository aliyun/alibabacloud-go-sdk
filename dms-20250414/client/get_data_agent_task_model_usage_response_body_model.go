// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataAgentTaskModelUsageResponseBodyData) *GetDataAgentTaskModelUsageResponseBody
	GetData() *GetDataAgentTaskModelUsageResponseBodyData
	SetErrorCode(v string) *GetDataAgentTaskModelUsageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataAgentTaskModelUsageResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataAgentTaskModelUsageResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDataAgentTaskModelUsageResponseBody
	GetSuccess() *string
}

type GetDataAgentTaskModelUsageResponseBody struct {
	// The summary data of model usage for DataAgent analysis tasks.
	Data *GetDataAgentTaskModelUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// DMS-DA-40411
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 67E910F2-***695C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataAgentTaskModelUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageResponseBody) GetData() *GetDataAgentTaskModelUsageResponseBodyData {
	return s.Data
}

func (s *GetDataAgentTaskModelUsageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataAgentTaskModelUsageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataAgentTaskModelUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAgentTaskModelUsageResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDataAgentTaskModelUsageResponseBody) SetData(v *GetDataAgentTaskModelUsageResponseBodyData) *GetDataAgentTaskModelUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBody) SetErrorCode(v string) *GetDataAgentTaskModelUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBody) SetErrorMessage(v string) *GetDataAgentTaskModelUsageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBody) SetRequestId(v string) *GetDataAgentTaskModelUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBody) SetSuccess(v string) *GetDataAgentTaskModelUsageResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataAgentTaskModelUsageResponseBodyData struct {
	AccelerationRatio       *float64 `json:"AccelerationRatio,omitempty" xml:"AccelerationRatio,omitempty"`
	RateLimitedSessionCount *int64   `json:"RateLimitedSessionCount,omitempty" xml:"RateLimitedSessionCount,omitempty"`
	TotalLlmWaitDuration    *float64 `json:"TotalLlmWaitDuration,omitempty" xml:"TotalLlmWaitDuration,omitempty"`
	TotalSessionCount       *int64   `json:"TotalSessionCount,omitempty" xml:"TotalSessionCount,omitempty"`
	// The peak TPM (tokens per minute) within the query time range, which is the maximum number of tokens consumed per minute.
	//
	// example:
	//
	// 42000
	PeakTpm *int64 `json:"peakTpm,omitempty" xml:"peakTpm,omitempty"`
	// The total number of model calls within the query time range.
	//
	// example:
	//
	// 1280
	TotalCallCount *int64 `json:"totalCallCount,omitempty" xml:"totalCallCount,omitempty"`
	// The total number of tokens consumed within the query time range.
	//
	// example:
	//
	// 3560000
	TotalTokenConsumed *int64 `json:"totalTokenConsumed,omitempty" xml:"totalTokenConsumed,omitempty"`
	// The number of models used within the query time range.
	//
	// example:
	//
	// 5
	UsedModels *int64 `json:"usedModels,omitempty" xml:"usedModels,omitempty"`
}

func (s GetDataAgentTaskModelUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetAccelerationRatio() *float64 {
	return s.AccelerationRatio
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetRateLimitedSessionCount() *int64 {
	return s.RateLimitedSessionCount
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetTotalLlmWaitDuration() *float64 {
	return s.TotalLlmWaitDuration
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetTotalSessionCount() *int64 {
	return s.TotalSessionCount
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetPeakTpm() *int64 {
	return s.PeakTpm
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetTotalCallCount() *int64 {
	return s.TotalCallCount
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetTotalTokenConsumed() *int64 {
	return s.TotalTokenConsumed
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) GetUsedModels() *int64 {
	return s.UsedModels
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetAccelerationRatio(v float64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.AccelerationRatio = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetRateLimitedSessionCount(v int64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.RateLimitedSessionCount = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetTotalLlmWaitDuration(v float64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.TotalLlmWaitDuration = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetTotalSessionCount(v int64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.TotalSessionCount = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetPeakTpm(v int64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.PeakTpm = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetTotalCallCount(v int64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.TotalCallCount = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetTotalTokenConsumed(v int64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.TotalTokenConsumed = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) SetUsedModels(v int64) *GetDataAgentTaskModelUsageResponseBodyData {
	s.UsedModels = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
