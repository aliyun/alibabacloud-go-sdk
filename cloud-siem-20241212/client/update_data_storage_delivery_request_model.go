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
	// 是否将后续新增且可用的纳管成员自动加入该日志类型的账号选择。true 表示开启，false 表示关闭；关闭不会移除已选择成员。省略保持原值。需同时传 LogCode、ConfigurationVersion、TargetProject，可与 SelectedAccountIds 一起更新，不可同时传 LogDeliveryStatus 或成员 SubUserId。实际投递还受多账号总开关和该日志类型投递开关控制。
	//
	// example:
	//
	// true
	AutoIncludeNewAccounts *bool `json:"AutoIncludeNewAccounts,omitempty" xml:"AutoIncludeNewAccounts,omitempty"`
	// 本次修改所依据的投递设置版本。先调用 GetDataStorage，取 Data.DeliverySettings.ConfigurationVersion 原值传入；示例 1 仅展示格式，不能作为固定版本使用。修改总开关、账号选择、自动纳入设置或指定成员投递时必传。版本缺失或已变化时返回 OperationDenied.CentralizedDeliveryConfigChanged，应重新查询后再提交。
	//
	// example:
	//
	// 1
	ConfigurationVersion *int64 `json:"ConfigurationVersion,omitempty" xml:"ConfigurationVersion,omitempty"`
	// The language of the response. Valid values:
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
	// - enable: Enables log delivery.
	//
	// - disable: Disables log delivery.
	//
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// 该日志类型的完整账号选择列表，元素为十进制 UID 字符串，必须包含 Data.DeliverySettings.DaAccountId（本账号），最多 10000 项。新增成员须在 GetDataStorage 的 Accounts 列表中且 Authorized=true。仅保留本账号时传本账号 UID，不能用空字符串或空数组代替。省略本字段保持原选择。HTTP 请求使用 array(simple) 的逗号分隔格式：SelectedAccountIds=1234567890123456,1234567890123457；单账号只传一个 UID。不使用 SelectedAccountIds.1/.2 或 JSON 数组字符串。需同时传 ConfigurationVersion、TargetProject，并省略 LogDeliveryStatus 和成员 SubUserId。
	//
	// example:
	//
	// 1234567890123456,1234567890123457
	SelectedAccountIds []*string `json:"SelectedAccountIds,omitempty" xml:"SelectedAccountIds,omitempty" type:"Repeated"`
	// 日志来源云账号 UID，使用十进制整数。省略或等于本账号 UID 时操作本账号投递；传成员 UID 时仅对该成员应用当前配置期望的投递状态，需同时传 LogDeliveryStatus、ConfigurationVersion、TargetProject。成员必须已在该类型选择范围中或已有投递记录。保存账号选择或自动纳入设置时应省略本字段，不可传其他成员 UID。
	//
	// example:
	//
	// 1234567890123457
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// 当前日志投递目标 SLS Project 的名称。先调用 GetDataStorage，取 Data.DeliverySettings.TargetProject 原值传入；用于校验本次修改的目标，不用于指定新的 Project。修改总开关、账号选择、自动纳入设置或指定成员投递时必传；缺失或与当前目标不同会返回 OperationDenied.CentralizedDeliveryConfigChanged。示例中的账号 UID 和地域需替换为查询结果。
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
