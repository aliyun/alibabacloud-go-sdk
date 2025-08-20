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
	return dara.Validate(s)
}

type ListSparkAppAttemptsResponseBodyData struct {
	// The queried attempts. Fields in the response parameter:
	//
	// 	- **AttemptId**: the attempt ID.
	//
	// 	- **State**: the state of the Spark application. Valid values:
	//
	//     	- **SUBMITTED**
	//
	//     	- **STARTING**
	//
	//     	- **RUNNING**
	//
	//     	- **FAILING**
	//
	//     	- **FAILED**
	//
	//     	- **KILLING**
	//
	//     	- **KILLED**
	//
	//     	- **SUCCEEDING**
	//
	//     	- **COMPLETED**
	//
	//     	- **FATAL**
	//
	//     	- **UNKNOWN**
	//
	// 	- **Message**: the alert message that is returned. If no alert is generated, null is returned.
	//
	// 	- **Data*	- the data of the Spark application template.
	//
	// 	- **EstimateExecutionCpuTimeInSeconds**: the amount of time that is required to consume CPU resources for running the Spark application. Unit: milliseconds.
	//
	// 	- **LogRootPath**: the storage path of log files.
	//
	// 	- **LastAttemptId**: the ID of the last attempt.
	//
	// 	- **WebUiAddress**: the web UI URL.
	//
	// 	- **SubmittedTimeInMillis**: the time when the Spark application was submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **StartedTimeInMillis**: the time when the Spark application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **TerminatedTimeInMillis**: the time when the Spark application task was terminated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **DBClusterId**: the ID of the cluster on which the Spark application runs.
	//
	// 	- **ResourceGroupName**: the name of the job resource group.
	//
	// 	- **DurationInMillis**: the amount of time that is required to run the Spark application. Unit: milliseconds.
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
	// The total number of entries returned.
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
	return dara.Validate(s)
}
