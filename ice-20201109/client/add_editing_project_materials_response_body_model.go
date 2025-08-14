// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveMaterials(v []*AddEditingProjectMaterialsResponseBodyLiveMaterials) *AddEditingProjectMaterialsResponseBody
	GetLiveMaterials() []*AddEditingProjectMaterialsResponseBodyLiveMaterials
	SetMediaInfos(v []*AddEditingProjectMaterialsResponseBodyMediaInfos) *AddEditingProjectMaterialsResponseBody
	GetMediaInfos() []*AddEditingProjectMaterialsResponseBodyMediaInfos
	SetProjectId(v string) *AddEditingProjectMaterialsResponseBody
	GetProjectId() *string
	SetProjectMaterials(v []*string) *AddEditingProjectMaterialsResponseBody
	GetProjectMaterials() []*string
	SetRequestId(v string) *AddEditingProjectMaterialsResponseBody
	GetRequestId() *string
}

type AddEditingProjectMaterialsResponseBody struct {
	// The materials associated with the live stream.
	LiveMaterials []*AddEditingProjectMaterialsResponseBodyLiveMaterials `json:"LiveMaterials,omitempty" xml:"LiveMaterials,omitempty" type:"Repeated"`
	// The media assets that meet the specified conditions.
	MediaInfos []*AddEditingProjectMaterialsResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// The ID of the online editing project.
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
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBody) GetLiveMaterials() []*AddEditingProjectMaterialsResponseBodyLiveMaterials {
	return s.LiveMaterials
}

func (s *AddEditingProjectMaterialsResponseBody) GetMediaInfos() []*AddEditingProjectMaterialsResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *AddEditingProjectMaterialsResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *AddEditingProjectMaterialsResponseBody) GetProjectMaterials() []*string {
	return s.ProjectMaterials
}

func (s *AddEditingProjectMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEditingProjectMaterialsResponseBody) SetLiveMaterials(v []*AddEditingProjectMaterialsResponseBodyLiveMaterials) *AddEditingProjectMaterialsResponseBody {
	s.LiveMaterials = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetMediaInfos(v []*AddEditingProjectMaterialsResponseBodyMediaInfos) *AddEditingProjectMaterialsResponseBody {
	s.MediaInfos = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetProjectId(v string) *AddEditingProjectMaterialsResponseBody {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetProjectMaterials(v []*string) *AddEditingProjectMaterialsResponseBody {
	s.ProjectMaterials = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetRequestId(v string) *AddEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddEditingProjectMaterialsResponseBodyLiveMaterials struct {
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
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBodyLiveMaterials) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyLiveMaterials) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) GetAppName() *string {
	return s.AppName
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) GetDomainName() *string {
	return s.DomainName
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) GetLiveUrl() *string {
	return s.LiveUrl
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) GetStreamName() *string {
	return s.StreamName
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) SetAppName(v string) *AddEditingProjectMaterialsResponseBodyLiveMaterials {
	s.AppName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) SetDomainName(v string) *AddEditingProjectMaterialsResponseBodyLiveMaterials {
	s.DomainName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) SetLiveUrl(v string) *AddEditingProjectMaterialsResponseBodyLiveMaterials {
	s.LiveUrl = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) SetStreamName(v string) *AddEditingProjectMaterialsResponseBodyLiveMaterials {
	s.StreamName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyLiveMaterials) Validate() error {
	return dara.Validate(s)
}

type AddEditingProjectMaterialsResponseBodyMediaInfos struct {
	// FileInfos
	FileInfoList []*AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information of the media assets.
	MediaBasicInfo *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// *****5cb2e35433198daae94a72*****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) GetFileInfoList() []*AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) GetMediaBasicInfo() *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) SetFileInfoList(v []*AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) *AddEditingProjectMaterialsResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) SetMediaBasicInfo(v *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) *AddEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) SetMediaId(v string) *AddEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) Validate() error {
	return dara.Validate(s)
}

type AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList struct {
	// The basic information of the file, including the duration and size.
	FileBasicInfo *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) Validate() error {
	return dara.Validate(s)
}

type AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
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
	// \\-Uploading
	//
	// \\-Normal
	//
	// \\-UploadFail
	//
	// \\-Disable
	//
	// \\-Deleted
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

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo struct {
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
	// audio
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The thumbnail URL of the media asset.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created.
	//
	// example:
	//
	// 2020-12-23T03:32:59Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the media asset was deleted.
	//
	// example:
	//
	// 2020-12-23T03:32:59Z
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
	// *****5cb2e35433198daae94a72*****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags of the media asset.
	//
	// example:
	//
	// sample_tag
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset.
	//
	// example:
	//
	// Video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time when the media asset was last modified.
	//
	// example:
	//
	// 2020-12-23T03:32:59Z
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
	// The sprite of the media asset.
	//
	// example:
	//
	// http://outin-example.oss-cn-shanghai.aliyuncs.com/test.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The status of the media asset. Valid values:
	//
	// \\- Init
	//
	// \\- Preparing
	//
	// \\- PrepareFail
	//
	// \\- Normal
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the media asset.
	//
	// example:
	//
	// default_title_2020-12-23T03:32:59Z
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding status of the media asset.
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

func (s AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
