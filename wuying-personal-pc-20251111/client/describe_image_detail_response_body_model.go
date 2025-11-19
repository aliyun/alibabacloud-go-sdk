// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageDetailResponseBody
	GetCode() *string
	SetData(v *DescribeImageDetailResponseBodyData) *DescribeImageDetailResponseBody
	GetData() *DescribeImageDetailResponseBodyData
	SetHttpStatusCode(v int32) *DescribeImageDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeImageDetailResponseBody
	GetRequestId() *string
	SetTraceId(v string) *DescribeImageDetailResponseBody
	GetTraceId() *string
}

type DescribeImageDetailResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeImageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId        *string                              `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeImageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageDetailResponseBody) GetData() *DescribeImageDetailResponseBodyData {
	return s.Data
}

func (s *DescribeImageDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeImageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeImageDetailResponseBody) SetCode(v string) *DescribeImageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageDetailResponseBody) SetData(v *DescribeImageDetailResponseBodyData) *DescribeImageDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageDetailResponseBody) SetHttpStatusCode(v int32) *DescribeImageDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeImageDetailResponseBody) SetRequestId(v string) *DescribeImageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageDetailResponseBody) SetTraceId(v string) *DescribeImageDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeImageDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageDetailResponseBodyData struct {
	CurrentImageCodeInfo     *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo `json:"CurrentImageCodeInfo,omitempty" xml:"CurrentImageCodeInfo,omitempty" type:"Struct"`
	DataDiskSize             *int32                                                   `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description              *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableStartUpConfig      *bool                                                    `json:"EnableStartUpConfig,omitempty" xml:"EnableStartUpConfig,omitempty"`
	GmtCreated               *string                                                  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	ImageId                  *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageScope               *string                                                  `json:"ImageScope,omitempty" xml:"ImageScope,omitempty"`
	IsGpu                    *bool                                                    `json:"IsGpu,omitempty" xml:"IsGpu,omitempty"`
	Name                     *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OsType                   *string                                                  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Permission               *string                                                  `json:"Permission,omitempty" xml:"Permission,omitempty"`
	Platform                 *string                                                  `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Progress                 *string                                                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ShareCodes               []*string                                                `json:"ShareCodes,omitempty" xml:"ShareCodes,omitempty" type:"Repeated"`
	ShareCodesIncludeDisable []*string                                                `json:"ShareCodesIncludeDisable,omitempty" xml:"ShareCodesIncludeDisable,omitempty" type:"Repeated"`
	Shared                   *bool                                                    `json:"Shared,omitempty" xml:"Shared,omitempty"`
	SourceDesktopType        *string                                                  `json:"SourceDesktopType,omitempty" xml:"SourceDesktopType,omitempty"`
	StartUpFileList          []*string                                                `json:"StartUpFileList,omitempty" xml:"StartUpFileList,omitempty" type:"Repeated"`
	Status                   *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDiskSize           *int32                                                   `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeImageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageDetailResponseBodyData) GetCurrentImageCodeInfo() *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	return s.CurrentImageCodeInfo
}

func (s *DescribeImageDetailResponseBodyData) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribeImageDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageDetailResponseBodyData) GetEnableStartUpConfig() *bool {
	return s.EnableStartUpConfig
}

func (s *DescribeImageDetailResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeImageDetailResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageDetailResponseBodyData) GetImageScope() *string {
	return s.ImageScope
}

func (s *DescribeImageDetailResponseBodyData) GetIsGpu() *bool {
	return s.IsGpu
}

func (s *DescribeImageDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeImageDetailResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *DescribeImageDetailResponseBodyData) GetPermission() *string {
	return s.Permission
}

func (s *DescribeImageDetailResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeImageDetailResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *DescribeImageDetailResponseBodyData) GetShareCodes() []*string {
	return s.ShareCodes
}

func (s *DescribeImageDetailResponseBodyData) GetShareCodesIncludeDisable() []*string {
	return s.ShareCodesIncludeDisable
}

func (s *DescribeImageDetailResponseBodyData) GetShared() *bool {
	return s.Shared
}

func (s *DescribeImageDetailResponseBodyData) GetSourceDesktopType() *string {
	return s.SourceDesktopType
}

func (s *DescribeImageDetailResponseBodyData) GetStartUpFileList() []*string {
	return s.StartUpFileList
}

func (s *DescribeImageDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageDetailResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeImageDetailResponseBodyData) SetCurrentImageCodeInfo(v *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) *DescribeImageDetailResponseBodyData {
	s.CurrentImageCodeInfo = v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetDataDiskSize(v int32) *DescribeImageDetailResponseBodyData {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetDescription(v string) *DescribeImageDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetEnableStartUpConfig(v bool) *DescribeImageDetailResponseBodyData {
	s.EnableStartUpConfig = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetGmtCreated(v string) *DescribeImageDetailResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetImageId(v string) *DescribeImageDetailResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetImageScope(v string) *DescribeImageDetailResponseBodyData {
	s.ImageScope = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetIsGpu(v bool) *DescribeImageDetailResponseBodyData {
	s.IsGpu = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetName(v string) *DescribeImageDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetOsType(v string) *DescribeImageDetailResponseBodyData {
	s.OsType = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetPermission(v string) *DescribeImageDetailResponseBodyData {
	s.Permission = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetPlatform(v string) *DescribeImageDetailResponseBodyData {
	s.Platform = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetProgress(v string) *DescribeImageDetailResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetShareCodes(v []*string) *DescribeImageDetailResponseBodyData {
	s.ShareCodes = v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetShareCodesIncludeDisable(v []*string) *DescribeImageDetailResponseBodyData {
	s.ShareCodesIncludeDisable = v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetShared(v bool) *DescribeImageDetailResponseBodyData {
	s.Shared = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetSourceDesktopType(v string) *DescribeImageDetailResponseBodyData {
	s.SourceDesktopType = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetStartUpFileList(v []*string) *DescribeImageDetailResponseBodyData {
	s.StartUpFileList = v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetStatus(v string) *DescribeImageDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) SetSystemDiskSize(v int32) *DescribeImageDetailResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeImageDetailResponseBodyData) Validate() error {
	if s.CurrentImageCodeInfo != nil {
		if err := s.CurrentImageCodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageDetailResponseBodyDataCurrentImageCodeInfo struct {
	CurrentPassword *string `json:"CurrentPassword,omitempty" xml:"CurrentPassword,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreated      *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageCode       *string `json:"ImageCode,omitempty" xml:"ImageCode,omitempty"`
	IsCopyPassword  *bool   `json:"IsCopyPassword,omitempty" xml:"IsCopyPassword,omitempty"`
	IsEncrypted     *bool   `json:"IsEncrypted,omitempty" xml:"IsEncrypted,omitempty"`
	IsFree          *bool   `json:"IsFree,omitempty" xml:"IsFree,omitempty"`
	PeriodDays      *int32  `json:"PeriodDays,omitempty" xml:"PeriodDays,omitempty"`
	RedeemCount     *int32  `json:"RedeemCount,omitempty" xml:"RedeemCount,omitempty"`
	RedeemQuota     *int32  `json:"RedeemQuota,omitempty" xml:"RedeemQuota,omitempty"`
}

func (s DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetCurrentPassword() *string {
	return s.CurrentPassword
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetImageCode() *string {
	return s.ImageCode
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetIsCopyPassword() *bool {
	return s.IsCopyPassword
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetIsEncrypted() *bool {
	return s.IsEncrypted
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetIsFree() *bool {
	return s.IsFree
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetPeriodDays() *int32 {
	return s.PeriodDays
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetRedeemCount() *int32 {
	return s.RedeemCount
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) GetRedeemQuota() *int32 {
	return s.RedeemQuota
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetCurrentPassword(v string) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.CurrentPassword = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetExpireTime(v string) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetGmtCreated(v string) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.GmtCreated = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetGmtModified(v string) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetImageCode(v string) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.ImageCode = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetIsCopyPassword(v bool) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.IsCopyPassword = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetIsEncrypted(v bool) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.IsEncrypted = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetIsFree(v bool) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.IsFree = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetPeriodDays(v int32) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.PeriodDays = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetRedeemCount(v int32) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.RedeemCount = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) SetRedeemQuota(v int32) *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo {
	s.RedeemQuota = &v
	return s
}

func (s *DescribeImageDetailResponseBodyDataCurrentImageCodeInfo) Validate() error {
	return dara.Validate(s)
}
