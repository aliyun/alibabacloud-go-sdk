// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListSmartJobsResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListSmartJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSmartJobsResponseBody
	GetRequestId() *string
	SetSmartJobList(v []*ListSmartJobsResponseBodySmartJobList) *ListSmartJobsResponseBody
	GetSmartJobList() []*ListSmartJobsResponseBodySmartJobList
	SetTotalCount(v string) *ListSmartJobsResponseBody
	GetTotalCount() *string
}

type ListSmartJobsResponseBody struct {
	// The maximum number of entries returned on a single page. The value is set to the maximum number of entries returned on each page except for the last page. Valid example: 10,10,5. Invalid example: 10,5,10.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// CBB6BC61D08
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****9262E3DA-07FA-4862-FCBB6BC61D08*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried intelligent jobs.
	SmartJobList []*ListSmartJobsResponseBodySmartJobList `json:"SmartJobList,omitempty" xml:"SmartJobList,omitempty" type:"Repeated"`
	// Optional. The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 110
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSmartJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmartJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListSmartJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSmartJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmartJobsResponseBody) GetSmartJobList() []*ListSmartJobsResponseBodySmartJobList {
	return s.SmartJobList
}

func (s *ListSmartJobsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListSmartJobsResponseBody) SetMaxResults(v string) *ListSmartJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSmartJobsResponseBody) SetNextToken(v string) *ListSmartJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSmartJobsResponseBody) SetRequestId(v string) *ListSmartJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartJobsResponseBody) SetSmartJobList(v []*ListSmartJobsResponseBodySmartJobList) *ListSmartJobsResponseBody {
	s.SmartJobList = v
	return s
}

func (s *ListSmartJobsResponseBody) SetTotalCount(v string) *ListSmartJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSmartJobsResponseBody) Validate() error {
	if s.SmartJobList != nil {
		for _, item := range s.SmartJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmartJobsResponseBodySmartJobList struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The job description.
	//
	// example:
	//
	// 测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The editing configurations.
	//
	// example:
	//
	// {"AudioConfig":{},"InputConfig":""}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configurations.
	InputConfig *ListSmartJobsResponseBodySmartJobListInputConfig `json:"InputConfig,omitempty" xml:"InputConfig,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job state.
	//
	// Valid values:
	//
	// 	- Finished: The job is complete.
	//
	// 	- Failed: The job failed.
	//
	// 	- Executing: The job is in progress.
	//
	// 	- Created: The job is created.
	//
	// example:
	//
	// Finished
	JobState *string `json:"JobState,omitempty" xml:"JobState,omitempty"`
	// The job type.
	//
	// Valid values:
	//
	// 	- ASR: ASR job.
	//
	// 	- DynamicChart: dynamic chart job.
	//
	// 	- TextToSpeech: intelligent audio production job.
	//
	// example:
	//
	// ASR
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The output configurations.
	OutputConfig *ListSmartJobsResponseBodySmartJobListOutputConfig `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty" type:"Struct"`
	// The job title.
	//
	// example:
	//
	// 测试标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// {"user":"data"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1084506228******
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSmartJobsResponseBodySmartJobList) String() string {
	return dara.Prettify(s)
}

func (s ListSmartJobsResponseBodySmartJobList) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBodySmartJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSmartJobsResponseBodySmartJobList) GetDescription() *string {
	return s.Description
}

func (s *ListSmartJobsResponseBodySmartJobList) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *ListSmartJobsResponseBodySmartJobList) GetInputConfig() *ListSmartJobsResponseBodySmartJobListInputConfig {
	return s.InputConfig
}

func (s *ListSmartJobsResponseBodySmartJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListSmartJobsResponseBodySmartJobList) GetJobState() *string {
	return s.JobState
}

func (s *ListSmartJobsResponseBodySmartJobList) GetJobType() *string {
	return s.JobType
}

func (s *ListSmartJobsResponseBodySmartJobList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListSmartJobsResponseBodySmartJobList) GetOutputConfig() *ListSmartJobsResponseBodySmartJobListOutputConfig {
	return s.OutputConfig
}

func (s *ListSmartJobsResponseBodySmartJobList) GetTitle() *string {
	return s.Title
}

func (s *ListSmartJobsResponseBodySmartJobList) GetUserData() *string {
	return s.UserData
}

func (s *ListSmartJobsResponseBodySmartJobList) GetUserId() *int64 {
	return s.UserId
}

func (s *ListSmartJobsResponseBodySmartJobList) SetCreateTime(v string) *ListSmartJobsResponseBodySmartJobList {
	s.CreateTime = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetDescription(v string) *ListSmartJobsResponseBodySmartJobList {
	s.Description = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetEditingConfig(v string) *ListSmartJobsResponseBodySmartJobList {
	s.EditingConfig = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetInputConfig(v *ListSmartJobsResponseBodySmartJobListInputConfig) *ListSmartJobsResponseBodySmartJobList {
	s.InputConfig = v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetJobId(v string) *ListSmartJobsResponseBodySmartJobList {
	s.JobId = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetJobState(v string) *ListSmartJobsResponseBodySmartJobList {
	s.JobState = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetJobType(v string) *ListSmartJobsResponseBodySmartJobList {
	s.JobType = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetModifiedTime(v string) *ListSmartJobsResponseBodySmartJobList {
	s.ModifiedTime = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetOutputConfig(v *ListSmartJobsResponseBodySmartJobListOutputConfig) *ListSmartJobsResponseBodySmartJobList {
	s.OutputConfig = v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetTitle(v string) *ListSmartJobsResponseBodySmartJobList {
	s.Title = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetUserData(v string) *ListSmartJobsResponseBodySmartJobList {
	s.UserData = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetUserId(v int64) *ListSmartJobsResponseBodySmartJobList {
	s.UserId = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) Validate() error {
	if s.InputConfig != nil {
		if err := s.InputConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OutputConfig != nil {
		if err := s.OutputConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSmartJobsResponseBodySmartJobListInputConfig struct {
	// The information about the input file.
	//
	// example:
	//
	// oss://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	// The keyword information.
	//
	// example:
	//
	// 测试关键词
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListSmartJobsResponseBodySmartJobListInputConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSmartJobsResponseBodySmartJobListInputConfig) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) GetInputFile() *string {
	return s.InputFile
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) SetInputFile(v string) *ListSmartJobsResponseBodySmartJobListInputConfig {
	s.InputFile = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) SetKeyword(v string) *ListSmartJobsResponseBodySmartJobListInputConfig {
	s.Keyword = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) Validate() error {
	return dara.Validate(s)
}

type ListSmartJobsResponseBodySmartJobListOutputConfig struct {
	// The Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS object.
	//
	// example:
	//
	// test-object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListSmartJobsResponseBodySmartJobListOutputConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSmartJobsResponseBodySmartJobListOutputConfig) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) GetBucket() *string {
	return s.Bucket
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) GetObject() *string {
	return s.Object
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) SetBucket(v string) *ListSmartJobsResponseBodySmartJobListOutputConfig {
	s.Bucket = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) SetObject(v string) *ListSmartJobsResponseBodySmartJobListOutputConfig {
	s.Object = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) Validate() error {
	return dara.Validate(s)
}
