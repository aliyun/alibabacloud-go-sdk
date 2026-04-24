// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeVoiceNarratorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetYikeVoiceNarratorJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *GetYikeVoiceNarratorJobResponseBody
	GetJobId() *string
	SetJobParams(v string) *GetYikeVoiceNarratorJobResponseBody
	GetJobParams() *string
	SetJobResult(v []*GetYikeVoiceNarratorJobResponseBodyJobResult) *GetYikeVoiceNarratorJobResponseBody
	GetJobResult() []*GetYikeVoiceNarratorJobResponseBodyJobResult
	SetJobStatus(v string) *GetYikeVoiceNarratorJobResponseBody
	GetJobStatus() *string
	SetRequestId(v string) *GetYikeVoiceNarratorJobResponseBody
	GetRequestId() *string
	SetUserData(v string) *GetYikeVoiceNarratorJobResponseBody
	GetUserData() *string
}

type GetYikeVoiceNarratorJobResponseBody struct {
	// example:
	//
	// WorkflowTaskFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// task_abc123def456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"TextType\\":2,\\"TextContent\\":\\"Today, Beijing held a press conference to announce plans to further optimize the city\\"s transportation network, including adding three new subway lines within the next three years....\\",\\"AspectRatio\\":\\"16:9\\", \\"Resolution\\":\\"720P\\", \\"OutputLanguages\\":[\\"CN\\",\\"YUE\\"]"}
	JobParams *string                                         `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	JobResult []*GetYikeVoiceNarratorJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Repeated"`
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// req_query_20260420_002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://cms.example.com/callback/video-task\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetYikeVoiceNarratorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeVoiceNarratorJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetJobParams() *string {
	return s.JobParams
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetJobResult() []*GetYikeVoiceNarratorJobResponseBodyJobResult {
	return s.JobResult
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeVoiceNarratorJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetErrorCode(v string) *GetYikeVoiceNarratorJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetJobId(v string) *GetYikeVoiceNarratorJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetJobParams(v string) *GetYikeVoiceNarratorJobResponseBody {
	s.JobParams = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetJobResult(v []*GetYikeVoiceNarratorJobResponseBodyJobResult) *GetYikeVoiceNarratorJobResponseBody {
	s.JobResult = v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetJobStatus(v string) *GetYikeVoiceNarratorJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetRequestId(v string) *GetYikeVoiceNarratorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) SetUserData(v string) *GetYikeVoiceNarratorJobResponseBody {
	s.UserData = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBody) Validate() error {
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

type GetYikeVoiceNarratorJobResponseBodyJobResult struct {
	// example:
	//
	// 01a6adbbd181437eb5030d3d93e0182d
	EditingProjectId *string `json:"EditingProjectId,omitempty" xml:"EditingProjectId,omitempty"`
	// example:
	//
	// 9d7e982012c671f1b803e7f6d75a6302
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// CN
	OutputLanguage *string `json:"OutputLanguage,omitempty" xml:"OutputLanguage,omitempty"`
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/videos/task_abc123def456.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetYikeVoiceNarratorJobResponseBodyJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeVoiceNarratorJobResponseBodyJobResult) GoString() string {
	return s.String()
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) GetEditingProjectId() *string {
	return s.EditingProjectId
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) GetOutputLanguage() *string {
	return s.OutputLanguage
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) SetEditingProjectId(v string) *GetYikeVoiceNarratorJobResponseBodyJobResult {
	s.EditingProjectId = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) SetMediaId(v string) *GetYikeVoiceNarratorJobResponseBodyJobResult {
	s.MediaId = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) SetOutputLanguage(v string) *GetYikeVoiceNarratorJobResponseBodyJobResult {
	s.OutputLanguage = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) SetOutputUrl(v string) *GetYikeVoiceNarratorJobResponseBodyJobResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikeVoiceNarratorJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}
