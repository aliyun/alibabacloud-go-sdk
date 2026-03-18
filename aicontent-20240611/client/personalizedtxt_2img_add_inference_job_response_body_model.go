// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddInferenceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Personalizedtxt2imgAddInferenceJobResponseBodyData) *Personalizedtxt2imgAddInferenceJobResponseBody
	GetData() *Personalizedtxt2imgAddInferenceJobResponseBodyData
	SetErrCode(v string) *Personalizedtxt2imgAddInferenceJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *Personalizedtxt2imgAddInferenceJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *Personalizedtxt2imgAddInferenceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *Personalizedtxt2imgAddInferenceJobResponseBody
	GetSuccess() *bool
}

type Personalizedtxt2imgAddInferenceJobResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgAddInferenceJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s Personalizedtxt2imgAddInferenceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) GetData() *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	return s.Data
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetData(v *Personalizedtxt2imgAddInferenceJobResponseBodyData) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetErrCode(v string) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetErrMessage(v string) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetRequestId(v string) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetSuccess(v bool) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.Success = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Personalizedtxt2imgAddInferenceJobResponseBodyData struct {
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

func (s Personalizedtxt2imgAddInferenceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetPromptId() *string {
	return s.PromptId
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetId(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetPromptId(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetResultImageUrl(v []*string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.ResultImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
