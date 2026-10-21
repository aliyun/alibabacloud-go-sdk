// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataStorageResponseBodyData) *GetDataStorageResponseBody
	GetData() *GetDataStorageResponseBodyData
	SetRequestId(v string) *GetDataStorageResponseBody
	GetRequestId() *string
}

type GetDataStorageResponseBody struct {
	// The detailed information returned.
	//
	// example:
	//
	// {"DataStorageRegionId":"cn-shanghai","DataStorageTotalCapacity":1000,"DataStorageUsedCapacity":41.2,"LogProject":"aliyun-cloudsiem-data-1234567890123456-cn-shanghai","SasLogStores":[{"LogName":"Client events","LogStoreName":"aegis-log-client","LogDeliveryStatus":"enable","LogStoreTtl":180,"LogDeliveryUpdateTime":"2026-10-19T14:00:00","LogDeliveryPermission":"allow","LogDeliveryGroup":"host","LogSearchConditions":"[{\\"__topic__\\":\\"aegis-log-client\\"}]","LogCode":"aegis-log-client","LogStoreExisted":true,"UsedCapacity":10.3,"AccountDeliveries":[{"AccountId":1234567890123456,"AccountName":"1234567890123456","LogDeliveryStatus":"enable","LogDeliveryUpdateTime":"2026-10-19T14:00:00Z","IsDaAccount":true}]}],"NormalizationLogStores":[{"LogStoreName":"host-activity","LogStoreTtl":180,"UsedCapacity":10.3}],"UnusedLogStores":[{"LogStoreName":"cloud-siem","LogStoreTtl":180,"UsedCapacity":10.3}],"RecordLogStores":[{"LogStoreName":"alert-record","LogStoreTtl":180,"UsedCapacity":10.3}],"NormalizationLogViews":[{"CategoryName":"Endpoint security","ActivityName":"Host activity","LogViewName":"host_activity","LogStoreName":"host-activity","DetectionRuleReferenceCount":1,"DetectionRuleReferenceProductIds":["sas"],"LogSearchConditions":"[{\\"SCHEMA\\":\\"host-activity\\"}]","LogViewExisted":true}],"ColdStorageUsedCapacity":0.0,"DataStorageUsedCapacityDetail":"{\\"purchasedHotStorageCapacity\\":1000,\\"usedHotStorageCapacity\\":41.2,\\"usedColdStorageCapacity\\":0.0,\\"usedHotStorageCapacityDetail\\":{\\"cn-shanghai\\":41.2},\\"usedColdStorageCapacityDetail\\":{\\"cn-shanghai\\":0.0}}","DataStorageRegionPermission":"allow","LogServiceDisabled":false,"LogProjectStateChangeAllowed":true,"LogProjectState":"ACTIVE","DeliverySettings":{"PrerequisiteEnabled":true,"CanConfigure":true,"Enabled":true,"Initialized":true,"DaAccountId":"1234567890123456","TargetProject":"aliyun-cloudsiem-data-1234567890123456-cn-shanghai","TargetRegion":"cn-shanghai","ConfigurationVersion":1,"Accounts":[{"AccountId":"1234567890123456","AccountName":"1234567890123456","DaAccount":true,"Authorized":true}],"Types":[{"LogCode":"aegis-log-client","SelectedAccountIds":["1234567890123456"],"AutoIncludeNewAccounts":true,"Executions":[]}]}}
	Data *GetDataStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 81FB0DEA-52C1-55A0-8631-8E1B9A9D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBody) GetData() *GetDataStorageResponseBodyData {
	return s.Data
}

func (s *GetDataStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataStorageResponseBody) SetData(v *GetDataStorageResponseBodyData) *GetDataStorageResponseBody {
	s.Data = v
	return s
}

func (s *GetDataStorageResponseBody) SetRequestId(v string) *GetDataStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataStorageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataStorageResponseBodyData struct {
	// The amount of cold storage used for user logs.
	//
	// example:
	//
	// 100.0
	ColdStorageUsedCapacity *float64 `json:"ColdStorageUsedCapacity,omitempty" xml:"ColdStorageUsedCapacity,omitempty"`
	// The storage region of user logs.
	//
	// example:
	//
	// cn-shanghai
	DataStorageRegionId *string `json:"DataStorageRegionId,omitempty" xml:"DataStorageRegionId,omitempty"`
	// Specifies whether the storage region can be modified. By default, the region cannot be modified. You can contact the product manager to reset the region, and the region can be reset only once. Valid values:
	//
	// - allow: The region can be modified.
	//
	// - deny: The region cannot be modified.
	//
	// example:
	//
	// deny
	DataStorageRegionPermission *string `json:"DataStorageRegionPermission,omitempty" xml:"DataStorageRegionPermission,omitempty"`
	// The storage capacity purchased for subscription.
	//
	// example:
	//
	// 100
	DataStorageTotalCapacity *int64 `json:"DataStorageTotalCapacity,omitempty" xml:"DataStorageTotalCapacity,omitempty"`
	// The amount of storage used in user log management.
	//
	// example:
	//
	// 100.0
	DataStorageUsedCapacity *float64 `json:"DataStorageUsedCapacity,omitempty" xml:"DataStorageUsedCapacity,omitempty"`
	// The storage usage details of the log management user.
	//
	// example:
	//
	// {\\"purchasedHotStorageCapacity\\":1000,\\"usedHotStorageCapacity\\":4.2,\\"usedHotStorageCapacityDetail\\":{\\"ap-southeast-1\\":4.2,\\"cn-shenzhen\\":0.0,\\"cn-shanghai\\":0.0}}
	DataStorageUsedCapacityDetail *string `json:"DataStorageUsedCapacityDetail,omitempty" xml:"DataStorageUsedCapacityDetail,omitempty"`
	// The multi-account delivery settings for the current storage resources. Returned when the query is successful. If the read operation fails, see DeliverySettingsError.
	//
	// example:
	//
	// {"PrerequisiteEnabled":true,"CanConfigure":true,"Enabled":true,"Initialized":true,"DaAccountId":"1234567890123456","TargetProject":"aliyun-cloudsiem-data-1234567890123456-cn-shanghai","TargetRegion":"cn-shanghai","ConfigurationVersion":1,"Accounts":[{"AccountId":"1234567890123456","AccountName":"1234567890123456","DaAccount":true,"Authorized":true}],"Types":[{"LogCode":"aegis-log-client","SelectedAccountIds":["1234567890123456"],"AutoIncludeNewAccounts":true,"Executions":[]}]}
	DeliverySettings *GetDataStorageResponseBodyDataDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The read error code for the optional multi-account delivery settings. If the read operation fails, DELIVERY_SETTINGS_UNAVAILABLE is returned. If the read operation is successful, this field is empty or not returned, and does not affect other storage fields.
	//
	// example:
	//
	// DELIVERY_SETTINGS_UNAVAILABLE
	DeliverySettingsError *string `json:"DeliverySettingsError,omitempty" xml:"DeliverySettingsError,omitempty"`
	// The name of the Simple Log Service project where user logs are stored.
	//
	// example:
	//
	// aliyun-cloudsiem-data-171835723111****-cn-shanghai
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The status of the log project.
	//
	// example:
	//
	// ACTIVE
	LogProjectState *string `json:"LogProjectState,omitempty" xml:"LogProjectState,omitempty"`
	// Specifies whether the status of the log project can be changed.
	//
	// example:
	//
	// true
	LogProjectStateChangeAllowed *bool `json:"LogProjectStateChangeAllowed,omitempty" xml:"LogProjectStateChangeAllowed,omitempty"`
	// Specifies whether Simple Log Service is disabled.
	//
	// example:
	//
	// false
	LogServiceDisabled *bool `json:"LogServiceDisabled,omitempty" xml:"LogServiceDisabled,omitempty"`
	// The details of the Logstores for normalized data.
	//
	// example:
	//
	// [{"LogStoreName":"host-activity","LogStoreTtl":180,"UsedCapacity":10.3}]
	NormalizationLogStores []*GetDataStorageResponseBodyDataNormalizationLogStores `json:"NormalizationLogStores,omitempty" xml:"NormalizationLogStores,omitempty" type:"Repeated"`
	// The details of the normalized datasets.
	//
	// example:
	//
	// [{"CategoryName":"Endpoint security","ActivityName":"Host activity","LogViewName":"host_activity","LogStoreName":"host-activity","DetectionRuleReferenceCount":1,"DetectionRuleReferenceProductIds":["sas"],"LogSearchConditions":"[{\\"SCHEMA\\":\\"host-activity\\"}]","LogViewExisted":true}]
	NormalizationLogViews []*GetDataStorageResponseBodyDataNormalizationLogViews `json:"NormalizationLogViews,omitempty" xml:"NormalizationLogViews,omitempty" type:"Repeated"`
	// The list of Record Logstores.
	//
	// example:
	//
	// [{"LogStoreName":"alert-record","LogStoreTtl":180,"UsedCapacity":10.3}]
	RecordLogStores []*GetDataStorageResponseBodyDataRecordLogStores `json:"RecordLogStores,omitempty" xml:"RecordLogStores,omitempty" type:"Repeated"`
	// The storage details of raw logs in Security Center.
	//
	// example:
	//
	// [{"LogName":"Client events","LogStoreName":"aegis-log-client","LogDeliveryStatus":"enable","LogStoreTtl":180,"LogDeliveryUpdateTime":"2026-10-19T14:00:00","LogDeliveryPermission":"allow","LogDeliveryGroup":"host","LogSearchConditions":"[{\\"__topic__\\":\\"aegis-log-client\\"}]","LogCode":"aegis-log-client","LogStoreExisted":true,"UsedCapacity":10.3,"AccountDeliveries":[{"AccountId":1234567890123456,"AccountName":"1234567890123456","LogDeliveryStatus":"enable","LogDeliveryUpdateTime":"2026-10-19T14:00:00Z","IsDaAccount":true}]}]
	SasLogStores []*GetDataStorageResponseBodyDataSasLogStores `json:"SasLogStores,omitempty" xml:"SasLogStores,omitempty" type:"Repeated"`
	// The list of legacy Logstores in SIEM V1.
	//
	// example:
	//
	// [{"LogStoreName":"cloud-siem","LogStoreTtl":180,"UsedCapacity":10.3}]
	UnusedLogStores []*GetDataStorageResponseBodyDataUnusedLogStores `json:"UnusedLogStores,omitempty" xml:"UnusedLogStores,omitempty" type:"Repeated"`
}

func (s GetDataStorageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyData) GetColdStorageUsedCapacity() *float64 {
	return s.ColdStorageUsedCapacity
}

func (s *GetDataStorageResponseBodyData) GetDataStorageRegionId() *string {
	return s.DataStorageRegionId
}

func (s *GetDataStorageResponseBodyData) GetDataStorageRegionPermission() *string {
	return s.DataStorageRegionPermission
}

func (s *GetDataStorageResponseBodyData) GetDataStorageTotalCapacity() *int64 {
	return s.DataStorageTotalCapacity
}

func (s *GetDataStorageResponseBodyData) GetDataStorageUsedCapacity() *float64 {
	return s.DataStorageUsedCapacity
}

func (s *GetDataStorageResponseBodyData) GetDataStorageUsedCapacityDetail() *string {
	return s.DataStorageUsedCapacityDetail
}

func (s *GetDataStorageResponseBodyData) GetDeliverySettings() *GetDataStorageResponseBodyDataDeliverySettings {
	return s.DeliverySettings
}

func (s *GetDataStorageResponseBodyData) GetDeliverySettingsError() *string {
	return s.DeliverySettingsError
}

func (s *GetDataStorageResponseBodyData) GetLogProject() *string {
	return s.LogProject
}

func (s *GetDataStorageResponseBodyData) GetLogProjectState() *string {
	return s.LogProjectState
}

func (s *GetDataStorageResponseBodyData) GetLogProjectStateChangeAllowed() *bool {
	return s.LogProjectStateChangeAllowed
}

func (s *GetDataStorageResponseBodyData) GetLogServiceDisabled() *bool {
	return s.LogServiceDisabled
}

func (s *GetDataStorageResponseBodyData) GetNormalizationLogStores() []*GetDataStorageResponseBodyDataNormalizationLogStores {
	return s.NormalizationLogStores
}

func (s *GetDataStorageResponseBodyData) GetNormalizationLogViews() []*GetDataStorageResponseBodyDataNormalizationLogViews {
	return s.NormalizationLogViews
}

func (s *GetDataStorageResponseBodyData) GetRecordLogStores() []*GetDataStorageResponseBodyDataRecordLogStores {
	return s.RecordLogStores
}

func (s *GetDataStorageResponseBodyData) GetSasLogStores() []*GetDataStorageResponseBodyDataSasLogStores {
	return s.SasLogStores
}

func (s *GetDataStorageResponseBodyData) GetUnusedLogStores() []*GetDataStorageResponseBodyDataUnusedLogStores {
	return s.UnusedLogStores
}

func (s *GetDataStorageResponseBodyData) SetColdStorageUsedCapacity(v float64) *GetDataStorageResponseBodyData {
	s.ColdStorageUsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageRegionId(v string) *GetDataStorageResponseBodyData {
	s.DataStorageRegionId = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageRegionPermission(v string) *GetDataStorageResponseBodyData {
	s.DataStorageRegionPermission = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageTotalCapacity(v int64) *GetDataStorageResponseBodyData {
	s.DataStorageTotalCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageUsedCapacity(v float64) *GetDataStorageResponseBodyData {
	s.DataStorageUsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageUsedCapacityDetail(v string) *GetDataStorageResponseBodyData {
	s.DataStorageUsedCapacityDetail = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDeliverySettings(v *GetDataStorageResponseBodyDataDeliverySettings) *GetDataStorageResponseBodyData {
	s.DeliverySettings = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDeliverySettingsError(v string) *GetDataStorageResponseBodyData {
	s.DeliverySettingsError = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogProject(v string) *GetDataStorageResponseBodyData {
	s.LogProject = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogProjectState(v string) *GetDataStorageResponseBodyData {
	s.LogProjectState = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogProjectStateChangeAllowed(v bool) *GetDataStorageResponseBodyData {
	s.LogProjectStateChangeAllowed = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogServiceDisabled(v bool) *GetDataStorageResponseBodyData {
	s.LogServiceDisabled = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetNormalizationLogStores(v []*GetDataStorageResponseBodyDataNormalizationLogStores) *GetDataStorageResponseBodyData {
	s.NormalizationLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetNormalizationLogViews(v []*GetDataStorageResponseBodyDataNormalizationLogViews) *GetDataStorageResponseBodyData {
	s.NormalizationLogViews = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetRecordLogStores(v []*GetDataStorageResponseBodyDataRecordLogStores) *GetDataStorageResponseBodyData {
	s.RecordLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetSasLogStores(v []*GetDataStorageResponseBodyDataSasLogStores) *GetDataStorageResponseBodyData {
	s.SasLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetUnusedLogStores(v []*GetDataStorageResponseBodyDataUnusedLogStores) *GetDataStorageResponseBodyData {
	s.UnusedLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) Validate() error {
	if s.DeliverySettings != nil {
		if err := s.DeliverySettings.Validate(); err != nil {
			return err
		}
	}
	if s.NormalizationLogStores != nil {
		for _, item := range s.NormalizationLogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NormalizationLogViews != nil {
		for _, item := range s.NormalizationLogViews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RecordLogStores != nil {
		for _, item := range s.RecordLogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SasLogStores != nil {
		for _, item := range s.SasLogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnusedLogStores != nil {
		for _, item := range s.UnusedLogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataStorageResponseBodyDataDeliverySettings struct {
	// The account directory of the current account and managed member accounts.
	//
	// example:
	//
	// [{"AccountId":"1234567890123456","AccountName":"1234567890123456","DaAccount":true,"Authorized":true}]
	Accounts []*GetDataStorageResponseBodyDataDeliverySettingsAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// Specifies whether the current account meets the conditions for delivery setting operations. This is determined based on authorization, target project availability, capacity orders, and the current account identity. The PrerequisiteEnabled parameter must also be met.
	//
	// example:
	//
	// true
	CanConfigure *bool `json:"CanConfigure,omitempty" xml:"CanConfigure,omitempty"`
	// The current version of the delivery settings. When you modify multi-account configurations or specify member delivery, pass this value back as ConfigurationVersion. You must query this value again after each modification and cannot use a fixed example value.
	//
	// example:
	//
	// 1
	ConfigurationVersion *int64 `json:"ConfigurationVersion,omitempty" xml:"ConfigurationVersion,omitempty"`
	// The UID of the current management and delivery target account, which is a decimal string. This is also the current account that is permanently retained in the account selection list.
	//
	// example:
	//
	// 1234567890123456
	DaAccountId *string `json:"DaAccountId,omitempty" xml:"DaAccountId,omitempty"`
	// The master switch for the saved multi-account centralized delivery. A value of true indicates that the switch is turned on. A value of false indicates that the switch is turned off.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether the initialization of the account scope and automatic inclusion settings is completed when the feature is enabled for the first time.
	//
	// example:
	//
	// true
	Initialized *bool `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	// Specifies whether the prerequisite of enabling multi-account security management in Security Center is met.
	//
	// example:
	//
	// true
	PrerequisiteEnabled *bool `json:"PrerequisiteEnabled,omitempty" xml:"PrerequisiteEnabled,omitempty"`
	// The name of the current Simple Log Service project for centralized delivery. When you modify multi-account settings or specify member delivery, pass this value back as the TargetProject request parameter.
	//
	// example:
	//
	// aliyun-cloudsiem-data-1234567890123456-cn-shanghai
	TargetProject *string `json:"TargetProject,omitempty" xml:"TargetProject,omitempty"`
	// The region ID of the current Simple Log Service project for centralized delivery.
	//
	// example:
	//
	// cn-shanghai
	TargetRegion *string `json:"TargetRegion,omitempty" xml:"TargetRegion,omitempty"`
	// The account selection, automatic inclusion settings, and existing delivery status categorized by log type.
	//
	// example:
	//
	// [{"LogCode":"aegis-log-client","SelectedAccountIds":["1234567890123456"],"AutoIncludeNewAccounts":true,"Executions":[]}]
	Types []*GetDataStorageResponseBodyDataDeliverySettingsTypes `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// The reason code for being unconfigurable. Valid values: DA_UNAUTHORIZED (the management account is not authorized), TARGET_PROJECT_UNAVAILABLE (the target project is unavailable), LOG_CAPACITY_UNAVAILABLE (no valid capacity order exists), and NOT_CURRENT_ACCOUNT (the operation is not performed by the current account). This field is empty or not returned if the conditions are met.
	//
	// example:
	//
	// DA_UNAUTHORIZED
	UnavailableReason *string `json:"UnavailableReason,omitempty" xml:"UnavailableReason,omitempty"`
}

func (s GetDataStorageResponseBodyDataDeliverySettings) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataDeliverySettings) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetAccounts() []*GetDataStorageResponseBodyDataDeliverySettingsAccounts {
	return s.Accounts
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetCanConfigure() *bool {
	return s.CanConfigure
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetConfigurationVersion() *int64 {
	return s.ConfigurationVersion
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetDaAccountId() *string {
	return s.DaAccountId
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetInitialized() *bool {
	return s.Initialized
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetPrerequisiteEnabled() *bool {
	return s.PrerequisiteEnabled
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetTargetProject() *string {
	return s.TargetProject
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetTargetRegion() *string {
	return s.TargetRegion
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetTypes() []*GetDataStorageResponseBodyDataDeliverySettingsTypes {
	return s.Types
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) GetUnavailableReason() *string {
	return s.UnavailableReason
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetAccounts(v []*GetDataStorageResponseBodyDataDeliverySettingsAccounts) *GetDataStorageResponseBodyDataDeliverySettings {
	s.Accounts = v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetCanConfigure(v bool) *GetDataStorageResponseBodyDataDeliverySettings {
	s.CanConfigure = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetConfigurationVersion(v int64) *GetDataStorageResponseBodyDataDeliverySettings {
	s.ConfigurationVersion = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetDaAccountId(v string) *GetDataStorageResponseBodyDataDeliverySettings {
	s.DaAccountId = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetEnabled(v bool) *GetDataStorageResponseBodyDataDeliverySettings {
	s.Enabled = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetInitialized(v bool) *GetDataStorageResponseBodyDataDeliverySettings {
	s.Initialized = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetPrerequisiteEnabled(v bool) *GetDataStorageResponseBodyDataDeliverySettings {
	s.PrerequisiteEnabled = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetTargetProject(v string) *GetDataStorageResponseBodyDataDeliverySettings {
	s.TargetProject = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetTargetRegion(v string) *GetDataStorageResponseBodyDataDeliverySettings {
	s.TargetRegion = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetTypes(v []*GetDataStorageResponseBodyDataDeliverySettingsTypes) *GetDataStorageResponseBodyDataDeliverySettings {
	s.Types = v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) SetUnavailableReason(v string) *GetDataStorageResponseBodyDataDeliverySettings {
	s.UnavailableReason = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettings) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Types != nil {
		for _, item := range s.Types {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataStorageResponseBodyDataDeliverySettingsAccounts struct {
	// The UID of the source Alibaba Cloud account, which is a decimal string.
	//
	// example:
	//
	// 1234567890123456
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The display name of the account. If the name is missing, the UID string is returned.
	//
	// example:
	//
	// 1234567890123456
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Specifies whether the account has the authorization required for log delivery. This parameter must be set to true when you add a selected member.
	//
	// example:
	//
	// true
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// Specifies whether the account is the current management and delivery target account, which is the current account.
	//
	// example:
	//
	// true
	DaAccount *bool `json:"DaAccount,omitempty" xml:"DaAccount,omitempty"`
}

func (s GetDataStorageResponseBodyDataDeliverySettingsAccounts) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataDeliverySettingsAccounts) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) GetAuthorized() *bool {
	return s.Authorized
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) GetDaAccount() *bool {
	return s.DaAccount
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) SetAccountId(v string) *GetDataStorageResponseBodyDataDeliverySettingsAccounts {
	s.AccountId = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) SetAccountName(v string) *GetDataStorageResponseBodyDataDeliverySettingsAccounts {
	s.AccountName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) SetAuthorized(v bool) *GetDataStorageResponseBodyDataDeliverySettingsAccounts {
	s.Authorized = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) SetDaAccount(v bool) *GetDataStorageResponseBodyDataDeliverySettingsAccounts {
	s.DaAccount = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsAccounts) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataDeliverySettingsTypes struct {
	// Specifies whether to automatically add subsequent new and available managed members to the selection scope of this log type.
	//
	// example:
	//
	// true
	AutoIncludeNewAccounts *bool `json:"AutoIncludeNewAccounts,omitempty" xml:"AutoIncludeNewAccounts,omitempty"`
	// The status list of existing source delivery records for this log type.
	//
	// example:
	//
	// []
	Executions []*GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The unique identifier of the log type, which corresponds to UpdateDataStorageDelivery.LogCode.
	//
	// example:
	//
	// aegis-log-client
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The array of UIDs of the selected source accounts for this log type, including the current account. The response is a JSON string array. When you write an array(simple) request, you must serialize it into comma-separated values.
	//
	// example:
	//
	// ["1234567890123456"]
	SelectedAccountIds []*string `json:"SelectedAccountIds,omitempty" xml:"SelectedAccountIds,omitempty" type:"Repeated"`
}

func (s GetDataStorageResponseBodyDataDeliverySettingsTypes) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataDeliverySettingsTypes) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) GetAutoIncludeNewAccounts() *bool {
	return s.AutoIncludeNewAccounts
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) GetExecutions() []*GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions {
	return s.Executions
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) GetLogCode() *string {
	return s.LogCode
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) GetSelectedAccountIds() []*string {
	return s.SelectedAccountIds
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) SetAutoIncludeNewAccounts(v bool) *GetDataStorageResponseBodyDataDeliverySettingsTypes {
	s.AutoIncludeNewAccounts = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) SetExecutions(v []*GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) *GetDataStorageResponseBodyDataDeliverySettingsTypes {
	s.Executions = v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) SetLogCode(v string) *GetDataStorageResponseBodyDataDeliverySettingsTypes {
	s.LogCode = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) SetSelectedAccountIds(v []*string) *GetDataStorageResponseBodyDataDeliverySettingsTypes {
	s.SelectedAccountIds = v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypes) Validate() error {
	if s.Executions != nil {
		for _, item := range s.Executions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions struct {
	// The UID of the source account for this delivery record, which is a decimal string.
	//
	// example:
	//
	// 1234567890123457
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// A reserved compatibility field. The current implementation does not populate historical execution error codes, so this field is usually empty or not returned. If execution fails, read the error code returned by the corresponding write operation.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The recorded delivery status. RUNNING indicates that the status is recorded as enabled after the last successful operation. STOPPED indicates that the status is recorded as disabled. This field does not indicate real-time log arrival.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time of the delivery status, which is a UNIX timestamp in milliseconds. This field is empty or not returned if the record has no update time.
	//
	// example:
	//
	// 1792399200000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) GetStatus() *string {
	return s.Status
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) SetAccountId(v string) *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions {
	s.AccountId = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) SetErrorCode(v string) *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions {
	s.ErrorCode = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) SetStatus(v string) *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions {
	s.Status = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) SetUpdateTime(v int64) *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions {
	s.UpdateTime = &v
	return s
}

func (s *GetDataStorageResponseBodyDataDeliverySettingsTypesExecutions) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataNormalizationLogStores struct {
	// The name of the Logstore that stores normalized data.
	//
	// example:
	//
	// vulnerability-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The storage duration of the normalized data.
	//
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The amount of hot storage used.
	//
	// example:
	//
	// 10.333
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetDataStorageResponseBodyDataNormalizationLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataNormalizationLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataNormalizationLogViews struct {
	// The log type of the normalized logs.
	//
	// example:
	//
	// API security risk log
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The category of the normalized logs.
	//
	// example:
	//
	// Security Category
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The number of times this normalized dataset is referenced in Threat Analysis and Response.
	//
	// example:
	//
	// 3
	DetectionRuleReferenceCount *int32 `json:"DetectionRuleReferenceCount,omitempty" xml:"DetectionRuleReferenceCount,omitempty"`
	// The list of products that reference this normalized dataset in Threat Analysis and Response.
	//
	// example:
	//
	// ["sas"]
	DetectionRuleReferenceProductIds []*string `json:"DetectionRuleReferenceProductIds,omitempty" xml:"DetectionRuleReferenceProductIds,omitempty" type:"Repeated"`
	// The query statement used to query log types in the normalized dataset.
	//
	// example:
	//
	// [{\\"SCHEMA\\":\\"AZURE_ACTIVE_DIRECTORY_AUDIT_ACTIVITY\\"}]
	LogSearchConditions *string `json:"LogSearchConditions,omitempty" xml:"LogSearchConditions,omitempty"`
	// The Logstore where Threat Analysis and Response stores normalized logs.
	//
	// example:
	//
	// risk-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// Specifies whether the normalized dataset exists. Valid values:
	//
	// - true: The dataset exists.
	//
	// - false: The dataset does not exist.
	//
	// example:
	//
	// true
	LogViewExisted *bool `json:"LogViewExisted,omitempty" xml:"LogViewExisted,omitempty"`
	// The name of the normalized dataset.
	//
	// example:
	//
	// risk_activity
	LogViewName *string `json:"LogViewName,omitempty" xml:"LogViewName,omitempty"`
}

func (s GetDataStorageResponseBodyDataNormalizationLogViews) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataNormalizationLogViews) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetDetectionRuleReferenceCount() *int32 {
	return s.DetectionRuleReferenceCount
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetDetectionRuleReferenceProductIds() []*string {
	return s.DetectionRuleReferenceProductIds
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogSearchConditions() *string {
	return s.LogSearchConditions
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogViewExisted() *bool {
	return s.LogViewExisted
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogViewName() *string {
	return s.LogViewName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetActivityName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.ActivityName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetCategoryName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.CategoryName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetDetectionRuleReferenceCount(v int32) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.DetectionRuleReferenceCount = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetDetectionRuleReferenceProductIds(v []*string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.DetectionRuleReferenceProductIds = v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogSearchConditions(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogSearchConditions = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogStoreName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogViewExisted(v bool) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogViewExisted = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogViewName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogViewName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataRecordLogStores struct {
	// The name of the Logstore.
	//
	// example:
	//
	// alert-record
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The TTL of the Logstore.
	//
	// example:
	//
	// 90
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The used capacity of the Logstore.
	//
	// example:
	//
	// 11.111
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetDataStorageResponseBodyDataRecordLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataRecordLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataRecordLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataRecordLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataRecordLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataSasLogStores struct {
	// The delivery records of this log type expanded by source account, including the current account and member accounts with existing records.
	//
	// example:
	//
	// [{"AccountId":1234567890123456,"AccountName":"1234567890123456","LogDeliveryStatus":"enable","LogDeliveryUpdateTime":"2026-10-19T14:00:00Z","IsDaAccount":true}]
	AccountDeliveries []*GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries `json:"AccountDeliveries,omitempty" xml:"AccountDeliveries,omitempty" type:"Repeated"`
	// The log code.
	//
	// example:
	//
	// sas-net-block
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The group to which the log belongs. Valid values:
	//
	// - host: Host logs.
	//
	// - security: Security logs.
	//
	// example:
	//
	// host
	LogDeliveryGroup *string `json:"LogDeliveryGroup,omitempty" xml:"LogDeliveryGroup,omitempty"`
	// Specifies whether the log delivery switch can be operated. Log delivery is unavailable if the service is not purchased. Valid values:
	//
	// - allow: Allowed.
	//
	// - deny: Not allowed.
	//
	// example:
	//
	// deny
	LogDeliveryPermission *string `json:"LogDeliveryPermission,omitempty" xml:"LogDeliveryPermission,omitempty"`
	// The log delivery status. Valid values:
	//
	// - enable: Log delivery is enabled.
	//
	// - disable: Log delivery is disabled.
	//
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// The time of the last log delivery operation. The format is YYYY-MM-DDTHH:mm:ss.
	//
	// example:
	//
	// 2025-07-16T15:10:29
	LogDeliveryUpdateTime *string `json:"LogDeliveryUpdateTime,omitempty" xml:"LogDeliveryUpdateTime,omitempty"`
	// The log name.
	//
	// example:
	//
	// Process Snapshot
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The default query conditions for the log. When multiple logs are stored in a single Logstore, you must use query conditions to perform a log query for a specific log.
	//
	// example:
	//
	// [{\\"__topic__\\":\\"sas-net-block\\"}]
	LogSearchConditions *string `json:"LogSearchConditions,omitempty" xml:"LogSearchConditions,omitempty"`
	// Indicates whether the Logstore where the log is stored exists. Valid values:
	//
	// - true: Exists.
	//
	// - false: Does not exist.
	//
	// example:
	//
	// true
	LogStoreExisted *bool `json:"LogStoreExisted,omitempty" xml:"LogStoreExisted,omitempty"`
	// The name of the Logstore where the log is stored.
	//
	// example:
	//
	// sas-security-log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The storage duration of the Logstore where the log is stored. The minimum storage duration is 30 days.
	//
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The amount of hot storage used.
	//
	// example:
	//
	// 10.333
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetDataStorageResponseBodyDataSasLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataSasLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetAccountDeliveries() []*GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries {
	return s.AccountDeliveries
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogCode() *string {
	return s.LogCode
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryGroup() *string {
	return s.LogDeliveryGroup
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryPermission() *string {
	return s.LogDeliveryPermission
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryStatus() *string {
	return s.LogDeliveryStatus
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryUpdateTime() *string {
	return s.LogDeliveryUpdateTime
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogName() *string {
	return s.LogName
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogSearchConditions() *string {
	return s.LogSearchConditions
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogStoreExisted() *bool {
	return s.LogStoreExisted
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetAccountDeliveries(v []*GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) *GetDataStorageResponseBodyDataSasLogStores {
	s.AccountDeliveries = v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogCode(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogCode = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryGroup(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryGroup = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryPermission(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryPermission = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryStatus(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryStatus = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryUpdateTime(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryUpdateTime = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogName(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogSearchConditions(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogSearchConditions = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogStoreExisted(v bool) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogStoreExisted = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataSasLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) Validate() error {
	if s.AccountDeliveries != nil {
		for _, item := range s.AccountDeliveries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries struct {
	// The UID of the source cloud account. This field is an int64 number.
	//
	// example:
	//
	// 1234567890123456
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The display name of the source account. If the name is not obtained, the account UID string is used.
	//
	// example:
	//
	// 1234567890123456
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the account is the current management or target account (the current account). A value of true indicates the current account. A value of false indicates a member source account.
	//
	// example:
	//
	// true
	IsDaAccount *bool `json:"IsDaAccount,omitempty" xml:"IsDaAccount,omitempty"`
	// The recorded delivery status of the source account. Valid values:
	//
	// - enable: Enabled.
	//
	// - disable: Disabled.
	//
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// The last update time of the delivery status for the source account. If no record is generated, this field is empty or not returned.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-10-19T14:00:00Z
	LogDeliveryUpdateTime *string `json:"LogDeliveryUpdateTime,omitempty" xml:"LogDeliveryUpdateTime,omitempty"`
}

func (s GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) GetAccountName() *string {
	return s.AccountName
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) GetIsDaAccount() *bool {
	return s.IsDaAccount
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) GetLogDeliveryStatus() *string {
	return s.LogDeliveryStatus
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) GetLogDeliveryUpdateTime() *string {
	return s.LogDeliveryUpdateTime
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) SetAccountId(v int64) *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries {
	s.AccountId = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) SetAccountName(v string) *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries {
	s.AccountName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) SetIsDaAccount(v bool) *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries {
	s.IsDaAccount = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) SetLogDeliveryStatus(v string) *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries {
	s.LogDeliveryStatus = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) SetLogDeliveryUpdateTime(v string) *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries {
	s.LogDeliveryUpdateTime = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStoresAccountDeliveries) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataUnusedLogStores struct {
	// The name of the Logstore.
	//
	// example:
	//
	// cloud-siem
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The data retention period.
	//
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The amount of hot storage used.
	//
	// example:
	//
	// 10.333
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetDataStorageResponseBodyDataUnusedLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataUnusedLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataUnusedLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataUnusedLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataUnusedLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) Validate() error {
	return dara.Validate(s)
}
