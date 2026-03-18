// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddModelTrainJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Personalizedtxt2imgAddModelTrainJobResponseBodyData) *Personalizedtxt2imgAddModelTrainJobResponseBody
	GetData() *Personalizedtxt2imgAddModelTrainJobResponseBodyData
	SetErrCode(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *Personalizedtxt2imgAddModelTrainJobResponseBody
	GetSuccess() *bool
}

type Personalizedtxt2imgAddModelTrainJobResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgAddModelTrainJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s Personalizedtxt2imgAddModelTrainJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) GetData() *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	return s.Data
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetData(v *Personalizedtxt2imgAddModelTrainJobResponseBodyData) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetErrCode(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetErrMessage(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetRequestId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetSuccess(v bool) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.Success = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Personalizedtxt2imgAddModelTrainJobResponseBodyData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 456
	Id                  *string                                                                `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrl            []*string                                                              `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	InferenceImageCount *int32                                                                 `json:"inferenceImageCount,omitempty" xml:"inferenceImageCount,omitempty"`
	InferenceJobList    []*Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList `json:"inferenceJobList,omitempty" xml:"inferenceJobList,omitempty" type:"Repeated"`
	// example:
	//
	// TRAINING
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
	// 可爱熊猫模型训练任务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// panda
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetInferenceImageCount() *int32 {
	return s.InferenceImageCount
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetInferenceJobList() []*Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	return s.InferenceJobList
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetName() *string {
	return s.Name
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) GetObjectType() *string {
	return s.ObjectType
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetInferenceImageCount(v int32) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.InferenceImageCount = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetInferenceJobList(v []*Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetModelId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetName(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) Validate() error {
	if s.InferenceJobList != nil {
		for _, item := range s.InferenceJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList struct {
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

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetPromptId() *string {
	return s.PromptId
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetCreateTime(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetJobStatus(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetJobTrainProgress(v float64) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetModelId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetPromptId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetResultImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.ResultImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) Validate() error {
	return dara.Validate(s)
}
