// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenMediaIds(v []*string) *BatchGetMediaInfosResponseBody
	GetForbiddenMediaIds() []*string
	SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody
	GetMediaInfos() []*BatchGetMediaInfosResponseBodyMediaInfos
	SetNonExistMediaIds(v []*string) *BatchGetMediaInfosResponseBody
	GetNonExistMediaIds() []*string
	SetNonExistReferenceIds(v []*string) *BatchGetMediaInfosResponseBody
	GetNonExistReferenceIds() []*string
	SetRequestId(v string) *BatchGetMediaInfosResponseBody
	GetRequestId() *string
}

type BatchGetMediaInfosResponseBody struct {
	// The IDs of the media assets that do not support the operation typically because you are not authorized to perform the operation. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	ForbiddenMediaIds []*string `json:"ForbiddenMediaIds,omitempty" xml:"ForbiddenMediaIds,omitempty" type:"Repeated"`
	// Details about media assets.
	MediaInfos []*BatchGetMediaInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// The IDs of the media assets that do not exist.
	NonExistMediaIds     []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	NonExistReferenceIds []*string `json:"NonExistReferenceIds,omitempty" xml:"NonExistReferenceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9E290613-04F4-47F4-795D30732077****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetMediaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBody) GetForbiddenMediaIds() []*string {
	return s.ForbiddenMediaIds
}

func (s *BatchGetMediaInfosResponseBody) GetMediaInfos() []*BatchGetMediaInfosResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *BatchGetMediaInfosResponseBody) GetNonExistMediaIds() []*string {
	return s.NonExistMediaIds
}

func (s *BatchGetMediaInfosResponseBody) GetNonExistReferenceIds() []*string {
	return s.NonExistReferenceIds
}

func (s *BatchGetMediaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetMediaInfosResponseBody) SetForbiddenMediaIds(v []*string) *BatchGetMediaInfosResponseBody {
	s.ForbiddenMediaIds = v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetNonExistMediaIds(v []*string) *BatchGetMediaInfosResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetNonExistReferenceIds(v []*string) *BatchGetMediaInfosResponseBody {
	s.NonExistReferenceIds = v
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
	// The ID of the media asset.
	//
	// example:
	//
	// 10a5fa364a5b71ef89246733a78e****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The basic information of the media asset.
	MediaInfo *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// The source file information.
	MezzanineInfo *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo `json:"MezzanineInfo,omitempty" xml:"MezzanineInfo,omitempty" type:"Struct"`
	// The information about the audio or video stream.
	PlayInfoList []*BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList `json:"PlayInfoList,omitempty" xml:"PlayInfoList,omitempty" type:"Repeated"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaInfo() *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	return s.MediaInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMezzanineInfo() *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	return s.MezzanineInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetPlayInfoList() []*BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	return s.PlayInfoList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaInfo(v *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMezzanineInfo(v *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MezzanineInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetPlayInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.PlayInfoList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) Validate() error {
	if s.MediaInfo != nil {
		if err := s.MediaInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MezzanineInfo != nil {
		if err := s.MezzanineInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PlayInfoList != nil {
		for _, item := range s.PlayInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfosMediaInfo struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 781111****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// CateName
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL of the media asset.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the media asset.
	//
	// example:
	//
	// Aliyun VOD Video Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the offline download feature is enabled. If you enable the offline download feature, users can download and play videos by using the ApsaraVideo Player on a local PC. For more information, see [Configure download settings](https://help.aliyun.com/document_detail/86107.html). Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 7753d144efd74d6c45fe0570****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The time when the media asset was last updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-06-26T06:38:48Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The period of time in which the audio file remains in the restored state.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the media asset. Valid values:
	//
	// 	- **Processing**
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The array of video snapshot URLs.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The status of the video. Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **UploadFail**
	//
	// 	- **UploadSucc**
	//
	// 	- **Transcoding**
	//
	// 	- **TranscodeFail**
	//
	// 	- **Blocked**
	//
	// 	- **Normal**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type. Valid values:
	//
	// 	- **Standard**: All media assets are stored as Standard objects.
	//
	// 	- **IA**: All media assets are stored as IA objects.
	//
	// 	- **Archive**: All media assets are stored as Archive objects.
	//
	// 	- **ColdArchive**: All media assets are stored as Cold Archive objects.
	//
	// 	- **SourceIA**: Only the source files are IA objects.
	//
	// 	- **SourceArchive**: Only the source files are Archive objects.
	//
	// 	- **SourceColdArchive**: Only the source file is stored as a Cold Archive object.
	//
	// 	- **Changing**: The storage class of the media asset is being changed.
	//
	// 	- **SourceChanging**: The storage class of the media asset is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The storage address of the media asset.
	//
	// example:
	//
	// outin-***.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the media asset. Separate tags with commas (,).
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// b4039216985f4312a5382a4ed****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the media asset.
	//
	// example:
	//
	// Aliyun VOD Video Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom parameters.
	//
	// example:
	//
	// {"Extend":"xxx","MessageCallback":"xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetAppId() *string {
	return s.AppId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetCateName() *string {
	return s.CateName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetDescription() *string {
	return s.Description
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetStatus() *string {
	return s.Status
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetStorageClass() *string {
	return s.StorageClass
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetTags() *string {
	return s.Tags
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetTitle() *string {
	return s.Title
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) GetUserData() *string {
	return s.UserData
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetAppId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.AppId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetCateId(v int64) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.CateId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetCateName(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.CateName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetCoverURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.CoverURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetCreationTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.CreationTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetDescription(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.Description = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetDownloadSwitch(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.DownloadSwitch = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetModificationTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.ModificationTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetReferenceId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.ReferenceId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetRestoreExpiration(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.RestoreExpiration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetRestoreStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.RestoreStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetSnapshots(v []*string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.Snapshots = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.Status = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetStorageClass(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.StorageClass = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetStorageLocation(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.StorageLocation = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetTags(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.Tags = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetTemplateGroupId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.TemplateGroupId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetTitle(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.Title = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) SetUserData(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo {
	s.UserData = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo struct {
	// The information about the audio stream.
	AudioStreamList []*BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Repeated"`
	// The bitrate of the file. Unit: Kbit/s.
	//
	// example:
	//
	// 771.2280
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the source file was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The duration of the file. Unit: seconds.
	//
	// example:
	//
	// 42.4930
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileMD5  *string `json:"FileMD5,omitempty" xml:"FileMD5,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 27ffc438-164h67f57ef-0005-6884-51a-1****.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS URL of the source file.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-shanghai.aliyuncs.com/27ffc438-164h67f57ef-0005-6884-51a-1****.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The frame rate of the file.
	//
	// example:
	//
	// 25.0000
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the file. Unit: pixels.
	//
	// example:
	//
	// 540
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 1f1a6fc03ca04814031b8a6559e****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 4096477
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The state of the file. Valid values:
	//
	// 	- **Uploading**: The file is being uploaded. This is the initial status.
	//
	// 	- **Normal**: The file is uploaded.
	//
	// 	- **UploadFail**: The file failed to be uploaded.
	//
	// 	- **Deleted**: The file is deleted.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the video streams.
	VideoStreamList []*BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Repeated"`
	// The width of the file. Unit: pixels.
	//
	// example:
	//
	// 960
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetAudioStreamList() []*BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	return s.AudioStreamList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetFileMD5() *string {
	return s.FileMD5
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetFileName() *string {
	return s.FileName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetFileURL() *string {
	return s.FileURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetHeight() *int64 {
	return s.Height
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetSize() *int64 {
	return s.Size
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetStatus() *string {
	return s.Status
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetVideoStreamList() []*BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	return s.VideoStreamList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) GetWidth() *int64 {
	return s.Width
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetAudioStreamList(v []*BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.AudioStreamList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetCreationTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.CreationTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetFileMD5(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.FileMD5 = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetFileName(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.FileName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetFileURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.FileURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetFps(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Fps = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetHeight(v int64) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Height = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetSize(v int64) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Size = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Status = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetVideoStreamList(v []*BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.VideoStreamList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) SetWidth(v int64) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfo) Validate() error {
	if s.AudioStreamList != nil {
		for _, item := range s.AudioStreamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoStreamList != nil {
		for _, item := range s.VideoStreamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList struct {
	// The bitrate.
	//
	// example:
	//
	// 62.885
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The output layout of the audio channels. Valid values:
	//
	// 	- **mono**
	//
	// 	- **stereo**
	//
	// example:
	//
	// mono
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of sound tracks.
	//
	// example:
	//
	// 1
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// mp4a
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the encoder.
	//
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration.
	//
	// example:
	//
	// 3.227574
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 1
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The sampling format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The start time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 0.000000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetChannels() *string {
	return s.Channels
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetChannelLayout(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.ChannelLayout = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetChannels(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.Channels = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetCodecLongName(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetCodecName(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetCodecTag(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetCodecTagString(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetCodecTimeBase(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetIndex(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.Index = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetLang(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetNumFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.NumFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetSampleFmt(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.SampleFmt = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetSampleRate(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.SampleRate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetStartTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) SetTimebase(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoAudioStreamList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList struct {
	// The average frame rate.
	//
	// example:
	//
	// 30.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate of the file. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format.
	//
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x31637661
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// avc1
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the encoder.
	//
	// example:
	//
	// 1/60
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR).
	//
	// example:
	//
	// 0:1
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration.
	//
	// example:
	//
	// 3.166667
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate of the output file.
	//
	// example:
	//
	// 30.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The HDR type of the video stream.
	//
	// example:
	//
	// HDR
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// Indicates whether the video stream contains bidirectional frames (B-frames).
	//
	// example:
	//
	// 0
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The height of the video stream.
	//
	// example:
	//
	// 320
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video stream. The value identifies the position of the video stream in all video streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The codec level.
	//
	// example:
	//
	// 30
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 0
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format of the video stream.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The codec profile.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The rotation angle of the video. Valid values: [0,360).
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	//
	// example:
	//
	// 0:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 0.000000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The horizontal resolution of the video.
	//
	// example:
	//
	// 568
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetDar() *string {
	return s.Dar
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetHDRType() *string {
	return s.HDRType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetHeight() *string {
	return s.Height
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetLevel() *string {
	return s.Level
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetProfile() *string {
	return s.Profile
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetRotate() *string {
	return s.Rotate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetSar() *string {
	return s.Sar
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) GetWidth() *string {
	return s.Width
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetAvgFPS(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.AvgFPS = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetCodecLongName(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetCodecName(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetCodecTag(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetCodecTagString(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetCodecTimeBase(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetDar(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Dar = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetFps(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Fps = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetHDRType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.HDRType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetHasBFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.HasBFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetHeight(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Height = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetIndex(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Index = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetLang(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetLevel(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Level = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetNumFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.NumFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetPixFmt(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.PixFmt = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetProfile(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Profile = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetRotate(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Rotate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetSar(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Sar = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetStartTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetTimebase(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) SetWidth(v string) *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMezzanineInfoVideoStreamList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList struct {
	// The color depth. This value is an integer.
	//
	// example:
	//
	// 8
	BitDepth *int32 `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	// The bitrate of the media stream. Unit: Kbit/s.
	//
	// example:
	//
	// 450.878
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The short name of the codec.
	//
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The creation time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-18T07:37:15Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The quality of the video stream. Valid values:
	//
	// 	- **FD**: low definition
	//
	// 	- **LD**: standard definition
	//
	// 	- **SD**: high definition
	//
	// 	- **HD**: ultra-high definition
	//
	// 	- **OD**: original definition
	//
	// 	- **2K**
	//
	// 	- **4K**
	//
	// 	- **SQ**: standard sound quality
	//
	// 	- **HQ**: high sound quality
	//
	// 	- **AUTO**: adaptive bitrate
	//
	// example:
	//
	// LD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The duration of the media stream. Unit: seconds.
	//
	// example:
	//
	// 9.0464
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether the media stream was encrypted. Valid values:
	//
	// 	- **0**: The media stream is not encrypted.
	//
	// 	- **1**: The media stream is encrypted.
	//
	// example:
	//
	// 1
	Encrypt *int64 `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// The encryption type of the media stream. Valid values:
	//
	// 	- **License**: decryption on local devices.
	//
	// >  If the encryption type is **License**, only ApsaraVideo Player SDK can be used to play videos.
	//
	// example:
	//
	// License
	EncryptMode *string `json:"EncryptMode,omitempty" xml:"EncryptMode,omitempty"`
	// The encryption type of the media stream. Valid values:
	//
	// 	- **AliyunVoDEncryption**: Alibaba Cloud proprietary cryptography
	//
	// 	- **HLSEncryption**: HTTP Live Streaming (HLS) encryption
	//
	// >  If the encryption type is AliyunVoDEncryption, only ApsaraVideo Player SDK can be used to play videos.
	//
	// example:
	//
	// AliyunVoDEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The format of the media stream.
	//
	// 	- If the media file is a video file, the valid values are **mp4*	- and **m3u8**.
	//
	// 	- If the media asset is an audio-only file, the value is **mp3**.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the media stream. Unit: frames per second (FPS).
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The HDR type of the media stream. Valid values:
	//
	// 	- HDR
	//
	// 	- HDR10
	//
	// 	- HLG
	//
	// 	- DolbyVision
	//
	// 	- HDRVivid
	//
	// 	- SDR+
	//
	// example:
	//
	// HLG
	HDRType *string `json:"HDRType,omitempty" xml:"HDRType,omitempty"`
	// The height of the media stream. Unit: pixels.
	//
	// example:
	//
	// 640
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The custom watermark information of the copyright watermark. This parameter is returned if you set `JobType` to `2`.
	//
	// example:
	//
	// CopyrightMarkTest
	JobExt *string `json:"JobExt,omitempty" xml:"JobExt,omitempty"`
	// The job ID for transcoding the media stream. This ID uniquely identifies a media stream.
	//
	// example:
	//
	// 80e9c6580e754a798c3c19c59b16****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the digital watermark. Valid values:
	//
	// 	- **1**: user-tracing watermark
	//
	// 	- **2**: copyright watermark
	//
	// example:
	//
	// 2
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The update time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-20T06:32:19Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The transcoding type. Valid values:
	//
	// 	- **0**: regular transcoding
	//
	// 	- **1.0**: Narrowband HD™ 1.0 transcoding
	//
	// 	- **2.0**: Narrowband HD™ 2.0 transcoding
	//
	// example:
	//
	// 0
	NarrowBandType *string `json:"NarrowBandType,omitempty" xml:"NarrowBandType,omitempty"`
	// The playback URL of the video stream.
	PlayURL *string `json:"PlayURL,omitempty" xml:"PlayURL,omitempty"`
	// The size of the media stream. Unit: bytes.
	//
	// example:
	//
	// 418112
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The specifications of transcoded audio and video streams. For more information about the valid values, see [Output specifications](~~124671#section-6bv-l0g-opq~~).
	//
	// example:
	//
	// H264.LD
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The status of the audio or video stream. Valid values:
	//
	// 	- **Normal**: The latest transcoded stream in each quality and format is in the Normal status.
	//
	// 	- **Invisible**: If multiple streams are transcoded in the same quality and format, the latest transcoded stream is in the Normal status and other streams are in the Invisible status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the media stream. If the media stream is a video stream, the value is **video**. If the media stream is an audio-only stream, the value is **audio**.
	//
	// example:
	//
	// video
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// fb0716154b21a4ecb5b70a26ccc8****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The ID of the transcoding template.
	//
	// example:
	//
	// a86a4338dd2e83da45154004a541****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the watermark that is associated with the media stream.
	//
	// example:
	//
	// dgfn26457856****
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	// The width of the media stream. Unit: pixels.
	//
	// example:
	//
	// 360
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetBitDepth() *int32 {
	return s.BitDepth
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetDefinition() *string {
	return s.Definition
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetEncrypt() *int64 {
	return s.Encrypt
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetEncryptMode() *string {
	return s.EncryptMode
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetEncryptType() *string {
	return s.EncryptType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetFormat() *string {
	return s.Format
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetHDRType() *string {
	return s.HDRType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetHeight() *int64 {
	return s.Height
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetJobExt() *string {
	return s.JobExt
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetJobId() *string {
	return s.JobId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetJobType() *int32 {
	return s.JobType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetNarrowBandType() *string {
	return s.NarrowBandType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetPlayURL() *string {
	return s.PlayURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetSize() *int64 {
	return s.Size
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetSpecification() *string {
	return s.Specification
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetStatus() *string {
	return s.Status
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetStreamType() *string {
	return s.StreamType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) GetWidth() *int64 {
	return s.Width
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetBitDepth(v int32) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.BitDepth = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetCodecName(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetCreationTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.CreationTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetDefinition(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Definition = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetEncrypt(v int64) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Encrypt = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetEncryptMode(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.EncryptMode = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetEncryptType(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.EncryptType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetFormat(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Format = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetFps(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Fps = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetHDRType(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.HDRType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetHeight(v int64) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Height = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetJobExt(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.JobExt = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetJobId(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.JobId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetJobType(v int32) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.JobType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetModificationTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.ModificationTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetNarrowBandType(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.NarrowBandType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetPlayURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.PlayURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetSize(v int64) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Size = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetSpecification(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Specification = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Status = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetStreamType(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.StreamType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetTemplateGroupId(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.TemplateGroupId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetTemplateId(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.TemplateId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetWatermarkId(v string) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.WatermarkId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) SetWidth(v int64) *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosPlayInfoList) Validate() error {
	return dara.Validate(s)
}
