// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchMediaProducingJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEditingBatchJobList(v []*ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) *ListBatchMediaProducingJobsResponseBody
	GetEditingBatchJobList() []*ListBatchMediaProducingJobsResponseBodyEditingBatchJobList
	SetMaxResults(v int32) *ListBatchMediaProducingJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListBatchMediaProducingJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBatchMediaProducingJobsResponseBody
	GetRequestId() *string
}

type ListBatchMediaProducingJobsResponseBody struct {
	// The queried quick video production jobs.
	EditingBatchJobList []*ListBatchMediaProducingJobsResponseBodyEditingBatchJobList `json:"EditingBatchJobList,omitempty" xml:"EditingBatchJobList,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 8EqYpQbZ6Eh7+Zz8DxVYoQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBatchMediaProducingJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchMediaProducingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchMediaProducingJobsResponseBody) GetEditingBatchJobList() []*ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	return s.EditingBatchJobList
}

func (s *ListBatchMediaProducingJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchMediaProducingJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchMediaProducingJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchMediaProducingJobsResponseBody) SetEditingBatchJobList(v []*ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) *ListBatchMediaProducingJobsResponseBody {
	s.EditingBatchJobList = v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBody) SetMaxResults(v int32) *ListBatchMediaProducingJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBody) SetNextToken(v string) *ListBatchMediaProducingJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBody) SetRequestId(v string) *ListBatchMediaProducingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBatchMediaProducingJobsResponseBodyEditingBatchJobList struct {
	// The time when the job was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-09T06:38:09Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-09T06:36:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The editing configurations. For more information, see [EditingConfig](~~2692547#1be9bba03b7qu~~).
	//
	// example:
	//
	// {
	//
	//   "MediaConfig": {
	//
	//       "Volume": 0
	//
	//   },
	//
	//   "SpeechConfig": {
	//
	//       "Volume": 1
	//
	//   },
	//
	//  "BackgroundMusicConfig": {
	//
	//       "Volume": 0.3
	//
	//   }
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The extended information of the job.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The input configurations.
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The ID of the quick video production job.
	//
	// example:
	//
	// ******7ecbee4c6d9b8474498e******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job type.
	//
	// Valid values:
	//
	// 	- Script: script-based editing job that mixes media assets.
	//
	// 	- Smart_Mix: intelligent editing job that mixes media assets.
	//
	// example:
	//
	// Script
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2023-06-09T06:37:58Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The output configurations. For more information, see [OutputConfig](~~2692547#447b928fcbuoa~~).
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://xxx.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Count": 20,
	//
	//   "MaxDuration": 15,
	//
	//   "Width": 1080,
	//
	//   "Height": 1920,
	//
	//   "Video": {"Crf": 27}
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The job state.
	//
	// Valid values:
	//
	// 	- Finished
	//
	// 	- Init
	//
	// 	- Failed
	//
	// 	- Processing
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user-defined data in the JSON format, which can be up to 512 bytes in length. You can specify a custom callback URL. For more information, see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) String() string {
	return dara.Prettify(s)
}

func (s ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GoString() string {
	return s.String()
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetExtend() *string {
	return s.Extend
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetInputConfig() *string {
	return s.InputConfig
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetJobType() *string {
	return s.JobType
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetStatus() *string {
	return s.Status
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) GetUserData() *string {
	return s.UserData
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetCompleteTime(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.CompleteTime = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetCreateTime(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.CreateTime = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetEditingConfig(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.EditingConfig = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetExtend(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.Extend = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetInputConfig(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.InputConfig = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetJobId(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.JobId = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetJobType(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.JobType = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetModifiedTime(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.ModifiedTime = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetOutputConfig(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.OutputConfig = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetStatus(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.Status = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) SetUserData(v string) *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList {
	s.UserData = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponseBodyEditingBatchJobList) Validate() error {
	return dara.Validate(s)
}
