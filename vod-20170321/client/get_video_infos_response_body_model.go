// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistReferenceIds(v []*string) *GetVideoInfosResponseBody
	GetNonExistReferenceIds() []*string
	SetNonExistVideoIds(v []*string) *GetVideoInfosResponseBody
	GetNonExistVideoIds() []*string
	SetRequestId(v string) *GetVideoInfosResponseBody
	GetRequestId() *string
	SetVideoList(v []*GetVideoInfosResponseBodyVideoList) *GetVideoInfosResponseBody
	GetVideoList() []*GetVideoInfosResponseBodyVideoList
}

type GetVideoInfosResponseBody struct {
	// The list of custom IDs that do not exist.
	NonExistReferenceIds []*string `json:"NonExistReferenceIds,omitempty" xml:"NonExistReferenceIds,omitempty" type:"Repeated"`
	// The list of audio or video IDs that do not exist.
	NonExistVideoIds []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the audio and video files.
	VideoList []*GetVideoInfosResponseBodyVideoList `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Repeated"`
}

func (s GetVideoInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponseBody) GetNonExistReferenceIds() []*string {
	return s.NonExistReferenceIds
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

func (s *GetVideoInfosResponseBody) SetNonExistReferenceIds(v []*string) *GetVideoInfosResponseBody {
	s.NonExistReferenceIds = v
	return s
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
	if s.VideoList != nil {
		for _, item := range s.VideoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoInfosResponseBodyVideoList struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 781111****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// cate1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL of the audio or video file.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the audio or video file was created. The time follows the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-06-26T05:38:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the audio or video file.
	//
	// example:
	//
	// Alibaba Cloud VOD video description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status of the offline download switch. If the offline download feature is enabled, mobile users can cache videos to their local devices for offline viewing by using ApsaraVideo Player. For more information, see [Offline download](https://help.aliyun.com/document_detail/86107.html). Valid values:
	//
	// - **on**: enabled. Offline download is allowed.
	//
	// - **off**: disabled. Offline download is not allowed.
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
	// The last time when the audio or video file was updated. The time follows the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-06-26T06:38:48Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The custom ID. The value can contain only lowercase letters, uppercase letters, digits, hyphens (-), and underscores (_), and must be 6 to 64 characters in length. The value is unique at the user level.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The expiration time of the media asset restoration.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the media asset. Valid values:
	//
	// - **Processing**: The media asset is being restored.
	//
	// - **Success**: The media asset is restored.
	//
	// - **Failed**: The media asset failed to be restored.
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The size of the audio or video source file. Unit: bytes.
	//
	// example:
	//
	// 453
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The array of video snapshot URLs.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The video status. Valid values:
	//
	// - **Uploading**: The video is being uploaded.
	//
	// - **UploadFail**: The video failed to be uploaded.
	//
	// - **UploadSucc**: The video has been uploaded.
	//
	// - **Transcoding**: The video is being transcoded.
	//
	// - **TranscodeFail**: The video failed to be transcoded.
	//
	// - **Blocked**: The video is blocked.
	//
	// - **Normal**: The video is in a normal state.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the media asset. Valid values:
	//
	// - **Standard**: standard.
	//
	// - **IA**: media asset Infrequent Access.
	//
	// - **Archive**: media asset Archive.
	//
	// - **ColdArchive**: media asset Cold Archive.
	//
	// - **SourceIA**: source file Infrequent Access.
	//
	// - **SourceArchive**: source file Archive.
	//
	// - **SourceColdArchive**: source file Cold Archive.
	//
	// - **Changing**: the media asset storage class is being changed.
	//
	// - **SourceChanging**: the source file storage class is being changed.
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
	// The transcoding template group ID.
	//
	// example:
	//
	// b4039216985f4312a5382a4ed****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the audio or video file.
	//
	// example:
	//
	// Alibaba Cloud VOD Video Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks and upload acceleration. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The audio or video ID.
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

func (s *GetVideoInfosResponseBodyVideoList) GetReferenceId() *string {
	return s.ReferenceId
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

func (s *GetVideoInfosResponseBodyVideoList) SetReferenceId(v string) *GetVideoInfosResponseBodyVideoList {
	s.ReferenceId = &v
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
