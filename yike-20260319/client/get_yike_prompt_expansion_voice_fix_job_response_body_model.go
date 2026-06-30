// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikePromptExpansionVoiceFixJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetEndTime() *string
	SetErrorCode(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetErrorMessage() *string
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
	SetStartTime(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetStartTime() *string
	SetUserData(v string) *GetYikePromptExpansionVoiceFixJobResponseBody
	GetUserData() *string
}

type GetYikePromptExpansionVoiceFixJobResponseBody struct {
	// The end time.
	//
	// example:
	//
	// 2026-01-30T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code. This parameter is returned when the task is in the Failed state.
	//
	// example:
	//
	// Forbidden
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task parameters, in JSON format. The following fields are included:
	//
	// - model (String, required): The model name. Example: happyhorse-1.0-r2v.
	//
	// - input (Object, required): The input data object.
	//
	//   - prompt (String, required): The prompt content. The maximum length is 10,000 characters.
	//
	//   - media (Array(Object), required): The list of media materials used to specify reference images and audio.
	//
	//     - type (String, required): The input media type. Valid values: reference_image and reference_audio.
	//
	//     - url (String, required): The URL of the input media.
	//
	// - parameters (Object, required): The video generation parameter object.
	//
	//   - duration (Integer, required): The video duration in seconds. The value must be a positive integer. Valid values: 5 to 15.
	//
	//   - ratio (String, required): The aspect ratio. Valid values: 16:9, 9:16, 4:3, 3:4, and 1:1.
	//
	//   - resolution (String, required): The video resolution. Valid values: 720P and 1080P.
	//
	//   - specialEdition (Bool, optional): The cost-effective edition parameter. This parameter can be set to true only when the resolution is 1080P.
	//
	//   - skipPromptEnhancer (Bool, optional): Specifies whether to skip prompt enhancement. Default value: false.
	//
	//   - skipVoiceResync (Bool, optional): Specifies whether to skip audio repair. Default value: false.
	//
	// example:
	//
	// {\\"model\\":\\"happyhorse-1.0-r2v\\",\\"input\\":{\\"prompt\\":\\"[Image 1]中身着红色旗袍的女性，镜头先以侧面中景勾勒旗袍修身剪裁与S型曲线，随即切换至低角度仰拍，捕捉她轻抬玉手展开[Image 2]中的折扇的同时，[Image 3]中的流苏耳坠随头部转动轻盈摆动的细节，最后推近至面部特写，定格在她指尖轻点扇骨、眼波流转间的含蓄风情，多视角全方位展现东方韵味。\\",\\"media\\":[{\\"type\\":\\"reference_image\\",\\"url\\":\\"https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20260424/mvzfud/hh-v2v-girl.jpg\\"},{\\"type\\":\\"reference_image\\",\\"url\\":\\"https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20260424/fvuihk/hh-v2v2-folding-fan.jpg\\"},{\\"type\\":\\"reference_image\\",\\"url\\":\\"https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20260424/imerii/hh-v2v-earrings.jpg\\"}]},\\"parameters\\":{\\"resolution\\":\\"720P\\",\\"ratio\\":\\"16:9\\",\\"duration\\":5,\\"specialEdition\\":false,\\"skipPromptEnhancer\\":false,\\"skipVoiceResync\\":false}}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The task result of the node.
	JobResult []*GetYikePromptExpansionVoiceFixJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Repeated"`
	// The task status. Valid values:
	//
	// - **Succeeded**: The task is processed.
	//
	// - **Failed**: The task failed.
	//
	// - **Running**: The task is being processed.
	//
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
	// The start time.
	//
	// example:
	//
	// 2025-06-24T00:55:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The custom user parameter.
	//
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

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
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

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetEndTime(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetErrorCode(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetErrorMessage(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.ErrorMessage = &v
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

func (s *GetYikePromptExpansionVoiceFixJobResponseBody) SetStartTime(v string) *GetYikePromptExpansionVoiceFixJobResponseBody {
	s.StartTime = &v
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
