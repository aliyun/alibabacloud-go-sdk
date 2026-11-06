// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoIncludeNewAccounts(v bool) *UpdateDataStorageDeliveryRequest
	GetAutoIncludeNewAccounts() *bool
	SetConfigurationVersion(v int64) *UpdateDataStorageDeliveryRequest
	GetConfigurationVersion() *int64
	SetLang(v string) *UpdateDataStorageDeliveryRequest
	GetLang() *string
	SetLogCode(v string) *UpdateDataStorageDeliveryRequest
	GetLogCode() *string
	SetLogDeliveryStatus(v string) *UpdateDataStorageDeliveryRequest
	GetLogDeliveryStatus() *string
	SetRegionId(v string) *UpdateDataStorageDeliveryRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataStorageDeliveryRequest
	GetRoleFor() *int64
	SetSelectedAccountIds(v []*string) *UpdateDataStorageDeliveryRequest
	GetSelectedAccountIds() []*string
	SetSubUserId(v int64) *UpdateDataStorageDeliveryRequest
	GetSubUserId() *int64
	SetTargetProject(v string) *UpdateDataStorageDeliveryRequest
	GetTargetProject() *string
}

type UpdateDataStorageDeliveryRequest struct {
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
	SelectedAccountIds []*string `json:"SelectedAccountIds,omitempty" xml:"SelectedAccountIds,omitempty" type:"Repeated"`
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

func (s UpdateDataStorageDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageDeliveryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageDeliveryRequest) GetAutoIncludeNewAccounts() *bool {
	return s.AutoIncludeNewAccounts
}

func (s *UpdateDataStorageDeliveryRequest) GetConfigurationVersion() *int64 {
	return s.ConfigurationVersion
}

func (s *UpdateDataStorageDeliveryRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataStorageDeliveryRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *UpdateDataStorageDeliveryRequest) GetLogDeliveryStatus() *string {
	return s.LogDeliveryStatus
}

func (s *UpdateDataStorageDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataStorageDeliveryRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataStorageDeliveryRequest) GetSelectedAccountIds() []*string {
	return s.SelectedAccountIds
}

func (s *UpdateDataStorageDeliveryRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *UpdateDataStorageDeliveryRequest) GetTargetProject() *string {
	return s.TargetProject
}

func (s *UpdateDataStorageDeliveryRequest) SetAutoIncludeNewAccounts(v bool) *UpdateDataStorageDeliveryRequest {
	s.AutoIncludeNewAccounts = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetConfigurationVersion(v int64) *UpdateDataStorageDeliveryRequest {
	s.ConfigurationVersion = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetLang(v string) *UpdateDataStorageDeliveryRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetLogCode(v string) *UpdateDataStorageDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetLogDeliveryStatus(v string) *UpdateDataStorageDeliveryRequest {
	s.LogDeliveryStatus = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetRegionId(v string) *UpdateDataStorageDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetRoleFor(v int64) *UpdateDataStorageDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetSelectedAccountIds(v []*string) *UpdateDataStorageDeliveryRequest {
	s.SelectedAccountIds = v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetSubUserId(v int64) *UpdateDataStorageDeliveryRequest {
	s.SubUserId = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetTargetProject(v string) *UpdateDataStorageDeliveryRequest {
	s.TargetProject = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
