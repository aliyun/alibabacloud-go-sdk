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
	// The IDs of media files that cannot be accessed.
	ForbiddenMediaIds []*string `json:"ForbiddenMediaIds,omitempty" xml:"ForbiddenMediaIds,omitempty" type:"Repeated"`
	// The IDs of media files that do not exist.
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	// The request ID.
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
	if s.UploadDetails != nil {
		for _, item := range s.UploadDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUploadDetailsResponseBodyUploadDetails struct {
	// The completion time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-04-28T09:45:07Z
	CompletionTime *string `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	// The creation time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
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
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 46
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The ID of the uploaded audio or video file.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The modification time. The time is in the _yyyy-MM-ddTHH:mm:ssZ_ format (UTC).
	//
	// example:
	//
	// 2019-04-28T09:43:12Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The video status. For the valid values and descriptions of video statuses, see the value list in [Status: video status](https://help.aliyun.com/document_detail/52839.html).
	//
	// example:
	//
	// Uploading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	//
	// example:
	//
	// Test file upload details
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The upload IP address.
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
	// The upload size. Unit: bytes.
	//
	// example:
	//
	// 346
	UploadSize *int64 `json:"UploadSize,omitempty" xml:"UploadSize,omitempty"`
	// The upload source.
	//
	// example:
	//
	// WebSDK
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// The upload task status. For the valid values and descriptions of upload statuses, see the value list in [Status: URL upload task status](https://help.aliyun.com/document_detail/52839.html).
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
