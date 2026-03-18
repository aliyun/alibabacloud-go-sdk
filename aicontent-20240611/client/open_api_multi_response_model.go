// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenApiMultiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OpenApiMultiResponseData) *OpenApiMultiResponse
	GetData() []*OpenApiMultiResponseData
	SetErrCode(v string) *OpenApiMultiResponse
	GetErrCode() *string
	SetErrMessage(v string) *OpenApiMultiResponse
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *OpenApiMultiResponse
	GetHttpStatusCode() *int32
	SetRequestId(v string) *OpenApiMultiResponse
	GetRequestId() *string
	SetSuccess(v bool) *OpenApiMultiResponse
	GetSuccess() *bool
}

type OpenApiMultiResponse struct {
	// example:
	//
	// []
	Data       []*OpenApiMultiResponseData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode    *string                     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string                     `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenApiMultiResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenApiMultiResponse) GoString() string {
	return s.String()
}

func (s *OpenApiMultiResponse) GetData() []*OpenApiMultiResponseData {
	return s.Data
}

func (s *OpenApiMultiResponse) GetErrCode() *string {
	return s.ErrCode
}

func (s *OpenApiMultiResponse) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *OpenApiMultiResponse) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OpenApiMultiResponse) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenApiMultiResponse) GetSuccess() *bool {
	return s.Success
}

func (s *OpenApiMultiResponse) SetData(v []*OpenApiMultiResponseData) *OpenApiMultiResponse {
	s.Data = v
	return s
}

func (s *OpenApiMultiResponse) SetErrCode(v string) *OpenApiMultiResponse {
	s.ErrCode = &v
	return s
}

func (s *OpenApiMultiResponse) SetErrMessage(v string) *OpenApiMultiResponse {
	s.ErrMessage = &v
	return s
}

func (s *OpenApiMultiResponse) SetHttpStatusCode(v int32) *OpenApiMultiResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenApiMultiResponse) SetRequestId(v string) *OpenApiMultiResponse {
	s.RequestId = &v
	return s
}

func (s *OpenApiMultiResponse) SetSuccess(v bool) *OpenApiMultiResponse {
	s.Success = &v
	return s
}

func (s *OpenApiMultiResponse) Validate() error {
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

type OpenApiMultiResponseData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 456
	Id                  *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrl            []*string                                   `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	InferenceImageCount *int32                                      `json:"inferenceImageCount,omitempty" xml:"inferenceImageCount,omitempty"`
	InferenceJobList    []*OpenApiMultiResponseDataInferenceJobList `json:"inferenceJobList,omitempty" xml:"inferenceJobList,omitempty" type:"Repeated"`
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

func (s OpenApiMultiResponseData) String() string {
	return dara.Prettify(s)
}

func (s OpenApiMultiResponseData) GoString() string {
	return s.String()
}

func (s *OpenApiMultiResponseData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *OpenApiMultiResponseData) GetId() *string {
	return s.Id
}

func (s *OpenApiMultiResponseData) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *OpenApiMultiResponseData) GetInferenceImageCount() *int32 {
	return s.InferenceImageCount
}

func (s *OpenApiMultiResponseData) GetInferenceJobList() []*OpenApiMultiResponseDataInferenceJobList {
	return s.InferenceJobList
}

func (s *OpenApiMultiResponseData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *OpenApiMultiResponseData) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *OpenApiMultiResponseData) GetModelId() *string {
	return s.ModelId
}

func (s *OpenApiMultiResponseData) GetName() *string {
	return s.Name
}

func (s *OpenApiMultiResponseData) GetObjectType() *string {
	return s.ObjectType
}

func (s *OpenApiMultiResponseData) SetCreateTime(v string) *OpenApiMultiResponseData {
	s.CreateTime = &v
	return s
}

func (s *OpenApiMultiResponseData) SetId(v string) *OpenApiMultiResponseData {
	s.Id = &v
	return s
}

func (s *OpenApiMultiResponseData) SetImageUrl(v []*string) *OpenApiMultiResponseData {
	s.ImageUrl = v
	return s
}

func (s *OpenApiMultiResponseData) SetInferenceImageCount(v int32) *OpenApiMultiResponseData {
	s.InferenceImageCount = &v
	return s
}

func (s *OpenApiMultiResponseData) SetInferenceJobList(v []*OpenApiMultiResponseDataInferenceJobList) *OpenApiMultiResponseData {
	s.InferenceJobList = v
	return s
}

func (s *OpenApiMultiResponseData) SetJobStatus(v string) *OpenApiMultiResponseData {
	s.JobStatus = &v
	return s
}

func (s *OpenApiMultiResponseData) SetJobTrainProgress(v float64) *OpenApiMultiResponseData {
	s.JobTrainProgress = &v
	return s
}

func (s *OpenApiMultiResponseData) SetModelId(v string) *OpenApiMultiResponseData {
	s.ModelId = &v
	return s
}

func (s *OpenApiMultiResponseData) SetName(v string) *OpenApiMultiResponseData {
	s.Name = &v
	return s
}

func (s *OpenApiMultiResponseData) SetObjectType(v string) *OpenApiMultiResponseData {
	s.ObjectType = &v
	return s
}

func (s *OpenApiMultiResponseData) Validate() error {
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

type OpenApiMultiResponseDataInferenceJobList struct {
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

func (s OpenApiMultiResponseDataInferenceJobList) String() string {
	return dara.Prettify(s)
}

func (s OpenApiMultiResponseDataInferenceJobList) GoString() string {
	return s.String()
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetId() *string {
	return s.Id
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetJobStatus() *string {
	return s.JobStatus
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetJobTrainProgress() *float64 {
	return s.JobTrainProgress
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetModelId() *string {
	return s.ModelId
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetPromptId() *string {
	return s.PromptId
}

func (s *OpenApiMultiResponseDataInferenceJobList) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetCreateTime(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.CreateTime = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetId(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.Id = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetJobStatus(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.JobStatus = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetJobTrainProgress(v float64) *OpenApiMultiResponseDataInferenceJobList {
	s.JobTrainProgress = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetModelId(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.ModelId = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetPromptId(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.PromptId = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetResultImageUrl(v []*string) *OpenApiMultiResponseDataInferenceJobList {
	s.ResultImageUrl = v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) Validate() error {
	return dara.Validate(s)
}
