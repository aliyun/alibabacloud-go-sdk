// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetYikeAssetMediaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredList(v []*string) *BatchGetYikeAssetMediaInfosResponseBody
	GetIgnoredList() []*string
	SetMediaInfos(v []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) *BatchGetYikeAssetMediaInfosResponseBody
	GetMediaInfos() []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfos
	SetRequestId(v string) *BatchGetYikeAssetMediaInfosResponseBody
	GetRequestId() *string
}

type BatchGetYikeAssetMediaInfosResponseBody struct {
	IgnoredList []*string                                            `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	MediaInfos  []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) GetMediaInfos() []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) SetIgnoredList(v []*string) *BatchGetYikeAssetMediaInfosResponseBody {
	s.IgnoredList = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) SetMediaInfos(v []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) *BatchGetYikeAssetMediaInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) SetRequestId(v string) *BatchGetYikeAssetMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBody) Validate() error {
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

type BatchGetYikeAssetMediaInfosResponseBodyMediaInfos struct {
	BizData        *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData        `json:"BizData,omitempty" xml:"BizData,omitempty" type:"Struct"`
	FileInfoList   []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	MediaBasicInfo *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// example:
	//
	// ******c48fb37407365d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) GetBizData() *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	return s.BizData
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) GetFileInfoList() []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) GetMediaBasicInfo() *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) SetBizData(v *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos {
	s.BizData = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) SetFileInfoList(v []*BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) SetMediaId(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfos) Validate() error {
	if s.BizData != nil {
		if err := s.BizData.Validate(); err != nil {
			return err
		}
	}
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

type BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData struct {
	// example:
	//
	// Label
	AuditBlockedLabel *string `json:"AuditBlockedLabel,omitempty" xml:"AuditBlockedLabel,omitempty"`
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// ID
	CreationJobId *string `json:"CreationJobId,omitempty" xml:"CreationJobId,omitempty"`
	// example:
	//
	// pd_0617169475
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// 1
	IsFavorite *string `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// example:
	//
	// 1
	IsLogicalDeleted *string `json:"IsLogicalDeleted,omitempty" xml:"IsLogicalDeleted,omitempty"`
	// example:
	//
	// Image
	MediaAssetSubType *string `json:"MediaAssetSubType,omitempty" xml:"MediaAssetSubType,omitempty"`
	// example:
	//
	// HistoricalUpload
	MediaAssetType *string `json:"MediaAssetType,omitempty" xml:"MediaAssetType,omitempty"`
	// example:
	//
	// pd_0617169475
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// example:
	//
	// f4a26390f02371f0a1f4e6e7c7586706
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// name
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// example:
	//
	// MainBody
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetAuditBlockedLabel() *string {
	return s.AuditBlockedLabel
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetCreationJobId() *string {
	return s.CreationJobId
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetFolderId() *string {
	return s.FolderId
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetIsFavorite() *string {
	return s.IsFavorite
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetIsLogicalDeleted() *string {
	return s.IsLogicalDeleted
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetMediaAssetSubType() *string {
	return s.MediaAssetSubType
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetMediaAssetType() *string {
	return s.MediaAssetType
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetProductionId() *string {
	return s.ProductionId
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetSourceId() *string {
	return s.SourceId
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetSourceName() *string {
	return s.SourceName
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) GetSourceType() *string {
	return s.SourceType
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetAuditBlockedLabel(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.AuditBlockedLabel = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetAuditStatus(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.AuditStatus = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetCreationJobId(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.CreationJobId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetFolderId(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.FolderId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetIsFavorite(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.IsFavorite = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetIsLogicalDeleted(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.IsLogicalDeleted = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetMediaAssetSubType(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.MediaAssetSubType = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetMediaAssetType(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.MediaAssetType = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetProductionId(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.ProductionId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetSourceId(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.SourceId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetSourceName(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.SourceName = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) SetSourceType(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData {
	s.SourceType = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosBizData) Validate() error {
	return dara.Validate(s)
}

type BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList struct {
	FileBasicInfo *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoList) Validate() error {
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// example:
	//
	// 30
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 200
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// example.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 191
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// jpg
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// example:
	//
	// 416
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// example:
	//
	// ICE
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// category
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/01e1271d-ff4f-4689-9c20-e1df81486859_open_live_cover.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2021-01-08T16:52:07Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// example:
	//
	// sample_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// *****64623a94eca8516569c8f*****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// tag1，tag2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2021-01-08T16:52:04Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// []
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// Init
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// example:
	//
	// UserData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBiz(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
