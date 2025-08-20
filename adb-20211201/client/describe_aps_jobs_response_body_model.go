// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAPSJobs(v []*DescribeApsJobsResponseBodyAPSJobs) *DescribeApsJobsResponseBody
	GetAPSJobs() []*DescribeApsJobsResponseBodyAPSJobs
	SetCode(v string) *DescribeApsJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeApsJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeApsJobsResponseBody
	GetMessage() *string
	SetPageNumber(v string) *DescribeApsJobsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeApsJobsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeApsJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApsJobsResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *DescribeApsJobsResponseBody
	GetTotalCount() *string
}

type DescribeApsJobsResponseBody struct {
	// The queried APS jobs.
	//
	// example:
	//
	// -
	APSJobs []*DescribeApsJobsResponseBodyAPSJobs `json:"APSJobs,omitempty" xml:"APSJobs,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// ok
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// ok
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-7F9D-5DBD-993E-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsJobsResponseBody) GetAPSJobs() []*DescribeApsJobsResponseBodyAPSJobs {
	return s.APSJobs
}

func (s *DescribeApsJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApsJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeApsJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApsJobsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeApsJobsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeApsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApsJobsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeApsJobsResponseBody) SetAPSJobs(v []*DescribeApsJobsResponseBodyAPSJobs) *DescribeApsJobsResponseBody {
	s.APSJobs = v
	return s
}

func (s *DescribeApsJobsResponseBody) SetCode(v string) *DescribeApsJobsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetHttpStatusCode(v int32) *DescribeApsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetMessage(v string) *DescribeApsJobsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetPageNumber(v string) *DescribeApsJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetPageSize(v string) *DescribeApsJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetRequestId(v string) *DescribeApsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetSuccess(v bool) *DescribeApsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApsJobsResponseBody) SetTotalCount(v string) *DescribeApsJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApsJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApsJobsResponseBodyAPSJobs struct {
	// The job ID.
	//
	// example:
	//
	// aps-******
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The name of the APS job.
	//
	// example:
	//
	// data-sync-******
	ApsJobName *string `json:"ApsJobName,omitempty" xml:"ApsJobName,omitempty"`
	// The time when the APS job was created.
	//
	// example:
	//
	// 2022-06-28 15:00:04
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The synchronization latency.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The destination cluster ID.
	//
	// example:
	//
	// amv-*******
	DestinationInstanceID *string `json:"DestinationInstanceID,omitempty" xml:"DestinationInstanceID,omitempty"`
	// The error message.
	//
	// example:
	//
	// OK
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The progress.
	//
	// example:
	//
	// 30: The progress is 30%.
	Projress *string `json:"Projress,omitempty" xml:"Projress,omitempty"`
	// The ID of the source instance or cluster.
	//
	// example:
	//
	// pc-******
	SourceInstanceID *string `json:"SourceInstanceID,omitempty" xml:"SourceInstanceID,omitempty"`
	// The status of the APS job.
	//
	// example:
	//
	// -
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// -
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
}

func (s DescribeApsJobsResponseBodyAPSJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobsResponseBodyAPSJobs) GoString() string {
	return s.String()
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetApsJobName() *string {
	return s.ApsJobName
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetDestinationInstanceID() *string {
	return s.DestinationInstanceID
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetProjress() *string {
	return s.Projress
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetSourceInstanceID() *string {
	return s.SourceInstanceID
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetStatus() *string {
	return s.Status
}

func (s *DescribeApsJobsResponseBodyAPSJobs) GetSubStatus() *string {
	return s.SubStatus
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetApsJobId(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.ApsJobId = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetApsJobName(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.ApsJobName = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetCreateTime(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.CreateTime = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetDelay(v int64) *DescribeApsJobsResponseBodyAPSJobs {
	s.Delay = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetDestinationInstanceID(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.DestinationInstanceID = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetErrMessage(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.ErrMessage = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetProjress(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.Projress = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetSourceInstanceID(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.SourceInstanceID = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetStatus(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.Status = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) SetSubStatus(v string) *DescribeApsJobsResponseBodyAPSJobs {
	s.SubStatus = &v
	return s
}

func (s *DescribeApsJobsResponseBodyAPSJobs) Validate() error {
	return dara.Validate(s)
}
