// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageAddInferenceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PersonalizedTextToImageAddInferenceJobResponseBodyData) *PersonalizedTextToImageAddInferenceJobResponseBody
	GetData() *PersonalizedTextToImageAddInferenceJobResponseBodyData
	SetErrCode(v string) *PersonalizedTextToImageAddInferenceJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PersonalizedTextToImageAddInferenceJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *PersonalizedTextToImageAddInferenceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PersonalizedTextToImageAddInferenceJobResponseBody
	GetSuccess() *bool
}

type PersonalizedTextToImageAddInferenceJobResponseBody struct {
	// example:
	//
	// []
	Data *PersonalizedTextToImageAddInferenceJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s PersonalizedTextToImageAddInferenceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) GetData() *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	return s.Data
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetData(v *PersonalizedTextToImageAddInferenceJobResponseBodyData) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.Data = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetErrCode(v string) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetErrMessage(v string) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetHttpStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetRequestId(v string) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetSuccess(v bool) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.Success = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PersonalizedTextToImageAddInferenceJobResponseBodyData struct {
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

func (s PersonalizedTextToImageAddInferenceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetPromptId() *string {
	return s.PromptId
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetCreateTime(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetId(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetJobStatus(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetJobTrainProgress(v float64) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetModelId(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetPromptId(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetResultImageUrl(v []*string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.ResultImageUrl = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
