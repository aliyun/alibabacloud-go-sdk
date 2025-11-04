// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveRecordJobs(v []*ListLiveRecordJobsResponseBodyLiveRecordJobs) *ListLiveRecordJobsResponseBody
	GetLiveRecordJobs() []*ListLiveRecordJobsResponseBodyLiveRecordJobs
	SetPageNo(v int64) *ListLiveRecordJobsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListLiveRecordJobsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListLiveRecordJobsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLiveRecordJobsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListLiveRecordJobsResponseBody
	GetTotalCount() *int64
}

type ListLiveRecordJobsResponseBody struct {
	// The list of live stream recording jobs.
	LiveRecordJobs []*ListLiveRecordJobsResponseBodyLiveRecordJobs `json:"LiveRecordJobs,omitempty" xml:"LiveRecordJobs,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A27DFFA4-F272-5563-8363-CB0BC42740BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 180
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveRecordJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRecordJobsResponseBody) GetLiveRecordJobs() []*ListLiveRecordJobsResponseBodyLiveRecordJobs {
	return s.LiveRecordJobs
}

func (s *ListLiveRecordJobsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLiveRecordJobsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLiveRecordJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRecordJobsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLiveRecordJobsResponseBody) SetLiveRecordJobs(v []*ListLiveRecordJobsResponseBodyLiveRecordJobs) *ListLiveRecordJobsResponseBody {
	s.LiveRecordJobs = v
	return s
}

func (s *ListLiveRecordJobsResponseBody) SetPageNo(v int64) *ListLiveRecordJobsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordJobsResponseBody) SetPageSize(v int64) *ListLiveRecordJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordJobsResponseBody) SetRequestId(v string) *ListLiveRecordJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRecordJobsResponseBody) SetSortBy(v string) *ListLiveRecordJobsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordJobsResponseBody) SetTotalCount(v int64) *ListLiveRecordJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveRecordJobsResponseBody) Validate() error {
	if s.LiveRecordJobs != nil {
		for _, item := range s.LiveRecordJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveRecordJobsResponseBodyLiveRecordJobs struct {
	// The time when the job was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-20T03:26:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the recording job.
	//
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the recording job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// https://example.com/imsnotify
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The storage address of the recording.
	RecordOutput *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty" type:"Struct"`
	// The state of the recording job.
	//
	// example:
	//
	// paused
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The URL of the live stream.
	StreamInput *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput `json:"StreamInput,omitempty" xml:"StreamInput,omitempty" type:"Struct"`
	// The ID of the recording template.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the recording template.
	//
	// example:
	//
	// test template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListLiveRecordJobsResponseBodyLiveRecordJobs) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordJobsResponseBodyLiveRecordJobs) GoString() string {
	return s.String()
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetName() *string {
	return s.Name
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetRecordOutput() *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput {
	return s.RecordOutput
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetStatus() *string {
	return s.Status
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetStreamInput() *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput {
	return s.StreamInput
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetCreateTime(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.CreateTime = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetJobId(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.JobId = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetName(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.Name = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetNotifyUrl(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.NotifyUrl = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetRecordOutput(v *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.RecordOutput = v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetStatus(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.Status = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetStreamInput(v *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.StreamInput = v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetTemplateId(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.TemplateId = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) SetTemplateName(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobs {
	s.TemplateName = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobs) Validate() error {
	if s.RecordOutput != nil {
		if err := s.RecordOutput.Validate(); err != nil {
			return err
		}
	}
	if s.StreamInput != nil {
		if err := s.StreamInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput struct {
	// The bucket name.
	//
	// example:
	//
	// imsbucket1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The endpoint of the storage service.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The type of the storage address.
	//
	// Valid values:
	//
	// 	- vod
	//
	// 	- oss
	//
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) GoString() string {
	return s.String()
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) GetBucket() *string {
	return s.Bucket
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) GetType() *string {
	return s.Type
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) SetBucket(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput {
	s.Bucket = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) SetEndpoint(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput {
	s.Endpoint = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) SetType(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput {
	s.Type = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsRecordOutput) Validate() error {
	return dara.Validate(s)
}

type ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput struct {
	// The type of the live stream URL.
	//
	// example:
	//
	// rtmp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://example-live.com/live/stream1
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) GoString() string {
	return s.String()
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) GetType() *string {
	return s.Type
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) GetUrl() *string {
	return s.Url
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) SetType(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput {
	s.Type = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) SetUrl(v string) *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput {
	s.Url = &v
	return s
}

func (s *ListLiveRecordJobsResponseBodyLiveRecordJobsStreamInput) Validate() error {
	return dara.Validate(s)
}
