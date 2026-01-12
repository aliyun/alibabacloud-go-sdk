// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProjectResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListProjectResponseBody
	GetCurrent() *int32
	SetData(v []*ListProjectResponseBodyData) *ListProjectResponseBody
	GetData() []*ListProjectResponseBodyData
	SetMessage(v string) *ListProjectResponseBody
	GetMessage() *string
	SetPages(v int32) *ListProjectResponseBody
	GetPages() *int32
	SetRequestId(v string) *ListProjectResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListProjectResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListProjectResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListProjectResponseBody
	GetTotal() *int32
}

type ListProjectResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                         `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*ListProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                         `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProjectResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListProjectResponseBody) GetData() []*ListProjectResponseBodyData {
	return s.Data
}

func (s *ListProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectResponseBody) GetPages() *int32 {
	return s.Pages
}

func (s *ListProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListProjectResponseBody) SetCode(v string) *ListProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectResponseBody) SetCurrent(v int32) *ListProjectResponseBody {
	s.Current = &v
	return s
}

func (s *ListProjectResponseBody) SetData(v []*ListProjectResponseBodyData) *ListProjectResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectResponseBody) SetMessage(v string) *ListProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectResponseBody) SetPages(v int32) *ListProjectResponseBody {
	s.Pages = &v
	return s
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectResponseBody) SetSize(v int32) *ListProjectResponseBody {
	s.Size = &v
	return s
}

func (s *ListProjectResponseBody) SetSuccess(v bool) *ListProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectResponseBody) SetTotal(v int32) *ListProjectResponseBody {
	s.Total = &v
	return s
}

func (s *ListProjectResponseBody) Validate() error {
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

type ListProjectResponseBodyData struct {
	BizUsage         *string                             `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	CreateTime       *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset          *ListProjectResponseBodyDataDataset `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Ext              *string                             `json:"Ext,omitempty" xml:"Ext,omitempty"`
	ExtInfo          map[string]interface{}              `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Id               *string                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                             `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                             `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                             `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source           *ListProjectResponseBodyDataSource  `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status           *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                             `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
	User             *ListProjectResponseBodyDataUser    `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyData) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListProjectResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProjectResponseBodyData) GetDataset() *ListProjectResponseBodyDataDataset {
	return s.Dataset
}

func (s *ListProjectResponseBodyData) GetExt() *string {
	return s.Ext
}

func (s *ListProjectResponseBodyData) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *ListProjectResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListProjectResponseBodyData) GetIntro() *string {
	return s.Intro
}

func (s *ListProjectResponseBodyData) GetMaterialCoverUrl() *string {
	return s.MaterialCoverUrl
}

func (s *ListProjectResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListProjectResponseBodyData) GetSource() *ListProjectResponseBodyDataSource {
	return s.Source
}

func (s *ListProjectResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListProjectResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListProjectResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListProjectResponseBodyData) GetUser() *ListProjectResponseBodyDataUser {
	return s.User
}

func (s *ListProjectResponseBodyData) SetBizUsage(v string) *ListProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *ListProjectResponseBodyData) SetCreateTime(v string) *ListProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProjectResponseBodyData) SetDataset(v *ListProjectResponseBodyDataDataset) *ListProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *ListProjectResponseBodyData) SetExt(v string) *ListProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *ListProjectResponseBodyData) SetExtInfo(v map[string]interface{}) *ListProjectResponseBodyData {
	s.ExtInfo = v
	return s
}

func (s *ListProjectResponseBodyData) SetId(v string) *ListProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListProjectResponseBodyData) SetIntro(v string) *ListProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *ListProjectResponseBodyData) SetMaterialCoverUrl(v string) *ListProjectResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *ListProjectResponseBodyData) SetModifiedTime(v string) *ListProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectResponseBodyData) SetSource(v *ListProjectResponseBodyDataSource) *ListProjectResponseBodyData {
	s.Source = v
	return s
}

func (s *ListProjectResponseBodyData) SetStatus(v string) *ListProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListProjectResponseBodyData) SetTitle(v string) *ListProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListProjectResponseBodyData) SetType(v string) *ListProjectResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListProjectResponseBodyData) SetUser(v *ListProjectResponseBodyDataUser) *ListProjectResponseBodyData {
	s.User = v
	return s
}

func (s *ListProjectResponseBodyData) Validate() error {
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
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	ErrorMessage    *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	OriginResultUrl *string                `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PoseUrl         *string                `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s ListProjectResponseBodyDataDataset) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyDataDataset) GetBuildResultUrl() map[string]interface{} {
	return s.BuildResultUrl
}

func (s *ListProjectResponseBodyDataDataset) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListProjectResponseBodyDataDataset) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProjectResponseBodyDataDataset) GetGlbModelUrl() *string {
	return s.GlbModelUrl
}

func (s *ListProjectResponseBodyDataDataset) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *ListProjectResponseBodyDataDataset) GetOriginResultUrl() *string {
	return s.OriginResultUrl
}

func (s *ListProjectResponseBodyDataDataset) GetOssKey() *string {
	return s.OssKey
}

func (s *ListProjectResponseBodyDataDataset) GetPoseUrl() *string {
	return s.PoseUrl
}

func (s *ListProjectResponseBodyDataDataset) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *ListProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *ListProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetCoverUrl(v string) *ListProjectResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetErrorMessage(v string) *ListProjectResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetGlbModelUrl(v string) *ListProjectResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetModelUrl(v string) *ListProjectResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetOriginResultUrl(v string) *ListProjectResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetOssKey(v string) *ListProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetPoseUrl(v string) *ListProjectResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) SetPreviewUrl(v string) *ListProjectResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataDataset) Validate() error {
	return dara.Validate(s)
}

type ListProjectResponseBodyDataSource struct {
	Clothes []*ListProjectResponseBodyDataSourceClothes `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	Files   []*ListProjectResponseBodyDataSourceFiles   `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	OssKey  *string                                     `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s ListProjectResponseBodyDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyDataSource) GetClothes() []*ListProjectResponseBodyDataSourceClothes {
	return s.Clothes
}

func (s *ListProjectResponseBodyDataSource) GetFiles() []*ListProjectResponseBodyDataSourceFiles {
	return s.Files
}

func (s *ListProjectResponseBodyDataSource) GetOssKey() *string {
	return s.OssKey
}

func (s *ListProjectResponseBodyDataSource) SetClothes(v []*ListProjectResponseBodyDataSourceClothes) *ListProjectResponseBodyDataSource {
	s.Clothes = v
	return s
}

func (s *ListProjectResponseBodyDataSource) SetFiles(v []*ListProjectResponseBodyDataSourceFiles) *ListProjectResponseBodyDataSource {
	s.Files = v
	return s
}

func (s *ListProjectResponseBodyDataSource) SetOssKey(v string) *ListProjectResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *ListProjectResponseBodyDataSource) Validate() error {
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
	return nil
}

type ListProjectResponseBodyDataSourceClothes struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectResponseBodyDataSourceClothes) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyDataSourceClothes) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyDataSourceClothes) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListProjectResponseBodyDataSourceClothes) GetName() *string {
	return s.Name
}

func (s *ListProjectResponseBodyDataSourceClothes) GetOssKey() *string {
	return s.OssKey
}

func (s *ListProjectResponseBodyDataSourceClothes) GetType() *string {
	return s.Type
}

func (s *ListProjectResponseBodyDataSourceClothes) SetCoverUrl(v string) *ListProjectResponseBodyDataSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceClothes) SetName(v string) *ListProjectResponseBodyDataSourceClothes {
	s.Name = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceClothes) SetOssKey(v string) *ListProjectResponseBodyDataSourceClothes {
	s.OssKey = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceClothes) SetType(v string) *ListProjectResponseBodyDataSourceClothes {
	s.Type = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceClothes) Validate() error {
	return dara.Validate(s)
}

type ListProjectResponseBodyDataSourceFiles struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Size     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListProjectResponseBodyDataSourceFiles) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyDataSourceFiles) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyDataSourceFiles) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListProjectResponseBodyDataSourceFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListProjectResponseBodyDataSourceFiles) GetId() *string {
	return s.Id
}

func (s *ListProjectResponseBodyDataSourceFiles) GetSize() *int64 {
	return s.Size
}

func (s *ListProjectResponseBodyDataSourceFiles) GetType() *string {
	return s.Type
}

func (s *ListProjectResponseBodyDataSourceFiles) GetUrl() *string {
	return s.Url
}

func (s *ListProjectResponseBodyDataSourceFiles) SetCoverUrl(v string) *ListProjectResponseBodyDataSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceFiles) SetFileName(v string) *ListProjectResponseBodyDataSourceFiles {
	s.FileName = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceFiles) SetId(v string) *ListProjectResponseBodyDataSourceFiles {
	s.Id = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceFiles) SetSize(v int64) *ListProjectResponseBodyDataSourceFiles {
	s.Size = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceFiles) SetType(v string) *ListProjectResponseBodyDataSourceFiles {
	s.Type = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceFiles) SetUrl(v string) *ListProjectResponseBodyDataSourceFiles {
	s.Url = &v
	return s
}

func (s *ListProjectResponseBodyDataSourceFiles) Validate() error {
	return dara.Validate(s)
}

type ListProjectResponseBodyDataUser struct {
	AliyunUid    *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Nickname     *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
}

func (s ListProjectResponseBodyDataUser) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyDataUser) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListProjectResponseBodyDataUser) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProjectResponseBodyDataUser) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListProjectResponseBodyDataUser) GetNickname() *string {
	return s.Nickname
}

func (s *ListProjectResponseBodyDataUser) SetAliyunUid(v string) *ListProjectResponseBodyDataUser {
	s.AliyunUid = &v
	return s
}

func (s *ListProjectResponseBodyDataUser) SetCreateTime(v string) *ListProjectResponseBodyDataUser {
	s.CreateTime = &v
	return s
}

func (s *ListProjectResponseBodyDataUser) SetModifiedTime(v string) *ListProjectResponseBodyDataUser {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectResponseBodyDataUser) SetNickname(v string) *ListProjectResponseBodyDataUser {
	s.Nickname = &v
	return s
}

func (s *ListProjectResponseBodyDataUser) Validate() error {
	return dara.Validate(s)
}
