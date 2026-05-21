// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAgentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetYikeAgentJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *GetYikeAgentJobResponseBody
	GetJobId() *string
	SetJobParams(v string) *GetYikeAgentJobResponseBody
	GetJobParams() *string
	SetJobResult(v []*GetYikeAgentJobResponseBodyJobResult) *GetYikeAgentJobResponseBody
	GetJobResult() []*GetYikeAgentJobResponseBodyJobResult
	SetJobStatus(v string) *GetYikeAgentJobResponseBody
	GetJobStatus() *string
	SetJobType(v string) *GetYikeAgentJobResponseBody
	GetJobType() *string
	SetRequestId(v string) *GetYikeAgentJobResponseBody
	GetRequestId() *string
	SetUserData(v string) *GetYikeAgentJobResponseBody
	GetUserData() *string
}

type GetYikeAgentJobResponseBody struct {
	// example:
	//
	// WorkflowTaskFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"TextType\\":2,\\"TextContent\\":\\"Today, Beijing held a press conference to announce plans to further optimize the city\\"s transportation network, including adding three new subway lines within the next three years....\\",\\"AspectRatio\\":\\"16:9\\", \\"Resolution\\":\\"720P\\", \\"OutputLanguages\\":[\\"CN\\",\\"YUE\\"]"}
	JobParams *string                                 `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	JobResult []*GetYikeAgentJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Repeated"`
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// VoiceNarrator
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://cms.example.com/callback/video-task\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetYikeAgentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetYikeAgentJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeAgentJobResponseBody) GetJobParams() *string {
	return s.JobParams
}

func (s *GetYikeAgentJobResponseBody) GetJobResult() []*GetYikeAgentJobResponseBodyJobResult {
	return s.JobResult
}

func (s *GetYikeAgentJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetYikeAgentJobResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *GetYikeAgentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeAgentJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetYikeAgentJobResponseBody) SetErrorCode(v string) *GetYikeAgentJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetJobId(v string) *GetYikeAgentJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetJobParams(v string) *GetYikeAgentJobResponseBody {
	s.JobParams = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetJobResult(v []*GetYikeAgentJobResponseBodyJobResult) *GetYikeAgentJobResponseBody {
	s.JobResult = v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetJobStatus(v string) *GetYikeAgentJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetJobType(v string) *GetYikeAgentJobResponseBody {
	s.JobType = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetRequestId(v string) *GetYikeAgentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) SetUserData(v string) *GetYikeAgentJobResponseBody {
	s.UserData = &v
	return s
}

func (s *GetYikeAgentJobResponseBody) Validate() error {
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

type GetYikeAgentJobResponseBodyJobResult struct {
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

func (s GetYikeAgentJobResponseBodyJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobResponseBodyJobResult) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobResponseBodyJobResult) GetEditingProjectId() *string {
	return s.EditingProjectId
}

func (s *GetYikeAgentJobResponseBodyJobResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAgentJobResponseBodyJobResult) GetOutputLanguage() *string {
	return s.OutputLanguage
}

func (s *GetYikeAgentJobResponseBodyJobResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikeAgentJobResponseBodyJobResult) SetEditingProjectId(v string) *GetYikeAgentJobResponseBodyJobResult {
	s.EditingProjectId = &v
	return s
}

func (s *GetYikeAgentJobResponseBodyJobResult) SetMediaId(v string) *GetYikeAgentJobResponseBodyJobResult {
	s.MediaId = &v
	return s
}

func (s *GetYikeAgentJobResponseBodyJobResult) SetOutputLanguage(v string) *GetYikeAgentJobResponseBodyJobResult {
	s.OutputLanguage = &v
	return s
}

func (s *GetYikeAgentJobResponseBodyJobResult) SetOutputUrl(v string) *GetYikeAgentJobResponseBodyJobResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikeAgentJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}
