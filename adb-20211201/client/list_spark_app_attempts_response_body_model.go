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
	// The list of retry information. Metric description:
	//
	// - **AttemptId**: the retry ID.
	//
	// - **State**: the execute status of the application. Valid values:
	//
	//     - **SUBMITTED**: commit.
	//
	//     - **STARTING**: starting.
	//
	//     - **RUNNING**: executing.
	//
	//     - **FAILING**: the node failed and the environment is being cleaned up.
	//
	//     - **FAILED**: failed.
	//
	//     - **KILLING**: aborting the task and cleaning up the environment.
	//
	//     - **KILLED**: the task is aborted.
	//
	//     - **SUCCEEDING**: the node execution is complete and the environment is being cleaned up.
	//
	//     - **COMPLETED**: the node execution is complete.
	//
	//     - **FATAL**: unexpected failure.
	//
	//     - **UNKNOWN**: unknown fault.
	//
	// - **Message**: the alerting message. This parameter is empty if no alerting is generated.
	//
	// - **Data**: the Spark application template data.
	//
	// - **EstimateExecutionCpuTimeInSeconds**: the CPU time consumed to execute the Spark application, in milliseconds (ms).
	//
	// - **LogRootPath**: the storage path of log files.
	//
	// - **LastAttemptId**: the ID of the last retry.
	//
	// - **WebUiAddress**: the web UI address.
	//
	// - **SubmittedTimeInMillis**: the time when the Spark application was committed. This value is a UNIX timestamp in milliseconds (ms).
	//
	// - **StartedTimeInMillis**: the time when the Spark application was created. This value is a UNIX timestamp in milliseconds (ms).
	//
	// - **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. This value is a UNIX timestamp in milliseconds (ms).
	//
	// - **TerminatedTimeInMillis**: the time when the Spark application stopped executing. This value is a UNIX timestamp in milliseconds (ms).
	//
	// - **DBClusterId**: the ID of the cluster that executes the Spark application.
	//
	// - **ResourceGroupName**: the name of the Job-type resource group.
	//
	// - **DurationInMillis**: the execution duration of the Spark application, in milliseconds (ms).
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
