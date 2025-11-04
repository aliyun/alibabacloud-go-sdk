// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaBasicInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMediaBasicInfosResponseBody
	GetMaxResults() *int32
	SetMediaInfos(v []*ListMediaBasicInfosResponseBodyMediaInfos) *ListMediaBasicInfosResponseBody
	GetMediaInfos() []*ListMediaBasicInfosResponseBodyMediaInfos
	SetNextToken(v string) *ListMediaBasicInfosResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMediaBasicInfosResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListMediaBasicInfosResponseBody
	GetTotalCount() *int64
}

type ListMediaBasicInfosResponseBody struct {
	// The maximum number of entries returned in the query.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The media assets that meet the specified conditions.
	MediaInfos []*ListMediaBasicInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// 8EqYpQbZ6Eh7+Zz8DxVYoQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******B7-7F87-4792-BFE9-63CD21******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMediaBasicInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaBasicInfosResponseBody) GetMediaInfos() []*ListMediaBasicInfosResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *ListMediaBasicInfosResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaBasicInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaBasicInfosResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMediaBasicInfosResponseBody) SetMaxResults(v int32) *ListMediaBasicInfosResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetMediaInfos(v []*ListMediaBasicInfosResponseBodyMediaInfos) *ListMediaBasicInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetNextToken(v string) *ListMediaBasicInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetRequestId(v string) *ListMediaBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetTotalCount(v int64) *ListMediaBasicInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) Validate() error {
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

type ListMediaBasicInfosResponseBodyMediaInfos struct {
	// FileInfos
	FileInfoList []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// BasicInfo
	MediaBasicInfo *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****019b82e24b37a1c2958dec38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) GetFileInfoList() []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) GetMediaBasicInfo() *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetFileInfoList(v []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetMediaId(v string) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) Validate() error {
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

type ListMediaBasicInfosResponseBodyMediaInfosFileInfoList struct {
	// The basic information of the file, including the duration and size.
	FileBasicInfo *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) Validate() error {
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// The bitrate.
	//
	// example:
	//
	// 1912.13
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the file was created.
	//
	// example:
	//
	// 2021-01-08T16:52:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The duration.
	//
	// example:
	//
	// 60.00000
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
	// 14340962
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
	// The OSS URL of the file.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example2.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
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
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 2021-01-08T16:52:07Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
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
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetCreateTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetModifiedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// The service to which the media asset belongs.
	//
	// example:
	//
	// ICE
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The business type of the media asset.
	//
	// example:
	//
	// opening
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 3049
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
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
	// 2021-01-08T16:52:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the media asset was deleted.
	//
	// example:
	//
	// 2021-01-08T16:52:07Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The description of the media asset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the media asset in another service.
	//
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****019b82e24b37a1c2958dec38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags of the media asset.
	//
	// example:
	//
	// tags,tags2
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
	// 2021-01-08T16:52:07Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The custom ID of the media asset. The ID is a string that contains 6 to 64 characters. Only letters, digits, hyphens (-), and underscores (_) are supported. The ID is unique among users.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The snapshots of the media asset.
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"3","iceJobId":"******f48f0e4154976b2b8c45******","location":"oss-cn-beijing","snapshotRegular":"example.jpg","templateId":"******e6a6440b29eb60bd7c******"}]
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source of the media asset.
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
	// The upload source of the media asset.
	//
	// example:
	//
	// general
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// The user data.
	//
	// example:
	//
	// userData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetUploadSource() *string {
	return s.UploadSource
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBiz(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCateId(v int64) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CateId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetReferenceId(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ReferenceId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUploadSource(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UploadSource = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
