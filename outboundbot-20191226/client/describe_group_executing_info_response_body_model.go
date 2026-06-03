// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupExecutingInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeGroupExecutingInfoResponseBody
	GetCode() *string
	SetExecutingInfo(v *DescribeGroupExecutingInfoResponseBodyExecutingInfo) *DescribeGroupExecutingInfoResponseBody
	GetExecutingInfo() *DescribeGroupExecutingInfoResponseBodyExecutingInfo
	SetGroupId(v string) *DescribeGroupExecutingInfoResponseBody
	GetGroupId() *string
	SetHttpStatusCode(v int32) *DescribeGroupExecutingInfoResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *DescribeGroupExecutingInfoResponseBody
	GetInstanceId() *string
	SetMessage(v string) *DescribeGroupExecutingInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeGroupExecutingInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeGroupExecutingInfoResponseBody
	GetSuccess() *bool
}

type DescribeGroupExecutingInfoResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	ExecutingInfo *DescribeGroupExecutingInfoResponseBodyExecutingInfo `json:"ExecutingInfo,omitempty" xml:"ExecutingInfo,omitempty" type:"Struct"`
	// example:
	//
	// b24d321a-2a74-4dd1-a0ba-4ab09cef6652
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// c46001bc-3ead-4bfd-9a69-4b5b66a4a3f4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGroupExecutingInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupExecutingInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupExecutingInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeGroupExecutingInfoResponseBody) GetExecutingInfo() *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	return s.ExecutingInfo
}

func (s *DescribeGroupExecutingInfoResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupExecutingInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeGroupExecutingInfoResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGroupExecutingInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeGroupExecutingInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupExecutingInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeGroupExecutingInfoResponseBody) SetCode(v string) *DescribeGroupExecutingInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetExecutingInfo(v *DescribeGroupExecutingInfoResponseBodyExecutingInfo) *DescribeGroupExecutingInfoResponseBody {
	s.ExecutingInfo = v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetGroupId(v string) *DescribeGroupExecutingInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetHttpStatusCode(v int32) *DescribeGroupExecutingInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetInstanceId(v string) *DescribeGroupExecutingInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetMessage(v string) *DescribeGroupExecutingInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetRequestId(v string) *DescribeGroupExecutingInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) SetSuccess(v bool) *DescribeGroupExecutingInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBody) Validate() error {
	if s.ExecutingInfo != nil {
		if err := s.ExecutingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupExecutingInfoResponseBodyExecutingInfo struct {
	AvgTalkTime *int32 `json:"AvgTalkTime,omitempty" xml:"AvgTalkTime,omitempty"`
	// example:
	//
	// 5
	CallFailedNum *int32 `json:"CallFailedNum,omitempty" xml:"CallFailedNum,omitempty"`
	// example:
	//
	// 5
	CallNum *int32 `json:"CallNum,omitempty" xml:"CallNum,omitempty"`
	// example:
	//
	// xxx
	CreatorName          *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	DurationDistribution *string `json:"DurationDistribution,omitempty" xml:"DurationDistribution,omitempty"`
	// example:
	//
	// 1640087774563
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 5
	FinishedNum *int32 `json:"FinishedNum,omitempty" xml:"FinishedNum,omitempty"`
	// example:
	//
	// 5
	HangUpByClientNum *int32 `json:"HangUpByClientNum,omitempty" xml:"HangUpByClientNum,omitempty"`
	// example:
	//
	// {}
	JobsProgress     *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress `json:"JobsProgress,omitempty" xml:"JobsProgress,omitempty" type:"Struct"`
	NoInteractionNum *int32                                                           `json:"NoInteractionNum,omitempty" xml:"NoInteractionNum,omitempty"`
	// example:
	//
	// 1640087774563
	StartTime             *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TalkTurnsDistribution *string `json:"TalkTurnsDistribution,omitempty" xml:"TalkTurnsDistribution,omitempty"`
	// example:
	//
	// 5
	TransferByIntentNum *int32 `json:"TransferByIntentNum,omitempty" xml:"TransferByIntentNum,omitempty"`
	// example:
	//
	// 5
	TransferByNoAnswer *int32 `json:"TransferByNoAnswer,omitempty" xml:"TransferByNoAnswer,omitempty"`
}

func (s DescribeGroupExecutingInfoResponseBodyExecutingInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupExecutingInfoResponseBodyExecutingInfo) GoString() string {
	return s.String()
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetAvgTalkTime() *int32 {
	return s.AvgTalkTime
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetCallFailedNum() *int32 {
	return s.CallFailedNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetCallNum() *int32 {
	return s.CallNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetDurationDistribution() *string {
	return s.DurationDistribution
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetFinishedNum() *int32 {
	return s.FinishedNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetHangUpByClientNum() *int32 {
	return s.HangUpByClientNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetJobsProgress() *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	return s.JobsProgress
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetNoInteractionNum() *int32 {
	return s.NoInteractionNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetTalkTurnsDistribution() *string {
	return s.TalkTurnsDistribution
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetTransferByIntentNum() *int32 {
	return s.TransferByIntentNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) GetTransferByNoAnswer() *int32 {
	return s.TransferByNoAnswer
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetAvgTalkTime(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.AvgTalkTime = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetCallFailedNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.CallFailedNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetCallNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.CallNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetCreatorName(v string) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.CreatorName = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetDurationDistribution(v string) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.DurationDistribution = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetEndTime(v int64) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetFinishedNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.FinishedNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetHangUpByClientNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.HangUpByClientNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetJobsProgress(v *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.JobsProgress = v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetNoInteractionNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.NoInteractionNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetStartTime(v int64) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetTalkTurnsDistribution(v string) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.TalkTurnsDistribution = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetTransferByIntentNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.TransferByIntentNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) SetTransferByNoAnswer(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfo {
	s.TransferByNoAnswer = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfo) Validate() error {
	if s.JobsProgress != nil {
		if err := s.JobsProgress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress struct {
	// example:
	//
	// 5
	CancelledNum *int32 `json:"CancelledNum,omitempty" xml:"CancelledNum,omitempty"`
	// example:
	//
	// 5
	ExecutingNum *int32 `json:"ExecutingNum,omitempty" xml:"ExecutingNum,omitempty"`
	// example:
	//
	// 5
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// example:
	//
	// 5
	PausedNum *int32 `json:"PausedNum,omitempty" xml:"PausedNum,omitempty"`
	// example:
	//
	// 5
	SchedulingNum *int32 `json:"SchedulingNum,omitempty" xml:"SchedulingNum,omitempty"`
	// example:
	//
	// 5
	TotalCompletedNum *int32 `json:"TotalCompletedNum,omitempty" xml:"TotalCompletedNum,omitempty"`
	// example:
	//
	// 5
	TotalJobs *int32 `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	// example:
	//
	// 5
	TotalNotAnsweredNum *int32 `json:"TotalNotAnsweredNum,omitempty" xml:"TotalNotAnsweredNum,omitempty"`
}

func (s DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GoString() string {
	return s.String()
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetCancelledNum() *int32 {
	return s.CancelledNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetExecutingNum() *int32 {
	return s.ExecutingNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetPausedNum() *int32 {
	return s.PausedNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetSchedulingNum() *int32 {
	return s.SchedulingNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetTotalCompletedNum() *int32 {
	return s.TotalCompletedNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetTotalJobs() *int32 {
	return s.TotalJobs
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) GetTotalNotAnsweredNum() *int32 {
	return s.TotalNotAnsweredNum
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetCancelledNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.CancelledNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetExecutingNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.ExecutingNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetFailedNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.FailedNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetPausedNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.PausedNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetSchedulingNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.SchedulingNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetTotalCompletedNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.TotalCompletedNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetTotalJobs(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.TotalJobs = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) SetTotalNotAnsweredNum(v int32) *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress {
	s.TotalNotAnsweredNum = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponseBodyExecutingInfoJobsProgress) Validate() error {
	return dara.Validate(s)
}
