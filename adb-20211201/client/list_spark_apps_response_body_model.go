// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSparkAppsResponseBodyData) *ListSparkAppsResponseBody
	GetData() *ListSparkAppsResponseBodyData
	SetPageNumber(v int64) *ListSparkAppsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSparkAppsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSparkAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListSparkAppsResponseBody
	GetTotalCount() *int64
}

type ListSparkAppsResponseBody struct {
	// The returned data.
	Data *ListSparkAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkAppsResponseBody) GetData() *ListSparkAppsResponseBodyData {
	return s.Data
}

func (s *ListSparkAppsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkAppsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSparkAppsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSparkAppsResponseBody) SetData(v *ListSparkAppsResponseBodyData) *ListSparkAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkAppsResponseBody) SetPageNumber(v int64) *ListSparkAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppsResponseBody) SetPageSize(v int64) *ListSparkAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppsResponseBody) SetRequestId(v string) *ListSparkAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkAppsResponseBody) SetTotalCount(v int64) *ListSparkAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSparkAppsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSparkAppsResponseBodyData struct {
	// The list of application information. Response parameter description:
	//
	// - **Data**: the Spark application template data.
	//
	// - **EstimateExecutionCpuTimeInSeconds**: the CPU time consumed to execute the Spark application, in milliseconds (ms).
	//
	// - **LogRootPath**: the storage path of log files.
	//
	// - **LastAttemptId**: the retry ID.
	//
	// - **WebUiAddress**: the Web UI address.
	//
	// - **SubmittedTimeInMillis**: the time when the Spark application was submitted, in UNIX timestamp format, in milliseconds (ms).
	//
	// - **StartedTimeInMillis**: the time when the Spark application was created, in UNIX timestamp format, in milliseconds (ms).
	//
	// - **LastUpdatedTimeInMillis**: the time when the Spark application was last updated, in UNIX timestamp format, in milliseconds (ms).
	//
	// - **TerminatedTimeInMillis**: the time when the Spark application stopped execution, in UNIX timestamp format, in milliseconds (ms).
	//
	// - **DBClusterId**: the ID of the cluster that executed the Spark application.
	//
	// - **ResourceGroupName**: the name of the job resource group.
	//
	// - **DurationInMillis**: the execution duration of the Spark application, in milliseconds (ms).
	AppInfoList []*SparkAppInfo `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
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
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkAppsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkAppsResponseBodyData) GetAppInfoList() []*SparkAppInfo {
	return s.AppInfoList
}

func (s *ListSparkAppsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkAppsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkAppsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSparkAppsResponseBodyData) SetAppInfoList(v []*SparkAppInfo) *ListSparkAppsResponseBodyData {
	s.AppInfoList = v
	return s
}

func (s *ListSparkAppsResponseBodyData) SetPageNumber(v int64) *ListSparkAppsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppsResponseBodyData) SetPageSize(v int64) *ListSparkAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppsResponseBodyData) SetTotalCount(v int64) *ListSparkAppsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSparkAppsResponseBodyData) Validate() error {
	if s.AppInfoList != nil {
		for _, item := range s.AppInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
