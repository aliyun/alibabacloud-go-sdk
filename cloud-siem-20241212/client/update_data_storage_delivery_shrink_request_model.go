// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageDeliveryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoIncludeNewAccounts(v bool) *UpdateDataStorageDeliveryShrinkRequest
	GetAutoIncludeNewAccounts() *bool
	SetConfigurationVersion(v int64) *UpdateDataStorageDeliveryShrinkRequest
	GetConfigurationVersion() *int64
	SetLang(v string) *UpdateDataStorageDeliveryShrinkRequest
	GetLang() *string
	SetLogCode(v string) *UpdateDataStorageDeliveryShrinkRequest
	GetLogCode() *string
	SetLogDeliveryStatus(v string) *UpdateDataStorageDeliveryShrinkRequest
	GetLogDeliveryStatus() *string
	SetRegionId(v string) *UpdateDataStorageDeliveryShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataStorageDeliveryShrinkRequest
	GetRoleFor() *int64
	SetSelectedAccountIdsShrink(v string) *UpdateDataStorageDeliveryShrinkRequest
	GetSelectedAccountIdsShrink() *string
	SetSubUserId(v int64) *UpdateDataStorageDeliveryShrinkRequest
	GetSubUserId() *int64
	SetTargetProject(v string) *UpdateDataStorageDeliveryShrinkRequest
	GetTargetProject() *string
}

type UpdateDataStorageDeliveryShrinkRequest struct {
	// Specifies whether to automatically add subsequent new and available managed members to the account selection for the log type. A value of true indicates that the feature is enabled. A value of false indicates that the feature is disabled. Disabling the feature does not remove selected members. If this parameter is omitted, the original value is retained. You must also specify LogCode, ConfigurationVersion, and TargetProject. You can update this parameter together with SelectedAccountIds. You cannot specify LogDeliveryStatus or the member SubUserId at the same time. The actual delivery is also controlled by the global switch for multi-account delivery and the delivery switch for the log type.
	//
	// example:
	//
	// true
	AutoIncludeNewAccounts *bool `json:"AutoIncludeNewAccounts,omitempty" xml:"AutoIncludeNewAccounts,omitempty"`
	// The version of the delivery settings on which the current modification is based. Call GetDataStorage first and pass the original value of Data.DeliverySettings.ConfigurationVersion. Example 1 only shows the format and cannot be used as a fixed version. This parameter is required when you modify the global switch, account selection, auto-inclusion settings, or specified member delivery. If the version is missing or has changed, the OperationDenied.CentralizedDeliveryConfigChanged error is returned. In this case, query the version again before you submit the request.
	//
	// example:
	//
	// 1
	ConfigurationVersion *int64 `json:"ConfigurationVersion,omitempty" xml:"ConfigurationVersion,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// This parameter is required. If this parameter is not specified, the API returns a 400 error.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The log code.
	//
	// This parameter is required.
	//
	// example:
	//
	// aegis-log-login
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The log delivery status. Valid values:
	//
	// - enable: Enables delivery.
	//
	// - disable: Disables delivery.
	//
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// The region where the data management center for threat analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: The assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The complete list of selected accounts for the log type. Each element is a decimal UID string. The list must include Data.DeliverySettings.DaAccountId (the current account) and can contain up to 10,000 items. A new member must be in the Accounts list returned by GetDataStorage and have Authorized set to true. If only the current account is retained, specify the UID of the current account. Do not use an empty string or an empty array. If this parameter is omitted, the original selection is retained. For HTTP requests, use the comma-separated format of array(simple): SelectedAccountIds=1234567890123456,1234567890123457. For a single account, specify only one UID. Do not use SelectedAccountIds.1/.2 or a JSON array string. You must also specify ConfigurationVersion and TargetProject, and omit LogDeliveryStatus and the member SubUserId.
	//
	// example:
	//
	// 1234567890123456,1234567890123457
	SelectedAccountIdsShrink *string `json:"SelectedAccountIds,omitempty" xml:"SelectedAccountIds,omitempty"`
	// The UID of the source Alibaba Cloud account, which is a decimal integer. If this parameter is omitted or set to the UID of the current account, the delivery settings of the current account are modified. If a member UID is specified, the expected delivery status in the current configuration is applied only to the member. In this case, you must also specify LogDeliveryStatus, ConfigurationVersion, and TargetProject. The member must be within the selection scope of the log type or have existing delivery records. You must omit this parameter when you save the account selection or auto-inclusion settings. You cannot specify the UID of another member.
	//
	// example:
	//
	// 1234567890123457
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The name of the current Simple Log Service project for log delivery. Invoke GetDataStorage first and pass the original value of Data.DeliverySettings.TargetProject. This parameter is used to verify the target of the current modification, not to specify a new project. This parameter is required when you modify the global switch, account selection, auto-inclusion settings, or specified member log delivery. If this parameter is missing or different from the current target, the OperationDenied.CentralizedDeliveryConfigChanged error is returned. Replace the account UID and region in the example with the query results.
	//
	// example:
	//
	// aliyun-cloudsiem-data-1234567890123456-cn-shanghai
	TargetProject *string `json:"TargetProject,omitempty" xml:"TargetProject,omitempty"`
}

func (s UpdateDataStorageDeliveryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageDeliveryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetAutoIncludeNewAccounts() *bool {
	return s.AutoIncludeNewAccounts
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetConfigurationVersion() *int64 {
	return s.ConfigurationVersion
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetLogDeliveryStatus() *string {
	return s.LogDeliveryStatus
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetSelectedAccountIdsShrink() *string {
	return s.SelectedAccountIdsShrink
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *UpdateDataStorageDeliveryShrinkRequest) GetTargetProject() *string {
	return s.TargetProject
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetAutoIncludeNewAccounts(v bool) *UpdateDataStorageDeliveryShrinkRequest {
	s.AutoIncludeNewAccounts = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetConfigurationVersion(v int64) *UpdateDataStorageDeliveryShrinkRequest {
	s.ConfigurationVersion = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetLang(v string) *UpdateDataStorageDeliveryShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetLogCode(v string) *UpdateDataStorageDeliveryShrinkRequest {
	s.LogCode = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetLogDeliveryStatus(v string) *UpdateDataStorageDeliveryShrinkRequest {
	s.LogDeliveryStatus = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetRegionId(v string) *UpdateDataStorageDeliveryShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetRoleFor(v int64) *UpdateDataStorageDeliveryShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetSelectedAccountIdsShrink(v string) *UpdateDataStorageDeliveryShrinkRequest {
	s.SelectedAccountIdsShrink = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetSubUserId(v int64) *UpdateDataStorageDeliveryShrinkRequest {
	s.SubUserId = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) SetTargetProject(v string) *UpdateDataStorageDeliveryShrinkRequest {
	s.TargetProject = &v
	return s
}

func (s *UpdateDataStorageDeliveryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
