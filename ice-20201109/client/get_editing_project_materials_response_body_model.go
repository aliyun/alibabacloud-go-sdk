// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveMaterials(v []*GetEditingProjectMaterialsResponseBodyLiveMaterials) *GetEditingProjectMaterialsResponseBody
	GetLiveMaterials() []*GetEditingProjectMaterialsResponseBodyLiveMaterials
	SetMediaInfos(v []*GetEditingProjectMaterialsResponseBodyMediaInfos) *GetEditingProjectMaterialsResponseBody
	GetMediaInfos() []*GetEditingProjectMaterialsResponseBodyMediaInfos
	SetProjectId(v string) *GetEditingProjectMaterialsResponseBody
	GetProjectId() *string
	SetProjectMaterials(v []*string) *GetEditingProjectMaterialsResponseBody
	GetProjectMaterials() []*string
	SetRequestId(v string) *GetEditingProjectMaterialsResponseBody
	GetRequestId() *string
}

type GetEditingProjectMaterialsResponseBody struct {
	// The materials associated with the live stream.
	LiveMaterials []*GetEditingProjectMaterialsResponseBodyLiveMaterials `json:"LiveMaterials,omitempty" xml:"LiveMaterials,omitempty" type:"Repeated"`
	// The media assets that meet the specified conditions.
	MediaInfos []*GetEditingProjectMaterialsResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// The project ID.
	//
	// example:
	//
	// *****67ae06542b9b93e0d1c387*****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The materials associated with the editing project. A live stream editing project will be associated with a regular editing project after the live streaming ends.
	//
	// example:
	//
	// *****9b145c5cafc2e057304fcd*****
	ProjectMaterials []*string `json:"ProjectMaterials,omitempty" xml:"ProjectMaterials,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******89-C21D-4B78-AE24-3788B8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBody) GetLiveMaterials() []*GetEditingProjectMaterialsResponseBodyLiveMaterials {
	return s.LiveMaterials
}

func (s *GetEditingProjectMaterialsResponseBody) GetMediaInfos() []*GetEditingProjectMaterialsResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *GetEditingProjectMaterialsResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectMaterialsResponseBody) GetProjectMaterials() []*string {
	return s.ProjectMaterials
}

func (s *GetEditingProjectMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEditingProjectMaterialsResponseBody) SetLiveMaterials(v []*GetEditingProjectMaterialsResponseBodyLiveMaterials) *GetEditingProjectMaterialsResponseBody {
	s.LiveMaterials = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetMediaInfos(v []*GetEditingProjectMaterialsResponseBodyMediaInfos) *GetEditingProjectMaterialsResponseBody {
	s.MediaInfos = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetProjectId(v string) *GetEditingProjectMaterialsResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetProjectMaterials(v []*string) *GetEditingProjectMaterialsResponseBody {
	s.ProjectMaterials = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetRequestId(v string) *GetEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyLiveMaterials struct {
	// The application name of the live stream.
	//
	// example:
	//
	// testrecord
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name of the live stream.
	//
	// example:
	//
	// test.alivecdn.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://test.alivecdn.com/testrecord/teststream
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// testrecord
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyLiveMaterials) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyLiveMaterials) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) GetAppName() *string {
	return s.AppName
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) GetDomainName() *string {
	return s.DomainName
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) GetLiveUrl() *string {
	return s.LiveUrl
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) GetStreamName() *string {
	return s.StreamName
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) SetAppName(v string) *GetEditingProjectMaterialsResponseBodyLiveMaterials {
	s.AppName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) SetDomainName(v string) *GetEditingProjectMaterialsResponseBodyLiveMaterials {
	s.DomainName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) SetLiveUrl(v string) *GetEditingProjectMaterialsResponseBodyLiveMaterials {
	s.LiveUrl = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) SetStreamName(v string) *GetEditingProjectMaterialsResponseBodyLiveMaterials {
	s.StreamName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyLiveMaterials) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMediaInfos struct {
	// The information about the file.
	FileInfoList []*GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information of the media asset.
	MediaBasicInfo *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// *****64623a94eca8516569c8fe*****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) GetFileInfoList() []*GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) GetMediaBasicInfo() *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) SetFileInfoList(v []*GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) *GetEditingProjectMaterialsResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) SetMediaBasicInfo(v *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) *GetEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) SetMediaId(v string) *GetEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList struct {
	// The basic information of the file, such as the duration and size.
	FileBasicInfo *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// The bitrate.
	//
	// example:
	//
	// 1132.68
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration.
	//
	// example:
	//
	// 216.206667
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// example:
	//
	// example.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 30611502
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file status.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// The file type.
	//
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The Object Storage Service (OSS) URL of the file.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The container format.
	//
	// example:
	//
	// mov,mp4,m4a,3gp,3g2,mj2
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height.
	//
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The region in which the file resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The width.
	//
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo struct {
	// The business type of the media asset.
	//
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category of the media asset.
	//
	// example:
	//
	// video
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The thumbnail URL of the media asset.
	//
	// example:
	//
	// http://sample-bucket.oss-cn-shanghai.aliyuncs.com/sample-corver.jpg?Expires=1628670610&OSSAccessKeyId=AK&Signature=signature
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created.
	//
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the media asset was deleted.
	//
	// example:
	//
	// 2020-12-26T04:11:08Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The description of the media asset.
	//
	// example:
	//
	// sample_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the media asset in another service.
	//
	// example:
	//
	// http://bucket.oss-cn-shanghai.aliyuncs.com/file.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// *****64623a94eca8516569c8f*****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags of the media asset.
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time when the media asset was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:08Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The snapshots of the media asset.
	//
	// example:
	//
	// null
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source of the media asset.
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The sprite of the media asset
	//
	// example:
	//
	// null
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The status of the media asset.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the media asset.
	//
	// example:
	//
	// file.mp4
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding status of the media asset.
	//
	// Valid values:
	//
	// 	- TranscodeSuccess: transcoding completed.
	//
	// 	- TranscodeFailed: transcoding failed.
	//
	// 	- Init: initializing.
	//
	// 	- Transcoding: transcoding in progress.
	//
	// example:
	//
	// Init
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// The user data.
	//
	// example:
	//
	// userData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
