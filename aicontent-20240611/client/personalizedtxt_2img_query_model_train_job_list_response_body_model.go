// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryModelTrainJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) *Personalizedtxt2imgQueryModelTrainJobListResponseBody
	GetData() []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData
	SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainJobListResponseBody
	GetSuccess() *bool
}

type Personalizedtxt2imgQueryModelTrainJobListResponseBody struct {
	// example:
	//
	// []
	Data []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) GetData() []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	return s.Data
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetData(v []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.Success = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type Personalizedtxt2imgQueryModelTrainJobListResponseBodyData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 456
	Id                  *string                                                                      `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrl            []*string                                                                    `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	InferenceImageCount *int32                                                                       `json:"inferenceImageCount,omitempty" xml:"inferenceImageCount,omitempty"`
	InferenceJobList    []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList `json:"inferenceJobList,omitempty" xml:"inferenceJobList,omitempty" type:"Repeated"`
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

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetInferenceImageCount() *int32 {
	return s.InferenceImageCount
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetInferenceJobList() []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	return s.InferenceJobList
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetName() *string {
	return s.Name
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GetObjectType() *string {
	return s.ObjectType
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetImageUrl(v []*string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetInferenceImageCount(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.InferenceImageCount = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetInferenceJobList(v []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetName(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetObjectType(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) Validate() error {
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

type Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList struct {
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

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetPromptId() *string {
	return s.PromptId
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetCreateTime(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetJobStatus(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetJobTrainProgress(v float64) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetPromptId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetResultImageUrl(v []*string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.ResultImageUrl = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) Validate() error {
	return dara.Validate(s)
}
