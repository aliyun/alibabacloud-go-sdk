// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody
	GetMediaInfos() []*BatchGetMediaInfosResponseBodyMediaInfos
	SetRequestId(v string) *BatchGetMediaInfosResponseBody
	GetRequestId() *string
}

type BatchGetMediaInfosResponseBody struct {
	// The queried media assets.
	MediaInfos []*BatchGetMediaInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetMediaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBody) GetMediaInfos() []*BatchGetMediaInfosResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *BatchGetMediaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetMediaInfosResponseBody) SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetRequestId(v string) *BatchGetMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBody) Validate() error {
	if s.MediaInfos != nil {
		for _, item := range s.MediaInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfos struct {
	// FileInfos
	FileInfoList []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information of the media asset.
	MediaBasicInfo *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// ******c48fb37407365d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetFileInfoList() []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaBasicInfo() *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetFileInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) Validate() error {
	if s.FileInfoList != nil {
		for _, item := range s.FileInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MediaBasicInfo != nil {
		if err := s.MediaBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoList struct {
	// The basic information of the file, including the duration and size.
	FileBasicInfo *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) Validate() error {
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
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
	// 200
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// example:
	//
	// example
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
	// 1080
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
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo struct {
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The business type of the media asset.
	//
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category of the media asset.
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
	// 2020-12-26T04:11:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the media asset was deleted.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The description of the media asset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the media asset in another service.
	//
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// MediaId
	//
	// example:
	//
	// ******c48fb37407365d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags of the media asset.
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset. Valid values:
	//
	// \\- image
	//
	// \\- video
	//
	// \\- audio
	//
	// \\- text
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time when the media asset was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:12Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The snapshots of the media asset.
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"3","iceJobId":"******f48f0e4154976b2b8c45******","location":"oss-cn-beijing","snapshotRegular":"example.jpg","templateId":"******e6a6440b29eb60bd7c******"}]
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source of the media asset. Valid values:
	//
	// \\- oss
	//
	// \\- vod
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The sprite.
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The status of the media asset.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the media asset.
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
	// userDataTest
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBiz(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
