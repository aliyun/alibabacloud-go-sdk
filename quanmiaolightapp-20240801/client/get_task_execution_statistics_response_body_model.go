// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskExecutionStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskExecutionStatisticsResponseBody
	GetCode() *string
	SetData(v *GetTaskExecutionStatisticsResponseBodyData) *GetTaskExecutionStatisticsResponseBody
	GetData() *GetTaskExecutionStatisticsResponseBodyData
	SetHttpStatusCode(v int32) *GetTaskExecutionStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTaskExecutionStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskExecutionStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskExecutionStatisticsResponseBody
	GetSuccess() *bool
}

type GetTaskExecutionStatisticsResponseBody struct {
	// example:
	//
	// Success
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTaskExecutionStatisticsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-GHIJ-KLMNOPQRST
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTaskExecutionStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskExecutionStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskExecutionStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskExecutionStatisticsResponseBody) GetData() *GetTaskExecutionStatisticsResponseBodyData {
	return s.Data
}

func (s *GetTaskExecutionStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTaskExecutionStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskExecutionStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskExecutionStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskExecutionStatisticsResponseBody) SetCode(v string) *GetTaskExecutionStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBody) SetData(v *GetTaskExecutionStatisticsResponseBodyData) *GetTaskExecutionStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBody) SetHttpStatusCode(v int32) *GetTaskExecutionStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBody) SetMessage(v string) *GetTaskExecutionStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBody) SetRequestId(v string) *GetTaskExecutionStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBody) SetSuccess(v bool) *GetTaskExecutionStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskExecutionStatisticsResponseBodyData struct {
	// example:
	//
	// 100
	CompletedCount *int64 `json:"completedCount,omitempty" xml:"completedCount,omitempty"`
	// example:
	//
	// 5
	RunningCount *int64                                                    `json:"runningCount,omitempty" xml:"runningCount,omitempty"`
	TpmPerMinute []*GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute `json:"tpmPerMinute,omitempty" xml:"tpmPerMinute,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	WaitingCount *int64 `json:"waitingCount,omitempty" xml:"waitingCount,omitempty"`
}

func (s GetTaskExecutionStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskExecutionStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskExecutionStatisticsResponseBodyData) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *GetTaskExecutionStatisticsResponseBodyData) GetRunningCount() *int64 {
	return s.RunningCount
}

func (s *GetTaskExecutionStatisticsResponseBodyData) GetTpmPerMinute() []*GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute {
	return s.TpmPerMinute
}

func (s *GetTaskExecutionStatisticsResponseBodyData) GetWaitingCount() *int64 {
	return s.WaitingCount
}

func (s *GetTaskExecutionStatisticsResponseBodyData) SetCompletedCount(v int64) *GetTaskExecutionStatisticsResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBodyData) SetRunningCount(v int64) *GetTaskExecutionStatisticsResponseBodyData {
	s.RunningCount = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBodyData) SetTpmPerMinute(v []*GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) *GetTaskExecutionStatisticsResponseBodyData {
	s.TpmPerMinute = v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBodyData) SetWaitingCount(v int64) *GetTaskExecutionStatisticsResponseBodyData {
	s.WaitingCount = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBodyData) Validate() error {
	if s.TpmPerMinute != nil {
		for _, item := range s.TpmPerMinute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute struct {
	// example:
	//
	// 3
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 2024-08-01 10:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) String() string {
	return dara.Prettify(s)
}

func (s GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) GoString() string {
	return s.String()
}

func (s *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) GetCount() *int64 {
	return s.Count
}

func (s *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) GetTime() *string {
	return s.Time
}

func (s *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) SetCount(v int64) *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute {
	s.Count = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) SetTime(v string) *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute {
	s.Time = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponseBodyDataTpmPerMinute) Validate() error {
	return dara.Validate(s)
}
