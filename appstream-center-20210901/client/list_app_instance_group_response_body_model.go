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
	// The delivery group information.
	AppInstanceGroupModels []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	// The page number of the displayed query results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page.
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
	// The total number of query results.
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
	if s.AppInstanceGroupModels != nil {
		for _, item := range s.AppInstanceGroupModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	// The access type.
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The number of subscription resources configured by the user. Minimum value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The application image ID.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The delivery group ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name.
	//
	// example:
	//
	// 办公应用
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The specification type of the delivery group.
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
	// Indicates whether the current image supports the unified policy.
	//
	// example:
	//
	// false
	AppPolicyImageCheck *bool `json:"AppPolicyImageCheck,omitempty" xml:"AppPolicyImageCheck,omitempty"`
	// The policy version.
	//
	// example:
	//
	// CENTER
	AppPolicyVersion *string `json:"AppPolicyVersion,omitempty" xml:"AppPolicyVersion,omitempty"`
	// The application information.
	Apps []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The authorization mode.
	//
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// The sales mode.
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The expiration time of the delivery group.
	//
	// example:
	//
	// 2022-04-27T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The creation time.
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
	// The resource group information.
	NodePool []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	// The office network ID.
	//
	// example:
	//
	// cn-beijing+dir-172301****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The OTA upgrade task information.
	OtaInfo *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	// The product type.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID of the delivery group. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved instance percentage, which is the ratio of unused sessions in the delivery group. Valid values: 0 to 99.
	//
	// example:
	//
	// 20
	ReserveAmountRatio *string `json:"ReserveAmountRatio,omitempty" xml:"ReserveAmountRatio,omitempty"`
	// The maximum number of reserved instances, which is the maximum number of unused sessions in the delivery group. Minimum value: 1.
	//
	// example:
	//
	// 5
	ReserveMaxAmount *int32 `json:"ReserveMaxAmount,omitempty" xml:"ReserveMaxAmount,omitempty"`
	// The minimum number of reserved instances, which is the minimum number of unused sessions in the delivery group. Minimum value: 1.
	//
	// example:
	//
	// 1
	ReserveMinAmount *int32 `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	// The resource status.
	//
	// example:
	//
	// AVAILABLE
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The list of resource tags.
	ResourceTags []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
	// The duration of no session connections, in minutes. When a resource remains in a no-session-connection state for the specified duration, automatic scale-in is triggered. Minimum value: 0.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The number of sessions created per scale-out operation. Minimum value: 1.
	//
	// example:
	//
	// 10
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper threshold of session usage (%). When the session usage exceeds this threshold, automatic scale-out is triggered. The formula for session usage is: session usage = number of sessions in use ÷ total number of sessions × 100%. Valid values: 0 to 99.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The session disconnection retention duration, in minutes. After an end user session is disconnected, the session is retained for the specified duration before being logged off. Set this value to `-1` to retain the session indefinitely. Valid values: -1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *string `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// Specifies whether to skip user authorization verification.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
	// The ID that uniquely corresponds to the delivery group ID.
	//
	// example:
	//
	// spec-8o18t8uc31qib0****
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// The delivery group status.
	//
	// example:
	//
	// PUBLISHED
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportUserGroupMixedAuth *bool   `json:"SupportUserGroupMixedAuth,omitempty" xml:"SupportUserGroupMixedAuth,omitempty"`
	// The list of resource tags.
	Tags []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// Mixed
	UserGroupAuthMode *string `json:"UserGroupAuthMode,omitempty" xml:"UserGroupAuthMode,omitempty"`
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

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSupportUserGroupMixedAuth() *bool {
	return s.SupportUserGroupMixedAuth
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetTags() []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	return s.Tags
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GetUserGroupAuthMode() *string {
	return s.UserGroupAuthMode
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

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSupportUserGroupMixedAuth(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SupportUserGroupMixedAuth = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetTags(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Tags = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetUserGroupAuthMode(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.UserGroupAuthMode = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) Validate() error {
	if s.Apps != nil {
		for _, item := range s.Apps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodePool != nil {
		for _, item := range s.NodePool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OtaInfo != nil {
		if err := s.OtaInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	// The application icon.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The application ID.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// 办公应用
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The application version name.
	//
	// example:
	//
	// 初始版本
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
	// The upper limit of idle sessions. When this value is specified, automatic scale-out is triggered only when the session usage exceeds `ScalingUsageThreshold` and the number of idle sessions in the delivery group is less than `MaxIdleAppInstanceAmount`. Otherwise, the delivery group is considered to have sufficient idle sessions and automatic scale-out is not triggered. This parameter allows flexible control over elastic scaling behavior and helps reduce costs.
	//
	// example:
	//
	// 3
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// The maximum number of resources that can be created during scale-out.
	//
	// example:
	//
	// 8
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// The total number of current subscription resources.
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The number of concurrent sessions, which is the number of sessions that a single resource can handle simultaneously. If too many sessions are connected simultaneously, the application experience may degrade. The valid values vary depending on the resource specification.
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The specification type ID of the purchased resource.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-g6922kced36hx****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// The resource specification name.
	//
	// example:
	//
	// 无影-通用型_4核8G
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
	// The resource count of subscription resources in use.
	//
	// example:
	//
	// 1
	NodeUsed *int32 `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	// The list of policy execution cycles.
	RecurrenceSchedules []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// The duration of no session connections, in minutes. When a resource remains in a no-session-connection state for the specified duration, automatic scale-in is triggered. Default value: 5.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The total number of elastic resources.
	//
	// example:
	//
	// 8
	ScalingNodeAmount *int32 `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	// The resource count of elastic resources in use.
	//
	// example:
	//
	// 4
	ScalingNodeUsed *int32 `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	// The number of resources created per scale-out operation. Valid values: 1 to 10.
	//
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper threshold of session usage (%). When the session usage exceeds this threshold, automatic scale-out is triggered. The formula for session usage is: `session usage = current number of sessions ÷ (total number of resources × concurrent sessions per resource) × 100%`.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The date when the policy expires. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// The date when the policy takes effect. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// The elastic policy type.
	//
	// > `NODE_SCALING_BY_USAGE` (usage-based scaling policy) applies only to `PrePaid` (subscription) resources. `NODE_SCALING_BY_SCHEDULE` (scheduled scaling policy) applies only to `PostPaid` (pay-as-you-go) resources.
	//
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// Indicates whether the resource prefetch policy is enabled.
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
	if s.RecurrenceSchedules != nil {
		for _, item := range s.RecurrenceSchedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	// The type of the policy execution cycle. You must specify both `RecurrenceType` and `RecurrenceValues`.
	//
	// example:
	//
	// Weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The list of values for the policy execution cycle.
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The list of time periods for the policy execution cycle.
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
	if s.TimerPeriods != nil {
		for _, item := range s.TimerPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	// The target resource count.
	//
	// example:
	//
	// 5
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The end time. Format: HH:mm.
	//
	// example:
	//
	// 11:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time. Format: HH:mm.
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
	// The new OTA version. An empty value indicates that no new version is available.
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
	// The OTA upgrade task ID.
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
	// The tag type.
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
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag type.
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
