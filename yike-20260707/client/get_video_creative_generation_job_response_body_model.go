// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreativeGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetVideoCreativeGenerationJobResponseBody
	GetCreateTime() *string
	SetErrorCode(v string) *GetVideoCreativeGenerationJobResponseBody
	GetErrorCode() *string
	SetFinishTime(v string) *GetVideoCreativeGenerationJobResponseBody
	GetFinishTime() *string
	SetJobId(v string) *GetVideoCreativeGenerationJobResponseBody
	GetJobId() *string
	SetJobResult(v []*GetVideoCreativeGenerationJobResponseBodyJobResult) *GetVideoCreativeGenerationJobResponseBody
	GetJobResult() []*GetVideoCreativeGenerationJobResponseBodyJobResult
	SetJobStatus(v string) *GetVideoCreativeGenerationJobResponseBody
	GetJobStatus() *string
	SetRequestId(v string) *GetVideoCreativeGenerationJobResponseBody
	GetRequestId() *string
	SetUserData(v string) *GetVideoCreativeGenerationJobResponseBody
	GetUserData() *string
}

type GetVideoCreativeGenerationJobResponseBody struct {
	// The time when the job was created. The time is displayed in UTC in the yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\" format.
	//
	// example:
	//
	// 2026-02-06T18:53:34.001+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. This field is returned only when the job status is Failed.
	//
	// example:
	//
	// WorkflowTaskFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The time when the job was completed. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-02-06T18:53:34.001+08:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task result. This field is valid only when the node status is Succeeded. Each element represents a single-language video output.
	JobResult []*GetVideoCreativeGenerationJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Repeated"`
	// The job status. Valid values: Running, Succeeded, and Failed.
	//
	// example:
	//
	// Succeeded
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user-defined parameter. This value is returned as is.
	//
	// example:
	//
	// 12374214127312673137132
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetVideoCreativeGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreativeGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetJobResult() []*GetVideoCreativeGenerationJobResponseBodyJobResult {
	return s.JobResult
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoCreativeGenerationJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetCreateTime(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetErrorCode(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetFinishTime(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetJobId(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetJobResult(v []*GetVideoCreativeGenerationJobResponseBodyJobResult) *GetVideoCreativeGenerationJobResponseBody {
	s.JobResult = v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetJobStatus(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetRequestId(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) SetUserData(v string) *GetVideoCreativeGenerationJobResponseBody {
	s.UserData = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBody) Validate() error {
	if s.JobResult != nil {
		for _, item := range s.JobResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoCreativeGenerationJobResponseBodyJobResult struct {
	// The ID of the online editing project.
	//
	// example:
	//
	// 01a6adbbd181437eb5030d3d93e0182d
	EditingProjectId *string `json:"EditingProjectId,omitempty" xml:"EditingProjectId,omitempty"`
	// The media asset ID of the produced video.
	//
	// example:
	//
	// 9d7e982012c671f1b803e7f6d75a6302
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The output language. Valid values: CN, EN, and YUE.
	//
	// example:
	//
	// CN
	OutputLanguage *string `json:"OutputLanguage,omitempty" xml:"OutputLanguage,omitempty"`
	// The download URL of the produced video. The URL is signed and time-sensitive.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetVideoCreativeGenerationJobResponseBodyJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreativeGenerationJobResponseBodyJobResult) GoString() string {
	return s.String()
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) GetEditingProjectId() *string {
	return s.EditingProjectId
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) GetOutputLanguage() *string {
	return s.OutputLanguage
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) SetEditingProjectId(v string) *GetVideoCreativeGenerationJobResponseBodyJobResult {
	s.EditingProjectId = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) SetMediaId(v string) *GetVideoCreativeGenerationJobResponseBodyJobResult {
	s.MediaId = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) SetOutputLanguage(v string) *GetVideoCreativeGenerationJobResponseBodyJobResult {
	s.OutputLanguage = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) SetOutputUrl(v string) *GetVideoCreativeGenerationJobResponseBodyJobResult {
	s.OutputUrl = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}
