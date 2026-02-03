// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddDocumentParam(v *SubmitVideoAnalysisTaskRequestAddDocumentParam) *SubmitVideoAnalysisTaskRequest
	GetAddDocumentParam() *SubmitVideoAnalysisTaskRequestAddDocumentParam
	SetAutoRoleRecognitionVideoUrl(v string) *SubmitVideoAnalysisTaskRequest
	GetAutoRoleRecognitionVideoUrl() *string
	SetDeduplicationId(v string) *SubmitVideoAnalysisTaskRequest
	GetDeduplicationId() *string
	SetExcludeGenerateOptions(v []*string) *SubmitVideoAnalysisTaskRequest
	GetExcludeGenerateOptions() []*string
	SetFaceIdentitySimilarityMinScore(v float32) *SubmitVideoAnalysisTaskRequest
	GetFaceIdentitySimilarityMinScore() *float32
	SetFrameSampleMethod(v *SubmitVideoAnalysisTaskRequestFrameSampleMethod) *SubmitVideoAnalysisTaskRequest
	GetFrameSampleMethod() *SubmitVideoAnalysisTaskRequestFrameSampleMethod
	SetGenerateOptions(v []*string) *SubmitVideoAnalysisTaskRequest
	GetGenerateOptions() []*string
	SetLanguage(v string) *SubmitVideoAnalysisTaskRequest
	GetLanguage() *string
	SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequest
	GetModelCustomPromptTemplate() *string
	SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *SubmitVideoAnalysisTaskRequest
	GetModelId() *string
	SetSnapshotInterval(v float64) *SubmitVideoAnalysisTaskRequest
	GetSnapshotInterval() *float64
	SetSplitInterval(v int32) *SubmitVideoAnalysisTaskRequest
	GetSplitInterval() *int32
	SetSplitType(v string) *SubmitVideoAnalysisTaskRequest
	GetSplitType() *string
	SetTextProcessTasks(v []*SubmitVideoAnalysisTaskRequestTextProcessTasks) *SubmitVideoAnalysisTaskRequest
	GetTextProcessTasks() []*SubmitVideoAnalysisTaskRequestTextProcessTasks
	SetVideoCaptionInfo(v *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) *SubmitVideoAnalysisTaskRequest
	GetVideoCaptionInfo() *SubmitVideoAnalysisTaskRequestVideoCaptionInfo
	SetVideoExtraInfo(v string) *SubmitVideoAnalysisTaskRequest
	GetVideoExtraInfo() *string
	SetVideoModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequest
	GetVideoModelCustomPromptTemplate() *string
	SetVideoModelId(v string) *SubmitVideoAnalysisTaskRequest
	GetVideoModelId() *string
	SetVideoRoles(v []*SubmitVideoAnalysisTaskRequestVideoRoles) *SubmitVideoAnalysisTaskRequest
	GetVideoRoles() []*SubmitVideoAnalysisTaskRequestVideoRoles
	SetVideoShotFaceIdentityCount(v int32) *SubmitVideoAnalysisTaskRequest
	GetVideoShotFaceIdentityCount() *int32
	SetVideoUrl(v string) *SubmitVideoAnalysisTaskRequest
	GetVideoUrl() *string
	SetVideoUrls(v []*string) *SubmitVideoAnalysisTaskRequest
	GetVideoUrls() []*string
}

type SubmitVideoAnalysisTaskRequest struct {
	AddDocumentParam            *SubmitVideoAnalysisTaskRequestAddDocumentParam `json:"addDocumentParam,omitempty" xml:"addDocumentParam,omitempty" type:"Struct"`
	AutoRoleRecognitionVideoUrl *string                                         `json:"autoRoleRecognitionVideoUrl,omitempty" xml:"autoRoleRecognitionVideoUrl,omitempty"`
	// example:
	//
	// 1
	DeduplicationId                *string                                          `json:"deduplicationId,omitempty" xml:"deduplicationId,omitempty"`
	ExcludeGenerateOptions         []*string                                        `json:"excludeGenerateOptions,omitempty" xml:"excludeGenerateOptions,omitempty" type:"Repeated"`
	FaceIdentitySimilarityMinScore *float32                                         `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethod              *SubmitVideoAnalysisTaskRequestFrameSampleMethod `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty" type:"Struct"`
	GenerateOptions                []*string                                        `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
	// example:
	//
	// chinese
	Language                  *string `json:"language,omitempty" xml:"language,omitempty"`
	ModelCustomPromptTemplate *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// PlotDetail
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 2
	SnapshotInterval *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval *int32 `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	// example:
	//
	// fixDuration
	SplitType                      *string                                           `json:"splitType,omitempty" xml:"splitType,omitempty"`
	TextProcessTasks               []*SubmitVideoAnalysisTaskRequestTextProcessTasks `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty" type:"Repeated"`
	VideoCaptionInfo               *SubmitVideoAnalysisTaskRequestVideoCaptionInfo   `json:"videoCaptionInfo,omitempty" xml:"videoCaptionInfo,omitempty" type:"Struct"`
	VideoExtraInfo                 *string                                           `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string                                           `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max-latest
	VideoModelId               *string                                     `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRoles                 []*SubmitVideoAnalysisTaskRequestVideoRoles `json:"videoRoles,omitempty" xml:"videoRoles,omitempty" type:"Repeated"`
	VideoShotFaceIdentityCount *int32                                      `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
	// example:
	//
	// http://xxxx.mp4
	VideoUrl  *string   `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
	VideoUrls []*string `json:"videoUrls,omitempty" xml:"videoUrls,omitempty" type:"Repeated"`
}

func (s SubmitVideoAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequest) GetAddDocumentParam() *SubmitVideoAnalysisTaskRequestAddDocumentParam {
	return s.AddDocumentParam
}

func (s *SubmitVideoAnalysisTaskRequest) GetAutoRoleRecognitionVideoUrl() *string {
	return s.AutoRoleRecognitionVideoUrl
}

func (s *SubmitVideoAnalysisTaskRequest) GetDeduplicationId() *string {
	return s.DeduplicationId
}

func (s *SubmitVideoAnalysisTaskRequest) GetExcludeGenerateOptions() []*string {
	return s.ExcludeGenerateOptions
}

func (s *SubmitVideoAnalysisTaskRequest) GetFaceIdentitySimilarityMinScore() *float32 {
	return s.FaceIdentitySimilarityMinScore
}

func (s *SubmitVideoAnalysisTaskRequest) GetFrameSampleMethod() *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	return s.FrameSampleMethod
}

func (s *SubmitVideoAnalysisTaskRequest) GetGenerateOptions() []*string {
	return s.GenerateOptions
}

func (s *SubmitVideoAnalysisTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *SubmitVideoAnalysisTaskRequest) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *SubmitVideoAnalysisTaskRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *SubmitVideoAnalysisTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitVideoAnalysisTaskRequest) GetSnapshotInterval() *float64 {
	return s.SnapshotInterval
}

func (s *SubmitVideoAnalysisTaskRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *SubmitVideoAnalysisTaskRequest) GetSplitType() *string {
	return s.SplitType
}

func (s *SubmitVideoAnalysisTaskRequest) GetTextProcessTasks() []*SubmitVideoAnalysisTaskRequestTextProcessTasks {
	return s.TextProcessTasks
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoCaptionInfo() *SubmitVideoAnalysisTaskRequestVideoCaptionInfo {
	return s.VideoCaptionInfo
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoModelCustomPromptTemplate() *string {
	return s.VideoModelCustomPromptTemplate
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoModelId() *string {
	return s.VideoModelId
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoRoles() []*SubmitVideoAnalysisTaskRequestVideoRoles {
	return s.VideoRoles
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoShotFaceIdentityCount() *int32 {
	return s.VideoShotFaceIdentityCount
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SubmitVideoAnalysisTaskRequest) GetVideoUrls() []*string {
	return s.VideoUrls
}

func (s *SubmitVideoAnalysisTaskRequest) SetAddDocumentParam(v *SubmitVideoAnalysisTaskRequestAddDocumentParam) *SubmitVideoAnalysisTaskRequest {
	s.AddDocumentParam = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetAutoRoleRecognitionVideoUrl(v string) *SubmitVideoAnalysisTaskRequest {
	s.AutoRoleRecognitionVideoUrl = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetDeduplicationId(v string) *SubmitVideoAnalysisTaskRequest {
	s.DeduplicationId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetExcludeGenerateOptions(v []*string) *SubmitVideoAnalysisTaskRequest {
	s.ExcludeGenerateOptions = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetFaceIdentitySimilarityMinScore(v float32) *SubmitVideoAnalysisTaskRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetFrameSampleMethod(v *SubmitVideoAnalysisTaskRequestFrameSampleMethod) *SubmitVideoAnalysisTaskRequest {
	s.FrameSampleMethod = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetGenerateOptions(v []*string) *SubmitVideoAnalysisTaskRequest {
	s.GenerateOptions = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetLanguage(v string) *SubmitVideoAnalysisTaskRequest {
	s.Language = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetModelId(v string) *SubmitVideoAnalysisTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetSnapshotInterval(v float64) *SubmitVideoAnalysisTaskRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetSplitInterval(v int32) *SubmitVideoAnalysisTaskRequest {
	s.SplitInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetSplitType(v string) *SubmitVideoAnalysisTaskRequest {
	s.SplitType = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetTextProcessTasks(v []*SubmitVideoAnalysisTaskRequestTextProcessTasks) *SubmitVideoAnalysisTaskRequest {
	s.TextProcessTasks = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoCaptionInfo(v *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) *SubmitVideoAnalysisTaskRequest {
	s.VideoCaptionInfo = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoExtraInfo(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoExtraInfo = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoModelId(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoRoles(v []*SubmitVideoAnalysisTaskRequestVideoRoles) *SubmitVideoAnalysisTaskRequest {
	s.VideoRoles = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoShotFaceIdentityCount(v int32) *SubmitVideoAnalysisTaskRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoUrl(v string) *SubmitVideoAnalysisTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) SetVideoUrls(v []*string) *SubmitVideoAnalysisTaskRequest {
	s.VideoUrls = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequest) Validate() error {
	if s.AddDocumentParam != nil {
		if err := s.AddDocumentParam.Validate(); err != nil {
			return err
		}
	}
	if s.FrameSampleMethod != nil {
		if err := s.FrameSampleMethod.Validate(); err != nil {
			return err
		}
	}
	if s.TextProcessTasks != nil {
		for _, item := range s.TextProcessTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoCaptionInfo != nil {
		if err := s.VideoCaptionInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VideoRoles != nil {
		for _, item := range s.VideoRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskRequestAddDocumentParam struct {
	DatasetId   *int64                                                  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	DatasetName *string                                                 `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	Document    *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument `json:"document,omitempty" xml:"document,omitempty" type:"Struct"`
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParam) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParam) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) GetDocument() *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	return s.Document
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) SetDatasetId(v int64) *SubmitVideoAnalysisTaskRequestAddDocumentParam {
	s.DatasetId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) SetDatasetName(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParam {
	s.DatasetName = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) SetDocument(v *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) *SubmitVideoAnalysisTaskRequestAddDocumentParam {
	s.Document = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParam) Validate() error {
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskRequestAddDocumentParamDocument struct {
	CategoryUuid *string                                                         `json:"categoryUuid,omitempty" xml:"categoryUuid,omitempty"`
	DocId        *string                                                         `json:"docId,omitempty" xml:"docId,omitempty"`
	Extend1      *string                                                         `json:"extend1,omitempty" xml:"extend1,omitempty"`
	Extend2      *string                                                         `json:"extend2,omitempty" xml:"extend2,omitempty"`
	Extend3      *string                                                         `json:"extend3,omitempty" xml:"extend3,omitempty"`
	Metadata     *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	Tags         []*string                                                       `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Title        *string                                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetDocId() *string {
	return s.DocId
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetExtend1() *string {
	return s.Extend1
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetExtend2() *string {
	return s.Extend2
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetExtend3() *string {
	return s.Extend3
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetMetadata() *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata {
	return s.Metadata
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetTags() []*string {
	return s.Tags
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) GetTitle() *string {
	return s.Title
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetCategoryUuid(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.CategoryUuid = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetDocId(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.DocId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetExtend1(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.Extend1 = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetExtend2(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.Extend2 = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetExtend3(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.Extend3 = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetMetadata(v *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.Metadata = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetTags(v []*string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.Tags = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) SetTitle(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument {
	s.Title = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocument) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata struct {
	KeyValues []*SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues `json:"keyValues,omitempty" xml:"keyValues,omitempty" type:"Repeated"`
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata) GetKeyValues() []*SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues {
	return s.KeyValues
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata) SetKeyValues(v []*SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata {
	s.KeyValues = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadata) Validate() error {
	if s.KeyValues != nil {
		for _, item := range s.KeyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) GetKey() *string {
	return s.Key
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) GetValue() *string {
	return s.Value
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) SetKey(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues {
	s.Key = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) SetValue(v string) *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues {
	s.Value = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestAddDocumentParamDocumentMetadataKeyValues) Validate() error {
	return dara.Validate(s)
}

type SubmitVideoAnalysisTaskRequestFrameSampleMethod struct {
	// example:
	//
	// 2
	Interval *float64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// standard
	MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
	// example:
	//
	// 768
	Pixel *int32 `json:"pixel,omitempty" xml:"pixel,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestFrameSampleMethod) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestFrameSampleMethod) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) GetInterval() *float64 {
	return s.Interval
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) GetMethodName() *string {
	return s.MethodName
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) GetPixel() *int32 {
	return s.Pixel
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) SetInterval(v float64) *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	s.Interval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) SetMethodName(v string) *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	s.MethodName = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) SetPixel(v int32) *SubmitVideoAnalysisTaskRequestFrameSampleMethod {
	s.Pixel = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestFrameSampleMethod) Validate() error {
	return dara.Validate(s)
}

type SubmitVideoAnalysisTaskRequestTextProcessTasks struct {
	ModelCustomPromptTemplate   *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	ModelId                     *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestTextProcessTasks) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestTextProcessTasks) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskRequestTextProcessTasks {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskRequestTextProcessTasks {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) SetModelId(v string) *SubmitVideoAnalysisTaskRequestTextProcessTasks {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestTextProcessTasks) Validate() error {
	return dara.Validate(s)
}

type SubmitVideoAnalysisTaskRequestVideoCaptionInfo struct {
	// example:
	//
	// oss:// | http://
	VideoCaptionFileUrl *string                                                        `json:"videoCaptionFileUrl,omitempty" xml:"videoCaptionFileUrl,omitempty"`
	VideoCaptions       []*SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions `json:"videoCaptions,omitempty" xml:"videoCaptions,omitempty" type:"Repeated"`
}

func (s SubmitVideoAnalysisTaskRequestVideoCaptionInfo) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestVideoCaptionInfo) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) GetVideoCaptionFileUrl() *string {
	return s.VideoCaptionFileUrl
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) GetVideoCaptions() []*SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions {
	return s.VideoCaptions
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) SetVideoCaptionFileUrl(v string) *SubmitVideoAnalysisTaskRequestVideoCaptionInfo {
	s.VideoCaptionFileUrl = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) SetVideoCaptions(v []*SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) *SubmitVideoAnalysisTaskRequestVideoCaptionInfo {
	s.VideoCaptions = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfo) Validate() error {
	if s.VideoCaptions != nil {
		for _, item := range s.VideoCaptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions struct {
	// example:
	//
	// 10000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 张三
	Speaker *string `json:"speaker,omitempty" xml:"speaker,omitempty"`
	// example:
	//
	// 1000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 你好
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) GetSpeaker() *string {
	return s.Speaker
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) GetText() *string {
	return s.Text
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) SetEndTime(v int64) *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions {
	s.EndTime = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) SetSpeaker(v string) *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions {
	s.Speaker = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) SetStartTime(v int64) *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions {
	s.StartTime = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) SetText(v string) *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions {
	s.Text = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoCaptionInfoVideoCaptions) Validate() error {
	return dara.Validate(s)
}

type SubmitVideoAnalysisTaskRequestVideoRoles struct {
	IsAutoRecognition *bool                                                    `json:"isAutoRecognition,omitempty" xml:"isAutoRecognition,omitempty"`
	RoleInfo          *string                                                  `json:"roleInfo,omitempty" xml:"roleInfo,omitempty"`
	RoleName          *string                                                  `json:"roleName,omitempty" xml:"roleName,omitempty"`
	TimeIntervals     []*SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals `json:"timeIntervals,omitempty" xml:"timeIntervals,omitempty" type:"Repeated"`
	Urls              []*string                                                `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s SubmitVideoAnalysisTaskRequestVideoRoles) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestVideoRoles) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) GetIsAutoRecognition() *bool {
	return s.IsAutoRecognition
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) GetRoleInfo() *string {
	return s.RoleInfo
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) GetTimeIntervals() []*SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals {
	return s.TimeIntervals
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) GetUrls() []*string {
	return s.Urls
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetIsAutoRecognition(v bool) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.IsAutoRecognition = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetRoleInfo(v string) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.RoleInfo = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetRoleName(v string) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.RoleName = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetTimeIntervals(v []*SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.TimeIntervals = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) SetUrls(v []*string) *SubmitVideoAnalysisTaskRequestVideoRoles {
	s.Urls = v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRoles) Validate() error {
	if s.TimeIntervals != nil {
		for _, item := range s.TimeIntervals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals struct {
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) SetEndTime(v int64) *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals {
	s.EndTime = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) SetStartTime(v int64) *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals {
	s.StartTime = &v
	return s
}

func (s *SubmitVideoAnalysisTaskRequestVideoRolesTimeIntervals) Validate() error {
	return dara.Validate(s)
}
