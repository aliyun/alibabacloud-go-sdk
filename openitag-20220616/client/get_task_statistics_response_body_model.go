// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskStatisticsResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskStatisticsResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskStatisticsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskStatisticsResponseBody
	GetSuccess() *bool
	SetTaskStatistics(v *TaskStatistic) *GetTaskStatisticsResponseBody
	GetTaskStatistics() *TaskStatistic
}

type GetTaskStatisticsResponseBody struct {
	// Return encoding. The default value is 0, indicating Normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation Succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Job statistics.
	TaskStatistics *TaskStatistic `json:"TaskStatistics,omitempty" xml:"TaskStatistics,omitempty"`
}

func (s GetTaskStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatisticsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskStatisticsResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskStatisticsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskStatisticsResponseBody) GetTaskStatistics() *TaskStatistic {
	return s.TaskStatistics
}

func (s *GetTaskStatisticsResponseBody) SetCode(v int32) *GetTaskStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetDetails(v string) *GetTaskStatisticsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetErrorCode(v string) *GetTaskStatisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetMessage(v string) *GetTaskStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetRequestId(v string) *GetTaskStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetSuccess(v bool) *GetTaskStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetTaskStatistics(v *TaskStatistic) *GetTaskStatisticsResponseBody {
	s.TaskStatistics = v
	return s
}

func (s *GetTaskStatisticsResponseBody) Validate() error {
	if s.TaskStatistics != nil {
		if err := s.TaskStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}
