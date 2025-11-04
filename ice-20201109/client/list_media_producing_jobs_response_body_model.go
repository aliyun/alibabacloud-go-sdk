// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaProducingJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListMediaProducingJobsResponseBody
	GetMaxResults() *string
	SetMediaProducingJobList(v []*ListMediaProducingJobsResponseBodyMediaProducingJobList) *ListMediaProducingJobsResponseBody
	GetMediaProducingJobList() []*ListMediaProducingJobsResponseBodyMediaProducingJobList
	SetNextToken(v string) *ListMediaProducingJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMediaProducingJobsResponseBody
	GetRequestId() *string
}

type ListMediaProducingJobsResponseBody struct {
	// The maximum number of entries returned.
	//
	// Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The queried media editing and production jobs.
	MediaProducingJobList []*ListMediaProducingJobsResponseBodyMediaProducingJobList `json:"MediaProducingJobList,omitempty" xml:"MediaProducingJobList,omitempty" type:"Repeated"`
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
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaProducingJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaProducingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListMediaProducingJobsResponseBody) GetMediaProducingJobList() []*ListMediaProducingJobsResponseBodyMediaProducingJobList {
	return s.MediaProducingJobList
}

func (s *ListMediaProducingJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaProducingJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaProducingJobsResponseBody) SetMaxResults(v string) *ListMediaProducingJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMediaProducingJobsResponseBody) SetMediaProducingJobList(v []*ListMediaProducingJobsResponseBodyMediaProducingJobList) *ListMediaProducingJobsResponseBody {
	s.MediaProducingJobList = v
	return s
}

func (s *ListMediaProducingJobsResponseBody) SetNextToken(v string) *ListMediaProducingJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaProducingJobsResponseBody) SetRequestId(v string) *ListMediaProducingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBody) Validate() error {
	if s.MediaProducingJobList != nil {
		for _, item := range s.MediaProducingJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMediaProducingJobsResponseBodyMediaProducingJobList struct {
	// The template material parameters.
	//
	// example:
	//
	// {"Text1":"text","Text0":"text","Media1":"mediaId","Media0":"mediaId"}
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-21T16:40:30Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-21T16:40:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The duration of the output file. Unit: seconds.
	//
	// example:
	//
	// 15.5
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the online editing job.
	//
	// example:
	//
	// ******8750b54e3c976a47da6f******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The media asset ID of the output file.
	//
	// example:
	//
	// 0ce4ea70f52471edab61f7e7d6786302
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the output file.
	//
	// example:
	//
	// http://your-bucket.oss-cn-shanghai.aliyuncs.com/your-video.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The returned message. Note: Pay attention to this parameter if the job failed.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2022-03-21T16:41:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ******faa3b542f5a6135217e3******
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The job state.
	//
	// example:
	//
	// Sucess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the online editing template.
	//
	// example:
	//
	// cb786a39c5d44cecb23d8c864facffc1
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The user-defined data in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListMediaProducingJobsResponseBodyMediaProducingJobList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaProducingJobsResponseBodyMediaProducingJobList) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetCode() *string {
	return s.Code
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetDuration() *float32 {
	return s.Duration
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetMediaURL() *string {
	return s.MediaURL
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetMessage() *string {
	return s.Message
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetStatus() *string {
	return s.Status
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) GetUserData() *string {
	return s.UserData
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetClipsParam(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ClipsParam = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCode(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Code = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCompleteTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.CompleteTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCreateTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.CreateTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetDuration(v float32) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Duration = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetJobId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.JobId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMediaId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.MediaId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMediaURL(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.MediaURL = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMessage(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Message = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetModifiedTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetProjectId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ProjectId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetStatus(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Status = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetTemplateId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.TemplateId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetUserData(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.UserData = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) Validate() error {
	return dara.Validate(s)
}
