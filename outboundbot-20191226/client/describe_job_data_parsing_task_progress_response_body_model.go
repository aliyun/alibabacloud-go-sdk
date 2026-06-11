// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobDataParsingTaskProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobDataParsingTaskProgressResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeJobDataParsingTaskProgressResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeJobDataParsingTaskProgressResponseBody
	GetMessage() *string
	SetProgress(v *DescribeJobDataParsingTaskProgressResponseBodyProgress) *DescribeJobDataParsingTaskProgressResponseBody
	GetProgress() *DescribeJobDataParsingTaskProgressResponseBodyProgress
	SetRequestId(v string) *DescribeJobDataParsingTaskProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobDataParsingTaskProgressResponseBody
	GetSuccess() *bool
}

type DescribeJobDataParsingTaskProgressResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Parsing progress.
	//
	// example:
	//
	// {}
	Progress *DescribeJobDataParsingTaskProgressResponseBodyProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Struct"`
	// ID of the request.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeJobDataParsingTaskProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobDataParsingTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) GetProgress() *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	return s.Progress
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) SetCode(v string) *DescribeJobDataParsingTaskProgressResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) SetHttpStatusCode(v int32) *DescribeJobDataParsingTaskProgressResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) SetMessage(v string) *DescribeJobDataParsingTaskProgressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) SetProgress(v *DescribeJobDataParsingTaskProgressResponseBodyProgress) *DescribeJobDataParsingTaskProgressResponseBody {
	s.Progress = v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) SetRequestId(v string) *DescribeJobDataParsingTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) SetSuccess(v bool) *DescribeJobDataParsingTaskProgressResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBody) Validate() error {
	if s.Progress != nil {
		if err := s.Progress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobDataParsingTaskProgressResponseBodyProgress struct {
	// Error code for a failed parsing task.
	//
	// example:
	//
	// Permission.JobStatus
	FailErrorCode *string `json:"FailErrorCode,omitempty" xml:"FailErrorCode,omitempty"`
	// Reason for failure.
	//
	// example:
	//
	// CreateCorpus
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// Download URL for the file that lists parsing failures. This field is deprecated.
	//
	// example:
	//
	// xxxx
	FeedbackUrl *string `json:"FeedbackUrl,omitempty" xml:"FeedbackUrl,omitempty"`
	// Number of jobs processed so far.
	//
	// example:
	//
	// 2
	HandledJobCount *int32 `json:"HandledJobCount,omitempty" xml:"HandledJobCount,omitempty"`
	// Status of the job.
	//
	// - Pending: The job is waiting to start.
	//
	// - InProgress: The job is running.
	//
	// - Finished: The job completed successfully.
	//
	// - PartialFinished: The job completed partially.
	//
	// - Failed: The job failed.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Total number of jobs in the task.
	//
	// example:
	//
	// 3
	TotalJobCount *int32 `json:"TotalJobCount,omitempty" xml:"TotalJobCount,omitempty"`
}

func (s DescribeJobDataParsingTaskProgressResponseBodyProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobDataParsingTaskProgressResponseBodyProgress) GoString() string {
	return s.String()
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) GetFailErrorCode() *string {
	return s.FailErrorCode
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) GetFailReason() *string {
	return s.FailReason
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) GetFeedbackUrl() *string {
	return s.FeedbackUrl
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) GetHandledJobCount() *int32 {
	return s.HandledJobCount
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) GetTotalJobCount() *int32 {
	return s.TotalJobCount
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) SetFailErrorCode(v string) *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	s.FailErrorCode = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) SetFailReason(v string) *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	s.FailReason = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) SetFeedbackUrl(v string) *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	s.FeedbackUrl = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) SetHandledJobCount(v int32) *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	s.HandledJobCount = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) SetStatus(v string) *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	s.Status = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) SetTotalJobCount(v int32) *DescribeJobDataParsingTaskProgressResponseBodyProgress {
	s.TotalJobCount = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponseBodyProgress) Validate() error {
	return dara.Validate(s)
}
