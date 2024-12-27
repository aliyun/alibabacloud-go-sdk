// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	return tea.Prettify(s)
}

func (s OpenApiMultiResponse) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s OpenApiMultiResponseData) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s OpenApiMultiResponseDataInferenceJobList) GoString() string {
	return s.String()
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

type OpenApiSingleResponse struct {
	// example:
	//
	// []
	Data       *OpenApiSingleResponseData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode    *string                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string                    `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenApiSingleResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenApiSingleResponse) GoString() string {
	return s.String()
}

func (s *OpenApiSingleResponse) SetData(v *OpenApiSingleResponseData) *OpenApiSingleResponse {
	s.Data = v
	return s
}

func (s *OpenApiSingleResponse) SetErrCode(v string) *OpenApiSingleResponse {
	s.ErrCode = &v
	return s
}

func (s *OpenApiSingleResponse) SetErrMessage(v string) *OpenApiSingleResponse {
	s.ErrMessage = &v
	return s
}

func (s *OpenApiSingleResponse) SetHttpStatusCode(v int32) *OpenApiSingleResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenApiSingleResponse) SetRequestId(v string) *OpenApiSingleResponse {
	s.RequestId = &v
	return s
}

func (s *OpenApiSingleResponse) SetSuccess(v bool) *OpenApiSingleResponse {
	s.Success = &v
	return s
}

type OpenApiSingleResponseData struct {
	// example:
	//
	// FINISHED
	ModelTrainStatus *string `json:"modelTrainStatus,omitempty" xml:"modelTrainStatus,omitempty"`
}

func (s OpenApiSingleResponseData) String() string {
	return tea.Prettify(s)
}

func (s OpenApiSingleResponseData) GoString() string {
	return s.String()
}

func (s *OpenApiSingleResponseData) SetModelTrainStatus(v string) *OpenApiSingleResponseData {
	s.ModelTrainStatus = &v
	return s
}

type Personalizedtxt2imgAddInferenceJobCmd struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int32 `json:"seed,omitempty" xml:"seed,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobCmd) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobCmd) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobCmd {
	s.ImageNumber = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobCmd {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobCmd {
	s.Prompt = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetSeed(v int32) *Personalizedtxt2imgAddInferenceJobCmd {
	s.Seed = &v
	return s
}

type Personalizedtxt2imgAddModelTrainJobCmd struct {
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 熊猫图片生成
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dog
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobCmd) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobCmd) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetName(v string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.ObjectType = &v
	return s
}

type Personalizedtxt2imgInferenceJobInfoDTO struct {
	// example:
	//
	// 123456
	CreateUserId *string `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 123456
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 123456
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 123456
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgInferenceJobInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgInferenceJobInfoDTO) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetCreateUserId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.CreateUserId = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetJobStatus(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetModelId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetResultImageUrl(v []*string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.ResultImageUrl = v
	return s
}

type Personalizedtxt2imgModelTrainJobInfoDTO struct {
	CreateTime       *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId     *string                                   `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	ImageUrl         []*string                                 `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" type:"Repeated"`
	InferenceJobList []*Personalizedtxt2imgInferenceJobInfoDTO `json:"InferenceJobList,omitempty" xml:"InferenceJobList,omitempty" type:"Repeated"`
	JobStatus        *string                                   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Name             *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType       *string                                   `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s Personalizedtxt2imgModelTrainJobInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgModelTrainJobInfoDTO) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetCreateTime(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetCreateUserId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.CreateUserId = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetImageUrl(v []*string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetInferenceJobList(v []*Personalizedtxt2imgInferenceJobInfoDTO) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetJobStatus(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetName(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetObjectType(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.Id = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13
	Grade        *string   `json:"grade,omitempty" xml:"grade,omitempty"`
	KeySentences []*string `json:"keySentences,omitempty" xml:"keySentences,omitempty" type:"Repeated"`
	KeyWords     []*string `json:"keyWords,omitempty" xml:"keyWords,omitempty" type:"Repeated"`
	// example:
	//
	// Understanding unique professions such as dog walkers, hotel test sleepers, and food tasters, including their job responsibilities and the benefits or challenges associated with each role.
	LearningObject *string `json:"learningObject,omitempty" xml:"learningObject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dog walker Dog walking, as a profession, originated in the US. Some may think that it\\"s a perfect job, because dog walkers won\\"t be imprisoned in an office. But it\\"s actually manual labour. At their busiest, dog walkers may have more than ten dogs to take care of in a day. Hotel test sleeper A hotel test sleeper, as the name suggests, has to write expert reviews about the facilities, locations, prices, dining and other services of hotels, in order to provide evaluations and guides for travelers. Hotel test sleepers don\\"t need to punch in for work and they get about ten thousand yuan as income every month. What a comfortable job! Food taster In ancient times, a food taster was a person who tasted foods (or drinks) to be served to someone else, to confirm that it was safe to eat. But now, those working as food tasters just get to taste various new foods and drinks aimed at specific regions across the world. They then give their opinions on these products to the companies and suggest improvements.
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	Textbook    *string `json:"textbook,omitempty" xml:"textbook,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6440xxxxxxxxxx5fafc98c421
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateRequest) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetGrade(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Grade = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetKeySentences(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.KeySentences = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetKeyWords(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.KeyWords = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetLearningObject(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.LearningObject = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTextContent(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.TextContent = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTextbook(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Textbook = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTopic(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Topic = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetUserId(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.UserId = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBody struct {
	Data *AITeacherExpansionPracticeTaskGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
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

func (s AITeacherExpansionPracticeTaskGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetData(v *AITeacherExpansionPracticeTaskGenerateResponseBodyData) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetErrCode(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetErrMessage(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetHttpStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetRequestId(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetSuccess(v bool) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.Success = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyData struct {
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	BackgroundDescription *string                                                        `json:"backgroundDescription,omitempty" xml:"backgroundDescription,omitempty"`
	RoleSet               *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet `json:"roleSet,omitempty" xml:"roleSet,omitempty" type:"Struct"`
	// example:
	//
	// Hey Jamie, do you know what a travel blogger does?
	StartSentence *string                                                              `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	TaskContent   []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent `json:"taskContent,omitempty" xml:"taskContent,omitempty" type:"Repeated"`
	// example:
	//
	// textbook_dialogue
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetBackgroundDescription(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.BackgroundDescription = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetRoleSet(v *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.RoleSet = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetStartSentence(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.StartSentence = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetTaskContent(v []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.TaskContent = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetTaskType(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.TaskType = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet struct {
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) SetAssistant(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	s.Assistant = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) SetUser(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	s.User = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent struct {
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) SetAssistant(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	s.Assistant = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) SetUser(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	s.User = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AITeacherExpansionPracticeTaskGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponse) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetHeaders(v map[string]*string) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.Headers = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetBody(v *AITeacherExpansionPracticeTaskGenerateResponseBody) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.Body = v
	return s
}

type AITeacherSyncPracticeTaskGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13
	Grade        *string   `json:"grade,omitempty" xml:"grade,omitempty"`
	KeySentences []*string `json:"keySentences,omitempty" xml:"keySentences,omitempty" type:"Repeated"`
	KeyWords     []*string `json:"keyWords,omitempty" xml:"keyWords,omitempty" type:"Repeated"`
	// example:
	//
	// Understanding unique professions such as dog walkers, hotel test sleepers, and food tasters, including their job responsibilities and the benefits or challenges associated with each role.
	LearningObject *string `json:"learningObject,omitempty" xml:"learningObject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dog walker Dog walking, as a profession, originated in the US. Some may think that it\\"s a perfect job, because dog walkers won\\"t be imprisoned in an office. But it\\"s actually manual labour. At their busiest, dog walkers may have more than ten dogs to take care of in a day. Hotel test sleeper A hotel test sleeper, as the name suggests, has to write expert reviews about the facilities, locations, prices, dining and other services of hotels, in order to provide evaluations and guides for travelers. Hotel test sleepers don\\"t need to punch in for work and they get about ten thousand yuan as income every month. What a comfortable job! Food taster In ancient times, a food taster was a person who tasted foods (or drinks) to be served to someone else, to confirm that it was safe to eat. But now, those working as food tasters just get to taste various new foods and drinks aimed at specific regions across the world. They then give their opinions on these products to the companies and suggest improvements.
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	Textbook    *string `json:"textbook,omitempty" xml:"textbook,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6440xxxxxxxxxx5fafc98c421
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateRequest) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetGrade(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Grade = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetKeySentences(v []*string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.KeySentences = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetKeyWords(v []*string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.KeyWords = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetLearningObject(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.LearningObject = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTextContent(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.TextContent = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTextbook(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Textbook = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTopic(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Topic = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetUserId(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.UserId = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponseBody struct {
	Data *AITeacherSyncPracticeTaskGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
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

func (s AITeacherSyncPracticeTaskGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetData(v *AITeacherSyncPracticeTaskGenerateResponseBodyData) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetErrCode(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetErrMessage(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetHttpStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetRequestId(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetSuccess(v bool) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.Success = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponseBodyData struct {
	TaskContent []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent `json:"taskContent,omitempty" xml:"taskContent,omitempty" type:"Repeated"`
	// example:
	//
	// textbook_question_answering
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) SetTaskContent(v []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	s.TaskContent = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) SetTaskType(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	s.TaskType = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent struct {
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) SetAssistant(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	s.Assistant = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) SetUser(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	s.User = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AITeacherSyncPracticeTaskGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponse) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetHeaders(v map[string]*string) *AITeacherSyncPracticeTaskGenerateResponse {
	s.Headers = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetBody(v *AITeacherSyncPracticeTaskGenerateResponseBody) *AITeacherSyncPracticeTaskGenerateResponse {
	s.Body = v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetData(v []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.Success = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData struct {
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"FreeConcurrencyCount,omitempty" xml:"FreeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"FreeCount,omitempty" xml:"FreeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// AI算法模型-个性化文生图-在线按量调用
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetFreeConcurrencyCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetFreeCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetServiceCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetServiceName(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.ServiceName = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.Body = v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetData(v []*AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.Success = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData struct {
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"FreeConcurrencyCount,omitempty" xml:"FreeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"FreeCount,omitempty" xml:"FreeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// AI算法模型-个性化文生图-在线按量调用
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetFreeConcurrencyCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetFreeCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetServiceCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetServiceName(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.ServiceName = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherExpansionDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string                                            `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	Records      []*ExecuteAITeacherExpansionDialogueRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *ExecuteAITeacherExpansionDialogueRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Let\\"s talk about traffic rules.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.Background = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.LanguageCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueRequestRecords) *ExecuteAITeacherExpansionDialogueRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRequest {
	s.RoleInfo = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.StartSentence = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.Topic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueRequestRoleInfo {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherExpansionDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueResponseBodyData) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherExpansionDialogueResponseBodyData struct {
	// example:
	//
	// 1
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// 1
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// example:
	//
	// true
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// true
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// example:
	//
	// 2
	QuestionIndex *int32 `json:"questionIndex,omitempty" xml:"questionIndex,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetChineseResult(v string) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetQuestionIndex(v int32) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.QuestionIndex = &v
	return s
}

type ExecuteAITeacherExpansionDialogueResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherExpansionDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetBody(v *ExecuteAITeacherExpansionDialogueResponseBody) *ExecuteAITeacherExpansionDialogueResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*ExecuteAITeacherExpansionDialogueRefineRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.Background = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.LanguageCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueRefineRequestRecords) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.RoleInfo = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.StartSentence = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.Topic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Jane, a caring mother
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Lily, a friendly student
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherExpansionDialogueRefineResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueRefineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineResponseBodyData struct {
	// example:
	//
	// Yes, I\\"ll be right there.
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) SetResult(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBodyData {
	s.Result = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherExpansionDialogueRefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueRefineResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetBody(v *ExecuteAITeacherExpansionDialogueRefineResponseBody) *ExecuteAITeacherExpansionDialogueRefineResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In this dialogue, you will be playing the role of Lily, a young girl. I will be Jane, Lily\\"s mother. We are in the kitchen, where I am preparing dinner. I am asking you about your food preferences, specifically if you like meat, fish, and milk. You like meat and milk, but you don\\"t like fish because of its smell. I explain to you the nutritional benefits of these foods and suggest alternatives for the ones you don\\"t like. Finally, I invite you to start eating.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	Records       []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords       `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about food.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.Background = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.RoleInfo = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.StartSentence = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.Topic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Jane, a caring mother
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Lily, a friendly student
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateResponseBodyData struct {
	// example:
	//
	// 太好了，谢谢你过来，莉莉。你喜欢吃肉吗？
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) SetResult(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData {
	s.Result = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherExpansionDialogueTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueTranslateResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetBody(v *ExecuteAITeacherExpansionDialogueTranslateResponseBody) *ExecuteAITeacherExpansionDialogueTranslateResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherGrammarCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i is good
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckRequest) SetContent(v string) *ExecuteAITeacherGrammarCheckRequest {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckRequest) SetUserId(v string) *ExecuteAITeacherGrammarCheckRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherGrammarCheckResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherGrammarCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherGrammarCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetData(v *ExecuteAITeacherGrammarCheckResponseBodyData) *ExecuteAITeacherGrammarCheckResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetErrCode(v string) *ExecuteAITeacherGrammarCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetErrMessage(v string) *ExecuteAITeacherGrammarCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetRequestId(v string) *ExecuteAITeacherGrammarCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetSuccess(v bool) *ExecuteAITeacherGrammarCheckResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherGrammarCheckResponseBodyData struct {
	// example:
	//
	// 主语 "I" 对应的动词应该是 "am" 而不是 "is"。
	Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// example:
	//
	// I am good.
	Correction *string `json:"correction,omitempty" xml:"correction,omitempty"`
	// example:
	//
	// Has_Error
	CorrectionStatus *string `json:"correctionStatus,omitempty" xml:"correctionStatus,omitempty"`
	ErrorReason      *string `json:"errorReason,omitempty" xml:"errorReason,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetAnalysis(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.Analysis = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetCorrection(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.Correction = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetCorrectionStatus(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.CorrectionStatus = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetErrorReason(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.ErrorReason = &v
	return s
}

type ExecuteAITeacherGrammarCheckResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherGrammarCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherGrammarCheckResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetBody(v *ExecuteAITeacherGrammarCheckResponseBody) *ExecuteAITeacherGrammarCheckResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherSyncDialogueRequest struct {
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherSyncDialogueRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string                                       `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	Records      []*ExecuteAITeacherSyncDialogueRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueRequestDialogueTasks) *ExecuteAITeacherSyncDialogueRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetLanguageCode(v string) *ExecuteAITeacherSyncDialogueRequest {
	s.LanguageCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetRecords(v []*ExecuteAITeacherSyncDialogueRequestRecords) *ExecuteAITeacherSyncDialogueRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetUserId(v string) *ExecuteAITeacherSyncDialogueRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherSyncDialogueRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherSyncDialogueRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetContent(v string) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetOrder(v int32) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetRole(v string) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherSyncDialogueResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherSyncDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherSyncDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetData(v *ExecuteAITeacherSyncDialogueResponseBodyData) *ExecuteAITeacherSyncDialogueResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetErrCode(v string) *ExecuteAITeacherSyncDialogueResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetErrMessage(v string) *ExecuteAITeacherSyncDialogueResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetRequestId(v string) *ExecuteAITeacherSyncDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetSuccess(v bool) *ExecuteAITeacherSyncDialogueResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherSyncDialogueResponseBodyData struct {
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// example:
	//
	// true
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// example:
	//
	// 2
	QuestionIndex *int32 `json:"questionIndex,omitempty" xml:"questionIndex,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetQuestionIndex(v int32) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.QuestionIndex = &v
	return s
}

type ExecuteAITeacherSyncDialogueResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherSyncDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetBody(v *ExecuteAITeacherSyncDialogueResponseBody) *ExecuteAITeacherSyncDialogueResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateRequest struct {
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	Records       []*ExecuteAITeacherSyncDialogueTranslateRequestRecords       `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherSyncDialogueTranslateRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetRecords(v []*ExecuteAITeacherSyncDialogueTranslateRequestRecords) *ExecuteAITeacherSyncDialogueTranslateRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetUserId(v string) *ExecuteAITeacherSyncDialogueTranslateRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetContent(v string) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetOrder(v int32) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetRole(v string) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherSyncDialogueTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherSyncDialogueTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetData(v *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetErrCode(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetErrMessage(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetRequestId(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetSuccess(v bool) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateResponseBodyData struct {
	// example:
	//
	// 太好了，谢谢你过来，莉莉。你喜欢吃肉吗？
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) SetResult(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBodyData {
	s.Result = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherSyncDialogueTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueTranslateResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetBody(v *ExecuteAITeacherSyncDialogueTranslateResponseBody) *ExecuteAITeacherSyncDialogueTranslateResponse {
	s.Body = v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*GetAITeacherExpansionDialogueSuggestionRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Let\\"s talk about traffic rules.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequest) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetBackground(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Background = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetDialogueTasks(v []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.DialogueTasks = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetLanguageCode(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.LanguageCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetRecords(v []*GetAITeacherExpansionDialogueSuggestionRequestRecords) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Records = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetRoleInfo(v *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.RoleInfo = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetStartSentence(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.StartSentence = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetTopic(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Topic = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetUserId(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.UserId = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetAssistant(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetAssistantTranslate(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetOrder(v int32) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetUser(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.User = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRecords) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetContent(v string) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Content = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetIsOffTopicControl(v bool) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetIsOnTopic(v bool) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetOrder(v int32) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Order = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetRole(v string) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Role = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) SetAssistant(v string) *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) SetUser(v string) *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	s.User = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionResponseBody struct {
	// example:
	//
	// []
	Data *GetAITeacherExpansionDialogueSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetAITeacherExpansionDialogueSuggestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetData(v *GetAITeacherExpansionDialogueSuggestionResponseBodyData) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetErrCode(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetErrMessage(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetHttpStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetRequestId(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetSuccess(v bool) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.Success = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionResponseBodyData struct {
	// example:
	//
	// 谢谢莉莉.你喜欢吃肉吗，莉莉？
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) SetChineseResult(v string) *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) SetEnglishResult(v string) *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	s.EnglishResult = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITeacherExpansionDialogueSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponse) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetHeaders(v map[string]*string) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.Headers = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetBody(v *GetAITeacherExpansionDialogueSuggestionResponseBody) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.Body = v
	return s
}

type GetAITeacherSyncDialogueSuggestionRequest struct {
	// This parameter is required.
	DialogueTasks []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*GetAITeacherSyncDialogueSuggestionRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequest) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetDialogueTasks(v []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) *GetAITeacherSyncDialogueSuggestionRequest {
	s.DialogueTasks = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetLanguageCode(v string) *GetAITeacherSyncDialogueSuggestionRequest {
	s.LanguageCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetRecords(v []*GetAITeacherSyncDialogueSuggestionRequestRecords) *GetAITeacherSyncDialogueSuggestionRequest {
	s.Records = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetUserId(v string) *GetAITeacherSyncDialogueSuggestionRequest {
	s.UserId = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetAssistant(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetAssistantTranslate(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetOrder(v int32) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetUser(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.User = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequestRecords) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetContent(v string) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Content = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetIsOffTopicControl(v bool) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetIsOnTopic(v bool) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetOrder(v int32) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Order = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetRole(v string) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Role = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionResponseBody struct {
	// example:
	//
	// []
	Data *GetAITeacherSyncDialogueSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetAITeacherSyncDialogueSuggestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetData(v *GetAITeacherSyncDialogueSuggestionResponseBodyData) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetErrCode(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetErrMessage(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetHttpStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetRequestId(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetSuccess(v bool) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.Success = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionResponseBodyData struct {
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	// example:
	//
	// 谢谢莉莉.你喜欢吃肉吗，莉莉？
	EnglishResult1 *string `json:"englishResult1,omitempty" xml:"englishResult1,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) SetEnglishResult(v string) *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) SetEnglishResult1(v string) *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	s.EnglishResult1 = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITeacherSyncDialogueSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponse) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetHeaders(v map[string]*string) *GetAITeacherSyncDialogueSuggestionResponse {
	s.Headers = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetBody(v *GetAITeacherSyncDialogueSuggestionResponseBody) *GetAITeacherSyncDialogueSuggestionResponse {
	s.Body = v
	return s
}

type PersonalizedTextToImageAddInferenceJobRequest struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// example:
	//
	// 1
	Strength *float64 `json:"strength,omitempty" xml:"strength,omitempty"`
	// example:
	//
	// 800
	TrainSteps *int32 `json:"trainSteps,omitempty" xml:"trainSteps,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetImageNumber(v int32) *PersonalizedTextToImageAddInferenceJobRequest {
	s.ImageNumber = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetImageUrl(v []*string) *PersonalizedTextToImageAddInferenceJobRequest {
	s.ImageUrl = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetPrompt(v string) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Prompt = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetSeed(v int64) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Seed = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetStrength(v float64) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Strength = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetTrainSteps(v int32) *PersonalizedTextToImageAddInferenceJobRequest {
	s.TrainSteps = &v
	return s
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
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponseBodyData) GoString() string {
	return s.String()
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

type PersonalizedTextToImageAddInferenceJobResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalizedTextToImageAddInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageAddInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetBody(v *PersonalizedTextToImageAddInferenceJobResponseBody) *PersonalizedTextToImageAddInferenceJobResponse {
	s.Body = v
	return s
}

type PersonalizedTextToImageQueryImageAssetRequest struct {
	// example:
	//
	// base64
	EncodeFormat *string `json:"encodeFormat,omitempty" xml:"encodeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000.png
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
}

func (s PersonalizedTextToImageQueryImageAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryImageAssetRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) SetEncodeFormat(v string) *PersonalizedTextToImageQueryImageAssetRequest {
	s.EncodeFormat = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) SetImageId(v string) *PersonalizedTextToImageQueryImageAssetRequest {
	s.ImageId = &v
	return s
}

type PersonalizedTextToImageQueryImageAssetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageQueryImageAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryImageAssetResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryImageAssetResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetStatusCode(v int32) *PersonalizedTextToImageQueryImageAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetBody(v interface{}) *PersonalizedTextToImageQueryImageAssetResponse {
	s.Body = v
	return s
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// girl-xxxx-xxxx-xxxx-xxxx
	InferenceJobId *string `json:"inferenceJobId,omitempty" xml:"inferenceJobId,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) SetInferenceJobId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest {
	s.InferenceJobId = &v
	return s
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
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GoString() string {
	return s.String()
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

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse struct {
	Headers    map[string]*string                                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetBody(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgAddInferenceJobRequest struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobRequest {
	s.ImageNumber = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobRequest {
	s.Prompt = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetSeed(v int64) *Personalizedtxt2imgAddInferenceJobRequest {
	s.Seed = &v
	return s
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponseBodyData) GoString() string {
	return s.String()
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

type Personalizedtxt2imgAddInferenceJobResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgAddInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgAddInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetBody(v *Personalizedtxt2imgAddInferenceJobResponseBody) *Personalizedtxt2imgAddInferenceJobResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgAddModelTrainJobRequest struct {
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 熊猫图片生成
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dog
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// example:
	//
	// 800
	TrainSteps *int32 `json:"trainSteps,omitempty" xml:"trainSteps,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetName(v string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetTrainSteps(v int32) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.TrainSteps = &v
	return s
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyData) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GoString() string {
	return s.String()
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

type Personalizedtxt2imgAddModelTrainJobResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgAddModelTrainJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetBody(v *Personalizedtxt2imgAddModelTrainJobResponseBody) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryImageAssetRequest struct {
	// example:
	//
	// base64
	EncodeFormat *string `json:"encodeFormat,omitempty" xml:"encodeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000.png
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// girl-xxxx-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
}

func (s Personalizedtxt2imgQueryImageAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryImageAssetRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetEncodeFormat(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.EncodeFormat = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetImageId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.ImageId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetModelId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetPromptId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.PromptId = &v
	return s
}

type Personalizedtxt2imgQueryImageAssetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryImageAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryImageAssetResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryImageAssetResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryImageAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetBody(v interface{}) *Personalizedtxt2imgQueryImageAssetResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryInferenceJobInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 180
	InferenceJobId *string `json:"inferenceJobId,omitempty" xml:"inferenceJobId,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoRequest) SetInferenceJobId(v string) *Personalizedtxt2imgQueryInferenceJobInfoRequest {
	s.InferenceJobId = &v
	return s
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GoString() string {
	return s.String()
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

type Personalizedtxt2imgQueryInferenceJobInfoResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryInferenceJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetBody(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GoString() string {
	return s.String()
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

type Personalizedtxt2imgQueryModelTrainJobListResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryModelTrainJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetBody(v *Personalizedtxt2imgQueryModelTrainJobListResponseBody) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusRequest) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainStatusRequest {
	s.ModelId = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetData(v *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.Success = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusResponseBodyData struct {
	// example:
	//
	// FINISHED
	ModelTrainStatus *string `json:"modelTrainStatus,omitempty" xml:"modelTrainStatus,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) SetModelTrainStatus(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData {
	s.ModelTrainStatus = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryModelTrainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetBody(v *Personalizedtxt2imgQueryModelTrainStatusResponseBody) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aicontent"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练问答对生成
//
// @param request - AITeacherExpansionPracticeTaskGenerateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AITeacherExpansionPracticeTaskGenerateResponse
func (client *Client) AITeacherExpansionPracticeTaskGenerateWithOptions(request *AITeacherExpansionPracticeTaskGenerateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AITeacherExpansionPracticeTaskGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.KeySentences)) {
		body["keySentences"] = request.KeySentences
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWords)) {
		body["keyWords"] = request.KeyWords
	}

	if !tea.BoolValue(util.IsUnset(request.LearningObject)) {
		body["learningObject"] = request.LearningObject
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["textContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Textbook)) {
		body["textbook"] = request.Textbook
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AITeacherExpansionPracticeTaskGenerate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/generateTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AITeacherExpansionPracticeTaskGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练问答对生成
//
// @param request - AITeacherExpansionPracticeTaskGenerateRequest
//
// @return AITeacherExpansionPracticeTaskGenerateResponse
func (client *Client) AITeacherExpansionPracticeTaskGenerate(request *AITeacherExpansionPracticeTaskGenerateRequest) (_result *AITeacherExpansionPracticeTaskGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AITeacherExpansionPracticeTaskGenerateResponse{}
	_body, _err := client.AITeacherExpansionPracticeTaskGenerateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步基础练问答对生成
//
// @param request - AITeacherSyncPracticeTaskGenerateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AITeacherSyncPracticeTaskGenerateResponse
func (client *Client) AITeacherSyncPracticeTaskGenerateWithOptions(request *AITeacherSyncPracticeTaskGenerateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AITeacherSyncPracticeTaskGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.KeySentences)) {
		body["keySentences"] = request.KeySentences
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWords)) {
		body["keyWords"] = request.KeyWords
	}

	if !tea.BoolValue(util.IsUnset(request.LearningObject)) {
		body["learningObject"] = request.LearningObject
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["textContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Textbook)) {
		body["textbook"] = request.Textbook
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AITeacherSyncPracticeTaskGenerate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/generateTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AITeacherSyncPracticeTaskGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步基础练问答对生成
//
// @param request - AITeacherSyncPracticeTaskGenerateRequest
//
// @return AITeacherSyncPracticeTaskGenerateResponse
func (client *Client) AITeacherSyncPracticeTaskGenerate(request *AITeacherSyncPracticeTaskGenerateRequest) (_result *AITeacherSyncPracticeTaskGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AITeacherSyncPracticeTaskGenerateResponse{}
	_body, _err := client.AITeacherSyncPracticeTaskGenerateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AliyunConsoleOpenApiQueryAliyunConsoleServcieList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunconsole/queryAliyunConsoleServcieList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieList() (_result *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse{}
	_body, _err := client.AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServiceListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AliyunConsoleOpenApiQueryAliyunConsoleServiceList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunconsole/queryAliyunConsoleServiceList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServiceList() (_result *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse{}
	_body, _err := client.AliyunConsoleOpenApiQueryAliyunConsoleServiceListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 进行拓展练对话
//
// @param request - ExecuteAITeacherExpansionDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueResponse
func (client *Client) ExecuteAITeacherExpansionDialogueWithOptions(request *ExecuteAITeacherExpansionDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherExpansionDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/executeExpansionTraining"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 进行拓展练对话
//
// @param request - ExecuteAITeacherExpansionDialogueRequest
//
// @return ExecuteAITeacherExpansionDialogueResponse
func (client *Client) ExecuteAITeacherExpansionDialogue(request *ExecuteAITeacherExpansionDialogueRequest) (_result *ExecuteAITeacherExpansionDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherExpansionDialogueResponse{}
	_body, _err := client.ExecuteAITeacherExpansionDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练根据上下文进行润色
//
// @param request - ExecuteAITeacherExpansionDialogueRefineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueRefineResponse
func (client *Client) ExecuteAITeacherExpansionDialogueRefineWithOptions(request *ExecuteAITeacherExpansionDialogueRefineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueRefineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherExpansionDialogueRefine"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/refineByContext"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueRefineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练根据上下文进行润色
//
// @param request - ExecuteAITeacherExpansionDialogueRefineRequest
//
// @return ExecuteAITeacherExpansionDialogueRefineResponse
func (client *Client) ExecuteAITeacherExpansionDialogueRefine(request *ExecuteAITeacherExpansionDialogueRefineRequest) (_result *ExecuteAITeacherExpansionDialogueRefineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherExpansionDialogueRefineResponse{}
	_body, _err := client.ExecuteAITeacherExpansionDialogueRefineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练语境翻译
//
// @param request - ExecuteAITeacherExpansionDialogueTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueTranslateResponse
func (client *Client) ExecuteAITeacherExpansionDialogueTranslateWithOptions(request *ExecuteAITeacherExpansionDialogueTranslateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherExpansionDialogueTranslate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/translate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练语境翻译
//
// @param request - ExecuteAITeacherExpansionDialogueTranslateRequest
//
// @return ExecuteAITeacherExpansionDialogueTranslateResponse
func (client *Client) ExecuteAITeacherExpansionDialogueTranslate(request *ExecuteAITeacherExpansionDialogueTranslateRequest) (_result *ExecuteAITeacherExpansionDialogueTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherExpansionDialogueTranslateResponse{}
	_body, _err := client.ExecuteAITeacherExpansionDialogueTranslateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语法检测
//
// @param request - ExecuteAITeacherGrammarCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherGrammarCheckResponse
func (client *Client) ExecuteAITeacherGrammarCheckWithOptions(request *ExecuteAITeacherGrammarCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherGrammarCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherGrammarCheck"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/common/grammarChecking"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherGrammarCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语法检测
//
// @param request - ExecuteAITeacherGrammarCheckRequest
//
// @return ExecuteAITeacherGrammarCheckResponse
func (client *Client) ExecuteAITeacherGrammarCheck(request *ExecuteAITeacherGrammarCheckRequest) (_result *ExecuteAITeacherGrammarCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherGrammarCheckResponse{}
	_body, _err := client.ExecuteAITeacherGrammarCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 进行同步练对话
//
// @param request - ExecuteAITeacherSyncDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherSyncDialogueResponse
func (client *Client) ExecuteAITeacherSyncDialogueWithOptions(request *ExecuteAITeacherSyncDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherSyncDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherSyncDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/executeSyncTraining"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherSyncDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 进行同步练对话
//
// @param request - ExecuteAITeacherSyncDialogueRequest
//
// @return ExecuteAITeacherSyncDialogueResponse
func (client *Client) ExecuteAITeacherSyncDialogue(request *ExecuteAITeacherSyncDialogueRequest) (_result *ExecuteAITeacherSyncDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherSyncDialogueResponse{}
	_body, _err := client.ExecuteAITeacherSyncDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步练语境翻译
//
// @param request - ExecuteAITeacherSyncDialogueTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherSyncDialogueTranslateResponse
func (client *Client) ExecuteAITeacherSyncDialogueTranslateWithOptions(request *ExecuteAITeacherSyncDialogueTranslateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherSyncDialogueTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherSyncDialogueTranslate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/translate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherSyncDialogueTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步练语境翻译
//
// @param request - ExecuteAITeacherSyncDialogueTranslateRequest
//
// @return ExecuteAITeacherSyncDialogueTranslateResponse
func (client *Client) ExecuteAITeacherSyncDialogueTranslate(request *ExecuteAITeacherSyncDialogueTranslateRequest) (_result *ExecuteAITeacherSyncDialogueTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherSyncDialogueTranslateResponse{}
	_body, _err := client.ExecuteAITeacherSyncDialogueTranslateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练小助手
//
// @param request - GetAITeacherExpansionDialogueSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITeacherExpansionDialogueSuggestionResponse
func (client *Client) GetAITeacherExpansionDialogueSuggestionWithOptions(request *GetAITeacherExpansionDialogueSuggestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAITeacherExpansionDialogueSuggestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAITeacherExpansionDialogueSuggestion"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/suggestion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAITeacherExpansionDialogueSuggestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练小助手
//
// @param request - GetAITeacherExpansionDialogueSuggestionRequest
//
// @return GetAITeacherExpansionDialogueSuggestionResponse
func (client *Client) GetAITeacherExpansionDialogueSuggestion(request *GetAITeacherExpansionDialogueSuggestionRequest) (_result *GetAITeacherExpansionDialogueSuggestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAITeacherExpansionDialogueSuggestionResponse{}
	_body, _err := client.GetAITeacherExpansionDialogueSuggestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步练小助手
//
// @param request - GetAITeacherSyncDialogueSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITeacherSyncDialogueSuggestionResponse
func (client *Client) GetAITeacherSyncDialogueSuggestionWithOptions(request *GetAITeacherSyncDialogueSuggestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAITeacherSyncDialogueSuggestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAITeacherSyncDialogueSuggestion"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/suggestion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAITeacherSyncDialogueSuggestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步练小助手
//
// @param request - GetAITeacherSyncDialogueSuggestionRequest
//
// @return GetAITeacherSyncDialogueSuggestionResponse
func (client *Client) GetAITeacherSyncDialogueSuggestion(request *GetAITeacherSyncDialogueSuggestionRequest) (_result *GetAITeacherSyncDialogueSuggestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAITeacherSyncDialogueSuggestionResponse{}
	_body, _err := client.GetAITeacherSyncDialogueSuggestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个预训练模型创建图片推理任务
//
// @param request - PersonalizedTextToImageAddInferenceJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageAddInferenceJobResponse
func (client *Client) PersonalizedTextToImageAddInferenceJobWithOptions(request *PersonalizedTextToImageAddInferenceJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PersonalizedTextToImageAddInferenceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageNumber)) {
		body["imageNumber"] = request.ImageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		body["seed"] = request.Seed
	}

	if !tea.BoolValue(util.IsUnset(request.Strength)) {
		body["strength"] = request.Strength
	}

	if !tea.BoolValue(util.IsUnset(request.TrainSteps)) {
		body["trainSteps"] = request.TrainSteps
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PersonalizedTextToImageAddInferenceJob"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/addPreModelInferenceJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PersonalizedTextToImageAddInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个预训练模型创建图片推理任务
//
// @param request - PersonalizedTextToImageAddInferenceJobRequest
//
// @return PersonalizedTextToImageAddInferenceJobResponse
func (client *Client) PersonalizedTextToImageAddInferenceJob(request *PersonalizedTextToImageAddInferenceJobRequest) (_result *PersonalizedTextToImageAddInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PersonalizedTextToImageAddInferenceJobResponse{}
	_body, _err := client.PersonalizedTextToImageAddInferenceJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/通过唯一的图片编号获取图片内容
//
// @param request - PersonalizedTextToImageQueryImageAssetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageQueryImageAssetResponse
func (client *Client) PersonalizedTextToImageQueryImageAssetWithOptions(request *PersonalizedTextToImageQueryImageAssetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PersonalizedTextToImageQueryImageAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeFormat)) {
		query["encodeFormat"] = request.EncodeFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["imageId"] = request.ImageId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PersonalizedTextToImageQueryImageAsset"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryImageAssetFromImageId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("any"),
	}
	_result = &PersonalizedTextToImageQueryImageAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/通过唯一的图片编号获取图片内容
//
// @param request - PersonalizedTextToImageQueryImageAssetRequest
//
// @return PersonalizedTextToImageQueryImageAssetResponse
func (client *Client) PersonalizedTextToImageQueryImageAsset(request *PersonalizedTextToImageQueryImageAssetRequest) (_result *PersonalizedTextToImageQueryImageAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PersonalizedTextToImageQueryImageAssetResponse{}
	_body, _err := client.PersonalizedTextToImageQueryImageAssetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询预制模型推理任务的状态
//
// @param request - PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
func (client *Client) PersonalizedTextToImageQueryPreModelInferenceJobInfoWithOptions(request *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InferenceJobId)) {
		query["inferenceJobId"] = request.InferenceJobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PersonalizedTextToImageQueryPreModelInferenceJobInfo"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryPreModelInferenceJobInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询预制模型推理任务的状态
//
// @param request - PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest
//
// @return PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
func (client *Client) PersonalizedTextToImageQueryPreModelInferenceJobInfo(request *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) (_result *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse{}
	_body, _err := client.PersonalizedTextToImageQueryPreModelInferenceJobInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个模型创建图片推理任务
//
// @param request - Personalizedtxt2imgAddInferenceJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgAddInferenceJobResponse
func (client *Client) Personalizedtxt2imgAddInferenceJobWithOptions(request *Personalizedtxt2imgAddInferenceJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgAddInferenceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageNumber)) {
		body["imageNumber"] = request.ImageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		body["seed"] = request.Seed
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgAddInferenceJob"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/addInferenceJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgAddInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个模型创建图片推理任务
//
// @param request - Personalizedtxt2imgAddInferenceJobRequest
//
// @return Personalizedtxt2imgAddInferenceJobResponse
func (client *Client) Personalizedtxt2imgAddInferenceJob(request *Personalizedtxt2imgAddInferenceJobRequest) (_result *Personalizedtxt2imgAddInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgAddInferenceJobResponse{}
	_body, _err := client.Personalizedtxt2imgAddInferenceJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/创建一个模型训练任务
//
// @param request - Personalizedtxt2imgAddModelTrainJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgAddModelTrainJobResponse
func (client *Client) Personalizedtxt2imgAddModelTrainJobWithOptions(request *Personalizedtxt2imgAddModelTrainJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgAddModelTrainJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.TrainSteps)) {
		body["trainSteps"] = request.TrainSteps
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgAddModelTrainJob"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/addModelTrainJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgAddModelTrainJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/创建一个模型训练任务
//
// @param request - Personalizedtxt2imgAddModelTrainJobRequest
//
// @return Personalizedtxt2imgAddModelTrainJobResponse
func (client *Client) Personalizedtxt2imgAddModelTrainJob(request *Personalizedtxt2imgAddModelTrainJobRequest) (_result *Personalizedtxt2imgAddModelTrainJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgAddModelTrainJobResponse{}
	_body, _err := client.Personalizedtxt2imgAddModelTrainJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/图片二进制内容获取
//
// @param request - Personalizedtxt2imgQueryImageAssetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryImageAssetResponse
func (client *Client) Personalizedtxt2imgQueryImageAssetWithOptions(request *Personalizedtxt2imgQueryImageAssetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryImageAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeFormat)) {
		query["encodeFormat"] = request.EncodeFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["imageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.PromptId)) {
		query["promptId"] = request.PromptId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryImageAsset"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryImageAsset"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("any"),
	}
	_result = &Personalizedtxt2imgQueryImageAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/图片二进制内容获取
//
// @param request - Personalizedtxt2imgQueryImageAssetRequest
//
// @return Personalizedtxt2imgQueryImageAssetResponse
func (client *Client) Personalizedtxt2imgQueryImageAsset(request *Personalizedtxt2imgQueryImageAssetRequest) (_result *Personalizedtxt2imgQueryImageAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryImageAssetResponse{}
	_body, _err := client.Personalizedtxt2imgQueryImageAssetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型推理任务的状态和结果信息
//
// @param request - Personalizedtxt2imgQueryInferenceJobInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryInferenceJobInfoResponse
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfoWithOptions(request *Personalizedtxt2imgQueryInferenceJobInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryInferenceJobInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InferenceJobId)) {
		query["inferenceJobId"] = request.InferenceJobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryInferenceJobInfo"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryInferenceJobInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgQueryInferenceJobInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型推理任务的状态和结果信息
//
// @param request - Personalizedtxt2imgQueryInferenceJobInfoRequest
//
// @return Personalizedtxt2imgQueryInferenceJobInfoResponse
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfo(request *Personalizedtxt2imgQueryInferenceJobInfoRequest) (_result *Personalizedtxt2imgQueryInferenceJobInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryInferenceJobInfoResponse{}
	_body, _err := client.Personalizedtxt2imgQueryInferenceJobInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型训练任务列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryModelTrainJobListResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainJobListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryModelTrainJobListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryModelTrainJobList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryModelTrainJobList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgQueryModelTrainJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型训练任务列表
//
// @return Personalizedtxt2imgQueryModelTrainJobListResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainJobList() (_result *Personalizedtxt2imgQueryModelTrainJobListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryModelTrainJobListResponse{}
	_body, _err := client.Personalizedtxt2imgQueryModelTrainJobListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/模型训练状态查询
//
// @param request - Personalizedtxt2imgQueryModelTrainStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryModelTrainStatusResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainStatusWithOptions(request *Personalizedtxt2imgQueryModelTrainStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryModelTrainStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["modelId"] = request.ModelId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryModelTrainStatus"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryModelTrainStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgQueryModelTrainStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/模型训练状态查询
//
// @param request - Personalizedtxt2imgQueryModelTrainStatusRequest
//
// @return Personalizedtxt2imgQueryModelTrainStatusResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainStatus(request *Personalizedtxt2imgQueryModelTrainStatusRequest) (_result *Personalizedtxt2imgQueryModelTrainStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryModelTrainStatusResponse{}
	_body, _err := client.Personalizedtxt2imgQueryModelTrainStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
