// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkAppAttemptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSparkAppAttemptsResponseBodyData) *ListSparkAppAttemptsResponseBody
	GetData() *ListSparkAppAttemptsResponseBodyData
	SetRequestId(v string) *ListSparkAppAttemptsResponseBody
	GetRequestId() *string
}

type ListSparkAppAttemptsResponseBody struct {
	// The returned data.
	Data *ListSparkAppAttemptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkAppAttemptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppAttemptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsResponseBody) GetData() *ListSparkAppAttemptsResponseBodyData {
	return s.Data
}

func (s *ListSparkAppAttemptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSparkAppAttemptsResponseBody) SetData(v *ListSparkAppAttemptsResponseBodyData) *ListSparkAppAttemptsResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkAppAttemptsResponseBody) SetRequestId(v string) *ListSparkAppAttemptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkAppAttemptsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSparkAppAttemptsResponseBodyData struct {
	// The retry information list. Response parameter description:
	//
	// - **AttemptId**: The retry ID.
	//
	// - **State**: The execution state of the application. Valid values:
	//
	//     - **SUBMITTED**: Submitted.
	//
	//     - **STARTING**: Starting.
	//
	//     - **RUNNING**: Running.
	//
	//     - **FAILING**: The task failed. Cleaning up the environment.
	//
	//     - **FAILED**: Failed.
	//
	//     - **KILLING**: Aborting the task. Cleaning up the environment.
	//
	//     - **KILLED**: The task aborted.
	//
	//     - **SUCCEEDING**: The task execution ended. Cleaning up the environment.
	//
	//     - **COMPLETED**: The task execution ended.
	//
	//     - **FATAL**: Unexpected failure.
	//
	//     - **UNKNOWN**: Unknown error.
	//
	// - **Message**: The alert information. If no alert is generated, an empty value is returned.
	//
	// - **Data**: The Spark application template data.
	//
	// - **EstimateExecutionCpuTimeInSeconds**: The CPU time consumed to execute the Spark application, in milliseconds.
	//
	// - **LogRootPath**: The storage path of the log file.
	//
	// - **LastAttemptId**: The ID of the last retry.
	//
	// - **WebUiAddress**: The Web UI address.
	//
	// - **SubmittedTimeInMillis**: The time when the Spark application was submitted. The value is a UNIX timestamp in milliseconds.
	//
	// - **StartedTimeInMillis**: The time when the Spark application was created. The value is a UNIX timestamp in milliseconds.
	//
	// - **LastUpdatedTimeInMillis**: The time when the Spark application was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// - **TerminatedTimeInMillis**: The time when the execution of the Spark application was stopped. The value is a UNIX timestamp in milliseconds.
	//
	// - **DBClusterId**: The ID of the cluster where the Spark application is executed.
	//
	// - **ResourceGroupName**: The name of the job resource group.
	//
	// - **DurationInMillis**: The execution duration of the Spark application, in milliseconds.
	AttemptInfoList []*SparkAttemptInfo `json:"AttemptInfoList,omitempty" xml:"AttemptInfoList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkAppAttemptsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppAttemptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsResponseBodyData) GetAttemptInfoList() []*SparkAttemptInfo {
	return s.AttemptInfoList
}

func (s *ListSparkAppAttemptsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkAppAttemptsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkAppAttemptsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSparkAppAttemptsResponseBodyData) SetAttemptInfoList(v []*SparkAttemptInfo) *ListSparkAppAttemptsResponseBodyData {
	s.AttemptInfoList = v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) SetPageNumber(v int64) *ListSparkAppAttemptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) SetPageSize(v int64) *ListSparkAppAttemptsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) SetTotalCount(v int64) *ListSparkAppAttemptsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) Validate() error {
	if s.AttemptInfoList != nil {
		for _, item := range s.AttemptInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
