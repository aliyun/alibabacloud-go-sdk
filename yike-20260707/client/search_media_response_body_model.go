// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaResponseBody
	GetCode() *string
	SetMediaInfoList(v []*SearchMediaResponseBodyMediaInfoList) *SearchMediaResponseBody
	GetMediaInfoList() []*SearchMediaResponseBodyMediaInfoList
	SetRequestId(v string) *SearchMediaResponseBody
	GetRequestId() *string
	SetScrollToken(v string) *SearchMediaResponseBody
	GetScrollToken() *string
	SetSuccess(v string) *SearchMediaResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaResponseBody
	GetTotal() *int64
}

type SearchMediaResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The collection of media assets that meet the specified criteria.
	MediaInfoList []*SearchMediaResponseBodyMediaInfoList `json:"MediaInfoList,omitempty" xml:"MediaInfoList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of media assets that meet the specified criteria.
	//
	// example:
	//
	// 163
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaResponseBody) GetMediaInfoList() []*SearchMediaResponseBodyMediaInfoList {
	return s.MediaInfoList
}

func (s *SearchMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaResponseBody) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *SearchMediaResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaResponseBody) SetCode(v string) *SearchMediaResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaResponseBody) SetMediaInfoList(v []*SearchMediaResponseBodyMediaInfoList) *SearchMediaResponseBody {
	s.MediaInfoList = v
	return s
}

func (s *SearchMediaResponseBody) SetRequestId(v string) *SearchMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaResponseBody) SetScrollToken(v string) *SearchMediaResponseBody {
	s.ScrollToken = &v
	return s
}

func (s *SearchMediaResponseBody) SetSuccess(v string) *SearchMediaResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaResponseBody) SetTotal(v int64) *SearchMediaResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaResponseBody) Validate() error {
	if s.MediaInfoList != nil {
		for _, item := range s.MediaInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaResponseBodyMediaInfoList struct {
	// The list of media files.
	FileInfoList []*SearchMediaResponseBodyMediaInfoListFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information of the media asset.
	MediaBasicInfo *SearchMediaResponseBodyMediaInfoListMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The dynamic data of the media asset.
	MediaDynamicInfo *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo `json:"MediaDynamicInfo,omitempty" xml:"MediaDynamicInfo,omitempty" type:"Struct"`
	// The media asset ID.
	//
	// example:
	//
	// 3b187b3620c8490886cfc2a9578c****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoList) GetFileInfoList() []*SearchMediaResponseBodyMediaInfoListFileInfoList {
	return s.FileInfoList
}

func (s *SearchMediaResponseBodyMediaInfoList) GetMediaBasicInfo() *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *SearchMediaResponseBodyMediaInfoList) GetMediaDynamicInfo() *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo {
	return s.MediaDynamicInfo
}

func (s *SearchMediaResponseBodyMediaInfoList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaResponseBodyMediaInfoList) SetFileInfoList(v []*SearchMediaResponseBodyMediaInfoListFileInfoList) *SearchMediaResponseBodyMediaInfoList {
	s.FileInfoList = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetMediaBasicInfo(v *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) *SearchMediaResponseBodyMediaInfoList {
	s.MediaBasicInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetMediaDynamicInfo(v *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo) *SearchMediaResponseBodyMediaInfoList {
	s.MediaDynamicInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetMediaId(v string) *SearchMediaResponseBodyMediaInfoList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) Validate() error {
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
	if s.MediaDynamicInfo != nil {
		if err := s.MediaDynamicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaResponseBodyMediaInfoListFileInfoList struct {
	// The basic file information, including duration and size.
	FileBasicInfo *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoList) GetFileBasicInfo() *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoList) SetFileBasicInfo(v *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) *SearchMediaResponseBodyMediaInfoListFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoList) Validate() error {
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo struct {
	// The bitrate.
	//
	// example:
	//
	// 1912.13
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-05-01T19:48Z
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
	// 1642650802***0527050.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 48524
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
	// https://test-bucket-***.oss-cn-shanghai.aliyuncs.com/sv/23d5cdd1-18180984899/23d5cdd1-1818098****.mp4
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
	// 480
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The image set information.
	//
	// example:
	//
	// {}
	ImagesInput *string `json:"ImagesInput,omitempty" xml:"ImagesInput,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-05-01T19:48Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The storage region of the file.
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

func (s SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetImagesInput() *string {
	return s.ImagesInput
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetBitrate(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetCreateTime(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetDuration(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileName(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileSize(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileStatus(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileType(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileUrl(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFormatName(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetHeight(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetImagesInput(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.ImagesInput = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetModifiedTime(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetRegion(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetWidth(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListMediaBasicInfo struct {
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
	// 10
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name.
	//
	// example:
	//
	// test-category-01
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The cover URL.
	//
	// example:
	//
	// https://test-bucket-***.oss-cn-shanghai.aliyuncs.com/cover/e694372e-4f5b-4821-ae09-efd064f2****_large_cover_url.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created.
	//
	// example:
	//
	// 2026-05-01T19:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The content description.
	//
	// example:
	//
	// This is a test video.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity ID.
	//
	// example:
	//
	// BaseMedia
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The media asset URL.
	//
	// example:
	//
	// https://test-bucket-***.oss-cn-shanghai.aliyuncs.com/sv/23d5cdd1-18180984899/23d5cdd1-1818098****.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// 3b187b3620c8490886cfc2a9578c****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags.
	//
	// example:
	//
	// tags,tags2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The media type of the media asset.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time when the media asset was last modified.
	//
	// example:
	//
	// 2026-05-01T20:48Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The snapshots.
	//
	// example:
	//
	// [{"bucket":"test-bucket-***","count":"3","location":"oss-cn-shanghai","snapshotRegular":"example.jpg","templateId":"******e6a6440b29eb60bd7c******"}]
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source of the media asset.
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The sprite images.
	//
	// example:
	//
	// [{"bucket":"test-bucket-***","count":"32","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The resource status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	//
	// example:
	//
	// test-title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding status.
	//
	// example:
	//
	// Init
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// The upload source of the media asset.
	//
	// example:
	//
	// oss
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// The user data.
	//
	// example:
	//
	// userData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCategoryName() *string {
	return s.CategoryName
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetUploadSource() *string {
	return s.UploadSource
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetBusinessType(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCategoryId(v int64) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCategoryName(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CategoryName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCoverURL(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCreateTime(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetDescription(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetEntityId(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.EntityId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetInputURL(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetMediaId(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetMediaTags(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetMediaType(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetModifiedTime(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetSnapshots(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetSource(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetSpriteImages(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetStatus(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetTitle(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetTranscodeStatus(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetUploadSource(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.UploadSource = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetUserData(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListMediaDynamicInfo struct {
	// The dynamic metadata.
	DynamicMetaData *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty" type:"Struct"`
}

func (s SearchMediaResponseBodyMediaInfoListMediaDynamicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListMediaDynamicInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo) GetDynamicMetaData() *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData {
	return s.DynamicMetaData
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo) SetDynamicMetaData(v *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo {
	s.DynamicMetaData = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfo) Validate() error {
	if s.DynamicMetaData != nil {
		if err := s.DynamicMetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData struct {
	// The dynamic metadata content.
	//
	// example:
	//
	// "{\\"ThirdPartyAssetStatus\\":\\"Success\\"}"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The entity ID.
	//
	// example:
	//
	// BaseMedia
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) GetData() *string {
	return s.Data
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) SetData(v string) *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData {
	s.Data = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) SetEntityId(v string) *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData {
	s.EntityId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaDynamicInfoDynamicMetaData) Validate() error {
	return dara.Validate(s)
}
