// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUploadProgressResponseBody
	GetRequestId() *string
	SetUploadProgress(v *GetUploadProgressResponseBodyUploadProgress) *GetUploadProgressResponseBody
	GetUploadProgress() *GetUploadProgressResponseBodyUploadProgress
}

type GetUploadProgressResponseBody struct {
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadProgress *GetUploadProgressResponseBodyUploadProgress `json:"UploadProgress,omitempty" xml:"UploadProgress,omitempty" type:"Struct"`
}

func (s GetUploadProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadProgressResponseBody) GetUploadProgress() *GetUploadProgressResponseBodyUploadProgress {
	return s.UploadProgress
}

func (s *GetUploadProgressResponseBody) SetRequestId(v string) *GetUploadProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadProgressResponseBody) SetUploadProgress(v *GetUploadProgressResponseBodyUploadProgress) *GetUploadProgressResponseBody {
	s.UploadProgress = v
	return s
}

func (s *GetUploadProgressResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUploadProgressResponseBodyUploadProgress struct {
	UploadProgressList []*GetUploadProgressResponseBodyUploadProgressUploadProgressList `json:"UploadProgressList,omitempty" xml:"UploadProgressList,omitempty" type:"Repeated"`
}

func (s GetUploadProgressResponseBodyUploadProgress) String() string {
	return dara.Prettify(s)
}

func (s GetUploadProgressResponseBodyUploadProgress) GoString() string {
	return s.String()
}

func (s *GetUploadProgressResponseBodyUploadProgress) GetUploadProgressList() []*GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	return s.UploadProgressList
}

func (s *GetUploadProgressResponseBodyUploadProgress) SetUploadProgressList(v []*GetUploadProgressResponseBodyUploadProgressUploadProgressList) *GetUploadProgressResponseBodyUploadProgress {
	s.UploadProgressList = v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgress) Validate() error {
	return dara.Validate(s)
}

type GetUploadProgressResponseBodyUploadProgressUploadProgressList struct {
	AppVersion     *string  `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AuthInfo       *string  `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	AuthTimestamp  *string  `json:"AuthTimestamp,omitempty" xml:"AuthTimestamp,omitempty"`
	BusinessType   *string  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ClientId       *string  `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DeviceModel    *string  `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DonePartsCount *int32   `json:"DonePartsCount,omitempty" xml:"DonePartsCount,omitempty"`
	FileCreateTime *string  `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	FileHash       *string  `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	FileName       *string  `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize       *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	PartSize       *int64   `json:"PartSize,omitempty" xml:"PartSize,omitempty"`
	Source         *string  `json:"Source,omitempty" xml:"Source,omitempty"`
	TerminalType   *string  `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
	TotalPart      *string  `json:"TotalPart,omitempty" xml:"TotalPart,omitempty"`
	UploadAddress  *string  `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	UploadId       *string  `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	UploadPoint    *string  `json:"UploadPoint,omitempty" xml:"UploadPoint,omitempty"`
	UploadRatio    *float32 `json:"UploadRatio,omitempty" xml:"UploadRatio,omitempty"`
	UploadSpeed    *float32 `json:"UploadSpeed,omitempty" xml:"UploadSpeed,omitempty"`
	VideoId        *string  `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetUploadProgressResponseBodyUploadProgressUploadProgressList) String() string {
	return dara.Prettify(s)
}

func (s GetUploadProgressResponseBodyUploadProgressUploadProgressList) GoString() string {
	return s.String()
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetAuthTimestamp() *string {
	return s.AuthTimestamp
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetClientId() *string {
	return s.ClientId
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetDonePartsCount() *int32 {
	return s.DonePartsCount
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetFileHash() *string {
	return s.FileHash
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetFileName() *string {
	return s.FileName
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetPartSize() *int64 {
	return s.PartSize
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetSource() *string {
	return s.Source
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetTerminalType() *string {
	return s.TerminalType
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetTotalPart() *string {
	return s.TotalPart
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetUploadId() *string {
	return s.UploadId
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetUploadPoint() *string {
	return s.UploadPoint
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetUploadRatio() *float32 {
	return s.UploadRatio
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetUploadSpeed() *float32 {
	return s.UploadSpeed
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) GetVideoId() *string {
	return s.VideoId
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetAppVersion(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.AppVersion = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetAuthInfo(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.AuthInfo = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetAuthTimestamp(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.AuthTimestamp = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetBusinessType(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.BusinessType = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetClientId(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.ClientId = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetDeviceModel(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.DeviceModel = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetDonePartsCount(v int32) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.DonePartsCount = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetFileCreateTime(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.FileCreateTime = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetFileHash(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.FileHash = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetFileName(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.FileName = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetFileSize(v int64) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.FileSize = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetPartSize(v int64) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.PartSize = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetSource(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.Source = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetTerminalType(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.TerminalType = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetTotalPart(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.TotalPart = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetUploadAddress(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.UploadAddress = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetUploadId(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.UploadId = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetUploadPoint(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.UploadPoint = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetUploadRatio(v float32) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.UploadRatio = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetUploadSpeed(v float32) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.UploadSpeed = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) SetVideoId(v string) *GetUploadProgressResponseBodyUploadProgressUploadProgressList {
	s.VideoId = &v
	return s
}

func (s *GetUploadProgressResponseBodyUploadProgressUploadProgressList) Validate() error {
	return dara.Validate(s)
}
