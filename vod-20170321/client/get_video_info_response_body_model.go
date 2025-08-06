// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVideoInfoResponseBody
	GetRequestId() *string
	SetVideo(v *GetVideoInfoResponseBodyVideo) *GetVideoInfoResponseBody
	GetVideo() *GetVideoInfoResponseBodyVideo
}

type GetVideoInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the audio or video file.
	Video *GetVideoInfoResponseBodyVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s GetVideoInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoInfoResponseBody) GetVideo() *GetVideoInfoResponseBodyVideo {
	return s.Video
}

func (s *GetVideoInfoResponseBody) SetRequestId(v string) *GetVideoInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoInfoResponseBody) SetVideo(v *GetVideoInfoResponseBodyVideo) *GetVideoInfoResponseBody {
	s.Video = v
	return s
}

func (s *GetVideoInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoInfoResponseBodyVideo struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The final review result of the audio or video file. Valid values:
	//
	// 	- **Normal**: pass
	//
	// 	- **Blocked**: blocked
	//
	// example:
	//
	// Normal
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category ID of the media file.
	//
	// example:
	//
	// 781111****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// Category name
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL of the media file.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The custom information about the media file.\\n\\n> This parameter has been deprecated. This parameter is no longer returned after you call the operation.
	//
	// example:
	//
	// {"aaa":"test"}
	CustomMediaInfo *string `json:"CustomMediaInfo,omitempty" xml:"CustomMediaInfo,omitempty"`
	// The description of the media file.
	//
	// example:
	//
	// Video description in ApsaraVideo VOD
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the offline download feature is enabled. If you enable the offline download feature, users can download and play videos by using the ApsaraVideo Player on a local PC. For more information, see [Configure download settings](https://help.aliyun.com/document_detail/86107.html). Valid values:
	//
	// 	- **on**: the offline download feature is enabled.
	//
	// 	- **off**: the offline download feature is not enabled.
	//
	// example:
	//
	// on
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	// The duration of the media file. Unit: seconds.
	//
	// example:
	//
	// 135.6
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the audio or video file was last updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T10:15:50Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The region where the media file is stored.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The period of time in which the object remains in the restored state.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the audio or video file. Valid values:
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
	// The size of the source file. Unit: bytes.
	//
	// example:
	//
	// 10897890
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The video snapshot URLs.
	Snapshots *GetVideoInfoResponseBodyVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The status of the media file. For more information about the operations that you can perform on files in different statuses and usage limits, see [Status: the status of a video](~~52839#title-vqg-8cz-7p8~~). Valid values:
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
	// The storage class of the audio or video file. Valid values:
	//
	// 	- **Standard**: All media resources are stored as Standard objects.
	//
	// 	- **IA**: All media resources are stored as IA objects.
	//
	// 	- **Archive**: All media resources are stored as Archive objects.
	//
	// 	- **ColdArchive**: All media resources are stored as Cold Archive objects.
	//
	// 	- **SourceIA**: Only the source files are IA objects.
	//
	// 	- **SourceArchive**: Only the source files are Archive objects.
	//
	// 	- **SourceColdArchive**: Only the source files are Cold Archive objects.
	//
	// 	- **Changing**: The storage class of the audio or video file is being changed.
	//
	// 	- **SourceChanging**: The storage class of the source file is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The storage address of the media file.
	//
	// example:
	//
	// out-201703232251****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the audio or video file. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// Tag 1,Tag 2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// 9ae2af636ca64835b0c10412f448****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the media file.
	//
	// example:
	//
	// Video title in ApsaraVideo VOD
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Custom settings. This is a JSON string that supports settings such as message callbacks and upload acceleration. For more information, please refer to [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// 9b73864d75f1d231e9001cd5f8****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoInfoResponseBodyVideo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfoResponseBodyVideo) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponseBodyVideo) GetAppId() *string {
	return s.AppId
}

func (s *GetVideoInfoResponseBodyVideo) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *GetVideoInfoResponseBodyVideo) GetCateId() *int64 {
	return s.CateId
}

func (s *GetVideoInfoResponseBodyVideo) GetCateName() *string {
	return s.CateName
}

func (s *GetVideoInfoResponseBodyVideo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetVideoInfoResponseBodyVideo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetVideoInfoResponseBodyVideo) GetCustomMediaInfo() *string {
	return s.CustomMediaInfo
}

func (s *GetVideoInfoResponseBodyVideo) GetDescription() *string {
	return s.Description
}

func (s *GetVideoInfoResponseBodyVideo) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *GetVideoInfoResponseBodyVideo) GetDuration() *float32 {
	return s.Duration
}

func (s *GetVideoInfoResponseBodyVideo) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetVideoInfoResponseBodyVideo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVideoInfoResponseBodyVideo) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *GetVideoInfoResponseBodyVideo) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *GetVideoInfoResponseBodyVideo) GetSize() *int64 {
	return s.Size
}

func (s *GetVideoInfoResponseBodyVideo) GetSnapshots() *GetVideoInfoResponseBodyVideoSnapshots {
	return s.Snapshots
}

func (s *GetVideoInfoResponseBodyVideo) GetStatus() *string {
	return s.Status
}

func (s *GetVideoInfoResponseBodyVideo) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetVideoInfoResponseBodyVideo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetVideoInfoResponseBodyVideo) GetTags() *string {
	return s.Tags
}

func (s *GetVideoInfoResponseBodyVideo) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *GetVideoInfoResponseBodyVideo) GetTitle() *string {
	return s.Title
}

func (s *GetVideoInfoResponseBodyVideo) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoInfoResponseBodyVideo) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoInfoResponseBodyVideo) SetAppId(v string) *GetVideoInfoResponseBodyVideo {
	s.AppId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetAuditStatus(v string) *GetVideoInfoResponseBodyVideo {
	s.AuditStatus = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCateId(v int64) *GetVideoInfoResponseBodyVideo {
	s.CateId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCateName(v string) *GetVideoInfoResponseBodyVideo {
	s.CateName = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCoverURL(v string) *GetVideoInfoResponseBodyVideo {
	s.CoverURL = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCreationTime(v string) *GetVideoInfoResponseBodyVideo {
	s.CreationTime = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCustomMediaInfo(v string) *GetVideoInfoResponseBodyVideo {
	s.CustomMediaInfo = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetDescription(v string) *GetVideoInfoResponseBodyVideo {
	s.Description = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetDownloadSwitch(v string) *GetVideoInfoResponseBodyVideo {
	s.DownloadSwitch = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetDuration(v float32) *GetVideoInfoResponseBodyVideo {
	s.Duration = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetModificationTime(v string) *GetVideoInfoResponseBodyVideo {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetRegionId(v string) *GetVideoInfoResponseBodyVideo {
	s.RegionId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetRestoreExpiration(v string) *GetVideoInfoResponseBodyVideo {
	s.RestoreExpiration = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetRestoreStatus(v string) *GetVideoInfoResponseBodyVideo {
	s.RestoreStatus = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetSize(v int64) *GetVideoInfoResponseBodyVideo {
	s.Size = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetSnapshots(v *GetVideoInfoResponseBodyVideoSnapshots) *GetVideoInfoResponseBodyVideo {
	s.Snapshots = v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetStatus(v string) *GetVideoInfoResponseBodyVideo {
	s.Status = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetStorageClass(v string) *GetVideoInfoResponseBodyVideo {
	s.StorageClass = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetStorageLocation(v string) *GetVideoInfoResponseBodyVideo {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetTags(v string) *GetVideoInfoResponseBodyVideo {
	s.Tags = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetTemplateGroupId(v string) *GetVideoInfoResponseBodyVideo {
	s.TemplateGroupId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetTitle(v string) *GetVideoInfoResponseBodyVideo {
	s.Title = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetUserData(v string) *GetVideoInfoResponseBodyVideo {
	s.UserData = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetVideoId(v string) *GetVideoInfoResponseBodyVideo {
	s.VideoId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) Validate() error {
	return dara.Validate(s)
}

type GetVideoInfoResponseBodyVideoSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s GetVideoInfoResponseBodyVideoSnapshots) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfoResponseBodyVideoSnapshots) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponseBodyVideoSnapshots) GetSnapshot() []*string {
	return s.Snapshot
}

func (s *GetVideoInfoResponseBodyVideoSnapshots) SetSnapshot(v []*string) *GetVideoInfoResponseBodyVideoSnapshots {
	s.Snapshot = v
	return s
}

func (s *GetVideoInfoResponseBodyVideoSnapshots) Validate() error {
	return dara.Validate(s)
}
