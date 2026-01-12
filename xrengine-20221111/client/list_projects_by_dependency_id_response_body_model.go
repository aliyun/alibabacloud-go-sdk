// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsByDependencyIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProjectsByDependencyIdResponseBody
	GetCode() *string
	SetData(v []*ListProjectsByDependencyIdResponseBodyData) *ListProjectsByDependencyIdResponseBody
	GetData() []*ListProjectsByDependencyIdResponseBodyData
	SetErrorName(v string) *ListProjectsByDependencyIdResponseBody
	GetErrorName() *string
	SetMessage(v string) *ListProjectsByDependencyIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProjectsByDependencyIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectsByDependencyIdResponseBody
	GetSuccess() *bool
}

type ListProjectsByDependencyIdResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListProjectsByDependencyIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                                       `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProjectsByDependencyIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProjectsByDependencyIdResponseBody) GetData() []*ListProjectsByDependencyIdResponseBodyData {
	return s.Data
}

func (s *ListProjectsByDependencyIdResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListProjectsByDependencyIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectsByDependencyIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsByDependencyIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectsByDependencyIdResponseBody) SetCode(v string) *ListProjectsByDependencyIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBody) SetData(v []*ListProjectsByDependencyIdResponseBodyData) *ListProjectsByDependencyIdResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsByDependencyIdResponseBody) SetErrorName(v string) *ListProjectsByDependencyIdResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBody) SetMessage(v string) *ListProjectsByDependencyIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBody) SetRequestId(v string) *ListProjectsByDependencyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBody) SetSuccess(v bool) *ListProjectsByDependencyIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBody) Validate() error {
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

type ListProjectsByDependencyIdResponseBodyData struct {
	BizUsage     *string                                                `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail  *ListProjectsByDependencyIdResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CreateTime   *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset      *ListProjectsByDependencyIdResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Ext          *string                                                `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id           *string                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string                                                `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ModifiedTime *string                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source       *ListProjectsByDependencyIdResponseBodyDataSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status       *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string                                                `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectsByDependencyIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetBuildDetail() *ListProjectsByDependencyIdResponseBodyDataBuildDetail {
	return s.BuildDetail
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetDataset() *ListProjectsByDependencyIdResponseBodyDataDataset {
	return s.Dataset
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetExt() *string {
	return s.Ext
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetIntro() *string {
	return s.Intro
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetSource() *ListProjectsByDependencyIdResponseBodyDataSource {
	return s.Source
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListProjectsByDependencyIdResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetBizUsage(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetBuildDetail(v *ListProjectsByDependencyIdResponseBodyDataBuildDetail) *ListProjectsByDependencyIdResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetCreateTime(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetDataset(v *ListProjectsByDependencyIdResponseBodyDataDataset) *ListProjectsByDependencyIdResponseBodyData {
	s.Dataset = v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetExt(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.Ext = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetId(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetIntro(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.Intro = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetModifiedTime(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetSource(v *ListProjectsByDependencyIdResponseBodyDataSource) *ListProjectsByDependencyIdResponseBodyData {
	s.Source = v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetStatus(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetTitle(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) SetType(v string) *ListProjectsByDependencyIdResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyData) Validate() error {
	if s.BuildDetail != nil {
		if err := s.BuildDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsByDependencyIdResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListProjectsByDependencyIdResponseBodyDataBuildDetail) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) GetCompletedTime() *string {
	return s.CompletedTime
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) GetEstimatedDuration() *int64 {
	return s.EstimatedDuration
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) GetRunningTime() *string {
	return s.RunningTime
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) SetCompletedTime(v string) *ListProjectsByDependencyIdResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) SetErrorMessage(v string) *ListProjectsByDependencyIdResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *ListProjectsByDependencyIdResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) SetRunningTime(v string) *ListProjectsByDependencyIdResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) SetSubmitTime(v string) *ListProjectsByDependencyIdResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataBuildDetail) Validate() error {
	return dara.Validate(s)
}

type ListProjectsByDependencyIdResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	GlbModelUrl     *string                `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	OriginResultUrl *string                `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PoseUrl         *string                `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s ListProjectsByDependencyIdResponseBodyDataDataset) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetBuildResultUrl() map[string]interface{} {
	return s.BuildResultUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetGlbModelUrl() *string {
	return s.GlbModelUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetOriginResultUrl() *string {
	return s.OriginResultUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetOssKey() *string {
	return s.OssKey
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetPoseUrl() *string {
	return s.PoseUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetCoverUrl(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetGlbModelUrl(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetModelUrl(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetOriginResultUrl(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetOssKey(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetPoseUrl(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) SetPreviewUrl(v string) *ListProjectsByDependencyIdResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataDataset) Validate() error {
	return dara.Validate(s)
}

type ListProjectsByDependencyIdResponseBodyDataSource struct {
	CreateTime   *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id           *string                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                        `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	ProjectId    *int64                                                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceFiles  []*ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
}

func (s ListProjectsByDependencyIdResponseBodyDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) GetId() *string {
	return s.Id
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) GetOssKey() *string {
	return s.OssKey
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) GetSourceFiles() []*ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	return s.SourceFiles
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) SetCreateTime(v string) *ListProjectsByDependencyIdResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) SetId(v string) *ListProjectsByDependencyIdResponseBodyDataSource {
	s.Id = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) SetModifiedTime(v string) *ListProjectsByDependencyIdResponseBodyDataSource {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) SetOssKey(v string) *ListProjectsByDependencyIdResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) SetProjectId(v int64) *ListProjectsByDependencyIdResponseBodyDataSource {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) SetSourceFiles(v []*ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) *ListProjectsByDependencyIdResponseBodyDataSource {
	s.SourceFiles = v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSource) Validate() error {
	if s.SourceFiles != nil {
		for _, item := range s.SourceFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Size     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GetId() *string {
	return s.Id
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GetSize() *int64 {
	return s.Size
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GetType() *string {
	return s.Type
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) GetUrl() *string {
	return s.Url
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) SetCoverUrl(v string) *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) SetFileName(v string) *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) SetId(v string) *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	s.Id = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) SetSize(v int64) *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	s.Size = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) SetType(v string) *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) SetUrl(v string) *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles {
	s.Url = &v
	return s
}

func (s *ListProjectsByDependencyIdResponseBodyDataSourceSourceFiles) Validate() error {
	return dara.Validate(s)
}
