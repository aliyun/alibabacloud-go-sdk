// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
	GetData() *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData
	SetErrCode(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
	GetSuccess() *bool
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody struct {
	// example:
	//
	// []
	Data *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GetData() *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	return s.Data
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetData(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetErrCode(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetErrMessage(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetHttpStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetRequestId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetSuccess(v bool) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.Success = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// promptId
	//
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetId() *string {
	return s.Id
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetPromptId() *string {
	return s.PromptId
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetCreateTime(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetJobStatus(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetJobTrainProgress(v float64) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetModelId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetPromptId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetResultImageUrl(v []*string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.ResultImageUrl = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
