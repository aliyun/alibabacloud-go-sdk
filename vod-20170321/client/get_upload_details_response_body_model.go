// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenMediaIds(v []*string) *GetUploadDetailsResponseBody
	GetForbiddenMediaIds() []*string
	SetNonExistMediaIds(v []*string) *GetUploadDetailsResponseBody
	GetNonExistMediaIds() []*string
	SetRequestId(v string) *GetUploadDetailsResponseBody
	GetRequestId() *string
	SetUploadDetails(v []*GetUploadDetailsResponseBodyUploadDetails) *GetUploadDetailsResponseBody
	GetUploadDetails() []*GetUploadDetailsResponseBodyUploadDetails
}

type GetUploadDetailsResponseBody struct {
	// The IDs of the media files that cannot be accessed.
	ForbiddenMediaIds []*string `json:"ForbiddenMediaIds,omitempty" xml:"ForbiddenMediaIds,omitempty" type:"Repeated"`
	// The IDs of the media files that do not exist.
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9E290613-04F4-47F4-795D30732077****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload details.
	UploadDetails []*GetUploadDetailsResponseBodyUploadDetails `json:"UploadDetails,omitempty" xml:"UploadDetails,omitempty" type:"Repeated"`
}

func (s GetUploadDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsResponseBody) GetForbiddenMediaIds() []*string {
	return s.ForbiddenMediaIds
}

func (s *GetUploadDetailsResponseBody) GetNonExistMediaIds() []*string {
	return s.NonExistMediaIds
}

func (s *GetUploadDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadDetailsResponseBody) GetUploadDetails() []*GetUploadDetailsResponseBodyUploadDetails {
	return s.UploadDetails
}

func (s *GetUploadDetailsResponseBody) SetForbiddenMediaIds(v []*string) *GetUploadDetailsResponseBody {
	s.ForbiddenMediaIds = v
	return s
}

func (s *GetUploadDetailsResponseBody) SetNonExistMediaIds(v []*string) *GetUploadDetailsResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *GetUploadDetailsResponseBody) SetRequestId(v string) *GetUploadDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadDetailsResponseBody) SetUploadDetails(v []*GetUploadDetailsResponseBodyUploadDetails) *GetUploadDetailsResponseBody {
	s.UploadDetails = v
	return s
}

func (s *GetUploadDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUploadDetailsResponseBodyUploadDetails struct {
	// The time when the upload job was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-04-28T09:45:07Z
	CompletionTime *string `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	// The time when the upload job was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-04-28T09:42:07Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The device model.
	//
	// example:
	//
	// Chrome
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// The size of the uploaded file. Unit: byte.
	//
	// example:
	//
	// 46
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The ID of the uploaded audio or video.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The time when the information about the media file was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-04-28T09:43:12Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The status of the video. For more information about the valid values and value description of the parameter, see the "Status: the status of a video" section of the [Basic structures](https://help.aliyun.com/document_detail/52839.html) topic.
	//
	// example:
	//
	// Uploading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the media file.
	//
	// example:
	//
	// Test details
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The IP address of the server that uploads the media file.
	//
	// example:
	//
	// 192.168.0.1
	UploadIP *string `json:"UploadIP,omitempty" xml:"UploadIP,omitempty"`
	// The upload ratio.
	//
	// example:
	//
	// 0.038
	UploadRatio *float32 `json:"UploadRatio,omitempty" xml:"UploadRatio,omitempty"`
	// The upload size. Unit: byte.
	//
	// example:
	//
	// 346
	UploadSize *int64 `json:"UploadSize,omitempty" xml:"UploadSize,omitempty"`
	// The method that is used to upload the media file.
	//
	// example:
	//
	// WebSDK
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// The status of the upload job. For more information about the valid values and value description of the parameter, see the "Status: the status of a URL-based upload job" section of the [Basic structures](https://help.aliyun.com/document_detail/52839.html) topic.
	//
	// example:
	//
	// Uploading
	UploadStatus *string `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
}

func (s GetUploadDetailsResponseBodyUploadDetails) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDetailsResponseBodyUploadDetails) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetCompletionTime() *string {
	return s.CompletionTime
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetMediaId() *string {
	return s.MediaId
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetStatus() *string {
	return s.Status
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetTitle() *string {
	return s.Title
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetUploadIP() *string {
	return s.UploadIP
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetUploadRatio() *float32 {
	return s.UploadRatio
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetUploadSize() *int64 {
	return s.UploadSize
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetUploadSource() *string {
	return s.UploadSource
}

func (s *GetUploadDetailsResponseBodyUploadDetails) GetUploadStatus() *string {
	return s.UploadStatus
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetCompletionTime(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.CompletionTime = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetCreationTime(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.CreationTime = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetDeviceModel(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.DeviceModel = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetFileSize(v int64) *GetUploadDetailsResponseBodyUploadDetails {
	s.FileSize = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetMediaId(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.MediaId = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetModificationTime(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.ModificationTime = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetStatus(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.Status = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetTitle(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.Title = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadIP(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadIP = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadRatio(v float32) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadRatio = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadSize(v int64) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadSize = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadSource(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadSource = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadStatus(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadStatus = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) Validate() error {
	return dara.Validate(s)
}
