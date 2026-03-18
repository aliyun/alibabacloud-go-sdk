// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryInferenceJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
	GetData() *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData
	SetErrCode(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
	GetSuccess() *bool
}

type Personalizedtxt2imgQueryInferenceJobInfoResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GetData() *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	return s.Data
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetData(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.Success = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData struct {
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
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetPromptId() *string {
	return s.PromptId
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetModelId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetPromptId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetResultImageUrl(v []*string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.ResultImageUrl = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
