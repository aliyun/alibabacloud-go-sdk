// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentMJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentMJobInfoResponseBody
	GetCode() *string
	SetData(v *GetAgentMJobInfoResponseBodyData) *GetAgentMJobInfoResponseBody
	GetData() *GetAgentMJobInfoResponseBodyData
	SetMessage(v string) *GetAgentMJobInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentMJobInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentMJobInfoResponseBody
	GetSuccess() *bool
}

type GetAgentMJobInfoResponseBody struct {
	// The response code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAgentMJobInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false/null: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentMJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentMJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentMJobInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentMJobInfoResponseBody) GetData() *GetAgentMJobInfoResponseBodyData {
	return s.Data
}

func (s *GetAgentMJobInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentMJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentMJobInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentMJobInfoResponseBody) SetCode(v string) *GetAgentMJobInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentMJobInfoResponseBody) SetData(v *GetAgentMJobInfoResponseBodyData) *GetAgentMJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentMJobInfoResponseBody) SetMessage(v string) *GetAgentMJobInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentMJobInfoResponseBody) SetRequestId(v string) *GetAgentMJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentMJobInfoResponseBody) SetSuccess(v bool) *GetAgentMJobInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentMJobInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentMJobInfoResponseBodyData struct {
	// The details of the task processing result.
	AgentMDetailResponse *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse `json:"AgentMDetailResponse,omitempty" xml:"AgentMDetailResponse,omitempty" type:"Struct"`
	// The end time of the scan range.
	//
	// example:
	//
	// 2026-08-26 20:00:00
	DataEndTime *string `json:"DataEndTime,omitempty" xml:"DataEndTime,omitempty"`
	// The start time of the scan range.
	//
	// example:
	//
	// 2026-08-26 19:00:00
	DataStartTime *string `json:"DataStartTime,omitempty" xml:"DataStartTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task status. Valid values:
	//
	// - queing: queuing.
	//
	// - readyAnalysis: pending analysis.
	//
	// - running: running.
	//
	// - error: failed.
	//
	// - finish: completed.
	//
	// - fileUploadUser: user-specified file upload completed.
	//
	// - fileUploadSystem: system-generated file upload completed.
	//
	// - expired: expired.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The actual end time of the task.
	//
	// example:
	//
	// 2026-08-26 20:00:00
	TaskEndTime *string `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	// The scheduled task ID.
	//
	// example:
	//
	// A6BEC8D-9A5B-4BE5-8432-4F635E***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The actual start time of the task.
	//
	// example:
	//
	// 2026-08-26 19:00:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s GetAgentMJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentMJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentMJobInfoResponseBodyData) GetAgentMDetailResponse() *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse {
	return s.AgentMDetailResponse
}

func (s *GetAgentMJobInfoResponseBodyData) GetDataEndTime() *string {
	return s.DataEndTime
}

func (s *GetAgentMJobInfoResponseBodyData) GetDataStartTime() *string {
	return s.DataStartTime
}

func (s *GetAgentMJobInfoResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetAgentMJobInfoResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetAgentMJobInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAgentMJobInfoResponseBodyData) GetTaskEndTime() *string {
	return s.TaskEndTime
}

func (s *GetAgentMJobInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgentMJobInfoResponseBodyData) GetTaskStartTime() *string {
	return s.TaskStartTime
}

func (s *GetAgentMJobInfoResponseBodyData) SetAgentMDetailResponse(v *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) *GetAgentMJobInfoResponseBodyData {
	s.AgentMDetailResponse = v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetDataEndTime(v string) *GetAgentMJobInfoResponseBodyData {
	s.DataEndTime = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetDataStartTime(v string) *GetAgentMJobInfoResponseBodyData {
	s.DataStartTime = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetId(v int64) *GetAgentMJobInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetMessage(v string) *GetAgentMJobInfoResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetStatus(v string) *GetAgentMJobInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetTaskEndTime(v string) *GetAgentMJobInfoResponseBodyData {
	s.TaskEndTime = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetTaskId(v string) *GetAgentMJobInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) SetTaskStartTime(v string) *GetAgentMJobInfoResponseBodyData {
	s.TaskStartTime = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyData) Validate() error {
	if s.AgentMDetailResponse != nil {
		if err := s.AgentMDetailResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentMJobInfoResponseBodyDataAgentMDetailResponse struct {
	// The execution summary.
	//
	// example:
	//
	// This quality inspection analyzed a total of 120 conversations..
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The list of result files. Each item contains complete file fields.
	SummaryUrls []*GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls `json:"SummaryUrls,omitempty" xml:"SummaryUrls,omitempty" type:"Repeated"`
}

func (s GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) GetSummary() *string {
	return s.Summary
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) GetSummaryUrls() []*GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls {
	return s.SummaryUrls
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) SetSummary(v string) *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse {
	s.Summary = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) SetSummaryUrls(v []*GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse {
	s.SummaryUrls = v
	return s
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponse) Validate() error {
	if s.SummaryUrls != nil {
		for _, item := range s.SummaryUrls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls struct {
	// The file name.
	//
	// example:
	//
	// SatisfactionAnalysis.xlsx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type.
	//
	// example:
	//
	// xlsx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file URL.
	//
	// example:
	//
	// http://******.oss-cn-hangzhou.aliyuncs.com/uploadTransfer/17****dline-express.zip
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) String() string {
	return dara.Prettify(s)
}

func (s GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) GoString() string {
	return s.String()
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) GetFileName() *string {
	return s.FileName
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) GetFileType() *string {
	return s.FileType
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) SetFileName(v string) *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls {
	s.FileName = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) SetFileType(v string) *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls {
	s.FileType = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) SetOssUrl(v string) *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls {
	s.OssUrl = &v
	return s
}

func (s *GetAgentMJobInfoResponseBodyDataAgentMDetailResponseSummaryUrls) Validate() error {
	return dara.Validate(s)
}
