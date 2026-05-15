// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikePromptExpansionVoiceFixJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetJobId() *string
	SetJobParams(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetJobParams() *string
	SetJobResult(v []*GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetJobResult() []*GetYikePromptExpansionVoiceFixJobResponseBodyJobResult
	SetJobStatus(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetJobStatus() *string
	SetRequestId(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetRequestId() *string
	SetUserData(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetUserData() *string
}

type GetYikePromptExpansionVoiceFixJobResponseBody struct {
	// example:
	//
	// Forbidden
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	JobId     *string                                                   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobParams *string                                                   `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	JobResult []*GetYikePromptExpansionVoiceFixJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Repeated"`
	// example:
	//
	// Succeeded
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"testKey":"testValue"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetYikePromptExpansionVoiceFixJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikePromptExpansionVoiceFixJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetJobParams() *string {
	return s.JobParams
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetJobResult() []*GetYikePromptExpansionVoiceFixJobResponseBodyJobResult {
	return s.JobResult
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetErrorCode(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetJobId(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetJobParams(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.JobParams = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetJobResult(v []*GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.JobResult = v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetJobStatus(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetRequestId(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetUserData(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.UserData = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) Validate() error {
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

type GetYikePromptExpansionVoiceFixJobResponseBodyJobResult struct {
	// Oss Url
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) GoString() string {
	return s.String()
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) SetOutputUrl(v string) *GetYikePromptExpansionVoiceFixJobResponseBodyJobResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}
