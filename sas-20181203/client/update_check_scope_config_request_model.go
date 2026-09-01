// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckScopeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfig(v string) *UpdateCheckScopeConfigRequest
	GetAutoConfig() *string
	SetAutoType(v int32) *UpdateCheckScopeConfigRequest
	GetAutoType() *int32
	SetConfigId(v string) *UpdateCheckScopeConfigRequest
	GetConfigId() *string
	SetResourceDirectoryAccountId(v int64) *UpdateCheckScopeConfigRequest
	GetResourceDirectoryAccountId() *int64
	SetType(v int32) *UpdateCheckScopeConfigRequest
	GetType() *int32
}

type UpdateCheckScopeConfigRequest struct {
	// The automatic scan configuration as a JSON string. The following fields are included:
	//
	// - **autoInclude**: specifies whether to enable automatic scan. Valid values: **true**: enabled. **false**: disabled.
	//
	// - **autoRule**: the enablement configuration.
	//
	// - **ruleOperator**: the enablement configuration rule. Set the value to **include**.
	//
	// - **operator**: the logical operator. Set the value to **or**.
	//
	// - **rule**: the rule.
	//
	// - **condition**: the rule condition. Valid values: **vendor**: vendor. **assetType**: level-1 asset type. **assetSubType**: level-2 asset type.
	//
	// > For more information, refer to the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation.
	//
	// example:
	//
	// "{\\"autoInclude\\":true,\\"autoRule\\":{\\"ruleOperator\\":\\"include\\",\\"operator\\":\\"or\\",\\"rule\\":[{\\"condition\\":\\"assetSubType\\",\\"ruleOperator\\":\\"include\\",\\"value\\":[{\\"vendor\\":\\"0\\",\\"assetType\\":\\"0\\",\\"assetSubType\\":\\"100\\"}]}]}}"
	AutoConfig *string `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The automatic scan configuration type. Valid values:
	//
	// - **0**: disable automatic scan
	//
	// - **1**: automatically scan newly added cloud assets
	//
	// example:
	//
	// 1
	AutoType *int32 `json:"AutoType,omitempty" xml:"AutoType,omitempty"`
	// The ID of the configuration.
	//
	// >Call the [GetCheckScopeConfig](~~GetCheckScopeConfig~~) operation to obtain this parameter.
	//
	// example:
	//
	// 00cfa8161da093089e6804ba6a33****
	ConfigId                   *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The scan scope configuration type. Valid values:
	//
	// - **1**: scan by instance
	//
	// - **3**: scan all
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCheckScopeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckScopeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckScopeConfigRequest) GetAutoConfig() *string {
	return s.AutoConfig
}

func (s *UpdateCheckScopeConfigRequest) GetAutoType() *int32 {
	return s.AutoType
}

func (s *UpdateCheckScopeConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCheckScopeConfigRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *UpdateCheckScopeConfigRequest) GetType() *int32 {
	return s.Type
}

func (s *UpdateCheckScopeConfigRequest) SetAutoConfig(v string) *UpdateCheckScopeConfigRequest {
	s.AutoConfig = &v
	return s
}

func (s *UpdateCheckScopeConfigRequest) SetAutoType(v int32) *UpdateCheckScopeConfigRequest {
	s.AutoType = &v
	return s
}

func (s *UpdateCheckScopeConfigRequest) SetConfigId(v string) *UpdateCheckScopeConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCheckScopeConfigRequest) SetResourceDirectoryAccountId(v int64) *UpdateCheckScopeConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *UpdateCheckScopeConfigRequest) SetType(v int32) *UpdateCheckScopeConfigRequest {
	s.Type = &v
	return s
}

func (s *UpdateCheckScopeConfigRequest) Validate() error {
	return dara.Validate(s)
}
