// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppInstanceShrinkRequest
	GetAppName() *string
	SetAppType(v string) *CreateAppInstanceShrinkRequest
	GetAppType() *string
	SetClientToken(v string) *CreateAppInstanceShrinkRequest
	GetClientToken() *string
	SetComponentsShrink(v string) *CreateAppInstanceShrinkRequest
	GetComponentsShrink() *string
	SetDBInstanceConfigShrink(v string) *CreateAppInstanceShrinkRequest
	GetDBInstanceConfigShrink() *string
	SetDBInstanceName(v string) *CreateAppInstanceShrinkRequest
	GetDBInstanceName() *string
	SetDashboardPassword(v string) *CreateAppInstanceShrinkRequest
	GetDashboardPassword() *string
	SetDashboardUsername(v string) *CreateAppInstanceShrinkRequest
	GetDashboardUsername() *string
	SetDatabase(v string) *CreateAppInstanceShrinkRequest
	GetDatabase() *string
	SetDatabasePassword(v string) *CreateAppInstanceShrinkRequest
	GetDatabasePassword() *string
	SetInitializeWithExistingData(v bool) *CreateAppInstanceShrinkRequest
	GetInitializeWithExistingData() *bool
	SetInstanceClass(v string) *CreateAppInstanceShrinkRequest
	GetInstanceClass() *string
	SetPublicEndpointEnabled(v bool) *CreateAppInstanceShrinkRequest
	GetPublicEndpointEnabled() *bool
	SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceShrinkRequest
	GetPublicNetworkAccessEnabled() *bool
	SetRAGEnabled(v bool) *CreateAppInstanceShrinkRequest
	GetRAGEnabled() *bool
	SetRegionId(v string) *CreateAppInstanceShrinkRequest
	GetRegionId() *string
	SetVSwitchId(v string) *CreateAppInstanceShrinkRequest
	GetVSwitchId() *string
}

type CreateAppInstanceShrinkRequest struct {
	// The name of the AI application to create.
	//
	// example:
	//
	// test-supabase
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Currently, only **supabase*	- is supported.
	//
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The client generates the value of this parameter to prevent duplicate requests from being submitted.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of components.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The database instance configuration.
	DBInstanceConfigShrink *string `json:"DBInstanceConfig,omitempty" xml:"DBInstanceConfig,omitempty"`
	// The instance ID of the RDS PostgreSQL database instance to which the AI application connects.
	//
	// <notice>Only **newly purchased empty RDS PostgreSQL instances*	- are supported. The major engine version must be **17**, and the minor engine version must be **20250630 or later**.
	//
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The password for the Supabase Dashboard.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// example:
	//
	// test_Password
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The username for the Supabase Dashboard.
	//
	// example:
	//
	// supabase
	DashboardUsername *string `json:"DashboardUsername,omitempty" xml:"DashboardUsername,omitempty"`
	// The database name.
	//
	// example:
	//
	// test_database_01
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The password for accessing the RDS database.
	//
	// The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// example:
	//
	// test_Password
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// Specifies whether to restore from existing PostgreSQL data. Valid values:
	//
	// - true: restores from existing data.
	//
	// - false (default): does not restore from existing data.
	//
	// example:
	//
	// false
	InitializeWithExistingData *bool `json:"InitializeWithExistingData,omitempty" xml:"InitializeWithExistingData,omitempty"`
	// The instance type. Currently, only **rdsai.supabase.basic*	- is supported.
	//
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Specifies whether to enable the public network connection endpoint. Valid values:
	//
	// - true (default): Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	PublicEndpointEnabled *bool `json:"PublicEndpointEnabled,omitempty" xml:"PublicEndpointEnabled,omitempty"`
	// Specifies whether to enable the Internet NAT gateway. Valid values:
	//
	// - **true**: enables the Internet NAT gateway.
	//
	// - **false*	- (default): does not enable the Internet NAT gateway.
	//
	// 	Notice: If an Internet NAT gateway is already enabled for the vSwitch specified by the **VSwitchId*	- parameter settings, set this parameter to **false**.
	//
	// example:
	//
	// false
	PublicNetworkAccessEnabled *bool `json:"PublicNetworkAccessEnabled,omitempty" xml:"PublicNetworkAccessEnabled,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// false
	RAGEnabled *bool `json:"RAGEnabled,omitempty" xml:"RAGEnabled,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-9dp2hkpm22gxscfgy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppInstanceShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppInstanceShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *CreateAppInstanceShrinkRequest) GetDBInstanceConfigShrink() *string {
	return s.DBInstanceConfigShrink
}

func (s *CreateAppInstanceShrinkRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateAppInstanceShrinkRequest) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *CreateAppInstanceShrinkRequest) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *CreateAppInstanceShrinkRequest) GetDatabase() *string {
	return s.Database
}

func (s *CreateAppInstanceShrinkRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *CreateAppInstanceShrinkRequest) GetInitializeWithExistingData() *bool {
	return s.InitializeWithExistingData
}

func (s *CreateAppInstanceShrinkRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateAppInstanceShrinkRequest) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *CreateAppInstanceShrinkRequest) GetPublicNetworkAccessEnabled() *bool {
	return s.PublicNetworkAccessEnabled
}

func (s *CreateAppInstanceShrinkRequest) GetRAGEnabled() *bool {
	return s.RAGEnabled
}

func (s *CreateAppInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAppInstanceShrinkRequest) SetAppName(v string) *CreateAppInstanceShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetAppType(v string) *CreateAppInstanceShrinkRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetClientToken(v string) *CreateAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetComponentsShrink(v string) *CreateAppInstanceShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDBInstanceConfigShrink(v string) *CreateAppInstanceShrinkRequest {
	s.DBInstanceConfigShrink = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDBInstanceName(v string) *CreateAppInstanceShrinkRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDashboardPassword(v string) *CreateAppInstanceShrinkRequest {
	s.DashboardPassword = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDashboardUsername(v string) *CreateAppInstanceShrinkRequest {
	s.DashboardUsername = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDatabase(v string) *CreateAppInstanceShrinkRequest {
	s.Database = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDatabasePassword(v string) *CreateAppInstanceShrinkRequest {
	s.DatabasePassword = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetInitializeWithExistingData(v bool) *CreateAppInstanceShrinkRequest {
	s.InitializeWithExistingData = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetInstanceClass(v string) *CreateAppInstanceShrinkRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetPublicEndpointEnabled(v bool) *CreateAppInstanceShrinkRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetPublicNetworkAccessEnabled(v bool) *CreateAppInstanceShrinkRequest {
	s.PublicNetworkAccessEnabled = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetRAGEnabled(v bool) *CreateAppInstanceShrinkRequest {
	s.RAGEnabled = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetRegionId(v string) *CreateAppInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetVSwitchId(v string) *CreateAppInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
