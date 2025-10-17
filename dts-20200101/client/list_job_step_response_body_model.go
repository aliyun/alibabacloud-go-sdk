// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobStepResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListJobStepResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListJobStepResponseBody
	GetHttpStatusCode() *int32
	SetJobSteps(v []*ListJobStepResponseBodyJobSteps) *ListJobStepResponseBody
	GetJobSteps() []*ListJobStepResponseBodyJobSteps
	SetRequestId(v string) *ListJobStepResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobStepResponseBody
	GetSuccess() *bool
	SetUseV2API(v bool) *ListJobStepResponseBody
	GetUseV2API() *bool
}

type ListJobStepResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	JobSteps       []*ListJobStepResponseBodyJobSteps `json:"JobSteps,omitempty" xml:"JobSteps,omitempty" type:"Repeated"`
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// True
	UseV2API *bool `json:"UseV2API,omitempty" xml:"UseV2API,omitempty"`
}

func (s ListJobStepResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobStepResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJobStepResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListJobStepResponseBody) GetJobSteps() []*ListJobStepResponseBodyJobSteps {
	return s.JobSteps
}

func (s *ListJobStepResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobStepResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobStepResponseBody) GetUseV2API() *bool {
	return s.UseV2API
}

func (s *ListJobStepResponseBody) SetCode(v string) *ListJobStepResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobStepResponseBody) SetHttpStatusCode(v int32) *ListJobStepResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJobStepResponseBody) SetJobSteps(v []*ListJobStepResponseBodyJobSteps) *ListJobStepResponseBody {
	s.JobSteps = v
	return s
}

func (s *ListJobStepResponseBody) SetRequestId(v string) *ListJobStepResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobStepResponseBody) SetSuccess(v bool) *ListJobStepResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobStepResponseBody) SetUseV2API(v bool) *ListJobStepResponseBody {
	s.UseV2API = &v
	return s
}

func (s *ListJobStepResponseBody) Validate() error {
	if s.JobSteps != nil {
		for _, item := range s.JobSteps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobStepResponseBodyJobSteps struct {
	// example:
	//
	// 2024-04-11T09:33:23Z
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// example:
	//
	// 01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2023-11-28T17:13:51Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// fullcheck find different records : 2372
	ErrMsg       *string                                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ErrorDetails []*ListJobStepResponseBodyJobStepsErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-03-15T02:15:14Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1
	IncLatencyMilliseconds *int64 `json:"IncLatencyMilliseconds,omitempty" xml:"IncLatencyMilliseconds,omitempty"`
	// example:
	//
	// -1
	IncLatencySeconds *int64 `json:"IncLatencySeconds,omitempty" xml:"IncLatencySeconds,omitempty"`
	// example:
	//
	// l02c1f7h179****
	JobStepId   *string `json:"JobStepId,omitempty" xml:"JobStepId,omitempty"`
	JobStepName *string `json:"JobStepName,omitempty" xml:"JobStepName,omitempty"`
	// example:
	//
	// 2025-01-03T02:26:14Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// true
	NeedAcceleration *bool `json:"NeedAcceleration,omitempty" xml:"NeedAcceleration,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	Serial *int32 `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// example:
	//
	// ○ Finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	SubJobCount *int32                                        `json:"SubJobCount,omitempty" xml:"SubJobCount,omitempty"`
	SubJobSteps []*ListJobStepResponseBodyJobStepsSubJobSteps `json:"SubJobSteps,omitempty" xml:"SubJobSteps,omitempty" type:"Repeated"`
	// example:
	//
	// full
	RedisPhaseType *string `json:"redisPhaseType,omitempty" xml:"redisPhaseType,omitempty"`
}

func (s ListJobStepResponseBodyJobSteps) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepResponseBodyJobSteps) GoString() string {
	return s.String()
}

func (s *ListJobStepResponseBodyJobSteps) GetBootTime() *string {
	return s.BootTime
}

func (s *ListJobStepResponseBodyJobSteps) GetCode() *string {
	return s.Code
}

func (s *ListJobStepResponseBodyJobSteps) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobStepResponseBodyJobSteps) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ListJobStepResponseBodyJobSteps) GetErrorDetails() []*ListJobStepResponseBodyJobStepsErrorDetails {
	return s.ErrorDetails
}

func (s *ListJobStepResponseBodyJobSteps) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListJobStepResponseBodyJobSteps) GetIncLatencyMilliseconds() *int64 {
	return s.IncLatencyMilliseconds
}

func (s *ListJobStepResponseBodyJobSteps) GetIncLatencySeconds() *int64 {
	return s.IncLatencySeconds
}

func (s *ListJobStepResponseBodyJobSteps) GetJobStepId() *string {
	return s.JobStepId
}

func (s *ListJobStepResponseBodyJobSteps) GetJobStepName() *string {
	return s.JobStepName
}

func (s *ListJobStepResponseBodyJobSteps) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListJobStepResponseBodyJobSteps) GetNeedAcceleration() *bool {
	return s.NeedAcceleration
}

func (s *ListJobStepResponseBodyJobSteps) GetProgress() *int32 {
	return s.Progress
}

func (s *ListJobStepResponseBodyJobSteps) GetSerial() *int32 {
	return s.Serial
}

func (s *ListJobStepResponseBodyJobSteps) GetState() *string {
	return s.State
}

func (s *ListJobStepResponseBodyJobSteps) GetSubJobCount() *int32 {
	return s.SubJobCount
}

func (s *ListJobStepResponseBodyJobSteps) GetSubJobSteps() []*ListJobStepResponseBodyJobStepsSubJobSteps {
	return s.SubJobSteps
}

func (s *ListJobStepResponseBodyJobSteps) GetRedisPhaseType() *string {
	return s.RedisPhaseType
}

func (s *ListJobStepResponseBodyJobSteps) SetBootTime(v string) *ListJobStepResponseBodyJobSteps {
	s.BootTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetCode(v string) *ListJobStepResponseBodyJobSteps {
	s.Code = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetCreateTime(v string) *ListJobStepResponseBodyJobSteps {
	s.CreateTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetErrMsg(v string) *ListJobStepResponseBodyJobSteps {
	s.ErrMsg = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetErrorDetails(v []*ListJobStepResponseBodyJobStepsErrorDetails) *ListJobStepResponseBodyJobSteps {
	s.ErrorDetails = v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetFinishTime(v string) *ListJobStepResponseBodyJobSteps {
	s.FinishTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetIncLatencyMilliseconds(v int64) *ListJobStepResponseBodyJobSteps {
	s.IncLatencyMilliseconds = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetIncLatencySeconds(v int64) *ListJobStepResponseBodyJobSteps {
	s.IncLatencySeconds = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetJobStepId(v string) *ListJobStepResponseBodyJobSteps {
	s.JobStepId = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetJobStepName(v string) *ListJobStepResponseBodyJobSteps {
	s.JobStepName = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetModifyTime(v string) *ListJobStepResponseBodyJobSteps {
	s.ModifyTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetNeedAcceleration(v bool) *ListJobStepResponseBodyJobSteps {
	s.NeedAcceleration = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetProgress(v int32) *ListJobStepResponseBodyJobSteps {
	s.Progress = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetSerial(v int32) *ListJobStepResponseBodyJobSteps {
	s.Serial = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetState(v string) *ListJobStepResponseBodyJobSteps {
	s.State = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetSubJobCount(v int32) *ListJobStepResponseBodyJobSteps {
	s.SubJobCount = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetSubJobSteps(v []*ListJobStepResponseBodyJobStepsSubJobSteps) *ListJobStepResponseBodyJobSteps {
	s.SubJobSteps = v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) SetRedisPhaseType(v string) *ListJobStepResponseBodyJobSteps {
	s.RedisPhaseType = &v
	return s
}

func (s *ListJobStepResponseBodyJobSteps) Validate() error {
	if s.ErrorDetails != nil {
		for _, item := range s.ErrorDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubJobSteps != nil {
		for _, item := range s.SubJobSteps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobStepResponseBodyJobStepsErrorDetails struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ****
	HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s ListJobStepResponseBodyJobStepsErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepResponseBodyJobStepsErrorDetails) GoString() string {
	return s.String()
}

func (s *ListJobStepResponseBodyJobStepsErrorDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobStepResponseBodyJobStepsErrorDetails) GetHelpUrl() *string {
	return s.HelpUrl
}

func (s *ListJobStepResponseBodyJobStepsErrorDetails) SetErrorCode(v string) *ListJobStepResponseBodyJobStepsErrorDetails {
	s.ErrorCode = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsErrorDetails) SetHelpUrl(v string) *ListJobStepResponseBodyJobStepsErrorDetails {
	s.HelpUrl = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsErrorDetails) Validate() error {
	return dara.Validate(s)
}

type ListJobStepResponseBodyJobStepsSubJobSteps struct {
	// example:
	//
	// 2025-01-02T02:00:21Z
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// example:
	//
	// 03
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2024-09-20T02:13:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// UncaughtException:java.lang.NullPointerException
	ErrMsg       *string                                                   `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ErrorDetails []*ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-03-15T02:15:14Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1
	IncLatencyMilliseconds *string `json:"IncLatencyMilliseconds,omitempty" xml:"IncLatencyMilliseconds,omitempty"`
	// example:
	//
	// 1
	IncLatencySeconds *int64 `json:"IncLatencySeconds,omitempty" xml:"IncLatencySeconds,omitempty"`
	// example:
	//
	// mj3z9w9s10am68o_0004_0000
	JobStepId *string `json:"JobStepId,omitempty" xml:"JobStepId,omitempty"`
	// example:
	//
	// test
	JobStepName *string `json:"JobStepName,omitempty" xml:"JobStepName,omitempty"`
	// example:
	//
	// 2024-08-22T02:04:35Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// true
	NeedAcceleration *bool `json:"NeedAcceleration,omitempty" xml:"NeedAcceleration,omitempty"`
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 123
	Serial *int32 `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobStepResponseBodyJobStepsSubJobSteps) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepResponseBodyJobStepsSubJobSteps) GoString() string {
	return s.String()
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetBootTime() *string {
	return s.BootTime
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetCode() *string {
	return s.Code
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetErrorDetails() []*ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails {
	return s.ErrorDetails
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetIncLatencyMilliseconds() *string {
	return s.IncLatencyMilliseconds
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetIncLatencySeconds() *int64 {
	return s.IncLatencySeconds
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetJobStepId() *string {
	return s.JobStepId
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetJobStepName() *string {
	return s.JobStepName
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetNeedAcceleration() *bool {
	return s.NeedAcceleration
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetProgress() *int32 {
	return s.Progress
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetSerial() *int32 {
	return s.Serial
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) GetState() *string {
	return s.State
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetBootTime(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.BootTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetCode(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.Code = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetCreateTime(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.CreateTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetErrMsg(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.ErrMsg = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetErrorDetails(v []*ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.ErrorDetails = v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetFinishTime(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.FinishTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetIncLatencyMilliseconds(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.IncLatencyMilliseconds = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetIncLatencySeconds(v int64) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.IncLatencySeconds = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetJobStepId(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.JobStepId = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetJobStepName(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.JobStepName = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetModifyTime(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.ModifyTime = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetNeedAcceleration(v bool) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.NeedAcceleration = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetProgress(v int32) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.Progress = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetSerial(v int32) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.Serial = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) SetState(v string) *ListJobStepResponseBodyJobStepsSubJobSteps {
	s.State = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobSteps) Validate() error {
	if s.ErrorDetails != nil {
		for _, item := range s.ErrorDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails struct {
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ****
	HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) GoString() string {
	return s.String()
}

func (s *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) GetHelpUrl() *string {
	return s.HelpUrl
}

func (s *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) SetErrorCode(v string) *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails {
	s.ErrorCode = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) SetHelpUrl(v string) *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails {
	s.HelpUrl = &v
	return s
}

func (s *ListJobStepResponseBodyJobStepsSubJobStepsErrorDetails) Validate() error {
	return dara.Validate(s)
}
