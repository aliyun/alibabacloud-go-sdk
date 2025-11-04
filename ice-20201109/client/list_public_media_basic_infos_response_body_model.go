// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicMediaBasicInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPublicMediaBasicInfosResponseBody
	GetMaxResults() *int32
	SetMediaInfos(v []*ListPublicMediaBasicInfosResponseBodyMediaInfos) *ListPublicMediaBasicInfosResponseBody
	GetMediaInfos() []*ListPublicMediaBasicInfosResponseBodyMediaInfos
	SetNextToken(v string) *ListPublicMediaBasicInfosResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPublicMediaBasicInfosResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPublicMediaBasicInfosResponseBody
	GetTotalCount() *int64
}

type ListPublicMediaBasicInfosResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The media assets that meet the specified conditions.
	MediaInfos []*ListPublicMediaBasicInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
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
	// The total number of media assets that meet the specified conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicMediaBasicInfosResponseBody) GetMediaInfos() []*ListPublicMediaBasicInfosResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *ListPublicMediaBasicInfosResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicMediaBasicInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublicMediaBasicInfosResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPublicMediaBasicInfosResponseBody) SetMaxResults(v int32) *ListPublicMediaBasicInfosResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetMediaInfos(v []*ListPublicMediaBasicInfosResponseBodyMediaInfos) *ListPublicMediaBasicInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetNextToken(v string) *ListPublicMediaBasicInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetRequestId(v string) *ListPublicMediaBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetTotalCount(v int64) *ListPublicMediaBasicInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) Validate() error {
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

type ListPublicMediaBasicInfosResponseBodyMediaInfos struct {
	// The file information of the media asset.
	FileInfoList []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information of the media asset.
	MediaBasicInfo *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****019b82e24b37a1c2958dec38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) GetFileInfoList() []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) GetMediaBasicInfo() *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetFileInfoList(v []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetMediaId(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) Validate() error {
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

type ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList struct {
	// The basic information of the file, such as the duration and size.
	FileBasicInfo *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) Validate() error {
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// The bitrate.
	//
	// example:
	//
	// 270112.12
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration.
	//
	// example:
	//
	// 10.040000
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
	// 338990717
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

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo struct {
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
	// 2021-01-08T16:52:04Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The description of the media asset.
	//
	// example:
	//
	// description
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
	// sticker-daily
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
	// 2021-01-08T16:52:04Z
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
	// userData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
