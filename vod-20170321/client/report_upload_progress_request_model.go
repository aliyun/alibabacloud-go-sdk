// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUploadProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *ReportUploadProgressRequest
	GetAppVersion() *string
	SetAuthInfo(v string) *ReportUploadProgressRequest
	GetAuthInfo() *string
	SetAuthTimestamp(v int64) *ReportUploadProgressRequest
	GetAuthTimestamp() *int64
	SetBusinessType(v string) *ReportUploadProgressRequest
	GetBusinessType() *string
	SetClientId(v string) *ReportUploadProgressRequest
	GetClientId() *string
	SetDeviceModel(v string) *ReportUploadProgressRequest
	GetDeviceModel() *string
	SetDonePartsCount(v int32) *ReportUploadProgressRequest
	GetDonePartsCount() *int32
	SetFileCreateTime(v string) *ReportUploadProgressRequest
	GetFileCreateTime() *string
	SetFileHash(v string) *ReportUploadProgressRequest
	GetFileHash() *string
	SetFileName(v string) *ReportUploadProgressRequest
	GetFileName() *string
	SetFileSize(v int64) *ReportUploadProgressRequest
	GetFileSize() *int64
	SetOwnerId(v int64) *ReportUploadProgressRequest
	GetOwnerId() *int64
	SetPartSize(v int64) *ReportUploadProgressRequest
	GetPartSize() *int64
	SetResourceOwnerAccount(v string) *ReportUploadProgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReportUploadProgressRequest
	GetResourceOwnerId() *int64
	SetSource(v string) *ReportUploadProgressRequest
	GetSource() *string
	SetTerminalType(v string) *ReportUploadProgressRequest
	GetTerminalType() *string
	SetTotalPart(v int32) *ReportUploadProgressRequest
	GetTotalPart() *int32
	SetUploadAddress(v string) *ReportUploadProgressRequest
	GetUploadAddress() *string
	SetUploadId(v string) *ReportUploadProgressRequest
	GetUploadId() *string
	SetUploadPoint(v string) *ReportUploadProgressRequest
	GetUploadPoint() *string
	SetUploadRatio(v float32) *ReportUploadProgressRequest
	GetUploadRatio() *float32
	SetUserId(v int64) *ReportUploadProgressRequest
	GetUserId() *int64
	SetVideoId(v string) *ReportUploadProgressRequest
	GetVideoId() *string
}

type ReportUploadProgressRequest struct {
	// This parameter is required.
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	AuthInfo *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	// This parameter is required.
	AuthTimestamp *int64 `json:"AuthTimestamp,omitempty" xml:"AuthTimestamp,omitempty"`
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	DeviceModel    *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DonePartsCount *int32  `json:"DonePartsCount,omitempty" xml:"DonePartsCount,omitempty"`
	// This parameter is required.
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// This parameter is required.
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PartSize             *int64  `json:"PartSize,omitempty" xml:"PartSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Source               *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	TerminalType  *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
	TotalPart     *int32  `json:"TotalPart,omitempty" xml:"TotalPart,omitempty"`
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// This parameter is required.
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// This parameter is required.
	UploadPoint *string  `json:"UploadPoint,omitempty" xml:"UploadPoint,omitempty"`
	UploadRatio *float32 `json:"UploadRatio,omitempty" xml:"UploadRatio,omitempty"`
	UserId      *int64   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VideoId     *string  `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ReportUploadProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportUploadProgressRequest) GoString() string {
	return s.String()
}

func (s *ReportUploadProgressRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ReportUploadProgressRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *ReportUploadProgressRequest) GetAuthTimestamp() *int64 {
	return s.AuthTimestamp
}

func (s *ReportUploadProgressRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ReportUploadProgressRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ReportUploadProgressRequest) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *ReportUploadProgressRequest) GetDonePartsCount() *int32 {
	return s.DonePartsCount
}

func (s *ReportUploadProgressRequest) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *ReportUploadProgressRequest) GetFileHash() *string {
	return s.FileHash
}

func (s *ReportUploadProgressRequest) GetFileName() *string {
	return s.FileName
}

func (s *ReportUploadProgressRequest) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ReportUploadProgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReportUploadProgressRequest) GetPartSize() *int64 {
	return s.PartSize
}

func (s *ReportUploadProgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReportUploadProgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReportUploadProgressRequest) GetSource() *string {
	return s.Source
}

func (s *ReportUploadProgressRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *ReportUploadProgressRequest) GetTotalPart() *int32 {
	return s.TotalPart
}

func (s *ReportUploadProgressRequest) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *ReportUploadProgressRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *ReportUploadProgressRequest) GetUploadPoint() *string {
	return s.UploadPoint
}

func (s *ReportUploadProgressRequest) GetUploadRatio() *float32 {
	return s.UploadRatio
}

func (s *ReportUploadProgressRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ReportUploadProgressRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *ReportUploadProgressRequest) SetAppVersion(v string) *ReportUploadProgressRequest {
	s.AppVersion = &v
	return s
}

func (s *ReportUploadProgressRequest) SetAuthInfo(v string) *ReportUploadProgressRequest {
	s.AuthInfo = &v
	return s
}

func (s *ReportUploadProgressRequest) SetAuthTimestamp(v int64) *ReportUploadProgressRequest {
	s.AuthTimestamp = &v
	return s
}

func (s *ReportUploadProgressRequest) SetBusinessType(v string) *ReportUploadProgressRequest {
	s.BusinessType = &v
	return s
}

func (s *ReportUploadProgressRequest) SetClientId(v string) *ReportUploadProgressRequest {
	s.ClientId = &v
	return s
}

func (s *ReportUploadProgressRequest) SetDeviceModel(v string) *ReportUploadProgressRequest {
	s.DeviceModel = &v
	return s
}

func (s *ReportUploadProgressRequest) SetDonePartsCount(v int32) *ReportUploadProgressRequest {
	s.DonePartsCount = &v
	return s
}

func (s *ReportUploadProgressRequest) SetFileCreateTime(v string) *ReportUploadProgressRequest {
	s.FileCreateTime = &v
	return s
}

func (s *ReportUploadProgressRequest) SetFileHash(v string) *ReportUploadProgressRequest {
	s.FileHash = &v
	return s
}

func (s *ReportUploadProgressRequest) SetFileName(v string) *ReportUploadProgressRequest {
	s.FileName = &v
	return s
}

func (s *ReportUploadProgressRequest) SetFileSize(v int64) *ReportUploadProgressRequest {
	s.FileSize = &v
	return s
}

func (s *ReportUploadProgressRequest) SetOwnerId(v int64) *ReportUploadProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReportUploadProgressRequest) SetPartSize(v int64) *ReportUploadProgressRequest {
	s.PartSize = &v
	return s
}

func (s *ReportUploadProgressRequest) SetResourceOwnerAccount(v string) *ReportUploadProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReportUploadProgressRequest) SetResourceOwnerId(v int64) *ReportUploadProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReportUploadProgressRequest) SetSource(v string) *ReportUploadProgressRequest {
	s.Source = &v
	return s
}

func (s *ReportUploadProgressRequest) SetTerminalType(v string) *ReportUploadProgressRequest {
	s.TerminalType = &v
	return s
}

func (s *ReportUploadProgressRequest) SetTotalPart(v int32) *ReportUploadProgressRequest {
	s.TotalPart = &v
	return s
}

func (s *ReportUploadProgressRequest) SetUploadAddress(v string) *ReportUploadProgressRequest {
	s.UploadAddress = &v
	return s
}

func (s *ReportUploadProgressRequest) SetUploadId(v string) *ReportUploadProgressRequest {
	s.UploadId = &v
	return s
}

func (s *ReportUploadProgressRequest) SetUploadPoint(v string) *ReportUploadProgressRequest {
	s.UploadPoint = &v
	return s
}

func (s *ReportUploadProgressRequest) SetUploadRatio(v float32) *ReportUploadProgressRequest {
	s.UploadRatio = &v
	return s
}

func (s *ReportUploadProgressRequest) SetUserId(v int64) *ReportUploadProgressRequest {
	s.UserId = &v
	return s
}

func (s *ReportUploadProgressRequest) SetVideoId(v string) *ReportUploadProgressRequest {
	s.VideoId = &v
	return s
}

func (s *ReportUploadProgressRequest) Validate() error {
	return dara.Validate(s)
}
