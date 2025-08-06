// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistVideoIds(v []*string) *GetVideoInfosResponseBody
	GetNonExistVideoIds() []*string
	SetRequestId(v string) *GetVideoInfosResponseBody
	GetRequestId() *string
	SetVideoList(v []*GetVideoInfosResponseBodyVideoList) *GetVideoInfosResponseBody
	GetVideoList() []*GetVideoInfosResponseBodyVideoList
}

type GetVideoInfosResponseBody struct {
	// The IDs of the videos that do not exist.
	NonExistVideoIds []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the audio or video files.
	VideoList []*GetVideoInfosResponseBodyVideoList `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Repeated"`
}

func (s GetVideoInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponseBody) GetNonExistVideoIds() []*string {
	return s.NonExistVideoIds
}

func (s *GetVideoInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoInfosResponseBody) GetVideoList() []*GetVideoInfosResponseBodyVideoList {
	return s.VideoList
}

func (s *GetVideoInfosResponseBody) SetNonExistVideoIds(v []*string) *GetVideoInfosResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *GetVideoInfosResponseBody) SetRequestId(v string) *GetVideoInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoInfosResponseBody) SetVideoList(v []*GetVideoInfosResponseBodyVideoList) *GetVideoInfosResponseBody {
	s.VideoList = v
	return s
}

func (s *GetVideoInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoInfosResponseBodyVideoList struct {
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
	// This is a category name.
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL of the audio or video file.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media file was created. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-06-26T05:38:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the audio or video file.
	//
	// example:
	//
	// This is a category description.
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
	// The duration of the audio or video file. Unit: seconds.
	//
	// example:
	//
	// 120
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the audio or video file was last updated. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-06-26T06:38:48Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The period of time in which the audio file remains in the restored state.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the audio file. Valid values:
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
	// 453
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The video snapshot URLs.
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
	// The storage class of the audio file. Valid values:
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
	// 	- **SourceColdArchive**: Only the source file is stored as a Cold Archive object.
	//
	// 	- **Changing**: The storage class of the video file is being changed.
	//
	// 	- **SourceChanging**: The storage class of the source file is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The storage address of the audio or video file.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the audio or video file. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// tag1, tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// b4039216985f4312a5382a4ed****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the audio or video file.
	//
	// example:
	//
	// Video tiltle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Custom settings. This is a JSON string that supports settings such as message callbacks and upload acceleration. For more information, please refer to [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the audio or video file.
	//
	// example:
	//
	// 7753d144efd74d6c45fe0570****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoInfosResponseBodyVideoList) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfosResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponseBodyVideoList) GetAppId() *string {
	return s.AppId
}

func (s *GetVideoInfosResponseBodyVideoList) GetCateId() *int64 {
	return s.CateId
}

func (s *GetVideoInfosResponseBodyVideoList) GetCateName() *string {
	return s.CateName
}

func (s *GetVideoInfosResponseBodyVideoList) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetVideoInfosResponseBodyVideoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetVideoInfosResponseBodyVideoList) GetDescription() *string {
	return s.Description
}

func (s *GetVideoInfosResponseBodyVideoList) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *GetVideoInfosResponseBodyVideoList) GetDuration() *float32 {
	return s.Duration
}

func (s *GetVideoInfosResponseBodyVideoList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetVideoInfosResponseBodyVideoList) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *GetVideoInfosResponseBodyVideoList) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *GetVideoInfosResponseBodyVideoList) GetSize() *int64 {
	return s.Size
}

func (s *GetVideoInfosResponseBodyVideoList) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *GetVideoInfosResponseBodyVideoList) GetStatus() *string {
	return s.Status
}

func (s *GetVideoInfosResponseBodyVideoList) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetVideoInfosResponseBodyVideoList) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetVideoInfosResponseBodyVideoList) GetTags() *string {
	return s.Tags
}

func (s *GetVideoInfosResponseBodyVideoList) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *GetVideoInfosResponseBodyVideoList) GetTitle() *string {
	return s.Title
}

func (s *GetVideoInfosResponseBodyVideoList) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoInfosResponseBodyVideoList) GetVideoId() *string {
	return s.VideoId
}

func (s *GetVideoInfosResponseBodyVideoList) SetAppId(v string) *GetVideoInfosResponseBodyVideoList {
	s.AppId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCateId(v int64) *GetVideoInfosResponseBodyVideoList {
	s.CateId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCateName(v string) *GetVideoInfosResponseBodyVideoList {
	s.CateName = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCoverURL(v string) *GetVideoInfosResponseBodyVideoList {
	s.CoverURL = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCreationTime(v string) *GetVideoInfosResponseBodyVideoList {
	s.CreationTime = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetDescription(v string) *GetVideoInfosResponseBodyVideoList {
	s.Description = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetDownloadSwitch(v string) *GetVideoInfosResponseBodyVideoList {
	s.DownloadSwitch = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetDuration(v float32) *GetVideoInfosResponseBodyVideoList {
	s.Duration = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetModificationTime(v string) *GetVideoInfosResponseBodyVideoList {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetRestoreExpiration(v string) *GetVideoInfosResponseBodyVideoList {
	s.RestoreExpiration = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetRestoreStatus(v string) *GetVideoInfosResponseBodyVideoList {
	s.RestoreStatus = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetSize(v int64) *GetVideoInfosResponseBodyVideoList {
	s.Size = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetSnapshots(v []*string) *GetVideoInfosResponseBodyVideoList {
	s.Snapshots = v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetStatus(v string) *GetVideoInfosResponseBodyVideoList {
	s.Status = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetStorageClass(v string) *GetVideoInfosResponseBodyVideoList {
	s.StorageClass = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetStorageLocation(v string) *GetVideoInfosResponseBodyVideoList {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetTags(v string) *GetVideoInfosResponseBodyVideoList {
	s.Tags = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetTemplateGroupId(v string) *GetVideoInfosResponseBodyVideoList {
	s.TemplateGroupId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetTitle(v string) *GetVideoInfosResponseBodyVideoList {
	s.Title = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetUserData(v string) *GetVideoInfosResponseBodyVideoList {
	s.UserData = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetVideoId(v string) *GetVideoInfosResponseBodyVideoList {
	s.VideoId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) Validate() error {
	return dara.Validate(s)
}
