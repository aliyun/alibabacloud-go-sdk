// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupModels(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListAppInstanceGroupResponseBody
	GetAppInstanceGroupModels() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels
	SetPageNumber(v int32) *ListAppInstanceGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppInstanceGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAppInstanceGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAppInstanceGroupResponseBody
	GetTotalCount() *int32
}

type ListAppInstanceGroupResponseBody struct {
	// The delivery groups.
	AppInstanceGroupModels []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBody) GetAppInstanceGroupModels() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	return s.AppInstanceGroupModels
}

func (s *ListAppInstanceGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppInstanceGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppInstanceGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageNumber(v int32) *ListAppInstanceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageSize(v int32) *ListAppInstanceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetRequestId(v string) *ListAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetTotalCount(v int32) *ListAppInstanceGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The number of subscription resources. Minimum value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The image ID of the app.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The resource type of the delivery group.
	//
	// example:
	//
	// __dynamic__
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-g3k5wa2ms2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// example:
	//
	// false
	AppPolicyImageCheck *bool `json:"AppPolicyImageCheck,omitempty" xml:"AppPolicyImageCheck,omitempty"`
	// example:
	//
	// CENTER
	AppPolicyVersion *string `json:"AppPolicyVersion,omitempty" xml:"AppPolicyVersion,omitempty"`
	// The apps.
	Apps []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- AppInstance: by session
	//
	// 	- Node: by resource
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the delivery group expires.
	//
	// example:
	//
	// 2022-04-27T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the delivery group was created.
	//
	// example:
	//
	// 2022-04-26T15:06:16.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The maximum number of instances. Minimum value: 1.
	//
	// example:
	//
	// 10
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// The minimum number of instances. Minimum value: 1.
	//
	// example:
	//
	// 1
	MinAmount *int32 `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	// The resource groups.
	NodePool []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing+dir-172301****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the operating system.
	//
	// Valid value:
	//
	// 	- Windows
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The information about the over-the-air (OTA) update task.
	OtaInfo *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The percentage of reserved instances. The value indicates the percentage of unused sessions in the delivery group. Valid values: 0 to 99.
	//
	// example:
	//
	// 20
	ReserveAmountRatio *string `json:"ReserveAmountRatio,omitempty" xml:"ReserveAmountRatio,omitempty"`
	// The maximum number of reserved instances. The value indicates the maximum number of unused sessions in the delivery group. Minimum value: 1.
	//
	// example:
	//
	// 5
	ReserveMaxAmount *int32 `json:"ReserveMaxAmount,omitempty" xml:"ReserveMaxAmount,omitempty"`
	// The minimum number of reserved instances. The value indicates the minimum number of unused sessions in the delivery group. Minimum value: 1.
	//
	// example:
	//
	// 1
	ReserveMinAmount *int32 `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	// The resource status.
	//
	// Valid values:
	//
	// 	- AVAILABLE
	//
	// 	- RELEASED
	//
	// 	- EXPIRED_IN_7_DAYS
	//
	// 	- UNAVAILABLE
	//
	// 	- UPGRADING
	//
	// 	- CREATING
	//
	// example:
	//
	// AVAILABLE
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The resource tags.
	ResourceTags []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
	// The duration for which no session is connected. Unit: minutes. If no session is connected in the resources after the specified duration elapses, auto scale-in is triggered. Minimum value: 0.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The number of sessions that are created each time the delivery group is scaled out. Minimum value: 1.
	//
	// example:
	//
	// 10
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper limit of session usage. If the session usage exceeds the specified upper limit, auto scale-out is triggered. The session usage rate is calculated by using the following formula: Session usage rate = Number of sessions in use/Total number of sessions × 100%. Valid values: 0 to 99.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The duration for which sessions are retained after disconnection. Unit: minutes. After an end user disconnects from a session, the session is closed only after the specified duration elapses. If you want to permanently retain sessions, set this parameter to `-1`. Valid values:-1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *string `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// Indicates whether user permission verification is skipped.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
	// The specification ID that uniquely corresponds to the ID of the delivery group.
	//
	// example:
	//
	// spec-8o18t8uc31qib0****
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// The status of the delivery group.
	//
	// Valid values:
	//
	// 	- PUBLISHED: The delivery group is published.
	//
	// 	- FAILED: The delivery group failed to be published.
	//
	// 	- MAINTAIN_FAILED: The delivery group failed to be updated.
	//
	// 	- EXPIRED: The delivery group is expired.
	//
	// 	- MAINTAINING: The delivery group is being updated.
	//
	// 	- CEASED: The delivery group has overdue payments.
	//
	// 	- EXPIRED_RECYCLING: The delivery group is expired and being recycled.
	//
	// 	- DEPLOYING: The delivery group is being published.
	//
	// example:
	//
	// PUBLISHED
	Status *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAccessType() *string {
	return s.AccessType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAmount() *int32 {
	return s.Amount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppPolicyImageCheck() *bool {
	return s.AppPolicyImageCheck
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppPolicyVersion() *string {
	return s.AppPolicyVersion
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetApps() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	return s.Apps
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAuthMode() *string {
	return s.AuthMode
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetChargeResourceMode() *string {
	return s.ChargeResourceMode
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetNodePool() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	return s.NodePool
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetOsType() *string {
	return s.OsType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetOtaInfo() *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	return s.OtaInfo
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetReserveAmountRatio() *string {
	return s.ReserveAmountRatio
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetReserveMaxAmount() *int32 {
	return s.ReserveMaxAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetReserveMinAmount() *int32 {
	return s.ReserveMinAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetResourceTags() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	return s.ResourceTags
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetScalingDownAfterIdleMinutes() *int32 {
	return s.ScalingDownAfterIdleMinutes
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetScalingStep() *int32 {
	return s.ScalingStep
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetScalingUsageThreshold() *string {
	return s.ScalingUsageThreshold
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSessionTimeout() *string {
	return s.SessionTimeout
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSkipUserAuthCheck() *bool {
	return s.SkipUserAuthCheck
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSpecId() *string {
	return s.SpecId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetStatus() *string {
	return s.Status
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetTags() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	return s.Tags
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAccessType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AccessType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyImageCheck(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyImageCheck = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAuthMode(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AuthMode = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeResourceMode(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeResourceMode = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMaxAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MaxAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMinAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MinAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOfficeSiteId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OfficeSiteId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveAmountRatio(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveAmountRatio = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMaxAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMaxAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMinAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMinAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceStatus = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceTags(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceTags = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSkipUserAuthCheck(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetTags(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Tags = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	// The app icon.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The app ID.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The app name.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The app version.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The name of the app version.
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppId() *string {
	return s.AppId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppName() *string {
	return s.AppName
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	// The number of resources purchased when the delivery group was created.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The maximum number of idle sessions. After you specify a value for this parameter, auto scale-out is triggered only if the number of idle sessions in the delivery group is smaller than the specified value and the session usage exceeds the value specified for `ScalingUsageThreshold`. Otherwise, the system determines that idle sessions in the delivery group are sufficient and does not perform auto scale-out.`` You can use this parameter to flexibly manage auto scaling and reduce costs.
	//
	// example:
	//
	// 3
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// The maximum number of resources that can be created for scale-out.
	//
	// example:
	//
	// 8
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// The total number of subscription resources.
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The maximum number of sessions that can be connected to a resource at the same time. If a resource connects to a large number of sessions at the same time, user experience can be compromised. The value range varies based on the resource specification. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The ID of the resource specification that you purchase.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-g6922kced36hx****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// The name of the resource specification.
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
	// The number of subscription resources that are in use.
	//
	// example:
	//
	// 1
	NodeUsed *int32 `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	// The intervals at which the scaling policy is executed.
	RecurrenceSchedules []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// The duration for which no session is connected. Unit: minutes. If no session is connected in the resources after the specified duration elapses, auto scale-in is triggered. Default value: 5.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The total number of scalable resources.
	//
	// example:
	//
	// 8
	ScalingNodeAmount *int32 `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	// The number of scalable resources that are in use.
	//
	// example:
	//
	// 4
	ScalingNodeUsed *int32 `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	// The number of resources that are created each time resources are scaled out. Valid values: 1 to 10.
	//
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper limit of session usage. If the session usage exceeds the specified upper limit, auto scale-out is triggered. The session usage is calculated by using the following formula: `Session usage = Number of current sessions/(Total number of resources × Number of concurrent sessions) × 100%`.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The expiration date of the scaling policy. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// The effective date of the scaling policy. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// The type of the scaling policy.
	//
	// >  `NODE_SCALING_BY_USAGE` is returned for this parameter only if ChargeType is set to `PrePaid`. `NODE_SCALING_BY_SCHEDULE` is returned for this parameter only if ChargeType is set to `PostPaid`.
	//
	// Valid values:
	//
	// 	- NODE_FIXED: No scalable resources are used.
	//
	// 	- NODE_SCALING_BY_SCHEDULE: Scheduled scaling is used.
	//
	// 	- NODE_SCALING_BY_USAGE: Resources are scaled based on usage.
	//
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// Indicates whether resource prefetch is enabled.
	//
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetAmount() *int32 {
	return s.Amount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetMaxIdleAppInstanceAmount() *int32 {
	return s.MaxIdleAppInstanceAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetMaxScalingAmount() *int32 {
	return s.MaxScalingAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeCapacity() *int32 {
	return s.NodeCapacity
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeTypeName() *string {
	return s.NodeTypeName
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeUsed() *int32 {
	return s.NodeUsed
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetRecurrenceSchedules() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	return s.RecurrenceSchedules
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingDownAfterIdleMinutes() *int32 {
	return s.ScalingDownAfterIdleMinutes
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingNodeAmount() *int32 {
	return s.ScalingNodeAmount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingNodeUsed() *int32 {
	return s.ScalingNodeUsed
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingStep() *int32 {
	return s.ScalingStep
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingUsageThreshold() *string {
	return s.ScalingUsageThreshold
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetStrategyDisableDate() *string {
	return s.StrategyDisableDate
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetStrategyEnableDate() *string {
	return s.StrategyEnableDate
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetWarmUp() *bool {
	return s.WarmUp
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxIdleAppInstanceAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeTypeName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeTypeName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	// The schedule type of the scaling policy. This parameter must be configured together with `RecurrenceValues`.``
	//
	// Valid value:
	//
	// 	- weekly: The scaling policy is executed on specific days each week.
	//
	// example:
	//
	// Weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The days of each week on which the scaling policy is executed.
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The time periods during which the scaling policy can be executed.
	TimerPeriods []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GetTimerPeriods() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	return s.TimerPeriods
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	// The number of destination resources.
	//
	// example:
	//
	// 5
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The time when the scaling policy ends. Format: HH:mm.
	//
	// example:
	//
	// 11:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the scaling policy starts. Format: HH:mm.
	//
	// example:
	//
	// 09:30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GetAmount() *int32 {
	return s.Amount
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
	// The new OTA version. A null value indicates that no new version is available.
	//
	// example:
	//
	// 0.0.1-D-20220630.11****
	NewOtaVersion *string `json:"NewOtaVersion,omitempty" xml:"NewOtaVersion,omitempty"`
	// The current OTA version.
	//
	// example:
	//
	// 0.0.1-D-20220615.11****
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// The ID of the OTA update task.
	//
	// example:
	//
	// ota-e49929gv8acz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GetNewOtaVersion() *string {
	return s.NewOtaVersion
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GetOtaVersion() *string {
	return s.OtaVersion
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag type. Valid values: Custom System
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag value.
	//
	// example:
	//
	// design
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) GetScope() *string {
	return s.Scope
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) SetKey(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	s.Key = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) SetScope(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	s.Scope = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) SetValue(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	s.Value = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GetKey() *string {
	return s.Key
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GetScope() *string {
	return s.Scope
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) SetKey(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	s.Key = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) SetScope(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	s.Scope = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) SetValue(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	s.Value = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) Validate() error {
	return dara.Validate(s)
}
