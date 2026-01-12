// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryProjectDetailResponseBody
	GetCode() *string
	SetData(v *QueryProjectDetailResponseBodyData) *QueryProjectDetailResponseBody
	GetData() *QueryProjectDetailResponseBodyData
	SetErrorName(v string) *QueryProjectDetailResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *QueryProjectDetailResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *QueryProjectDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryProjectDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryProjectDetailResponseBody
	GetSuccess() *bool
}

type QueryProjectDetailResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                             `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                              `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryProjectDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryProjectDetailResponseBody) GetData() *QueryProjectDetailResponseBodyData {
	return s.Data
}

func (s *QueryProjectDetailResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *QueryProjectDetailResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryProjectDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryProjectDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProjectDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryProjectDetailResponseBody) SetCode(v string) *QueryProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProjectDetailResponseBody) SetData(v *QueryProjectDetailResponseBodyData) *QueryProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectDetailResponseBody) SetErrorName(v string) *QueryProjectDetailResponseBody {
	s.ErrorName = &v
	return s
}

func (s *QueryProjectDetailResponseBody) SetHttpCode(v int32) *QueryProjectDetailResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryProjectDetailResponseBody) SetMessage(v string) *QueryProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProjectDetailResponseBody) SetRequestId(v string) *QueryProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectDetailResponseBody) SetSuccess(v bool) *QueryProjectDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProjectDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryProjectDetailResponseBodyData struct {
	BizUsage         *string                                        `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *QueryProjectDetailResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CreateMode       *string                                        `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset          *QueryProjectDetailResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                          `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                        `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                        `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                        `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                        `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source           *QueryProjectDetailResponseBodyDataSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status           *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryProjectDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyData) GetBizUsage() *string {
	return s.BizUsage
}

func (s *QueryProjectDetailResponseBodyData) GetBuildDetail() *QueryProjectDetailResponseBodyDataBuildDetail {
	return s.BuildDetail
}

func (s *QueryProjectDetailResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *QueryProjectDetailResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectDetailResponseBodyData) GetDataset() *QueryProjectDetailResponseBodyDataDataset {
	return s.Dataset
}

func (s *QueryProjectDetailResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *QueryProjectDetailResponseBodyData) GetDependencies() *string {
	return s.Dependencies
}

func (s *QueryProjectDetailResponseBodyData) GetExt() *string {
	return s.Ext
}

func (s *QueryProjectDetailResponseBodyData) GetId() *string {
	return s.Id
}

func (s *QueryProjectDetailResponseBodyData) GetIntro() *string {
	return s.Intro
}

func (s *QueryProjectDetailResponseBodyData) GetMaterialCoverUrl() *string {
	return s.MaterialCoverUrl
}

func (s *QueryProjectDetailResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryProjectDetailResponseBodyData) GetSource() *QueryProjectDetailResponseBodyDataSource {
	return s.Source
}

func (s *QueryProjectDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryProjectDetailResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *QueryProjectDetailResponseBodyData) GetType() *string {
	return s.Type
}

func (s *QueryProjectDetailResponseBodyData) SetBizUsage(v string) *QueryProjectDetailResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetBuildDetail(v *QueryProjectDetailResponseBodyDataBuildDetail) *QueryProjectDetailResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetCreateMode(v string) *QueryProjectDetailResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetCreateTime(v string) *QueryProjectDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetDataset(v *QueryProjectDetailResponseBodyDataDataset) *QueryProjectDetailResponseBodyData {
	s.Dataset = v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetDeleted(v bool) *QueryProjectDetailResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetDependencies(v string) *QueryProjectDetailResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetExt(v string) *QueryProjectDetailResponseBodyData {
	s.Ext = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetId(v string) *QueryProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetIntro(v string) *QueryProjectDetailResponseBodyData {
	s.Intro = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetMaterialCoverUrl(v string) *QueryProjectDetailResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetModifiedTime(v string) *QueryProjectDetailResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetSource(v *QueryProjectDetailResponseBodyDataSource) *QueryProjectDetailResponseBodyData {
	s.Source = v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetStatus(v string) *QueryProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetTitle(v string) *QueryProjectDetailResponseBodyData {
	s.Title = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) SetType(v string) *QueryProjectDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryProjectDetailResponseBodyData) Validate() error {
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

type QueryProjectDetailResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s QueryProjectDetailResponseBodyDataBuildDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) GetCompletedTime() *string {
	return s.CompletedTime
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) GetEstimatedDuration() *int64 {
	return s.EstimatedDuration
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) GetRunningTime() *string {
	return s.RunningTime
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) SetCompletedTime(v string) *QueryProjectDetailResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) SetErrorMessage(v string) *QueryProjectDetailResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *QueryProjectDetailResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) SetRunningTime(v string) *QueryProjectDetailResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) SetSubmitTime(v string) *QueryProjectDetailResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataBuildDetail) Validate() error {
	return dara.Validate(s)
}

type QueryProjectDetailResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                           `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                          `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                            `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                          `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	Id              *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string                                          `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                          `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                          `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                          `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *QueryProjectDetailResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                          `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                          `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s QueryProjectDetailResponseBodyDataDataset) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetBuildResultUrl() map[string]interface{} {
	return s.BuildResultUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetDeleted() *bool {
	return s.Deleted
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetGlbModelUrl() *string {
	return s.GlbModelUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetId() *string {
	return s.Id
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetOriginResultUrl() *string {
	return s.OriginResultUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetOssKey() *string {
	return s.OssKey
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetPolicy() *QueryProjectDetailResponseBodyDataDatasetPolicy {
	return s.Policy
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetPoseUrl() *string {
	return s.PoseUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *QueryProjectDetailResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetCoverUrl(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetCreateTime(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetDeleted(v bool) *QueryProjectDetailResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetErrorCode(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetErrorMessage(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetGlbModelUrl(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetId(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.Id = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetModelUrl(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetModifiedTime(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetOriginResultUrl(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetOssKey(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetPolicy(v *QueryProjectDetailResponseBodyDataDatasetPolicy) *QueryProjectDetailResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetPoseUrl(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) SetPreviewUrl(v string) *QueryProjectDetailResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDataset) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryProjectDetailResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s QueryProjectDetailResponseBodyDataDatasetPolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) GetAccessId() *string {
	return s.AccessId
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) GetDir() *string {
	return s.Dir
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) GetExpire() *string {
	return s.Expire
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) GetHost() *string {
	return s.Host
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) GetPolicy() *string {
	return s.Policy
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) GetSignature() *string {
	return s.Signature
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) SetAccessId(v string) *QueryProjectDetailResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) SetDir(v string) *QueryProjectDetailResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) SetExpire(v string) *QueryProjectDetailResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) SetHost(v string) *QueryProjectDetailResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) SetPolicy(v string) *QueryProjectDetailResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) SetSignature(v string) *QueryProjectDetailResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataDatasetPolicy) Validate() error {
	return dara.Validate(s)
}

type QueryProjectDetailResponseBodyDataSource struct {
	Clothes      []*QueryProjectDetailResponseBodyDataSourceClothes `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Files        []*QueryProjectDetailResponseBodyDataSourceFiles   `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	Id           *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                            `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *QueryProjectDetailResponseBodyDataSourcePolicy    `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s QueryProjectDetailResponseBodyDataSource) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataSource) GetClothes() []*QueryProjectDetailResponseBodyDataSourceClothes {
	return s.Clothes
}

func (s *QueryProjectDetailResponseBodyDataSource) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectDetailResponseBodyDataSource) GetDeleted() *bool {
	return s.Deleted
}

func (s *QueryProjectDetailResponseBodyDataSource) GetFiles() []*QueryProjectDetailResponseBodyDataSourceFiles {
	return s.Files
}

func (s *QueryProjectDetailResponseBodyDataSource) GetId() *string {
	return s.Id
}

func (s *QueryProjectDetailResponseBodyDataSource) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryProjectDetailResponseBodyDataSource) GetOssKey() *string {
	return s.OssKey
}

func (s *QueryProjectDetailResponseBodyDataSource) GetPolicy() *QueryProjectDetailResponseBodyDataSourcePolicy {
	return s.Policy
}

func (s *QueryProjectDetailResponseBodyDataSource) SetClothes(v []*QueryProjectDetailResponseBodyDataSourceClothes) *QueryProjectDetailResponseBodyDataSource {
	s.Clothes = v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetCreateTime(v string) *QueryProjectDetailResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetDeleted(v bool) *QueryProjectDetailResponseBodyDataSource {
	s.Deleted = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetFiles(v []*QueryProjectDetailResponseBodyDataSourceFiles) *QueryProjectDetailResponseBodyDataSource {
	s.Files = v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetId(v string) *QueryProjectDetailResponseBodyDataSource {
	s.Id = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetModifiedTime(v string) *QueryProjectDetailResponseBodyDataSource {
	s.ModifiedTime = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetOssKey(v string) *QueryProjectDetailResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) SetPolicy(v *QueryProjectDetailResponseBodyDataSourcePolicy) *QueryProjectDetailResponseBodyDataSource {
	s.Policy = v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSource) Validate() error {
	if s.Clothes != nil {
		for _, item := range s.Clothes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryProjectDetailResponseBodyDataSourceClothes struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryProjectDetailResponseBodyDataSourceClothes) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataSourceClothes) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) GetId() *string {
	return s.Id
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) GetName() *string {
	return s.Name
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) GetOssKey() *string {
	return s.OssKey
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) GetType() *string {
	return s.Type
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) SetCoverUrl(v string) *QueryProjectDetailResponseBodyDataSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) SetId(v string) *QueryProjectDetailResponseBodyDataSourceClothes {
	s.Id = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) SetName(v string) *QueryProjectDetailResponseBodyDataSourceClothes {
	s.Name = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) SetOssKey(v string) *QueryProjectDetailResponseBodyDataSourceClothes {
	s.OssKey = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) SetType(v string) *QueryProjectDetailResponseBodyDataSourceClothes {
	s.Type = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceClothes) Validate() error {
	return dara.Validate(s)
}

type QueryProjectDetailResponseBodyDataSourceFiles struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Size     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryProjectDetailResponseBodyDataSourceFiles) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataSourceFiles) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) GetFileName() *string {
	return s.FileName
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) GetId() *string {
	return s.Id
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) GetSize() *int64 {
	return s.Size
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) GetType() *string {
	return s.Type
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) GetUrl() *string {
	return s.Url
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) SetCoverUrl(v string) *QueryProjectDetailResponseBodyDataSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) SetFileName(v string) *QueryProjectDetailResponseBodyDataSourceFiles {
	s.FileName = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) SetId(v string) *QueryProjectDetailResponseBodyDataSourceFiles {
	s.Id = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) SetSize(v int64) *QueryProjectDetailResponseBodyDataSourceFiles {
	s.Size = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) SetType(v string) *QueryProjectDetailResponseBodyDataSourceFiles {
	s.Type = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) SetUrl(v string) *QueryProjectDetailResponseBodyDataSourceFiles {
	s.Url = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourceFiles) Validate() error {
	return dara.Validate(s)
}

type QueryProjectDetailResponseBodyDataSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s QueryProjectDetailResponseBodyDataSourcePolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponseBodyDataSourcePolicy) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) GetAccessId() *string {
	return s.AccessId
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) GetDir() *string {
	return s.Dir
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) GetExpire() *string {
	return s.Expire
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) GetHost() *string {
	return s.Host
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) GetPolicy() *string {
	return s.Policy
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) GetSignature() *string {
	return s.Signature
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) SetAccessId(v string) *QueryProjectDetailResponseBodyDataSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) SetDir(v string) *QueryProjectDetailResponseBodyDataSourcePolicy {
	s.Dir = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) SetExpire(v string) *QueryProjectDetailResponseBodyDataSourcePolicy {
	s.Expire = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) SetHost(v string) *QueryProjectDetailResponseBodyDataSourcePolicy {
	s.Host = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) SetPolicy(v string) *QueryProjectDetailResponseBodyDataSourcePolicy {
	s.Policy = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) SetSignature(v string) *QueryProjectDetailResponseBodyDataSourcePolicy {
	s.Signature = &v
	return s
}

func (s *QueryProjectDetailResponseBodyDataSourcePolicy) Validate() error {
	return dara.Validate(s)
}
