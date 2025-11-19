// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessibleImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAccessibleImagesResponseBody
	GetCode() *string
	SetData(v []*DescribeAccessibleImagesResponseBodyData) *DescribeAccessibleImagesResponseBody
	GetData() []*DescribeAccessibleImagesResponseBodyData
	SetHttpStatusCode(v int32) *DescribeAccessibleImagesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeAccessibleImagesResponseBody
	GetRequestId() *string
	SetTraceId(v string) *DescribeAccessibleImagesResponseBody
	GetTraceId() *string
}

type DescribeAccessibleImagesResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeAccessibleImagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId        *string                                     `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAccessibleImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessibleImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessibleImagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAccessibleImagesResponseBody) GetData() []*DescribeAccessibleImagesResponseBodyData {
	return s.Data
}

func (s *DescribeAccessibleImagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAccessibleImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessibleImagesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeAccessibleImagesResponseBody) SetCode(v string) *DescribeAccessibleImagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBody) SetData(v []*DescribeAccessibleImagesResponseBodyData) *DescribeAccessibleImagesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccessibleImagesResponseBody) SetHttpStatusCode(v int32) *DescribeAccessibleImagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBody) SetRequestId(v string) *DescribeAccessibleImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBody) SetTraceId(v string) *DescribeAccessibleImagesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessibleImagesResponseBodyData struct {
	ActivityType             *string                                                       `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	AuthTime                 *string                                                       `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	CanUpdate                *bool                                                         `json:"CanUpdate,omitempty" xml:"CanUpdate,omitempty"`
	CurrentImageCodeInfo     *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo `json:"CurrentImageCodeInfo,omitempty" xml:"CurrentImageCodeInfo,omitempty" type:"Struct"`
	DataDiskSize             *int32                                                        `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DeployMode               *string                                                       `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description              *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	DockerImageSize          *int32                                                        `json:"DockerImageSize,omitempty" xml:"DockerImageSize,omitempty"`
	EnableStartUpConfig      *bool                                                         `json:"EnableStartUpConfig,omitempty" xml:"EnableStartUpConfig,omitempty"`
	GmtCreated               *string                                                       `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	ImageId                  *string                                                       `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageScope               *string                                                       `json:"ImageScope,omitempty" xml:"ImageScope,omitempty"`
	ImageSource              *string                                                       `json:"ImageSource,omitempty" xml:"ImageSource,omitempty"`
	ImageType                *string                                                       `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	IsGpu                    *bool                                                         `json:"IsGpu,omitempty" xml:"IsGpu,omitempty"`
	IsLinggou                *string                                                       `json:"IsLinggou,omitempty" xml:"IsLinggou,omitempty"`
	IsWorkstation            *bool                                                         `json:"IsWorkstation,omitempty" xml:"IsWorkstation,omitempty"`
	Name                     *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem          *string                                                       `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	OsType                   *string                                                       `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Permission               *string                                                       `json:"Permission,omitempty" xml:"Permission,omitempty"`
	Platform                 *string                                                       `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Progress                 *string                                                       `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ReceiverType             *string                                                       `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ShareCodes               []*string                                                     `json:"ShareCodes,omitempty" xml:"ShareCodes,omitempty" type:"Repeated"`
	ShareCodesIncludeDisable []*string                                                     `json:"ShareCodesIncludeDisable,omitempty" xml:"ShareCodesIncludeDisable,omitempty" type:"Repeated"`
	Shared                   *bool                                                         `json:"Shared,omitempty" xml:"Shared,omitempty"`
	SharedBy                 *string                                                       `json:"SharedBy,omitempty" xml:"SharedBy,omitempty"`
	SourceDesktopId          *string                                                       `json:"SourceDesktopId,omitempty" xml:"SourceDesktopId,omitempty"`
	SourceDesktopType        *string                                                       `json:"SourceDesktopType,omitempty" xml:"SourceDesktopType,omitempty"`
	SourceImageId            *string                                                       `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
	SourceImageName          *string                                                       `json:"SourceImageName,omitempty" xml:"SourceImageName,omitempty"`
	StartUpFileList          []*string                                                     `json:"StartUpFileList,omitempty" xml:"StartUpFileList,omitempty" type:"Repeated"`
	Status                   *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportPublish           *bool                                                         `json:"SupportPublish,omitempty" xml:"SupportPublish,omitempty"`
	SystemDiskSize           *int32                                                        `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	ValidationCode           *string                                                       `json:"ValidationCode,omitempty" xml:"ValidationCode,omitempty"`
}

func (s DescribeAccessibleImagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessibleImagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccessibleImagesResponseBodyData) GetActivityType() *string {
	return s.ActivityType
}

func (s *DescribeAccessibleImagesResponseBodyData) GetAuthTime() *string {
	return s.AuthTime
}

func (s *DescribeAccessibleImagesResponseBodyData) GetCanUpdate() *bool {
	return s.CanUpdate
}

func (s *DescribeAccessibleImagesResponseBodyData) GetCurrentImageCodeInfo() *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	return s.CurrentImageCodeInfo
}

func (s *DescribeAccessibleImagesResponseBodyData) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribeAccessibleImagesResponseBodyData) GetDeployMode() *string {
	return s.DeployMode
}

func (s *DescribeAccessibleImagesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccessibleImagesResponseBodyData) GetDockerImageSize() *int32 {
	return s.DockerImageSize
}

func (s *DescribeAccessibleImagesResponseBodyData) GetEnableStartUpConfig() *bool {
	return s.EnableStartUpConfig
}

func (s *DescribeAccessibleImagesResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAccessibleImagesResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAccessibleImagesResponseBodyData) GetImageScope() *string {
	return s.ImageScope
}

func (s *DescribeAccessibleImagesResponseBodyData) GetImageSource() *string {
	return s.ImageSource
}

func (s *DescribeAccessibleImagesResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeAccessibleImagesResponseBodyData) GetIsGpu() *bool {
	return s.IsGpu
}

func (s *DescribeAccessibleImagesResponseBodyData) GetIsLinggou() *string {
	return s.IsLinggou
}

func (s *DescribeAccessibleImagesResponseBodyData) GetIsWorkstation() *bool {
	return s.IsWorkstation
}

func (s *DescribeAccessibleImagesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeAccessibleImagesResponseBodyData) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *DescribeAccessibleImagesResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *DescribeAccessibleImagesResponseBodyData) GetPermission() *string {
	return s.Permission
}

func (s *DescribeAccessibleImagesResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeAccessibleImagesResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *DescribeAccessibleImagesResponseBodyData) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *DescribeAccessibleImagesResponseBodyData) GetShareCodes() []*string {
	return s.ShareCodes
}

func (s *DescribeAccessibleImagesResponseBodyData) GetShareCodesIncludeDisable() []*string {
	return s.ShareCodesIncludeDisable
}

func (s *DescribeAccessibleImagesResponseBodyData) GetShared() *bool {
	return s.Shared
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSharedBy() *string {
	return s.SharedBy
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSourceDesktopId() *string {
	return s.SourceDesktopId
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSourceDesktopType() *string {
	return s.SourceDesktopType
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSourceImageId() *string {
	return s.SourceImageId
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSourceImageName() *string {
	return s.SourceImageName
}

func (s *DescribeAccessibleImagesResponseBodyData) GetStartUpFileList() []*string {
	return s.StartUpFileList
}

func (s *DescribeAccessibleImagesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSupportPublish() *bool {
	return s.SupportPublish
}

func (s *DescribeAccessibleImagesResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeAccessibleImagesResponseBodyData) GetValidationCode() *string {
	return s.ValidationCode
}

func (s *DescribeAccessibleImagesResponseBodyData) SetActivityType(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ActivityType = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetAuthTime(v string) *DescribeAccessibleImagesResponseBodyData {
	s.AuthTime = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetCanUpdate(v bool) *DescribeAccessibleImagesResponseBodyData {
	s.CanUpdate = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetCurrentImageCodeInfo(v *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) *DescribeAccessibleImagesResponseBodyData {
	s.CurrentImageCodeInfo = v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetDataDiskSize(v int32) *DescribeAccessibleImagesResponseBodyData {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetDeployMode(v string) *DescribeAccessibleImagesResponseBodyData {
	s.DeployMode = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetDescription(v string) *DescribeAccessibleImagesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetDockerImageSize(v int32) *DescribeAccessibleImagesResponseBodyData {
	s.DockerImageSize = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetEnableStartUpConfig(v bool) *DescribeAccessibleImagesResponseBodyData {
	s.EnableStartUpConfig = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetGmtCreated(v string) *DescribeAccessibleImagesResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetImageId(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetImageScope(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ImageScope = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetImageSource(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ImageSource = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetImageType(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetIsGpu(v bool) *DescribeAccessibleImagesResponseBodyData {
	s.IsGpu = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetIsLinggou(v string) *DescribeAccessibleImagesResponseBodyData {
	s.IsLinggou = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetIsWorkstation(v bool) *DescribeAccessibleImagesResponseBodyData {
	s.IsWorkstation = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetName(v string) *DescribeAccessibleImagesResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetOperationSystem(v string) *DescribeAccessibleImagesResponseBodyData {
	s.OperationSystem = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetOsType(v string) *DescribeAccessibleImagesResponseBodyData {
	s.OsType = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetPermission(v string) *DescribeAccessibleImagesResponseBodyData {
	s.Permission = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetPlatform(v string) *DescribeAccessibleImagesResponseBodyData {
	s.Platform = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetProgress(v string) *DescribeAccessibleImagesResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetReceiverType(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ReceiverType = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetShareCodes(v []*string) *DescribeAccessibleImagesResponseBodyData {
	s.ShareCodes = v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetShareCodesIncludeDisable(v []*string) *DescribeAccessibleImagesResponseBodyData {
	s.ShareCodesIncludeDisable = v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetShared(v bool) *DescribeAccessibleImagesResponseBodyData {
	s.Shared = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSharedBy(v string) *DescribeAccessibleImagesResponseBodyData {
	s.SharedBy = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSourceDesktopId(v string) *DescribeAccessibleImagesResponseBodyData {
	s.SourceDesktopId = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSourceDesktopType(v string) *DescribeAccessibleImagesResponseBodyData {
	s.SourceDesktopType = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSourceImageId(v string) *DescribeAccessibleImagesResponseBodyData {
	s.SourceImageId = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSourceImageName(v string) *DescribeAccessibleImagesResponseBodyData {
	s.SourceImageName = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetStartUpFileList(v []*string) *DescribeAccessibleImagesResponseBodyData {
	s.StartUpFileList = v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetStatus(v string) *DescribeAccessibleImagesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSupportPublish(v bool) *DescribeAccessibleImagesResponseBodyData {
	s.SupportPublish = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetSystemDiskSize(v int32) *DescribeAccessibleImagesResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) SetValidationCode(v string) *DescribeAccessibleImagesResponseBodyData {
	s.ValidationCode = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyData) Validate() error {
	if s.CurrentImageCodeInfo != nil {
		if err := s.CurrentImageCodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo struct {
	CurrentPassword *string `json:"CurrentPassword,omitempty" xml:"CurrentPassword,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreated      *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageCode       *string `json:"ImageCode,omitempty" xml:"ImageCode,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	IsCopyPassword  *bool   `json:"IsCopyPassword,omitempty" xml:"IsCopyPassword,omitempty"`
	IsEncrypted     *bool   `json:"IsEncrypted,omitempty" xml:"IsEncrypted,omitempty"`
	IsFree          *bool   `json:"IsFree,omitempty" xml:"IsFree,omitempty"`
	PeriodDays      *int32  `json:"PeriodDays,omitempty" xml:"PeriodDays,omitempty"`
	RedeemCount     *int32  `json:"RedeemCount,omitempty" xml:"RedeemCount,omitempty"`
	RedeemQuota     *int32  `json:"RedeemQuota,omitempty" xml:"RedeemQuota,omitempty"`
}

func (s DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetCurrentPassword() *string {
	return s.CurrentPassword
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetImageCode() *string {
	return s.ImageCode
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetIsCopyPassword() *bool {
	return s.IsCopyPassword
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetIsEncrypted() *bool {
	return s.IsEncrypted
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetIsFree() *bool {
	return s.IsFree
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetPeriodDays() *int32 {
	return s.PeriodDays
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetRedeemCount() *int32 {
	return s.RedeemCount
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) GetRedeemQuota() *int32 {
	return s.RedeemQuota
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetCurrentPassword(v string) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.CurrentPassword = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetExpireTime(v string) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetGmtCreated(v string) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetGmtModified(v string) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetImageCode(v string) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.ImageCode = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetImageId(v string) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetIsCopyPassword(v bool) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.IsCopyPassword = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetIsEncrypted(v bool) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.IsEncrypted = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetIsFree(v bool) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.IsFree = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetPeriodDays(v int32) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.PeriodDays = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetRedeemCount(v int32) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.RedeemCount = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) SetRedeemQuota(v int32) *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo {
	s.RedeemQuota = &v
	return s
}

func (s *DescribeAccessibleImagesResponseBodyDataCurrentImageCodeInfo) Validate() error {
	return dara.Validate(s)
}
