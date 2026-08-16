// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupModels(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) *GetAppInstanceGroupResponseBody
	GetAppInstanceGroupModels() *GetAppInstanceGroupResponseBodyAppInstanceGroupModels
	SetRequestId(v string) *GetAppInstanceGroupResponseBody
	GetRequestId() *string
}

type GetAppInstanceGroupResponseBody struct {
	// AppInstanceGroupModels
	AppInstanceGroupModels *GetAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBody) GetAppInstanceGroupModels() *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	return s.AppInstanceGroupModels
}

func (s *GetAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) *GetAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *GetAppInstanceGroupResponseBody) SetRequestId(v string) *GetAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBody) Validate() error {
	if s.AppInstanceGroupModels != nil {
		if err := s.AppInstanceGroupModels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
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
	// The application image name.
	//
	// example:
	//
	// OfficeApps
	AppCenterImageName *string `json:"AppCenterImageName,omitempty" xml:"AppCenterImageName,omitempty"`
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
	// The instance type of the delivery group.
	//
	// example:
	//
	// __dynamic__
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The name of the instance type of the delivery group.
	//
	// example:
	//
	// test001
	AppInstanceTypeName *string `json:"AppInstanceTypeName,omitempty" xml:"AppInstanceTypeName,omitempty"`
	// The application policy ID.
	//
	// example:
	//
	// pg-g3k5wa2ms2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// The list of application information.
	Apps []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
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
	NodePool []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
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
	OtaInfo *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
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
	// The percentage of reserved instances, which is the ratio of unused sessions in the delivery group. Valid values: 0 to 99.
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
	// The idle duration without session connections, in minutes. When a resource remains without session connections for the specified duration, automatic scale-in is triggered. Minimum value: 0.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The number of sessions created during each scale-out operation. Minimum value: 1.
	//
	// example:
	//
	// 10
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper threshold of session usage (%). When the session usage exceeds this threshold, automatic scale-out is triggered. The formula for session usage is: Session usage = Number of sessions in use ÷ Total number of sessions × 100%. Valid values: 0 to 99.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The session disconnection retention duration, in minutes. After a session is disconnected from the end user, the session is retained for the specified duration before being logged off. Set this parameter to `-1` to retain the session indefinitely. Valid values: -1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *string `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The session type.
	//
	// example:
	//
	// NORMAL
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
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
	Tags []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// Mixed
	UserGroupAuthMode *string `json:"UserGroupAuthMode,omitempty" xml:"UserGroupAuthMode,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAccessType() *string {
	return s.AccessType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAmount() *int32 {
	return s.Amount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppCenterImageName() *string {
	return s.AppCenterImageName
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceTypeName() *string {
	return s.AppInstanceTypeName
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetApps() []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	return s.Apps
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAuthMode() *string {
	return s.AuthMode
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetChargeResourceMode() *string {
	return s.ChargeResourceMode
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetNodePool() []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	return s.NodePool
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetOsType() *string {
	return s.OsType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetOtaInfo() *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	return s.OtaInfo
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetProductType() *string {
	return s.ProductType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetReserveAmountRatio() *string {
	return s.ReserveAmountRatio
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetReserveMaxAmount() *int32 {
	return s.ReserveMaxAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetReserveMinAmount() *int32 {
	return s.ReserveMinAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetScalingDownAfterIdleMinutes() *int32 {
	return s.ScalingDownAfterIdleMinutes
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetScalingStep() *int32 {
	return s.ScalingStep
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetScalingUsageThreshold() *string {
	return s.ScalingUsageThreshold
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSessionTimeout() *string {
	return s.SessionTimeout
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSessionType() *string {
	return s.SessionType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSkipUserAuthCheck() *bool {
	return s.SkipUserAuthCheck
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSpecId() *string {
	return s.SpecId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetStatus() *string {
	return s.Status
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetSupportUserGroupMixedAuth() *bool {
	return s.SupportUserGroupMixedAuth
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetTags() []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	return s.Tags
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GetUserGroupAuthMode() *string {
	return s.UserGroupAuthMode
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAccessType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AccessType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceTypeName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceTypeName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAuthMode(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AuthMode = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeResourceMode(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeResourceMode = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMaxAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MaxAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMinAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MinAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOfficeSiteId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OfficeSiteId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveAmountRatio(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveAmountRatio = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMaxAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMaxAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMinAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMinAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceStatus(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceStatus = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingDownAfterIdleMinutes(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingStep(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingStep = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingUsageThreshold(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSkipUserAuthCheck(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSupportUserGroupMixedAuth(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SupportUserGroupMixedAuth = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetTags(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Tags = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetUserGroupAuthMode(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.UserGroupAuthMode = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) Validate() error {
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

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
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

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppId() *string {
	return s.AppId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	// The number of resources purchased when the delivery group was created.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The upper limit of idle sessions. When this value is specified, automatic scale-out is triggered only when the session usage exceeds `ScalingUsageThreshold` and the number of idle sessions in the delivery group is less than `MaxIdleAppInstanceAmount`. Otherwise, the idle sessions are considered sufficient and no automatic scale-out occurs. This parameter allows flexible control over elastic scaling behavior and helps reduce costs.
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
	// The number of concurrent sessions, which is the number of sessions that can be simultaneously connected to a single resource. If too many sessions are connected simultaneously, the application experience may degrade. The valid values vary by resource specification:
	//
	// - appstreaming.general.4c8g: 1 to 2.
	//
	// - appstreaming.general.8c16g: 1 to 4.
	//
	// - appstreaming.vgpu.8c16g.4g: 1 to 4.
	//
	// - appstreaming.vgpu.8c31g.16g: 1 to 4.
	//
	// - appstreaming.vgpu.14c93g.12g: 1 to 6.
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The instance type ID of the purchased resource.
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
	RecurrenceSchedules []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// The idle duration without session connections, in minutes. When a resource remains without session connections for the specified duration, automatic scale-in is triggered. Default value: 5.
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
	// The number of resources created during each scale-out operation. Valid values: 1 to 10.
	//
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper threshold of session usage (%). When the session usage exceeds this threshold, automatic scale-out is triggered. The formula for session usage is: `Session usage = Current number of sessions ÷ (Total number of resources × Concurrent sessions per resource) × 100%`.
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

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetAmount() *int32 {
	return s.Amount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetMaxIdleAppInstanceAmount() *int32 {
	return s.MaxIdleAppInstanceAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetMaxScalingAmount() *int32 {
	return s.MaxScalingAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeCapacity() *int32 {
	return s.NodeCapacity
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeTypeName() *string {
	return s.NodeTypeName
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetNodeUsed() *int32 {
	return s.NodeUsed
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetRecurrenceSchedules() []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	return s.RecurrenceSchedules
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingDownAfterIdleMinutes() *int32 {
	return s.ScalingDownAfterIdleMinutes
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingNodeAmount() *int32 {
	return s.ScalingNodeAmount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingNodeUsed() *int32 {
	return s.ScalingNodeUsed
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingStep() *int32 {
	return s.ScalingStep
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetScalingUsageThreshold() *string {
	return s.ScalingUsageThreshold
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetStrategyDisableDate() *string {
	return s.StrategyDisableDate
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetStrategyEnableDate() *string {
	return s.StrategyEnableDate
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetStrategyType() *string {
	return s.StrategyType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GetWarmUp() *bool {
	return s.WarmUp
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxIdleAppInstanceAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeTypeName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeTypeName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) Validate() error {
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

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	// The type of the policy execution cycle. You must specify both `RecurrenceType` and `RecurrenceValues`.
	//
	// example:
	//
	// Weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The list of values for the policy execution cycle.
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The list of time periods for the policy execution cycle.
	TimerPeriods []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GetTimerPeriods() []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	return s.TimerPeriods
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) Validate() error {
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

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
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

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GetAmount() *int32 {
	return s.Amount
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
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

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GetNewOtaVersion() *string {
	return s.NewOtaVersion
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GetOtaVersion() *string {
	return s.OtaVersion
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags struct {
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

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GetKey() *string {
	return s.Key
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GetScope() *string {
	return s.Scope
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) GetValue() *string {
	return s.Value
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) SetKey(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	s.Key = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) SetScope(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	s.Scope = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) SetValue(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags {
	s.Value = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsTags) Validate() error {
	return dara.Validate(s)
}
